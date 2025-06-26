package server

import (
	"log"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/nosborn/federation-1999/internal/collections"
	"github.com/nosborn/federation-1999/internal/debug"
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/monitoring"
	"github.com/nosborn/federation-1999/internal/server/database"
	"github.com/nosborn/federation-1999/internal/server/global"
	"github.com/nosborn/federation-1999/internal/text"
)

const (
	companyCost        = 2000000 // 2,000,000 IG
	companyTimerPeriod = 20 * time.Second
)

var (
	allCompanies = collections.NewOrderedCollection[*Company]()
	companyIndex = collections.NewNameIndex[*Company]()
)

type Company struct {
	Balance int32 // Company's balance
	CEO     *Player
	Name    string
	Shares  int32 // Number of shares issued
	expend  int32 // Expenditure to date
	income  int32 // Income to date
	timer   *time.Timer

	curFactory *Factory
	factories  [model.MAX_FACTORY]*Factory
}

func AllCompanies() []*Company {
	return allCompanies.All()
}

func FindCompany(name string) (*Company, bool) {
	return companyIndex.Find(name)
}

func NewCompany(ceo *Player, name string) *Company {
	c := &Company{
		CEO:    ceo,
		Name:   name,
		Shares: 2,
	}
	if err := allCompanies.Insert(c); err != nil {
		log.Panic("PANIC: Duplicate company added: ", err)
	}
	companyIndex.Insert(c.Name, c)
	return c
}

func NewCompanyFromDB(ceo *Player, dbCompany model.DBCompany, dbFactories [model.MAX_FACTORY]model.DBFactory) *Company {
	c := &Company{
		Balance: dbCompany.Balance,
		CEO:     ceo,
		Name:    text.CStringToString(dbCompany.Name[:]),
		Shares:  dbCompany.Capital,
		expend:  dbCompany.Expend,
		income:  dbCompany.Income,
	}
	if err := allCompanies.Insert(c); err != nil {
		log.Panic("PANIC: Duplicate company added: ", err)
	}
	companyIndex.Insert(c.Name, c)

	var needsSave bool

	for i := range dbFactories {
		planetName := text.CStringToString(dbFactories[i].Planet[:])
		if planetName != "" {
			planet, ok := FindPlanet(planetName)
			if !ok {
				log.Printf("Liquidating factory on %s", planetName)
				needsSave = true
				continue
			}
			f := NewFactoryFromDB(c, i, &dbFactories[i], planet)
			f.Planet().AddCompany(c)
			c.factories[i] = f
		}
	}

	if needsSave {
		c.Save(database.SaveLater)
	}
	return c
}

func (c *Company) ChangeBalance(amount int32) {
	changeBalance(&c.Balance, amount)
}

func (c *Company) ClearBooks() {
	c.expend = 0
	c.income = 0
}

func (c *Company) Destroy() {
	c.Stop()

	for i := range model.MAX_FACTORY {
		if c.factories[i] != nil {
			c.factories[i].Planet().RemoveCompany(c)
		}
	}
}

// destroyFactories
// destroyFactory

func (c *Company) Display(caller *Player) {
	caller.Outputm(text.DisplayCompany,
		c.Name,
		humanize.Comma(int64(c.Shares)),
		humanize.Comma(int64(c.income)),
		humanize.Comma(int64(c.expend)),
		humanize.Comma(int64(c.Profit())),
		humanize.Comma(int64(c.Balance)))

	noFactories := true

	// for (size_t i = 0; i < MAX_FACTORY; ++i) {
	// 	if (factories[i] != NULL) {
	// 		const size_t messageIndex =
	// 		factories[i]->planet()->isClosed() ? 2 : 3;
	//
	// 		theCaller->output(factories[i]->product().message[messageIndex],
	// 		i + 1,
	// 		factories[i]->product().name,
	// 		factories[i]->planet()->name());
	//
	// 		noFactories = false;
	// 	}
	// }

	if noFactories {
		caller.Outputm(text.MN1058)
	}
}

