package server

import (
	"testing"

	"github.com/nosborn/federation-1999/internal/model"
)

func TestCompanySerializationRoundtrip(t *testing.T) {
	testCEO := &Player{
		uid:  666000,
		name: "TestCEO",
		rank: model.RankIndustrialist,
	}

	original := &Company{
		CEO:     testCEO,
		Name:    "TestCompany",
		Balance: 50000,
		Shares:  1000,
		income:  15000,
		expend:  12000,
	}

	var dbCompany model.DBCompany
	var dbFactories [model.MAX_FACTORY]model.DBFactory
	original.Serialize(&dbCompany, &dbFactories)

	restored := NewCompanyFromDB(testCEO, dbCompany, dbFactories)

	if restored.Name != original.Name {
		t.Errorf("Name mismatch: original %s, restored %s", original.Name, restored.Name)
	}

	if restored.Balance != original.Balance {
		t.Errorf("Balance mismatch: original %d, restored %d", original.Balance, restored.Balance)
	}

	if restored.Shares != original.Shares {
		t.Errorf("Shares mismatch: original %d, restored %d", original.Shares, restored.Shares)
	}

	if restored.income != original.income {
		t.Errorf("Income mismatch: original %d, restored %d", original.income, restored.income)
	}

	if restored.expend != original.expend {
		t.Errorf("Expend mismatch: original %d, restored %d", original.expend, restored.expend)
	}

	if restored.CEO != testCEO {
		t.Errorf("CEO mismatch: expected %v, restored %v", testCEO, restored.CEO)
	}

	// Check that both have the same number of non-nil factories
	originalCount := 0
	restoredCount := 0
	for i := range original.factories {
		if original.factories[i] != nil {
			originalCount++
		}
		if restored.factories[i] != nil {
			restoredCount++
		}
	}
	if originalCount != restoredCount {
		t.Errorf("Non-nil factories count mismatch: original %d, restored %d", originalCount, restoredCount)
	}

	t.Logf("✓ Company roundtrip successful")
}
