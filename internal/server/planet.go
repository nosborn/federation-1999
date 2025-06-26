package server

import (
	"math/rand/v2"

	"github.com/dustin/go-humanize"
	"github.com/nosborn/federation-1999/internal/collections"
	"github.com/nosborn/federation-1999/internal/debug"
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/server/build"
	"github.com/nosborn/federation-1999/internal/server/core"
	"github.com/nosborn/federation-1999/internal/server/database"
	"github.com/nosborn/federation-1999/internal/server/exchange"
	"github.com/nosborn/federation-1999/internal/text"
)

type Planet struct {
	companies     []*Company
	disaffection  int32 // population disaffect as a percentage
	education     int32 // education investment
	energy        int32 // energy investment
	exchange      *exchange.Exchange
	exchangeLocNo uint32 // exchange location number
	health        int32  // population healthcare
	hospitalLocNo uint32 // hospital location number
	infstr        int32  // infrastructure investment
	intSecurity   int32  // internal security
	landingLocNo  uint32 // landing pad location number
	level         model.Level
	name          string
	orbitLocNo    uint32 // orbit location number
	population    int32
	routeFlag     uint32
	socialSec     int32 // social security
	synonym       string
	system        Systemer
}

var planetIndex = collections.NewNameIndex[*Planet]()

func FindPlanet(name string) (*Planet, bool) {
	return planetIndex.Find(name)
	// for s := range allSystems.Values() {
	// 	p, ok := s.FindPlanet(name)
	// 	if ok {
	// 		return p, ok
	// 	}
	// }
	// return nil, false
}

func NewCorePlanet(system Systemer, data *core.Planet) *Planet {
	p := &Planet{
		name:          data.Name,
		synonym:       data.Synonym,
		system:        system,
		level:         model.Level(data.Level),
		population:    int32(data.Population),
		landingLocNo:  uint32(data.Landing),
		orbitLocNo:    uint32(data.Orbit),
		exchangeLocNo: uint32(data.Exchange),
		hospitalLocNo: uint32(data.Hospital),
		routeFlag:     data.RouteFlag,
	}
	planetIndex.Insert(p.name, p)
	if p.synonym != "" {
		planetIndex.Insert(p.synonym, p)
	}
	if p.level != model.LevelNoProduction && p.level != model.LevelCapital {
		p.exchange = exchange.NewCoreExchange(p, data)
	}
	return p
}

func NewPlayerPlanet(s *PlayerSystem, dbp model.DBPlanet) *Planet {
	p := newPlanetFromDB(s, dbp)
	planetIndex.Insert(p.name, p)
	return p
}

func newPlanetFromDB(s *PlayerSystem, dbp model.DBPlanet) *Planet {
	p := &Planet{
		disaffection: dbp.Disaffection,       // population disaffect as a percentage
		education:    dbp.Education,          // education investment
		energy:       dbp.Energy,             // energy investment
		health:       dbp.Health,             // population healthcare
		infstr:       dbp.Infstr,             // infrastructure investment
		intSecurity:  dbp.IntSecurity,        // internal security
		level:        model.Level(dbp.Level), // planet's development level
		name:         s.Name(),
		population:   dbp.Population, // current population index
		socialSec:    dbp.SocialSec,  // social security
		system:       s,
	}
	if p.level != model.LevelNoProduction && p.level != model.LevelCapital {
		p.exchange = exchange.NewExchangeFromDB(p, dbp)
	}
	return p
}

func (p *Planet) ActualStock(commodity model.Commodity) int32 {
	if p.exchange == nil {
		return 0
	}
	return p.exchange.ActualStock(commodity)
}

func (p *Planet) AddCompany(c *Company) {
	for i := range p.companies {
		if p.companies[i] == c {
			return
		}
	}
	p.companies = append(p.companies, c)
}

func (p *Planet) AddGoods(commodity model.Commodity, quantity int32) {
	p.exchange.AddStock(commodity, quantity)
}

func (p *Planet) Allocate(points int32, commodity model.Commodity) bool {
	if p.exchange == nil {
		return false
	}
	return p.exchange.Allocate(points, commodity)
}