// displayFactoriesHook

func (c *Company) DisplayPlanetHook(planet *Planet, caller *Player) bool {
	debug.Trace("Company.DisplayPlanetHook(%s)", planet.Name())

	if !c.CEO.IsPlaying() {
		return false
	}

	for i, f := range c.factories {
		if f == nil || f.Planet() != planet {
			continue
		}
		caller.Outputm(f.product().Message[0], c.Name, i+1, f.product().Name)
	}
	return true
}

func (c *Company) Employees(onPlanet *Planet) int32 {
	if c.CEO.IsPlaying() {
		return 0
	}
	var workers int32
	for i := range c.factories {
		if c.factories[i] != nil && c.factories[i].Planet() == onPlanet {
			workers += c.factories[i].Workers()
		}
	}
	return workers
}

func (c *Company) Expend(amount int32) {
	changeBalance(&c.Balance, -amount)
	changeBalance(&c.expend, amount)
}

func (c *Company) FindFactory(number int32) (*Factory, bool) {
	if number == -1 { // Use existing current factory
		if c.curFactory == nil {
			c.CEO.Outputm(text.MN393)
			return nil, false
		}
	} else { // Select a factory.
		c.curFactory = nil

		if number < 1 || number > model.MAX_FACTORY {
			c.CEO.Outputm(text.MN379)
			return nil, false
		}

		for i := range model.MAX_FACTORY {
			if c.factories[i] != nil && c.factories[i].number == number-1 {
				c.curFactory = c.factories[i]
				break
			}
		}

		if c.curFactory == nil {
			c.CEO.Outputm(text.MN1053)
			return nil, false
		}
	}
	return c.curFactory, true
}

func (c *Company) Income(amount int32) {
	changeBalance(&c.Balance, amount)
	changeBalance(&c.income, amount)
	c.CEO.CheckForPromotion()
}

// issueShares

func (c *Company) Profit() int32 {
	return c.income - c.expend
}

func (c *Company) Save(when database.SaveWhen) {
	debug.Trace("Company.Save(%s,%d)", c.Name, when)
	c.CEO.Save(when)
}

func (c *Company) Serialize(dbc *model.DBCompany, dbf *[12]model.DBFactory) {
	for i := range model.MAX_FACTORY {
		if c.factories[i] != nil {
			c.factories[i].Serialize(&dbf[i])
		}
	}
	copy(dbc.Name[:], c.Name)
	dbc.Balance = c.Balance
	dbc.Capital = c.Shares
	dbc.Income = c.income
	dbc.Expend = c.expend
}

func (c *Company) Start() {
	if c.timer == nil {
		c.timer = time.AfterFunc(companyTimerPeriod, c.timerHandler)
	}
}

// Stop factories when the owner leaves the game.
func (c *Company) Stop() {
	if c.timer != nil {
		c.timer.Stop()
		c.timer = nil
	}
	// TODO
}

// Stop factories when a planet closes or starts a build.
func (c *Company) StopFactories(onPlanet *Planet) {
	for i := range model.MAX_FACTORY {
		if c.factories[i] != nil && c.factories[i].Planet() == onPlanet {
			c.factories[i].Stop()
		}
	}
}

func (c *Company) timerHandler() {
	global.Lock()
	defer global.Unlock()

	monitoring.CompanyTimerTickTotal.Inc()

	defer database.CommitDatabase()
	c.timerProc()
}

func (c *Company) timerProc() {
	log.Printf("Company timer proc for %s", c.Name)

	c.timer = nil

	if !c.CEO.IsPlaying() { // Belt and braces!
		return
	}

	needsSave := false
	for _, f := range c.factories {
		if f != nil && f.Run() {
			needsSave = true
		}
	}
	if needsSave {
		c.CEO.Save(database.SaveNow)
	}

	c.timer = time.AfterFunc(companyTimerPeriod, c.timerHandler)
}
