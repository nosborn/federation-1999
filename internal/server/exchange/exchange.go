package exchange

import (
	"math/rand/v2"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/monitoring"
	"github.com/nosborn/federation-1999/internal/server/core"
	"github.com/nosborn/federation-1999/internal/server/database"
	"github.com/nosborn/federation-1999/internal/server/global"
	"github.com/nosborn/federation-1999/internal/server/goods"
	"github.com/nosborn/federation-1999/internal/text"
)

const (
	exchangeTimerPeriod = 90 * time.Second
)

const (
	EXCHANGE_MIN_STOCKPILE = 20
	EXCHANGE_MAX_STOCKPILE = 10000
)

type Bin struct {
	production  int32 // total production per cycle
	produce     bool  // produce this cycle?
	stock       int32 // stock on hand
	stockpile   int32 // stockpile level
	consumption int32 // total consumption per cycle
	markup      int32 // percentage markup on this commodity
	produced    int32 // production so far this cycle
	consumed    int32 // consumption so far this cycle
}

type Exchange struct {
	deficitLimit int32
	exTime       int32
	flux         int32
	goods        [52]Bin
	jobs         [4]model.PlanetJob
	markup       int32
	parent       ExchangeParent
	production   [5][2]int32
	timer        *time.Timer
}

type ExchangeParent interface {
	Balance() int32
	Education() int32
	Energy() int32
	ExchangeTicks() int
	Expenditure(int32)
	FeedWorkthings() int32
	Income(int32, bool)
	Investment(model.CommodityGroup) int32
	IsClosed() bool
	IsOwnerPlaying() bool
	Name() string
	Save(database.SaveWhen)
	SystemName() string
}

type ExchangeUser interface {
	Outputm(text.MsgNum, ...any)
}

// type PlanetJob struct { // used to store owner generated milkruns
// 	name      string          // planet to deliver to
// 	commodity model.Commodity // type of goods to deliver
// 	carriage  int32           // IG/ton for the hauler
// }

var Production [52]int

func NewCoreExchange(parent ExchangeParent, data *core.Planet) *Exchange {
	e := &Exchange{
		deficitLimit: -100,
		markup:       10,
		parent:       parent,
	}
	for i := range e.goods {
		e.goods[i] = Bin{
			production:  int32(data.Goods[i][0]),
			produce:     (data.Goods[i][1] == 1),
			stock:       int32(data.Goods[i][2]),
			stockpile:   int32(data.Goods[i][3]),
			consumption: int32(data.Goods[i][4]),
			markup:      int32(data.Goods[i][5]),
		}
	}
	e.Start() // This is perhaps a little too much knowledge!
	return e
}

func NewExchange(parent ExchangeParent, goods [52][6]int32) *Exchange {
	// log.Printf("NewExchange(%s)", parent.name)

	e := Exchange{
		deficitLimit: -100, // FIXME
		exTime:       0,    // FIXME
		flux:         0,    // FIXME
		markup:       10,
		parent:       parent,
	}
	for i := range e.goods {
		e.goods[i].production = goods[i][0]
		if goods[i][1] == 0 {
			e.goods[i].produce = false
		} else {
			e.goods[i].produce = true
		}
		e.goods[i].stock = goods[i][2]
		e.goods[i].stockpile = goods[i][3]
		e.goods[i].consumption = goods[i][4]
		e.goods[i].markup = goods[i][5]
	}
	return &e
}

func NewExchangeFromDB(parent ExchangeParent, dbp model.DBPlanet) *Exchange {
	e := Exchange{
		deficitLimit: -1000,
		exTime:       dbp.ExTime,
		flux:         dbp.Flux,
		markup:       dbp.Markup,
		parent:       parent,
		production:   dbp.Production,
	}
	// for i := range dbp.Jobs {
	// 	// It seems we don't actually do anything here.
	// }
	for i := range e.goods {
		e.goods[i] = Bin{
			production:  dbp.Goods[i].Production,
			produce:     (dbp.Goods[i].Produce == 1),
			stock:       dbp.Goods[i].Stock,
			stockpile:   dbp.Goods[i].Stockpile,
			consumption: dbp.Goods[i].Consumption,
			markup:      dbp.Goods[i].Markup,
			produced:    dbp.Goods[i].Produced,
			consumed:    dbp.Goods[i].Consumed,
		}
	}
	return &e
}

func (e *Exchange) ActualStock(commodity model.Commodity) int32 {
	return e.goods[commodity].stock
}

func (e *Exchange) AddStock(commodity model.Commodity, amount int32) {
	e.goods[commodity].stock += amount
}