func (p *Planet) AllocatedProductionPoints() int32 {
	if p.exchange == nil {
		return 0
	}
	return p.exchange.AllocatedProductionPoints()
}

func (p *Planet) AvailableQuantity(commodity model.Commodity) int32 {
	if p.exchange == nil {
		return 0
	}
	return p.exchange.AvailableQuantity(commodity)
}

func (p *Planet) Balance() int32 {
	return p.system.Balance()
}

func (p *Planet) BuyFromfactory(commodity model.Commodity, quantity int32) (int32, bool) {
	if p.exchange == nil {
		return 0, false
	}
	return p.exchange.BuyFromFactory(commodity, quantity)
}

func (p *Planet) BuyGoods(commodity model.Commodity, quantity int32) int32 {
	if p.exchange == nil {
		return 0
	}
	price := p.exchange.BuyingPrice(commodity) * quantity
	p.Expenditure(price)
	if p.exchange.ActualStock(commodity) >= 0 {
		p.exchange.AddStock(commodity, quantity/2)
	} else {
		p.exchange.AddStock(commodity, quantity)
	}
	return price
}

func (p *Planet) BuyingPrice(commodity model.Commodity) int32 {
	if p.exchange == nil {
		return 0
	}
	return p.exchange.BuyingPrice(commodity)
}

// checkPriceHook

func (p *Planet) CompleteBuild(project model.Project) {
	switch project {
	case build.EDUCATION:
		p.education++
	case build.ENERGY:
		p.energy++
	case build.HEALTH:
		p.health++
	case build.INFRA:
		p.infstr++
	case build.SECURITY:
		p.intSecurity++
	default: // make linter happy
	}
}

func (p *Planet) Deallocate(commodity model.Commodity, caller *Player) bool {
	if p.exchange == nil {
		caller.Outputm(text.NO_EXCHANGE_ON_PLANET, p.name)
		return false
	}
	if p.exchange.Deallocate(commodity) {
		caller.Outputm(text.MN328)
		return false
	}
	caller.Outputm(text.MN625)
	return true
}

func (p *Planet) Destroy() {
	// TODO
	if p.synonym != "" {
		planetIndex.Remove(p.synonym)
	}
	planetIndex.Remove(p.name)
}

func (p *Planet) Digest(group model.CommodityGroup, caller *Player) {
	if p.exchange == nil {
		caller.Outputm(text.NO_EXCHANGE_ON_PLANET, p.name)
		return
	}
	p.exchange.Digest(group, caller)
}

