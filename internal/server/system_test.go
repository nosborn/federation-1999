package server

// import (
// 	"testing"
// 	"time"
//
// 	"github.com/nosborn/federation-1999/internal/model"
// 	"github.com/nosborn/federation-1999/internal/server/planet"
// )
//
// func TestSystemSerializationRoundtrip(t *testing.T) {
// 	testInitializeGameWorld()
//
// 	testPlayer := &Player{
// 		uid:  666000,
// 		name: "TestPlayer",
// 		rank: model.RankIndustrialist,
// 	}
//
// 	original := &System{
// 		name:        "TestSystem",
// 		balance:     75000,
// 		taxRate:     20,
// 		touristTime: 450,
// 		Flags:       model.PLT_T4_HOLDER,
// 		lastOnline:  int32(time.Now().AddDate(0, 0, -3).Unix()),
// 		duchy:       SolDuchy,
// 		loadState:   SystemOnline,
// 		owner:       testPlayer,
// 	}
//
// 	testPlanet := &Planet{
// 		level:        model.LevelIndustrial,
// 		population:   8500,
// 		energy:       28,
// 		education:    22,
// 		socialSec:    15,
// 		infstr:       35,
// 		health:       25,
// 		intSecurity:  18,
// 		disaffection: 6,
// 		system:       original,
// 	}
//
// 	testExchange := &Exchange{
// 		exTime: 90,
// 		flux:   -25,
// 		markup: 18,
// 		parent: testPlanet,
// 	}
//
// 	for i := range 52 {
// 		testExchange.goods[i] = Bin{
// 			production:  int32(3000 + i*30),
// 			produce:     (i%4 == 0),
// 			stock:       int32(1000 + i*10),
// 			stockpile:   int32(500 + i*8),
// 			consumption: int32(700 + i*5),
// 			markup:      int32(25 + i),
// 			produced:    int32(400 + i*6),
// 			consumed:    int32(300 + i*8),
// 		}
// 	}
//
// 	testPlanet.exchange = testExchange
// 	original.Planets = []*Planet{testPlanet}
//
// 	var dbPlanet model.DBPlanet
// 	original.Serialize(&dbPlanet)
//
// 	restored := NewPlayerSystemFromDB(testPlayer, "TestSystem", dbPlanet)
//
// 	if restored.balance != original.balance {
// 		t.Errorf("Balance mismatch: original %d, restored %d", original.balance, restored.balance)
// 	}
//
// 	if restored.taxRate != original.taxRate {
// 		t.Errorf("TaxRate mismatch: original %d, restored %d", original.taxRate, restored.taxRate)
// 	}
//
// 	if restored.lastOnline != original.lastOnline {
// 		t.Errorf("LastOnline mismatch: original %d, restored %d", original.lastOnline, restored.lastOnline)
// 	}
//
// 	expectedFlags := dbPlanet.Flags &^ model.PLT_CLOSED
// 	if restored.Flags != expectedFlags {
// 		t.Errorf("Flags mismatch after PLT_CLOSED clear: expected %d, restored %d", expectedFlags, restored.Flags)
// 	}
//
// 	if restored.owner != testPlayer {
// 		t.Errorf("owner mismatch: expected %v, restored %v", testPlayer, restored.owner)
// 	}
//
// 	if restored.name != original.name {
// 		t.Errorf("Name mismatch: original %s, restored %s", original.name, restored.name)
// 	}
//
// 	if len(restored.Planets) != 1 {
// 		t.Fatalf("Expected 1 planet, got %d", len(restored.Planets))
// 	}
//
// 	restoredPlanet := restored.Planets[0]
// 	if restoredPlanet.level != testPlanet.level {
// 		t.Errorf("Planet level mismatch: original %d, restored %d", testPlanet.level, restoredPlanet.level)
// 	}
//
// 	if restoredPlanet.population != testPlanet.population {
// 		t.Errorf("Planet population mismatch: original %d, restored %d", testPlanet.population, restoredPlanet.population)
// 	}
//
// 	if restoredPlanet.exchange == nil {
// 		t.Fatal("Exchange was not created")
// 	}
//
// 	if restoredPlanet.exchange.exTime != testExchange.exTime {
// 		t.Errorf("Exchange exTime mismatch: original %d, restored %d", testExchange.exTime, restoredPlanet.exchange.exTime)
// 	}
//
// 	if restoredPlanet.exchange.goods[0].production != testExchange.goods[0].production {
// 		t.Errorf("First commodity production mismatch: original %d, restored %d",
// 			testExchange.goods[0].production, restoredPlanet.exchange.goods[0].production)
// 	}
//
// 	t.Logf("✓ System roundtrip successful")
// }
//
// func TestSystemSerializationOfflineState(t *testing.T) {
// 	testInitializeGameWorld()
//
// 	// Test that offline/unloading systems get PLT_CLOSED flag added
// 	original := &System{
// 		name:        "OfflineSystem",
// 		balance:     50000,
// 		taxRate:     15,
// 		touristTime: 200,
// 		Flags:       model.PLT_T4_HOLDER, // Start without PLT_CLOSED
// 		lastOnline:  int32(time.Now().AddDate(0, 0, -5).Unix()),
// 		duchy:       SolDuchy,
// 		loadState:   SystemOffline, // This should trigger PLT_CLOSED
// 	}
//
// 	// Create a capital planet (no exchange)
// 	testPlanet := &Planet{
// 		level:        model.LevelCapital,
// 		population:   12000,
// 		energy:       45,
// 		education:    40,
// 		socialSec:    30,
// 		infstr:       55,
// 		health:       40,
// 		intSecurity:  35,
// 		disaffection: 1,
// 		exchange:     nil, // Capital has no exchange
// 	}
//
// 	original.Planets = []*Planet{testPlanet}
//
// 	// Serialize to DBPlanet
// 	var dbPlanet model.DBPlanet
// 	original.Serialize(&dbPlanet)
//
// 	// Verify that PLT_CLOSED flag was added due to offline state
// 	if (dbPlanet.Flags & model.PLT_CLOSED) == 0 {
// 		t.Error("Expected PLT_CLOSED flag to be set for offline system")
// 	}
//
// 	// Verify original flags are preserved
// 	if (dbPlanet.Flags & model.PLT_T4_HOLDER) == 0 {
// 		t.Error("Expected original PLT_T4_HOLDER flag to be preserved")
// 	}
//
// 	// Verify other system fields
// 	if dbPlanet.Tax != original.taxRate {
// 		t.Errorf("Tax mismatch: original %d, serialized %d", original.taxRate, dbPlanet.Tax)
// 	}
//
// 	t.Logf("✓ Offline system correctly adds PLT_CLOSED flag")
// }
//
// func TestSystemSerializationUnloadingState(t *testing.T) {
// 	testInitializeGameWorld()
//
// 	// Test that unloading systems also get PLT_CLOSED flag
// 	original := &System{
// 		name:        "UnloadingSystem",
// 		balance:     25000,
// 		taxRate:     12,
// 		touristTime: 100,
// 		Flags:       0, // No flags initially
// 		lastOnline:  int32(time.Now().AddDate(0, 0, -5).Unix()),
// 		duchy:       SolDuchy,
// 		loadState:   SystemUnloading, // This should also trigger PLT_CLOSED
// 	}
//
// 	// Create a technological planet with exchange
// 	testPlanet := &Planet{
// 		level:        model.LevelTechnological,
// 		population:   6000,
// 		energy:       35,
// 		education:    30,
// 		socialSec:    20,
// 		infstr:       40,
// 		health:       28,
// 		intSecurity:  22,
// 		disaffection: 4,
// 	}
//
// 	// Create minimal exchange
// 	testExchange := &Exchange{
// 		exTime: 50,
// 		flux:   10,
// 		markup: 10,
// 	}
//
// 	// Set up minimal commodity data
// 	for i := range 52 {
// 		testExchange.goods[i] = Bin{
// 			production: int32(100 + i),
// 			stock:      int32(200 + i),
// 		}
// 	}
//
// 	testPlanet.exchange = testExchange
// 	original.Planets = []*Planet{testPlanet}
//
// 	// Serialize to DBPlanet
// 	var dbPlanet model.DBPlanet
// 	original.Serialize(&dbPlanet)
//
// 	// Verify that PLT_CLOSED flag was added due to unloading state
// 	if (dbPlanet.Flags & model.PLT_CLOSED) == 0 {
// 		t.Error("Expected PLT_CLOSED flag to be set for unloading system")
// 	}
//
// 	t.Logf("✓ Unloading system correctly adds PLT_CLOSED flag")
// }
//
// func TestSetTaxRate(t *testing.T) {
// 	s := &System{
// 		name:    "TestSystem",
// 		taxRate: 15,
// 	}
//
// 	// Test normal case - setting valid tax rate
// 	s.SetTaxRate(20)
// 	if s.taxRate != 20 {
// 		t.Errorf("Expected taxRate to be 20, got %d", s.taxRate)
// 	}
//
// 	// Test clamping negative values to 0
// 	s.SetTaxRate(-5)
// 	if s.taxRate != 0 {
// 		t.Errorf("Expected taxRate to be clamped to 0, got %d", s.taxRate)
// 	}
//
// 	// Test clamping values above 30 to 30
// 	s.SetTaxRate(50)
// 	if s.taxRate != 30 {
// 		t.Errorf("Expected taxRate to be clamped to 30, got %d", s.taxRate)
// 	}
//
// 	// Test boundary values
// 	s.SetTaxRate(0)
// 	if s.taxRate != 0 {
// 		t.Errorf("Expected taxRate to be 0, got %d", s.taxRate)
// 	}
//
// 	s.SetTaxRate(30)
// 	if s.taxRate != 30 {
// 		t.Errorf("Expected taxRate to be 30, got %d", s.taxRate)
// 	}
//
// 	t.Logf("✓ SetTaxRate tests passed")
// }