func (e *Exchange) Allocate(points int32, commodity model.Commodity) bool {
	// size_t slot;
	//
	//	for (slot = 0; slot < DIM_OF(m_production); slot++) {
	//		if (m_production[slot][1] == commodity) {
	//			break;
	//		}
	//	}
	//
	//	if (slot == DIM_OF(m_production)) {
	//		for (slot = 0; slot < DIM_OF(m_production); slot++) {
	//			if (m_production[slot][0] == 0) {
	//				break;
	//			}
	//		}
	//	}
	//
	//	if (slot == DIM_OF(m_production)) {
	//		return false;
	//	}
	//
	// m_production[slot][0]         += points;
	// m_production[slot][1]          = commodity;
	// m_goods[commodity].production += points;

	return true
}

func (e *Exchange) AllocatedProductionPoints() int32 {
	var allocated int32
	for i := range e.production {
		if e.production[i][0] > 0 {
			allocated += e.production[i][0]
		}
	}
	return allocated
}

// Computes the max quantity an exchange is selling of a commodity to players.
func (e *Exchange) AvailableQuantity(commodity model.Commodity) int32 {
	if e.goods[commodity].stock <= e.goods[commodity].stockpile {
		return 0
	}
	if e.goods[commodity].stock < e.goods[commodity].consumption/2 {
		return 0
	}
	return e.goods[commodity].stock - e.goods[commodity].stockpile
}

func (e *Exchange) BuyFromFactory(commodity model.Commodity, quantity int32) (int32, bool) {
	price := e.BuyingPrice(commodity)
	e.goods[commodity].stock += quantity
	return price, true
}

func (e *Exchange) BuyingPrice(commodity model.Commodity) int32 {
	// Exchange buys at 10 IG regardless if the treasury is broke.
	if e.parent.Balance() <= 0 {
		return 10
	}

	markdown := ((100 - e.markup - e.goods[commodity].markup) * int32(goods.GoodsArray[commodity].BasePrice)) / 100
	var price int32

	switch {
	case e.goods[commodity].stock < 0:
		price = markdown * 2
	case e.goods[commodity].stock <= e.goods[commodity].stockpile:
		price = (markdown * 3) / 2
	case e.goods[commodity].stock < e.goods[commodity].consumption:
		price = markdown
	default:
		return 10
	}

	price += (price * e.flux) / 100
	return price
}

func (e *Exchange) Deallocate(commodity model.Commodity) bool {
	found := false
	for i := range e.production {
		if e.production[i][0] > 0 && e.production[i][1] == int32(commodity) { // FIXME: int()
			e.goods[commodity].production -= e.production[i][0]
			e.production[i] = [2]int32{0, 0}
			found = true
		}
	}
	return found
}

func (e *Exchange) Digest(group model.CommodityGroup, caller ExchangeUser) {
	prices := "?????"
	switch e.flux / 8 {
	case -5:
		prices = "-----"
	case -4:
		prices = "----"
	case -3:
		prices = "---"
	case -2:
		prices = "--"
	case -1:
		prices = "-"
	case 0:
		prices = "..."
	case 1:
		prices = "+"
	case 2:
		prices = "++"
	case 3:
		prices = "+++"
	case 4:
		prices = "++++"
	case 5:
		prices = "+++++"
	}
	caller.Outputm(text.DigestHeader, e.parent.Name(), e.exTime, e.markup, prices)

	investment := e.parent.Investment(group)
	for i := range e.goods {
		if goods.GoodsArray[i].Group != group {
			continue
		}
		var producing string
		if e.goods[i].production+investment == 0 {
			producing = "---"
		} else {
			if e.goods[i].produce {
				producing = "Yes"
			} else {
				producing = " No"
			}
		}
		caller.Outputm(text.DigestLine,
			goods.GoodsArray[i].Name,
			humanize.Comma(int64(e.goods[i].stockpile)),
			e.goods[i].production,
			investment,
			e.goods[i].consumption,
			humanize.Comma(int64(e.goods[i].stock)),
			producing,
			e.goods[i].markup)
	}
}

func (e *Exchange) DisplayProduction(caller ExchangeUser) {
	for i := range e.production {
		if e.production[i][0] > 0 {
			caller.Outputm(text.MN643, i+1, e.production[i][0], goods.GoodsArray[e.production[i][1]].Name)
		} else {
			caller.Outputm(text.MN644, i+1)
		}
	}
}

func (e *Exchange) Expenditure(amount int32) {
	e.parent.Expenditure(amount)
}

