package server

import (
	"testing"
)

func TestSolDuchyDataIntegrity(t *testing.T) {
	t.Run("BasicProperties", func(t *testing.T) {
		if SolDuchy.Name() != "Sol" {
			t.Errorf("Expected name 'Sol', got '%s'", SolDuchy.Name())
		}

		expectedTaxRate := int32(25)
		if rate := SolDuchy.TaxRate(); rate != expectedTaxRate {
			t.Errorf("Expected tax rate %d, got %d", expectedTaxRate, rate)
		}

		expectedCustomsRate := int32(10)
		if rate := SolDuchy.CustomsRate(); rate != expectedCustomsRate {
			t.Errorf("Expected customs rate %d, got %d", expectedCustomsRate, rate)
		}
	})

	t.Run("SystemMembership", func(t *testing.T) {
		members := SolDuchy.AllMembers()
		expectedSystems := 3
		if len(members) != expectedSystems {
			t.Errorf("Expected %d systems, got %d", expectedSystems, len(members))
		}

		expectedOrder := []string{"Sol", "Arena", "Snark"}
		for i, expectedName := range expectedOrder {
			if i >= len(members) {
				t.Errorf("Missing system at index %d, expected '%s'", i, expectedName)
				continue
			}
			if members[i].Name() != expectedName {
				t.Errorf("System at index %d: expected '%s', got '%s'", i, expectedName, members[i].Name())
			}
		}
	})

	t.Run("CapitalSystem", func(t *testing.T) {
		capital := SolDuchy.CapitalSystem()
		if capital == nil {
			t.Error("Capital system should not be nil")
		} else if capital.Name() != "Sol" {
			t.Errorf("Expected capital system 'Sol', got '%s'", capital.Name())
		}
	})

	t.Run("IsSol", func(t *testing.T) {
		if !SolDuchy.IsSol() {
			t.Error("SolDuchy.IsSol() should return true")
		}

		if SolDuchy.IsHorsell() {
			t.Error("SolDuchy.IsHorsell() should return false")
		}
	})

	t.Run("FixedState", func(t *testing.T) {
		if SolDuchy.Owner() != nil {
			t.Error("SolDuchy should always have nil owner")
		}

		if SolDuchy.Embargo() != nil {
			t.Errorf("SolDuchy should have empty embargo, got '%v'", SolDuchy.Embargo())
		}

		if SolDuchy.Favoured() != nil {
			t.Errorf("SolDuchy should have empty favoured duchy, got '%v'", SolDuchy.Favoured())
		}

		if rate := SolDuchy.FavouredRate(); rate != 0 {
			t.Errorf("SolDuchy should have zero favoured rate, got %d", rate)
		}
	})

	t.Run("Visibility", func(t *testing.T) {
		if SolDuchy.IsHidden() {
			t.Error("SolDuchy should not be hidden")
		}

		if SolDuchy.IsClosed() {
			t.Error("SolDuchy should not be closed")
		}
	})

	// t.Run("AllDuchiesRegistration", func(t *testing.T) { -- FIXME
	// 	duchy, found := FindDuchy("Sol")
	// 	if !found {
	// 		t.Error("Sol duchy should be findable in allDuchies")
	// 	} else if duchy != SolDuchy {
	// 		t.Error("FindDuchy should return the same SolDuchy instance")
	// 	}
	//
	// 	allDuchies := AllDuchies()
	// 	if len(allDuchies) == 0 {
	// 		t.Error("AllDuchies() should not be empty")
	// 	} else if allDuchies[0] != SolDuchy {
	// 		t.Error("SolDuchy should be the first duchy in AllDuchies()")
	// 	}
	// })
}
