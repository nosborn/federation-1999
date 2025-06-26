package server

import (
	"testing"

	"github.com/nosborn/federation-1999/internal/model"
)

func TestDuchySerializationRoundtrip(t *testing.T) {
	embargo := &Duchy{
		name: "TestEmbargo",
	}
	duchyIndex.Insert(embargo.name, embargo)
	favoured := &Duchy{
		name: "TestFavoured",
	}
	duchyIndex.Insert(favoured.name, favoured)
	// Create original duchy with test values
	original := &Duchy{
		customsRate:  15,
		embargo:      embargo,
		favoured:     favoured,
		favouredRate: 5,
		name:         "TestDuchy",
		taxRate:      10,
	}

	// Serialize to DBDuchy
	var dbDuchy model.DBDuchy
	original.Serialize(&dbDuchy)

	// Create test player to own the duchy
	testDuke := &Player{
		uid:  666000,
		name: "TestDuke",
		rank: model.RankDuke,
	}

	// Roundtrip: Create duchy from DBDuchy data
	// NewPlayerDuchy creates the basic duchy structure
	restored := NewPlayerDuchy(testDuke, "TestDuchy", dbDuchy.CustomsRate, dbDuchy.FavouredRate)

	// Use SetEmbargo/SetFavoured to restore string fields from DBDuchy
	embargoName := string(dbDuchy.Embargo[:])
	if len(embargoName) > 0 && embargoName[0] != 0 {
		// Find null terminator and trim
		for i, b := range embargoName {
			if b == 0 {
				embargoName = embargoName[:i]
				break
			}
		}
		if embargoName != "" {
			restored.SetEmbargo(embargoName, nil)
		}
	}

	favouredName := string(dbDuchy.Favoured[:])
	if len(favouredName) > 0 && favouredName[0] != 0 {
		// Find null terminator and trim
		for i, b := range favouredName {
			if b == 0 {
				favouredName = favouredName[:i]
				break
			}
		}
		if favouredName != "" {
			restored.SetFavoured(favouredName, nil)
		}
	}

	// Compare roundtrip results with original
	if restored.customsRate != original.customsRate {
		t.Errorf("CustomsRate roundtrip failed: original %d, restored %d",
			original.customsRate, restored.customsRate)
	}

	if restored.favouredRate != original.favouredRate {
		t.Errorf("FavouredRate roundtrip failed: original %d, restored %d",
			original.favouredRate, restored.favouredRate)
	}

	if restored.embargo.Name() != original.embargo.Name() {
		t.Errorf("Embargo roundtrip failed: original %q, restored %q",
			original.embargo.Name(), restored.embargo.Name())
	}

	if restored.favoured != original.favoured {
		t.Errorf("Favoured roundtrip failed: original %v, restored %v",
			original.favoured, restored.favoured)
	}

	if restored.name != original.name {
		t.Errorf("Name roundtrip failed: original %q, restored %q",
			original.name, restored.name)
	}

	if restored.taxRate != original.taxRate {
		t.Errorf("TaxRate roundtrip failed: original %d, restored %d",
			original.taxRate, restored.taxRate)
	}

	t.Logf("✓ Duchy roundtrip successful")
}

func TestDuchySerializationEmptyStrings(t *testing.T) {
	// Test duchy with empty embargo and favoured fields
	original := &Duchy{
		customsRate:  10,
		embargo:      nil, // Empty embargo
		favoured:     nil, // Empty favoured
		favouredRate: 0,
		name:         "EmptyDuchy",
		taxRate:      25,
	}

	// Serialize to DBDuchy
	var dbDuchy model.DBDuchy
	original.Serialize(&dbDuchy)

	// Verify fields
	if dbDuchy.CustomsRate != original.customsRate {
		t.Errorf("CustomsRate mismatch: original %d, serialized %d",
			original.customsRate, dbDuchy.CustomsRate)
	}

	if dbDuchy.FavouredRate != original.favouredRate {
		t.Errorf("FavouredRate mismatch: original %d, serialized %d",
			original.favouredRate, dbDuchy.FavouredRate)
	}

	// Empty strings should serialize as null-filled byte arrays
	// Check that the first byte is 0 (null terminator for empty string)
	if dbDuchy.Favoured[0] != 0 {
		t.Errorf("Expected empty favoured to have null first byte, got %d", dbDuchy.Favoured[0])
	}

	if dbDuchy.Embargo[0] != 0 {
		t.Errorf("Expected empty embargo to have null first byte, got %d", dbDuchy.Embargo[0])
	}

	t.Logf("✓ Empty string fields serialized correctly")
}