func (e *Exchange) ForcedSellingPrice(commodity model.Commodity) int32 {
	var price int32
	markup := ((100 + e.markup + e.goods[commodity].markup) * int32(goods.GoodsArray[commodity].BasePrice)) / 100

	switch {
	case e.goods[commodity].stock <= e.goods[commodity].stockpile:
		price = (markup * 4) / 3 // + 1/3rd of markup
	case e.goods[commodity].stock < e.goods[commodity].consumption:
		price = markup
	default:
		price = (markup * 2) / 3 // - 1/3rd of markup
	}

	price += (price * e.flux) / 100
	return price
}

// func (e *Exchange) GetCommodityMarkup(commodity int) int32 {
// 	return e.goods[commodity].markup
// }

// func (e *Exchange) GetMarkup() int32 {
// 	return e.markup
// }

func (e *Exchange) GetStockpileLevel(commodity model.Commodity) int32 {
	return e.goods[commodity].stockpile
}

func (e *Exchange) Income(amount int32, isTaxable bool) {
	e.parent.Income(amount, isTaxable)
}

func (e *Exchange) RemoveStock(commodity model.Commodity, amount int32) {
	e.goods[commodity].stock -= amount
}

func (e *Exchange) SellToFactory(commodity model.Commodity, wanted int32) (int32, int32, bool) {
	available := min(e.goods[commodity].stock, wanted)
	if available <= 0 {
		return 0, 0, false
	}
	price := e.SellingPrice(commodity)
	if price == 0 { // Buy it out of the stockpile.
		price = ((100 + e.goods[commodity].consumption) * int32(goods.GoodsArray[commodity].BasePrice)) / 100
	}
	e.goods[commodity].stock -= available
	// Hmmm... no income for the Exchange?
	return available, price, true
}

func (e *Exchange) SellingPrice(commodity model.Commodity) int32 {
	if e.goods[commodity].stock <= e.goods[commodity].stockpile {
		return 0
	}

	markup := ((100 + e.markup + e.goods[commodity].markup) * int32(goods.GoodsArray[commodity].BasePrice)) / 100
	price := int32(0)
	if e.goods[commodity].stock < e.goods[commodity].consumption {
		price = markup
	} else {
		price = (markup * 2) / 3
	}
	price += (price * e.flux) / 100
	return price
}

func (e *Exchange) Serialize(dbp *model.DBPlanet) {
	for i := range e.jobs {
		dbp.Jobs[i] = model.DBPlanetJob{
			Commodity: int32(e.jobs[i].Commodity),
			Carriage:  e.jobs[i].Carriage,
		}
		copy(dbp.Jobs[i].Name[:], e.jobs[i].Name)
	}
	dbp.ExTime = e.exTime
	dbp.Flux = e.flux
	dbp.Markup = e.markup
	for i := range e.production {
		dbp.Production[i][0] = e.production[i][0]
		dbp.Production[i][1] = e.production[i][1]
	}
	for i := range e.goods {
		produce := 0
		if e.goods[i].produce {
			produce = 1
		}
		dbp.Goods[i] = model.DBBin{
			Production:  e.goods[i].production,
			Produce:     int32(produce),
			Stock:       e.goods[i].stock,
			Stockpile:   e.goods[i].stockpile,
			Consumption: e.goods[i].consumption,
			Markup:      e.goods[i].markup,
			Produced:    e.goods[i].produced,
			Consumed:    e.goods[i].consumed,
		}
	}
}

// func (e *Exchange) SetCommodityMarkup(commodity int, level int32) {
// 	e.goods[commodity].markup = level
// }

// func (e *Exchange) SetMarkup(level int32) {
// 	e.markup = level
// }

func (e *Exchange) SetMilkrun(slot int, destination string, commodity model.Commodity, carriage int32) {
	e.jobs[slot-1].Name = destination
	e.jobs[slot-1].Commodity = commodity
	e.jobs[slot-1].Carriage = carriage
}

func (e *Exchange) SetStockpileLevel(commodity model.Commodity, level int32) {
	e.goods[commodity].stockpile = max(0, min(level, EXCHANGE_MAX_STOCKPILE))
}

func (e *Exchange) Start() {
	if e.timer == nil {
		e.timer = time.AfterFunc(exchangeTimerPeriod, e.timerHandler)
	}
}

func (e *Exchange) StockDeficit() int32 {
	deficit := int32(0)
	for i := range e.goods {
		if e.goods[i].stock < 0 {
			deficit -= (e.goods[i].stock * int32(goods.GoodsArray[i].BasePrice))
		}
	}
	return deficit
}

func (e *Exchange) StockValue() int32 {
	value := int32(0)
	for i := range e.goods {
		if e.goods[i].stock > 0 {
			value += e.goods[i].stock * int32(goods.GoodsArray[i].BasePrice)
		}
	}
	return value
}

