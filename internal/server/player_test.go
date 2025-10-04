package server

import (
	"testing"

	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/server/sol"
)

func TestPlayerSerializationRoundtrip(t *testing.T) {
	// Initialize loader to avoid nil pointer crash
	NewLoader(func() {})

	// Create a Player with known test values for all fields
	original := &Player{
		uid:          666000,
		name:         "TestDuke",
		Desc:         "A test duke character",
		mood:         "Happy",
		sex:          model.SexMale,
		rank:         model.RankDuke,
		balance:      75000,
		loan:         5000,
		reward:       2500,
		Shipped:      42,
		Games:        15,
		Flags0:       0x12345678,
		Flags1:       0x87654321,
		deaths:       2,
		CustomRank:   3,
		TradeCredits: 8,
		Warper:       1,
		LocNo:        1000,
		LastTrade:    1640995200,
		ShipLoc:      sol.EarthLandingArea,
		Count:        [4]int32{10, 20, 30, 40},
		LastOn:       1640995300,
		Str:          PlayerStat{Max: 50, Cur: 45},
		Sta:          PlayerStat{Max: 60, Cur: 55},
		Int:          PlayerStat{Max: 70, Cur: 65},
		Dex:          PlayerStat{Max: 40, Cur: 35},
		curSysName:   "TestSystem",
	}

	// Create proper player duchy using the constructor (registers in allDuchies)
	// testDuchy := NewPlayerDuchy(original, "TestDuchy", 18, 12, 6)
	testDuchy := &Duchy{
		name:         "TestDuchy",
		owner:        original,
		taxRate:      18,
		customsRate:  12,
		favouredRate: 6,
	}

	// Create test system with known values - system belongs to Duke's duchy
	testSystem := &PlayerSystem{
		System: System{
			name:        "TestSystem",
			balance:     50000,
			taxRate:     15,
			touristTime: 300,
			flags:       model.PLT_T4_HOLDER,
			lastOnline:  1640995400,
			duchy:       testDuchy,
			owner:       original,
			planets:     []*Planet{{level: model.LevelCapital, population: 15000, energy: 45}},
		},
	}

	// Wire up ownership
	original.ownSystem = testSystem
	original.ownDuchy = testDuchy

	// Mock storage
	original.storage = &Storage{}
	original.storage.Warehouse[0] = &model.Warehouse{
		Planet: "TestSystem",
		Bay:    [20]model.Cargo{},
	}

	// Serialize to DBPersona
	dbPersona := original.Serialize()

	testInitializeGameWorld()

	// Create new Player from DBPersona (roundtrip)
	restored := NewPlayerFromDBPersona(dbPersona, 0)

	// Verify Player-specific fields survived the roundtrip
	if restored.UID() != original.UID() {
		t.Errorf("UID mismatch: original %d, restored %d", original.UID(), restored.UID())
	}

	if restored.Name() != original.Name() {
		t.Errorf("Name mismatch: original %q, restored %q", original.Name(), restored.Name())
	}

	if restored.sex != original.sex {
		t.Errorf("sex mismatch: original %c, restored %c", original.sex, restored.sex)
	}

	if restored.rank != original.rank {
		t.Errorf("rank mismatch: original %d, restored %d", original.rank, restored.rank)
	}

	if restored.balance != original.balance {
		t.Errorf("Balance mismatch: original %d, restored %d", original.balance, restored.balance)
	}

	if restored.loan != original.loan {
		t.Errorf("Loan mismatch: original %d, restored %d", original.loan, restored.loan)
	}

	if restored.Str.Max != original.Str.Max || restored.Str.Cur != original.Str.Cur {
		t.Errorf("Str mismatch: original %+v, restored %+v", original.Str, restored.Str)
	}

	if restored.Sta.Max != original.Sta.Max || restored.Sta.Cur != original.Sta.Cur {
		t.Errorf("Sta mismatch: original %+v, restored %+v", original.Sta, restored.Sta)
	}

	if restored.Int.Max != original.Int.Max || restored.Int.Cur != original.Int.Cur {
		t.Errorf("Int mismatch: original %+v, restored %+v", original.Int, restored.Int)
	}

	if restored.Dex.Max != original.Dex.Max || restored.Dex.Cur != original.Dex.Cur {
		t.Errorf("Dex mismatch: original %+v, restored %+v", original.Dex, restored.Dex)
	}

	if restored.Shipped != original.Shipped {
		t.Errorf("Shipped mismatch: original %d, restored %d", original.Shipped, restored.Shipped)
	}

	if restored.Games != original.Games {
		t.Errorf("Games mismatch: original %d, restored %d", original.Games, restored.Games)
	}

	// Spot-check: Verify System fields were restored (proves System.Serialize() worked)
	if restored.ownSystem == nil {
		t.Fatal("Expected restored player to have ownSystem, got nil")
	}

	if restored.ownSystem.TaxRate() != testSystem.TaxRate() {
		t.Errorf("System taxRate not restored: expected %d, got %d",
			testSystem.TaxRate(), restored.ownSystem.TaxRate())
	}

	if restored.ownSystem.Balance() != testSystem.Balance() {
		t.Errorf("System balance not restored: expected %d, got %d",
			testSystem.Balance(), restored.ownSystem.Balance())
	}

	// Spot-check: Verify Duchy fields were restored (proves Duchy.Serialize() worked)
	if restored.ownDuchy == nil {
		t.Fatal("Expected restored player to have ownDuchy, got nil")
	}

	if restored.ownDuchy.customsRate != testDuchy.customsRate {
		t.Errorf("Duchy customsRate not restored: expected %d, got %d",
			testDuchy.customsRate, restored.ownDuchy.customsRate)
	}

	if restored.ownDuchy.favoured != testDuchy.favoured {
		t.Errorf("Duchy favoured not restored: expected %v, got %v",
			testDuchy.favoured, restored.ownDuchy.favoured)
	}

	t.Logf("✓ Player fields roundtrip successful")
	t.Logf("✓ System fields roundtrip successful (via delegation)")
	t.Logf("✓ Duchy fields roundtrip successful (via delegation)")
}