func (p *Planet) Display(caller *Player, owner bool, duke bool) {
	debug.Precondition(caller != nil)

	// Common header details.
	if p.system.Duchy().IsSol() {
		caller.outputf("Report for %s - Sol Colony\n", p.name)
	} else {
		caller.outputf("Report for %s - Duchy of %s\n", p.name, p.system.Duchy().Name())
	}
	// if (m_system->isT4Holder()) {
	// 	theCaller->output("Holder of the Tungsten Tourist Trap Trophy.\n");
	// }

	if p.system.IsClosed() {
		caller.outputf(" Development level: %s\n Overlord: %s\n Status: Closed for business\n",
			p.LevelDescription(),
			p.system.Overlord())

		if caller.Rank() >= model.RankHostess && p.system.LastOnline() > 0 {
			caller.outputf(" Last open: %d\n", p.system.LastOnline()+200501)
		}

		if p.system.Owner() == caller || caller.Rank() >= model.RankHostess {
			switch {
			case p.system.IsLoading():
				queuePosition := 0 // FIXME: p.system.loaderQueuePosition()
				if queuePosition > 0 {
					caller.Outputm(text.DisplayPlanetQueuePosition, queuePosition)
				}
			case p.system.IsOffline():
				caller.Output(" Off-line: Yes\n")
			default:
				caller.Output(" Off-line: No\n")
			}
		}
	} else {
		caller.outputf(
			" Development level: %s\n Population level: %s\n Turnover tax base rate: %d%%\n Overlord: %s\n",
			p.LevelDescription(),
			humanize.Comma(int64(p.population)),
			p.system.TaxRate(),
			p.system.Overlord())
	}

	//
	if owner || duke {
		if p.HasExchange() {
			caller.Outputm(text.MN627, p.exchange.StockValue(), p.Balance(), p.exchange.StockDeficit())
		} else {
			caller.Outputm(text.MN627_NoExchange, p.Balance())
		}
	}

	//
	if owner {
		caller.Outputm(text.MN628,
			p.energy,
			p.education,
			p.intSecurity,
			p.socialSec,
			p.infstr,
			p.health,
			p.system.TouristTime(),
			p.disaffection)

		if p.level < model.LevelLeisure && p.level != model.LevelNoProduction {
			caller.Outputm(text.MN268, p.Value(), int32(p.level)*getConfig(CFG_PLANET_PROMO))
		}

		if p.HasExchange() {
			caller.Outputm(text.MN631)
			flag := true
			// for count := range 4 {
			// 	// FIXME:
			// 	if (jobs[count].name[0] != '\0') {
			// 	 flag = false;
			// 	 theCaller->output(mn632,
			// 	 count + 1,
			// 	 goods_array[jobs[count].commodity].name,
			// 	 jobs[count].name,
			// 	 jobs[count].carriage);
			// 	}
			// }
			if flag {
				caller.Outputm(text.MN633)
			}
		}
	}

	//
	if p.HasExchange() {
		var employed, employment int32
		if p.population > 0 {
			employed = p.population - p.LabourPool()
			employment = (100 * employed) / p.population
		}

		caller.outputf(" Employment level: %s (%d%%)\n Minimum wage: %d IG\n",
			humanize.Comma(int64(employed)), employment, p.MinimumWage())

		noFactories := true
		caller.Outputm(text.MN634)

		// for (CCI iter = m_companies.begin(); iter != m_companies.end(); iter++) {
		// 	if ((*iter)->displayPlanetHook(this, theCaller)) {
		// 		noFactories = false;
		// 	}
		// }

		if noFactories {
			caller.Outputm(text.MN633)
		}
	}

	//
	if p.level == model.LevelCapital && !p.system.IsSol() {
		caller.Output("\n")
		for _, s := range p.system.Duchy().AllMembers() {
			if s.Name() == p.system.Duchy().Name() {
				continue
			}
			var closed string
			if s.IsClosed() {
				closed = "(Closed)"
			}
			caller.outputf(" %-15s  %-19s  %s\n",
				s.Planets()[0].Name(),
				s.Planets()[0].LevelDescription(),
				closed)
		}
	}
}

func (p *Planet) DisplayFactories(caller *Player) {
	if p.exchange == nil {
		caller.Outputm(text.NO_EXCHANGE_ON_PLANET, p.name)
		return
	}
	noOutput := true
	// for (CCI iter = m_companies.begin(); iter != m_companies.end(); iter++) {
	// 	if ((*iter)->displayFactoriesHook(this, theCaller)) {
	// 		noOutput = false;
	// 	}
	// }
	if noOutput {
		caller.Output("There isn't any!\n")
	}
}

func (p *Planet) DisplayProduction(caller *Player) {
	if p.exchange == nil {
		caller.Outputm(text.NO_EXCHANGE_ON_PLANET, p.name)
		return
	}
	caller.Outputm(text.MN642)
	p.exchange.DisplayProduction(caller)
	caller.Outputm(text.MN645, p.AllocatedProductionPoints(), p.ProductionPoints())
}

func (p *Planet) Education() int32 {
	return p.education
}

func (p *Planet) Energy() int32 {
	return p.energy
}

func (p *Planet) ExchangeTicks() int {
	return p.system.ExchangeTicks()
}

func (p *Planet) Expenditure(amount int32) {
	p.system.Expenditure(amount)
}