func (e *Exchange) Stop() {
	if e.timer != nil {
		e.timer.Stop()
		e.timer = nil
	}
}

func (e *Exchange) WorkthingConsumption(munchies []model.Commodity, population int32) int32 {
	population /= 40
	if population <= 0 {
		return 0
	}
	income := int32(0)
	for _, commodity := range munchies {
		if e.ActualStock(commodity) > 0 {
			consumption := rand.Int32N(population) //nolint:gosec // "It's Just A Game"
			income += (consumption * e.ForcedSellingPrice(commodity))
			e.RemoveStock(commodity, consumption)
		}
	}
	return income
}

func (e *Exchange) run(ticks int, isOwnerPlaying bool, energy, education int32) int32 {
	var income int32

	// Production...
	for i := range e.goods {
		// Update stock in hand.
		if e.goods[i].produce {
			var production int32
			switch goods.GoodsArray[i].Group {
			case model.GroupMining, model.GroupIndustrial:
				production = e.goods[i].production + energy
			case model.GroupTechnological, model.GroupLeisure:
				production = e.goods[i].production + education
			default:
				production = e.goods[i].production
			}

			target := (production * (e.exTime + int32(ticks))) / 100
			if produce := target - e.goods[i].produced; produce > 0 {
				if isOwnerPlaying {
					produce += (produce / 2) // 150%
				}
				e.goods[i].stock += produce
				e.goods[i].produced = target // DON'T just add produce!
				income -= produce * ((e.BuyingPrice(model.Commodity(i)) * 9) / 10)
				Production[i] += int(produce)
			}
		}

		// Consumption...
		target := (e.goods[i].consumption * (e.exTime + int32(ticks))) / 100
		if consume := target - e.goods[i].consumed; consume > 0 {
			if isOwnerPlaying {
				consume += (consume / 2) // 150%
			}

			e.goods[i].stock -= consume
			e.goods[i].consumed = target // DON'T just subtract consume!

			if e.goods[i].stock > 0 { // Has stock to sell.
				price := e.SellingPrice(model.Commodity(i))
				if price == 0 {
					price = int32(goods.GoodsArray[i].BasePrice) * 2
				}
				income += (consume * price)
				Production[i] -= int(consume)
			}
		}
	}
	return income
}

func (e *Exchange) timerHandler() {
	global.Lock()
	defer global.Unlock()

	monitoring.ExchangeTimerTickTotal.WithLabelValues(e.parent.SystemName(), e.parent.Name()).Inc()

	defer database.CommitDatabase()
	e.timerProc()
}

func (e *Exchange) timerProc() {
	// log.Printf("timerProc for %s exchange", e.parent.Name())

	e.timer = nil

	if e.parent.IsClosed() {
		return
	}

	ticks := e.parent.ExchangeTicks()
	if ticks > 0 {
		income := int32(0)

		isOwnerPlaying := e.parent.IsOwnerPlaying()
		energy := e.parent.Energy()
		education := e.parent.Education()

		// Run to end of cycle.
		if int(e.exTime)+ticks > 100 {
			income += e.run(int(100-e.exTime), isOwnerPlaying, energy, education)
			income += e.parent.FeedWorkthings()

			ticks -= int(100 - e.exTime)

			e.exTime = 0
			e.flux += rand.Int32N(61) - 30 //nolint:gosec // "It's Just A Game"

			// Clear the per-cycle counters.
			for i := range e.goods {
				e.goods[i].produced = 0
				e.goods[i].consumed = 0
			}
		}

		// Run part of cycle.
		if ticks > 0 {
			income += e.run(ticks, isOwnerPlaying, energy, education)

			e.exTime += int32(ticks)
			e.flux += rand.Int32N(11) - 5 //nolint:gosec // "It's Just A Game"
		}

		// Keep flux within limits.
		if e.flux > 40 {
			e.flux = 40
		} else if e.flux < -40 {
			e.flux = -40
		}

		// See if we want to turn the production off.
		for i := range e.goods {
			switch {
			case e.goods[i].stock < 10:
				e.goods[i].produce = true

				// Keep deficits within limits.
				if e.goods[i].stock < e.deficitLimit {
					e.goods[i].stock = e.deficitLimit
				}
			case e.goods[i].stock >= e.goods[i].stockpile*2 || e.goods[i].stock >= EXCHANGE_MAX_STOCKPILE:
				e.goods[i].produce = false
			default:
				e.goods[i].produce = true
			}
		}

		e.parent.Income(income, true)
		e.parent.Save(database.SaveLater)
	}

	e.timer = time.AfterFunc(exchangeTimerPeriod, e.timerHandler)
}
