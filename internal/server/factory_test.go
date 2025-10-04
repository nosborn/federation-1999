package server

import (
	"testing"

	"github.com/nosborn/federation-1999/internal/model"
)

func TestFactorySerializationRoundtrip(t *testing.T) {
	testCEO := &Player{
		uid:  666000,
		name: "TestCEO",
		rank: model.RankIndustrialist,
	}

	testCompany := &Company{
		CEO:  testCEO,
		Name: "TestCompany",
	}

	testPlanet := &Planet{
		name:  "TestPlanet",
		level: model.LevelIndustrial,
	}

	testWarehouse := &model.Warehouse{
		Planet: testPlanet.Name(),
	}

	testStorage := &Storage{}
	testStorage.Warehouse[0] = testWarehouse
	testCEO.storage = testStorage

	original := &Factory{
		carriage:   500,
		commodity:  model.CommodityMasers,
		company:    testCompany,
		contracted: 150,
		cycle:      7,
		delivered:  200,
		delivery:   model.DeliverWarehouse,
		expend:     8000,
		income:     12000,
		layoff:     25,
		number:     3,
		opStock:    300,
		planet:     testPlanet,
		toFactory:  100,
		wages:      45,
		stock:      [6]int32{100, 200, 150, 75, 50, 25},
		warehouse:  testWarehouse,
	}

	var dbFactory model.DBFactory
	original.Serialize(&dbFactory)

	restored := NewFactoryFromDB(testCompany, 3, &dbFactory, testPlanet)

	if restored.carriage != original.carriage {
		t.Errorf("Carriage mismatch: original %d, restored %d", original.carriage, restored.carriage)
	}

	if restored.commodity != original.commodity {
		t.Errorf("Commodity mismatch: original %d, restored %d", original.commodity, restored.commodity)
	}

	if restored.company != original.company {
		t.Errorf("Company mismatch: expected %v, restored %v", original.company, restored.company)
	}

	if restored.contracted != original.contracted {
		t.Errorf("Contracted mismatch: original %d, restored %d", original.contracted, restored.contracted)
	}

	if restored.cycle != original.cycle {
		t.Errorf("Cycle mismatch: original %d, restored %d", original.cycle, restored.cycle)
	}

	if restored.delivered != original.delivered {
		t.Errorf("Delivered mismatch: original %d, restored %d", original.delivered, restored.delivered)
	}

	if restored.delivery != original.delivery {
		t.Errorf("Delivery mismatch: original %d, restored %d", original.delivery, restored.delivery)
	}

	if restored.expend != original.expend {
		t.Errorf("Expend mismatch: original %d, restored %d", original.expend, restored.expend)
	}

	if restored.income != original.income {
		t.Errorf("Income mismatch: original %d, restored %d", original.income, restored.income)
	}

	if restored.layoff != original.layoff {
		t.Errorf("Layoff mismatch: original %d, restored %d", original.layoff, restored.layoff)
	}

	if restored.number != original.number {
		t.Errorf("Number mismatch: original %d, restored %d", original.number, restored.number)
	}

	if restored.opStock != original.opStock {
		t.Errorf("OpStock mismatch: original %d, restored %d", original.opStock, restored.opStock)
	}

	if restored.planet != original.planet {
		t.Errorf("Planet mismatch: expected %v, restored %v", original.planet, restored.planet)
	}

	if restored.toFactory != original.toFactory {
		t.Errorf("ToFactory mismatch: original %d, restored %d", original.toFactory, restored.toFactory)
	}

	if restored.wages != original.wages {
		t.Errorf("Wages mismatch: original %d, restored %d", original.wages, restored.wages)
	}

	// Check stock array
	for i := range original.stock {
		if restored.stock[i] != original.stock[i] {
			t.Errorf("Stock[%d] mismatch: original %d, restored %d", i, original.stock[i], restored.stock[i])
		}
	}

	t.Logf("✓ Factory roundtrip successful")
}