func (p *Planet) FeedWorkthings() int32 {
	if p.exchange == nil {
		return 0
	}

	var munchies []model.Commodity
	switch p.level {
	case model.LevelAgricultural:
		munchies = []model.Commodity{
			model.CommodityTools,
			model.CommodityPowerPacks,
			model.CommodityControllers,
			model.CommodityPharms,
			model.CommodityPropellants,
			model.CommodityRads,
			model.CommodityGold,
			model.CommodityGames,
		}
	case model.LevelMining:
		munchies = []model.Commodity{
			model.CommodityPetros,
			model.CommodityCereals,
			model.CommodityTextiles,
			model.CommodityMeat,
			model.CommodityFruit,
			model.CommoditySoya,
			model.CommodityKatydidics,
			model.CommodityMusiks,
			model.CommodityHolos,
		}
	case model.LevelIndustrial:
		munchies = []model.Commodity{
			model.CommodityElectros,
			model.CommodityMechParts,
			model.CommodityHides,
			model.CommodityTextiles,
			model.CommodityMeat,
			model.CommodityFruit,
			model.CommodityStock,
			model.CommodityCrystals,
			model.CommoditySensAmps,
			model.CommodityHolos,
			model.CommodityUnivators,
			model.CommoditySims,
		}
	case model.LevelTechnological:

		munchies = []model.Commodity{
			model.CommodityPolymers,
			model.CommodityPetros,
			model.CommodityNitros,
			model.CommodityMunitions,
			model.CommodityMechParts,
			model.CommodityCereals,
			model.CommodityWoods,
			model.CommodityHides,
			model.CommodityMeat,
			model.CommoditySpices,
			model.CommodityFruit,
			model.CommoditySoya,
			model.CommodityGold,
			model.CommodityStudios,
			model.CommodityLibraries,
		}
	case model.LevelLeisure:
		munchies = []model.Commodity{
			model.CommoditySynths,
			model.CommodityPharms,
			model.CommodityRNA,
			model.CommodityCereals,
			model.CommodityWoods,
			model.CommodityHides,
			model.CommoditySpices,
			model.CommodityFurs,
			model.CommodityCrystals,
			model.CommodityGold,
			model.CommodityHypnotapes,
			model.CommoditySensAmps,
		}
	default:
		return 0
	}

	pop := p.population / 40
	if pop <= 0 {
		return 0
	}

	income := int32(0)
	for i := range munchies {
		if p.exchange.ActualStock(munchies[i]) > 0 {
			consumption := rand.Int32N(pop) //nolint:gosec // "It's Just A Game"
			income += (consumption * p.exchange.ForcedSellingPrice(munchies[i]))
			p.exchange.RemoveStock(munchies[i], consumption)
		}
	}
	return income
}

func (p *Planet) GetStockpileLevel(commodity model.Commodity) int32 {
	if p.exchange == nil {
		return 0
	}
	return p.exchange.GetStockpileLevel(commodity)
}

func (p *Planet) HasExchange() bool {
	switch p.level {
	case model.LevelNoProduction, model.LevelCapital:
		return false
	default:
		return p.exchange != nil
	}
}

func (p *Planet) Income(amount int32, taxable bool) {
	p.system.Income(amount, taxable)
}

func (p *Planet) Investment(group model.CommodityGroup) int32 {
	switch group {
	case model.GroupMining, model.GroupIndustrial:
		return p.energy
	case model.GroupTechnological, model.GroupLeisure:
		return p.education
	default:
		return 0
	}
}

func (p *Planet) IsClosed() bool {
	return p.system.IsClosed()
}

func (p *Planet) IsHidden() bool {
	return p.system.IsHidden()
}

func (p *Planet) IsOwnerPlaying() bool {
	return p.system.IsOwnerPlaying()
}

func (p *Planet) LabourPool() int32 {
	pool := p.population
	for i := range p.companies {
		pool -= p.companies[i].Employees(p)
	}
	return pool
}

