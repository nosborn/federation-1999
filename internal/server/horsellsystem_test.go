package server

// import (
// 	"testing"
//
// 	"github.com/nosborn/federation-1999/internal/model"
// )
//
// func TestHorsellTemplateDataIntegrity(t *testing.T) {
// 	t.Run("LocationCount", func(t *testing.T) {
// 		expectedLocations := 88 - SPACESHIP_SIZE
// 		if n := len(HorsellLocations); n != expectedLocations {
// 			t.Errorf("Expected %d locations, got %d", expectedLocations, n)
// 		}
// 	})
//
// 	t.Run("ObjectCount", func(t *testing.T) {
// 		expectedObjects := 7
// 		if n := len(HorsellObjects); n != expectedObjects {
// 			t.Errorf("Expected %d objects, got %d", expectedObjects, n)
// 		}
// 	})
//
// 	// t.Run("PlanetCount", func(t *testing.T) {
// 	// 	expectedPlanets := 1
// 	// 	if n := len(HorsellPlanets); n != expectedPlanets {
// 	// 		t.Errorf("Expected %d planets, got %d", expectedPlanets, n)
// 	// 	}
// 	// })
//
// 	t.Run("LocationSequence", func(t *testing.T) {
// 		for i, loc := range HorsellLocations {
// 			expectedNumber := uint32(i + 1 + SPACESHIP_SIZE)
// 			if loc.Number != expectedNumber {
// 				t.Errorf("Location at index %d has number %d, expected %d", i, loc.Number, expectedNumber)
// 			}
// 		}
// 	})
//
// 	maxLocation := uint32(len(HorsellLocations) + SPACESHIP_SIZE)
//
// 	t.Run("MovTabBounds", func(t *testing.T) {
// 		minLoc := uint32(SPACESHIP_SIZE + 1)
// 		for i, loc := range HorsellLocations {
// 			for j, movTabEntry := range loc.MovTab {
// 				if movTabEntry != 0 {
// 					if movTabEntry < minLoc || movTabEntry > maxLocation {
// 						t.Errorf("Location %d MovTab[%d] = %d is outside bounds %d-%d",
// 							i, j, movTabEntry, minLoc, maxLocation)
// 					}
// 				}
// 			}
// 		}
// 	})
//
// 	t.Run("ObjectLocationBounds", func(t *testing.T) {
// 		minLoc := uint32(SPACESHIP_SIZE + 1)
// 		for i, obj := range HorsellObjects {
// 			if obj.minLocNo != 0 && (obj.minLocNo < minLoc || obj.minLocNo > maxLocation) {
// 				t.Errorf("Object %d (%s) minLocNo %d is outside bounds %d-%d",
// 					i, obj.name, obj.minLocNo, minLoc, maxLocation)
// 			}
//
// 			if obj.maxLocNo != 0 && (obj.maxLocNo < minLoc || obj.maxLocNo > maxLocation) {
// 				t.Errorf("Object %d (%s) maxLocNo %d is outside bounds %d-%d",
// 					i, obj.name, obj.maxLocNo, minLoc, maxLocation)
// 			}
//
// 			if obj.minLocNo != 0 && obj.maxLocNo != 0 && obj.minLocNo > obj.maxLocNo {
// 				t.Errorf("Object %d (%s) minLocNo %d > maxLocNo %d",
// 					i, obj.name, obj.minLocNo, obj.maxLocNo)
// 			}
//
// 			if (obj.minLocNo == 0) != (obj.maxLocNo == 0) {
// 				t.Errorf("Object %d (%s) has inconsistent location bounds: minLocNo=%d, maxLocNo=%d",
// 					i, obj.name, obj.minLocNo, obj.maxLocNo)
// 			}
// 		}
// 	})
// }
//
// func TestNewHorsellSystem(t *testing.T) {
// 	testName := "Horsell0"
// 	duchy := NewHorsellDuchy(testName)
// 	player := createTestPlayer(666000, "TestPlayer", model.RankTrader)
// 	system := NewHorsellSystem(duchy, player)
//
// 	t.Run("Name", func(t *testing.T) {
// 		if system.Name() != testName {
// 			t.Errorf("Expected name '%s', got '%s'", testName, system.Name())
// 		}
// 	})
//
// 	t.Run("Duchy", func(t *testing.T) {
// 		if system.duchy != duchy {
// 			t.Error("System should belong to the provided duchy")
// 		}
// 	})
//
// 	t.Run("CopiedData", func(t *testing.T) {
// 		if len(system.Events) != 0 {
// 			t.Errorf("Expected 0 events, got %d", len(system.Events))
// 		}
//
// 		expectedLocations := len(HorsellLocations)
// 		if len(system.Locations) != expectedLocations {
// 			t.Errorf("Expected %d locations copied from template, got %d", expectedLocations, len(system.Locations))
// 		}
//
// 		expectedObjects := len(HorsellObjects)
// 		if len(system.Objects) != expectedObjects {
// 			t.Errorf("Expected %d objects copied from template, got %d", expectedObjects, len(system.Objects))
// 		}
//
// 		expectedPlanets := len(HorsellPlanets)
// 		if len(system.Planets) != expectedPlanets {
// 			t.Errorf("Expected %d planets copied from template, got %d", expectedPlanets, len(system.Planets))
// 		}
// 	})
//
// 	t.Run("DataIsolation", func(t *testing.T) {
// 		// Verify that modifying the system data doesn't affect the templates
// 		if len(system.Locations) > 0 && len(HorsellLocations) > 0 {
// 			originalTemplateValue := HorsellLocations[0].number
// 			system.Locations[0].number = 999
// 			if HorsellLocations[0].number != originalTemplateValue {
// 				t.Error("Template data was modified when system data changed - not properly isolated")
// 			}
// 			// Restore for other tests
// 			system.Locations[0].number = originalTemplateValue
// 		}
// 	})
// }
