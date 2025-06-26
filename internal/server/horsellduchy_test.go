package server

import (
	"strings"
	"testing"
)

func TestNewHorsellDuchy(t *testing.T) {
	testName := "Horsell0"
	duchy := NewHorsellDuchy(testName)

	t.Run("BasicProperties", func(t *testing.T) {
		if duchy.Name() != testName {
			t.Errorf("Expected name '%s', got '%s'", testName, duchy.Name())
		}

		expectedTaxRate := int32(0)
		if rate := duchy.TaxRate(); rate != expectedTaxRate {
			t.Errorf("Expected tax rate %d, got %d", expectedTaxRate, rate)
		}

		expectedCustomsRate := int32(10)
		if rate := duchy.CustomsRate(); rate != expectedCustomsRate {
			t.Errorf("Expected customs rate %d, got %d", expectedCustomsRate, rate)
		}
	})

	t.Run("SystemMembership", func(t *testing.T) {
		members := duchy.AllMembers()
		expectedSystems := 0
		if len(members) != expectedSystems {
			t.Errorf("Expected %d systems, got %d", expectedSystems, len(members))
		}
	})

	t.Run("IsHorsell", func(t *testing.T) {
		if !duchy.IsHorsell() {
			t.Error("Horsell duchy should return true for IsHorsell()")
		}

		if duchy.IsSol() {
			t.Error("Horsell duchy should return false for IsSol()")
		}
	})

	t.Run("InitialState", func(t *testing.T) {
		if duchy.Owner() != nil {
			t.Error("New Horsell duchy should have nil owner")
		}

		if duchy.Embargo() != nil {
			t.Errorf("New Horsell duchy should have empty embargo, got '%v'", duchy.Embargo())
		}

		if duchy.Favoured() != nil {
			t.Errorf("New Horsell duchy should have empty favoured duchy, got '%v'", duchy.Favoured())
		}

		if rate := duchy.FavouredRate(); rate != 0 {
			t.Errorf("New Horsell duchy should have zero favoured rate, got %d", rate)
		}
	})

	t.Run("Visibility", func(t *testing.T) {
		if !duchy.IsHidden() {
			t.Error("Horsell duchy should be hidden")
		}

		if !duchy.IsClosed() {
			t.Error("Horsell duchy should be closed")
		}
	})

	t.Run("NotInAllDuchies", func(t *testing.T) {
		_, found := FindDuchy(testName)
		if found {
			t.Errorf("Newly created duchy '%s' should not be findable in allDuchies", testName)
		}

		allDuchies := AllDuchies()
		for _, d := range allDuchies {
			if d == duchy {
				t.Error("Newly created Horsell duchy should not be present in AllDuchies()")
				break
			}
		}
	})

	t.Run("NamePattern", func(t *testing.T) {
		if !strings.HasPrefix(duchy.Name(), "Horsell") {
			t.Errorf("Duchy name '%s' should start with 'Horsell'", duchy.Name())
		}
	})
}