func (p *Planet) LevelDescription() string {
	switch p.level {
	case model.LevelAgricultural:
		return text.Msg(text.Level_Agricultural)
	case model.LevelMining:
		return text.Msg(text.Level_Mining)
	case model.LevelIndustrial:
		return text.Msg(text.Level_Industrial)
	case model.LevelTechnological:
		return text.Msg(text.Level_Technological)
	case model.LevelLeisure:
		return text.Msg(text.Level_Leisure)
	case model.LevelCapital:
		return text.Msg(text.Level_Capital)
	default:
		return text.Msg(text.Level_NoProduction)
	}
}

func (p *Planet) MakeNoProduction() {
	if p.exchange != nil {
		// delete m_exchange; -- FIXME
		p.exchange = nil
	}
	p.level = model.LevelNoProduction
	p.population = 1000
	p.exchange = nil
	p.energy = 0
	p.education = 0
	p.socialSec = 0
	p.infstr = 0
	p.health = 0
	p.intSecurity = 0
	p.disaffection = 0
}

func (p *Planet) MinimumWage() int32 {
	wage := int32((90 + (int(p.level) * 10)) / 2)
	if p.socialSec > 0 {
		wage += (p.socialSec * 2)
	}
	return wage
}

func (p *Planet) Name() string { // CHECK THIS
	return p.name
}

func (p *Planet) Owner() *Player {
	return p.system.Owner()
}

func (p *Planet) ProductionPoints() int32 {
	return int32(15 + (5 * int(p.level))) // LevelAgricultural had better be 1!
}

func (p *Planet) RemoveCompany(c *Company) {
	for i := range p.companies {
		if p.companies[i] == c {
			p.companies = append(p.companies[:i], p.companies[i+1:]...)
			return
		}
	}
}

func (p *Planet) RemoveGoods(commodity model.Commodity, quantity int32) {
	// TODO
}

func (p *Planet) ReturnGoods(job *model.Work) {
	p.AddGoods(job.Pallet.Type, job.Pallet.Quantity)
	p.Income((job.Pallet.Quantity*job.Value)/2, false) // Return 50%
}

func (p *Planet) RouteFlag() uint32 {
	return p.routeFlag
}

func (p *Planet) Save(when database.SaveWhen) {
	debug.Check(p.system != nil)
	debug.Trace("Planet.Save(%s)", p.system.Name())
	p.system.Save(when)
}

func (p *Planet) SellToFactory(commodity model.Commodity, wanted int32) (int32, int32, bool) {
	if p.exchange == nil {
		return 0, 0, false
	}
	return p.exchange.SellToFactory(commodity, wanted)
}

func (p *Planet) SellingPrice(commodity model.Commodity) int32 {
	if p.exchange == nil {
		return 0
	}
	return p.exchange.SellingPrice(commodity)
}

func (p *Planet) Serialize(dbp *model.DBPlanet) {
	if p.exchange != nil {
		p.exchange.Serialize(dbp)
	}

	dbp.Level = int32(p.level)
	dbp.Population = p.population
	dbp.Energy = p.energy
	dbp.Education = p.education
	dbp.SocialSec = p.socialSec
	dbp.Infstr = p.infstr
	dbp.Health = p.health
	dbp.IntSecurity = p.intSecurity
	dbp.Disaffection = p.disaffection
}

// setMarkup

func (p *Planet) SetMilkrun(slot int, destination string, commodity model.Commodity, carriage int32) {
	if p.exchange == nil {
		return
	}
	p.exchange.SetMilkrun(slot, destination, commodity, carriage)
}

func (p *Planet) SetStockpileLevel(commodity model.Commodity, level int32) {
	if p.exchange == nil {
		return
	}
	p.exchange.SetStockpileLevel(commodity, level)
}

func (p *Planet) StartExchange() {
	if p.exchange != nil {
		p.exchange.Start()
	}
}

func (p *Planet) StartOfDay() {
	// TODO
}

func (p *Planet) StopExchange() {
	if p.exchange != nil {
		p.exchange.Stop()
	}
}

func (p *Planet) StopFactories() {
	// TODO
}

func (p *Planet) System() Systemer {
	return p.system
}

func (p *Planet) SystemName() string {
	return p.system.Name()
}

func (p *Planet) Value() int32 {
	return p.energy + p.education + p.infstr + p.health + p.intSecurity
}
