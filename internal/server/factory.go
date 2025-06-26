package server

import (
	"log"

	"github.com/nosborn/federation-1999/internal/debug"
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/server/goods"
	"github.com/nosborn/federation-1999/internal/text"
)

type Factory struct {
	carriage   int32
	commodity  model.Commodity
	company    *Company
	contracted int32
	cycle      int32
	delivered  int32
	delivery   model.Delivery
	expend     int32
	income     int32
	layoff     int32
	number     int32
	opStock    int32
	planet     *Planet
	stock      [model.MAX_INPUTS]int32
	toFactory  int32
	wages      int32
	warehouse  *model.Warehouse
	workers    int32

	// Probably not needed:
	// uid        uint32
}

var FactoryProduction [52]int

func GetFactory(id model.FactoryID) *Factory {
	company, ok := FindCompany(id.Owner)
	if !ok {
		return nil
	}
	factory := company.factories[id.Number]
	if factory == nil {
		return nil
	}
	return factory
}

func NewFactoryFromDB(company *Company, number int, dbFactory *model.DBFactory, planet *Planet) *Factory {
	f := Factory{
		carriage:   dbFactory.Carriage,
		commodity:  model.Commodity(dbFactory.Product),
		company:    company,
		contracted: dbFactory.Contracted,
		cycle:      dbFactory.Cycle,
		delivered:  dbFactory.Delivered,
		delivery:   model.Delivery(dbFactory.Delivery),
		expend:     dbFactory.Expend,
		income:     dbFactory.Income,
		layoff:     dbFactory.Layoff,
		number:     int32(number),
		opStock:    dbFactory.OpStock,
		planet:     planet,
		toFactory:  dbFactory.ToFactory,
		wages:      dbFactory.Wages,
		warehouse:  company.CEO.FindWarehouse(planet.Name()),
	}
	if f.delivery == model.DeliverWarehouse && f.warehouse == nil {
		f.delivery = model.DeliverExchange
	}
	for i := range f.stock {
		f.stock[i] = dbFactory.Stock[i]
	}
	return &f
}

func (f *Factory) Clear(caller *Player) {
	if f.opStock == 0 {
		if caller != nil {
			caller.Outputm(text.MN1048)
		}
		return
	}
	f.flush()
	if caller != nil {
		caller.Outputm(text.MN1049)
	}
}

func (f *Factory) Company() *Company {
	return f.company
}

func (f *Factory) Commodity() model.Commodity {
	return f.commodity
}

func (f *Factory) Display(caller *Player) {
	caller.Outputm(goods.GoodsArray[f.commodity].Message[1],
		f.company.Name,
		f.number+1,
		goods.GoodsArray[f.commodity].Name,
		f.planet.Name())

	caller.Outputm(text.DisplayFactory1,
		goods.GoodsArray[f.commodity].Labour,
		f.workers,
		f.wages,
		f.layoff)

	if f.numInputs() > 0 {
		caller.Outputm(text.DisplayFactory2)
		for i := range f.numInputs() {
			if i == 3 {
				caller.Output("\n")
			}
			caller.Outputm(text.DisplayFactoryInput,
				goods.GoodsArray[f.input(i).Commodity].Name,
				f.input(i).Quantity,
				f.stock[i])
		}
		caller.Output("\n")
	}

	caller.Outputm(text.DisplayFactory3,
		f.opStock,
		f.cycle,
		f.income,
		f.expend)

	switch f.delivery {
	case model.DeliverExchange:
		caller.Outputm(text.MN1063)
	case model.DeliverFactory:
		caller.Outputm(text.MN1065,
			f.company.Name,
			f.toFactory,
			f.contracted-f.delivered,
			f.carriage)
	case model.DeliverWarehouse:
		caller.Outputm(text.MN1064)
	}
}

func (f *Factory) Expend(amount int32) {
	if amount > 1000000 /*&& !testFeaturesEnabled*/ {
		log.Printf("%d expenditure %s #%d", amount, f.company.Name, f.number+1)
		return
	}
	changeBalance(&f.expend, amount)
	f.company.Expend(amount)
}

func (f *Factory) Expenditure() int32 {
	return f.expend
}

func (f *Factory) Income(amount int32) {
	changeBalance(&f.income, amount)
	f.company.Income(amount)
}

func (f *Factory) IsClosed() bool {
	return f.planet.IsClosed()
}

func (f *Factory) LinkWarehouse(w *model.Warehouse) {
	f.warehouse = w

	// if f.warehouse == nil && f.delivery == deliverWarehouse {
	// 	f.delivery = deliverExchange
	// }
}

func (f *Factory) Planet() *Planet {
	return f.planet
}

func (f *Factory) ReturnGoods(job *model.Work) {
	f.delivered -= job.Pallet.Quantity
	f.opStock += job.Pallet.Quantity
	f.Income((job.Pallet.Quantity * job.Value) / 2) // Return 50%
}

