package exchange

// import (
// 	"testing"
//
// 	"github.com/nosborn/federation-1999/internal/model"
// 	"github.com/nosborn/federation-1999/internal/server/planet"
// )
//
// func TestExchangeSerializationRoundtrip(t *testing.T) {
// 	// Create test planet parent (required by NewExchangeFromDB)
// 	testPlanet := &Planet{
// 		level: model.LevelIndustrial,
// 		name:  "TestPlanet",
// 	}
//
// 	// Create original exchange with test data
// 	original := &Exchange{
// 		exTime: 75,
// 		flux:   -15,
// 		markup: 12,
// 		parent: testPlanet,
// 		production: [5][2]int32{
// 			{100, 50},
// 			{200, 75},
// 			{150, 25},
// 			{300, 100},
// 			{250, 80},
// 		},
// 		jobs: [4]PlanetJob{
// 			{name: "TestJob1", commodity: model.CommodityElectros, carriage: 500},
// 			{name: "TestJob2", commodity: model.CommodityPharms, carriage: 750},
// 			{name: "", commodity: 0, carriage: 0}, // Empty slot
// 			{name: "TestJob4", commodity: model.CommodityTextiles, carriage: 1000},
// 		},
// 	}
//
// 	// Set up test data for ALL 52 commodities with unique values
// 	for i := range 52 {
// 		original.goods[i] = Bin{
// 			production:  int32(1000 + i*10), // Unique values to detect off-by-one errors
// 			produce:     (i%2 == 0),         // Alternating true/false
// 			stock:       int32(500 + i*5),
// 			stockpile:   int32(200 + i*3),
// 			consumption: int32(300 + i*2),
// 			markup:      int32(15 + i),
// 			produced:    int32(250 + i*4),
// 			consumed:    int32(150 + i*6),
// 		}
// 	}
//
// 	// Serialize to DBPlanet
// 	var dbPlanet model.DBPlanet
// 	original.Serialize(&dbPlanet)
//
// 	// Roundtrip: Create exchange from DBPlanet using NewExchangeFromDB
// 	restored := NewExchangeFromDB(testPlanet, dbPlanet)
//
// 	// Compare roundtrip results with original
// 	if restored.exTime != original.exTime {
// 		t.Errorf("ExTime roundtrip failed: original %d, restored %d", original.exTime, restored.exTime)
// 	}
//
// 	if restored.flux != original.flux {
// 		t.Errorf("Flux roundtrip failed: original %d, restored %d", original.flux, restored.flux)
// 	}
//
// 	if restored.markup != original.markup {
// 		t.Errorf("Markup roundtrip failed: original %d, restored %d", original.markup, restored.markup)
// 	}
//
// 	// Verify production arrays roundtrip correctly
// 	for i := range original.production {
// 		for j := range original.production[i] {
// 			if restored.production[i][j] != original.production[i][j] {
// 				t.Errorf("Production[%d][%d] roundtrip failed: original %d, restored %d",
// 					i, j, original.production[i][j], restored.production[i][j])
// 			}
// 		}
// 	}
//
// 	// Verify jobs are serialized to DBPlanet (but don't roundtrip - that's intentional)
// 	for i := range original.jobs {
// 		if dbPlanet.Jobs[i].Commodity != int32(original.jobs[i].commodity) {
// 			t.Errorf("Job[%d] commodity serialization failed: original %d, serialized %d",
// 				i, original.jobs[i].commodity, dbPlanet.Jobs[i].Commodity)
// 		}
//
// 		if dbPlanet.Jobs[i].Carriage != original.jobs[i].carriage {
// 			t.Errorf("Job[%d] carriage serialization failed: original %d, serialized %d",
// 				i, original.jobs[i].carriage, dbPlanet.Jobs[i].Carriage)
// 		}
//
// 		jobName := string(dbPlanet.Jobs[i].Name[:])
// 		expectedName := original.jobs[i].name
// 		if len(expectedName) > 0 {
// 			// Find null terminator
// 			for j, b := range jobName {
// 				if b == 0 {
// 					jobName = jobName[:j]
// 					break
// 				}
// 			}
// 			if jobName != expectedName {
// 				t.Errorf("Job[%d] name serialization failed: original %q, serialized %q",
// 					i, expectedName, jobName)
// 			}
// 		}
// 	}
// 	// Note: Jobs are serialized but intentionally not deserialized by NewExchangeFromDB
//
// 	// CRITICAL: Verify ALL 52 commodities roundtrip correctly
// 	for i := range 52 {
// 		if restored.goods[i].production != original.goods[i].production {
// 			t.Errorf("Goods[%d] production roundtrip failed: original %d, restored %d",
// 				i, original.goods[i].production, restored.goods[i].production)
// 		}
//
// 		if restored.goods[i].produce != original.goods[i].produce {
// 			t.Errorf("Goods[%d] produce roundtrip failed: original %t, restored %t",
// 				i, original.goods[i].produce, restored.goods[i].produce)
// 		}
//
// 		if restored.goods[i].stock != original.goods[i].stock {
// 			t.Errorf("Goods[%d] stock roundtrip failed: original %d, restored %d",
// 				i, original.goods[i].stock, restored.goods[i].stock)
// 		}
//
// 		if restored.goods[i].stockpile != original.goods[i].stockpile {
// 			t.Errorf("Goods[%d] stockpile roundtrip failed: original %d, restored %d",
// 				i, original.goods[i].stockpile, restored.goods[i].stockpile)
// 		}
//
// 		if restored.goods[i].consumption != original.goods[i].consumption {
// 			t.Errorf("Goods[%d] consumption roundtrip failed: original %d, restored %d",
// 				i, original.goods[i].consumption, restored.goods[i].consumption)
// 		}
//
// 		if restored.goods[i].markup != original.goods[i].markup {
// 			t.Errorf("Goods[%d] markup roundtrip failed: original %d, restored %d",
// 				i, original.goods[i].markup, restored.goods[i].markup)
// 		}
//
// 		if restored.goods[i].produced != original.goods[i].produced {
// 			t.Errorf("Goods[%d] produced roundtrip failed: original %d, restored %d",
// 				i, original.goods[i].produced, restored.goods[i].produced)
// 		}
//
// 		if restored.goods[i].consumed != original.goods[i].consumed {
// 			t.Errorf("Goods[%d] consumed roundtrip failed: original %d, restored %d",
// 				i, original.goods[i].consumed, restored.goods[i].consumed)
// 		}
// 	}
//
// 	t.Logf("✓ Exchange roundtrip successful - all 52 commodities verified")
// }
