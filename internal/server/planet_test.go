package server

// import (
// 	"testing"
//
// 	"github.com/nosborn/federation-1999/internal/model"
// 	"github.com/nosborn/federation-1999/internal/server/exchange"
// 	"github.com/nosborn/federation-1999/internal/server/planet"
// )
//
// func TestPlanetSerializationRoundtrip(t *testing.T) {
// 	testDuchy := &Duchy{
// 		name: "TestDuchy",
// 	}
//
// 	testSystem := &System{
// 		name:  "TestSystem",
// 		duchy: testDuchy,
// 	}
//
// 	original := &Planet{
// 		level:        model.LevelIndustrial, // Valid level for exchanges
// 		population:   7500,
// 		energy:       25,
// 		education:    18,
// 		socialSec:    12,
// 		infstr:       30,
// 		health:       22,
// 		intSecurity:  15,
// 		disaffection: 8,
// 		system:       testSystem,
// 	}
// 	testExchange := &exchange.Exchange{
// 		exTime: 85,
// 		flux:   20,
// 		markup: 15,
// 		parent: original,
// 	}
// 	for i := range 52 {
// 		testExchange.goods[i] = exchange.Bin{
// 			production:  int32(2000 + i*20),
// 			produce:     (i%3 == 0),
// 			stock:       int32(800 + i*8),
// 			stockpile:   int32(400 + i*6),
// 			consumption: int32(600 + i*4),
// 			markup:      int32(20 + i),
// 			produced:    int32(300 + i*5),
// 			consumed:    int32(200 + i*7),
// 		}
// 	}
// 	original.exchange = testExchange
//
// 	// Serialize to DBPlanet
// 	var dbPlanet model.DBPlanet
// 	original.Serialize(&dbPlanet)
//
// 	// Roundtrip: Create planet from DBPlanet using NewPlayerPlanet
// 	restored := newPlanetFromDB(testSystem, dbPlanet)
//
// 	// Compare planet-specific fields
// 	if restored.level != original.level {
// 		t.Errorf("Level mismatch: original %d, restored %d", original.level, restored.level)
// 	}
//
// 	if restored.population != original.population {
// 		t.Errorf("Population mismatch: original %d, restored %d", original.population, restored.population)
// 	}
//
// 	if restored.energy != original.energy {
// 		t.Errorf("Energy mismatch: original %d, restored %d", original.energy, restored.energy)
// 	}
//
// 	if restored.education != original.education {
// 		t.Errorf("Education mismatch: original %d, restored %d", original.education, restored.education)
// 	}
//
// 	if restored.socialSec != original.socialSec {
// 		t.Errorf("SocialSec mismatch: original %d, restored %d", original.socialSec, restored.socialSec)
// 	}
//
// 	if restored.infstr != original.infstr {
// 		t.Errorf("Infstr mismatch: original %d, restored %d", original.infstr, restored.infstr)
// 	}
//
// 	if restored.health != original.health {
// 		t.Errorf("Health mismatch: original %d, restored %d", original.health, restored.health)
// 	}
//
// 	if restored.intSecurity != original.intSecurity {
// 		t.Errorf("IntSecurity mismatch: original %d, restored %d", original.intSecurity, restored.intSecurity)
// 	}
//
// 	if restored.disaffection != original.disaffection {
// 		t.Errorf("Disaffection mismatch: original %d, restored %d", original.disaffection, restored.disaffection)
// 	}
//
// 	// Verify exchange was created and basic fields match
// 	if restored.exchange == nil {
// 		t.Fatal("Exchange was not created for Industrial level planet")
// 	}
//
// 	if restored.exchange.exTime != original.exchange.exTime {
// 		t.Errorf("Exchange exTime mismatch: original %d, restored %d", original.exchange.exTime, restored.exchange.exTime)
// 	}
//
// 	if restored.exchange.Flux() != original.exchange.Flux() {
// 		t.Errorf("Exchange flux mismatch: original %d, restored %d", original.exchange.Flux(), restored.exchange.Flux())
// 	}
//
// 	if restored.exchange.markup != original.exchange.markup {
// 		t.Errorf("Exchange markup mismatch: original %d, restored %d", original.exchange.markup, restored.exchange.markup)
// 	}
//
// 	// Spot-check first and last commodity to verify exchange roundtrip
// 	if restored.exchange.goods[0].production != original.exchange.goods[0].production {
// 		t.Errorf("First commodity production mismatch: original %d, restored %d",
// 			original.exchange.goods[0].production, restored.exchange.goods[0].production)
// 	}
//
// 	if restored.exchange.goods[51].production != original.exchange.goods[51].production {
// 		t.Errorf("Last commodity production mismatch: original %d, restored %d",
// 			original.exchange.goods[51].production, restored.exchange.goods[51].production)
// 	}
//
// 	t.Logf("✓ Industrial planet roundtrip successful")
// }
//
// func TestPlanetSerializationNoExchange(t *testing.T) {
// 	// Create test system (parent required by NewPlayerPlanet)
// 	testSystem := &System{
// 		name:  "TestSystem",
// 		duchy: SolDuchy,
// 	}
//
// 	// Test planet without exchange (tests the nil exchange path)
// 	original := &Planet{
// 		level:        model.LevelCapital, // This level doesn't get exchanges
// 		population:   15000,
// 		energy:       40,
// 		education:    35,
// 		socialSec:    25,
// 		infstr:       50,
// 		health:       35,
// 		intSecurity:  30,
// 		disaffection: 2,
// 		exchange:     nil, // No exchange
// 		system:       testSystem,
// 	}
//
// 	// Serialize to DBPlanet
// 	var dbPlanet model.DBPlanet
// 	original.Serialize(&dbPlanet)
//
// 	// Roundtrip: Create planet from DBPlanet using NewPlayerPlanet
// 	restored := NewPlayerPlanet(testSystem, dbPlanet)
//
// 	// Compare all planet fields for lossless roundtrip
// 	if restored.level != original.level {
// 		t.Errorf("Level mismatch: original %d, restored %d", original.level, restored.level)
// 	}
//
// 	if restored.population != original.population {
// 		t.Errorf("Population mismatch: original %d, restored %d", original.population, restored.population)
// 	}
//
// 	if restored.energy != original.energy {
// 		t.Errorf("Energy mismatch: original %d, restored %d", original.energy, restored.energy)
// 	}
//
// 	if restored.education != original.education {
// 		t.Errorf("Education mismatch: original %d, restored %d", original.education, restored.education)
// 	}
//
// 	if restored.socialSec != original.socialSec {
// 		t.Errorf("SocialSec mismatch: original %d, restored %d", original.socialSec, restored.socialSec)
// 	}
//
// 	if restored.infstr != original.infstr {
// 		t.Errorf("Infstr mismatch: original %d, restored %d", original.infstr, restored.infstr)
// 	}
//
// 	if restored.health != original.health {
// 		t.Errorf("Health mismatch: original %d, restored %d", original.health, restored.health)
// 	}
//
// 	if restored.intSecurity != original.intSecurity {
// 		t.Errorf("IntSecurity mismatch: original %d, restored %d", original.intSecurity, restored.intSecurity)
// 	}
//
// 	if restored.disaffection != original.disaffection {
// 		t.Errorf("Disaffection mismatch: original %d, restored %d", original.disaffection, restored.disaffection)
// 	}
//
// 	// Test the nil exchange path - both should be nil after roundtrip
// 	if original.exchange != restored.exchange { // Both should be nil
// 		t.Errorf("Exchange state mismatch: original %v, restored %v", original.exchange, restored.exchange)
// 	}
//
// 	t.Logf("✓ Planet without exchange roundtrip successful")
// }