func (f *Factory) Run() bool {
	debug.Trace("Running %s %d", f.company.Name, f.number+1)

	// Don't run if the planet is closed.
	if f.planet.IsClosed() {
		return false
	}

	// Advance factory cycle.
	f.cycle++
	switch f.cycle % 100 {
	case 0:
		f.cycle = 0
	case 20, 40, 60, 80:
		// no-op
	default:
		return true
	}

	// Nobody works on No Production or Capital planets.
	if f.planet.HasExchange() {
		f.workers = 0
		return true
	}

	// No workers if there's no money to pay them.
	if f.company.Balance <= 0 || f.wages == 0 {
		f.workers = 0
		return true
	}

	// The workers won't get out of bed for less than the minimum wage.
	if f.wages < f.planet.MinimumWage() {
		f.workers = 0
		return true
	}

	// Recruit some workers if necessary.
	if uint32(f.workers) < f.product().Labour {
		f.getLabour()
	}

	// No workers? No production...
	if f.workers == 0 {
		return true
	}

	//
	inc := 100 + (f.planet.infstr * 5)

	// Try to stock up.
	for i := range f.numInputs() {
		amount := (inc*f.input(i).Quantity)/100 - f.stock[i]
		if amount > 0 {
			f.stock[i] += f.getInputs(f.input(i).Commodity, amount)
		}
	}

	// See if we have sufficient raw materials.
	for i := range f.numInputs() {
		if f.stock[i] < (inc*f.input(i).Quantity/5)/100 {
			debug.Trace("Insufficient stock of input %d", i)

			// Insufficient raw materials
			if f.wages <= 0 {
				f.workers = 0
			} else {
				// The workforce can exceed what's needed here if layoff pay
				// is higher than the normal wages.
				f.workers = (f.workers * f.layoff) / f.wages
				debug.Trace("Reduced workers to %d", f.workers)

				// And a bit more paranoia for the same reason...
				if f.workers < 0 {
					f.workers = 0
				} else if f.workers > int32(f.product().Labour) {
					f.workers = int32(f.product().Labour)
				}

				f.Expend((f.workers * f.layoff) / 5)
			}

			return true
		}
	}

	// Produce something!
	for i := range f.numInputs() {
		f.stock[i] -= (inc * f.input(i).Quantity / 5) / 100
	}

	f.Expend((f.workers * f.wages) / 5)

	production := (inc * (20 * f.workers) / int32(f.product().Labour)) / 100
	f.opStock += production
	FactoryProduction[f.commodity] += int(production)

	if f.opStock >= 150 {
		f.flush()
	}

	return true
}

func (f *Factory) Serialize(dbf *model.DBFactory) {
	dbf.Product = int32(f.commodity)
	dbf.Wages = f.wages
	dbf.Layoff = f.layoff
	dbf.Stock = f.stock
	dbf.Cycle = f.cycle
	dbf.Income = f.income
	dbf.Expend = f.expend
	dbf.Delivery = int32(f.delivery)
	dbf.Contracted = f.contracted
	dbf.Delivered = f.delivered
	dbf.Carriage = f.carriage
	dbf.ToFactory = f.toFactory
	dbf.OpStock = f.opStock
}

func (f *Factory) Stop() {
	f.workers = 0
}

func (f *Factory) UnlinkWarehouse(w *model.Warehouse) {
	if w == f.warehouse {
		f.LinkWarehouse(nil)
	}
}

func (f *Factory) Workers() int32 {
	return f.workers
}

func (f *Factory) flush() {
	// If unable to deliver after 10 tries - prevent stock build up.
	if f.opStock >= 400 {
		f.delivery = model.DeliverWarehouse
	}

	// TODO
}

// Gets inputs, either from the player's warehouse, or from the exchange.
func (f *Factory) getInputs(input model.Commodity, amount int32) int32 {
	needed := amount
	var found int32

	if f.warehouse != nil {
		for i := 0; i < 20 && needed > 0; i++ {
			if f.warehouse.Bay[i].Type == input {
				found += f.warehouse.Bay[i].Quantity
				needed -= f.warehouse.Bay[i].Quantity
				f.warehouse.Bay[i].Quantity = 0
			}
		}
		if found > 0 {
			cost := found * int32(goods.GoodsArray[input].BasePrice)
			f.Expend(cost)
			f.company.CEO.ChangeBalance(cost)
		}
		if needed <= 0 {
			return found
		}
	}

	// Try and get it from the exchange - factories can buy from stockpile.
	available, price, ok := f.planet.SellToFactory(input, needed)
	if ok {
		found += available
		f.Expend(available * price)
	}

	return found
}

// Attempts to recruit labour from the local labour exchange, taking into
// account wage levels and the size of the labour pool.
func (f *Factory) getLabour() {
	// Figure out what we need and what we've got.
	workersNeeded := int32(f.product().Labour) - f.workers
	workersAvailable := f.planet.LabourPool()

	// If there's not enough workers available then we'll poach 10% of the
	// workers from every factory on the planet paying less than this one.
	if workersNeeded > workersAvailable {
		for _, op := range Players {
			if !op.IsPlaying() || op.company == nil {
				continue
			}
			for _, of := range op.company.factories {
				if of == nil || of.planet != f.planet {
					continue
				}
				if of.wages < f.wages {
					workersPoached := of.workers / 10
					workersAvailable += workersPoached
					of.workers -= workersPoached
				}
			}
		}
	}

	// If we got enough then give the factory whatever it needs...
	if workersAvailable >= workersNeeded {
		f.workers = int32(f.product().Labour)
		return
	}

	// ...otherwise make do with what's available.
	f.workers += workersAvailable
}

func (f *Factory) input(index int) goods.InputCommodity {
	// debug.Check(index >= 0 && index < numInputs());
	return goods.GoodsArray[f.commodity].Inputs[index]
}

func (f *Factory) numInputs() int {
	return len(goods.GoodsArray[f.commodity].Inputs)
}

func (f *Factory) product() goods.TradeGoods {
	return goods.GoodsArray[f.commodity]
}
