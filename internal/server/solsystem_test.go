package server

import (
	"testing"
)

func TestSolSystemDataIntegrity(t *testing.T) {
	t.Run("Name", func(t *testing.T) {
		if SolSystem.Name() != "Sol" {
			t.Errorf("Expected name 'Sol', got '%s'", SolSystem.Name())
		}
	})

	t.Run("Duchy", func(t *testing.T) {
		if SolSystem.Duchy() != SolDuchy {
			t.Error("SolSystem should belong to SolDuchy")
		}
	})

	t.Run("EventCount", func(t *testing.T) {
		expectedEvents := 54
		if n := len(SolSystem.events); n != expectedEvents {
			t.Errorf("Expected %d events, got %d", expectedEvents, n)
		}
	})

	t.Run("LocationCount", func(t *testing.T) {
		expectedLocations := 698
		if n := len(SolSystem.locations); n != expectedLocations {
			t.Errorf("Expected %d locations, got %d", expectedLocations, n)
		}
	})

	t.Run("ObjectCount", func(t *testing.T) {
		expectedObjects := 90
		if n := len(SolSystem.objects); n != expectedObjects {
			t.Errorf("Expected %d objects, got %d", expectedObjects, n)
		}
	})

	t.Run("PlanetCount", func(t *testing.T) {
		expectedPlanets := 7
		if n := len(SolSystem.planets); n != expectedPlanets {
			t.Errorf("Expected %d planets, got %d", expectedPlanets, n)
		}
	})

	t.Run("EventLocationSequence", func(t *testing.T) {
		for i, loc := range SolSystem.locations {
			expectedNumber := uint32(i + 1)
			if loc.number != expectedNumber {
				t.Errorf("Location at index %d has number %d, expected %d", i, loc.number, expectedNumber)
			}
		}
	})

	maxLocation := uint32(len(SolSystem.locations))

	t.Run("MovTabBounds", func(t *testing.T) {
		for i, loc := range SolSystem.locations {
			for j, movTabEntry := range loc.MovTab {
				if movTabEntry != 0 {
					if i < SPACESHIP_SIZE {
						if movTabEntry < 1 || movTabEntry > SPACESHIP_SIZE {
							t.Errorf("Location %d MovTab[%d] = %d is outside bounds 1-%d",
								i, j, movTabEntry, SPACESHIP_SIZE)
						}
					} else {
						minLoc := uint32(SPACESHIP_SIZE + 1)
						if movTabEntry < minLoc || movTabEntry > maxLocation {
							t.Errorf("Location %d MovTab[%d] = %d is outside bounds %d-%d",
								i, j, movTabEntry, minLoc, maxLocation)
						}
					}
				}
			}
		}
	})

	t.Run("ObjectLocationBounds", func(t *testing.T) {
		minLoc := uint32(SPACESHIP_SIZE + 1)
		for i, obj := range SolSystem.objects {
			if obj.minLocNo != 0 && (obj.minLocNo < minLoc || obj.minLocNo > maxLocation) {
				t.Errorf("Object %d (%s) minLocNo %d is outside bounds %d-%d",
					i, obj.name, obj.minLocNo, minLoc, maxLocation)
			}

			if obj.maxLocNo != 0 && (obj.maxLocNo < minLoc || obj.maxLocNo > maxLocation) {
				t.Errorf("Object %d (%s) maxLocNo %d is outside bounds %d-%d",
					i, obj.name, obj.maxLocNo, minLoc, maxLocation)
			}

			if obj.minLocNo != 0 && obj.maxLocNo != 0 && obj.minLocNo > obj.maxLocNo {
				t.Errorf("Object %d (%s) minLocNo %d > maxLocNo %d",
					i, obj.name, obj.minLocNo, obj.maxLocNo)
			}

			if (obj.minLocNo == 0) != (obj.maxLocNo == 0) {
				t.Errorf("Object %d (%s) has inconsistent location bounds: minLocNo=%d, maxLocNo=%d",
					i, obj.name, obj.minLocNo, obj.maxLocNo)
			}
		}
	})
}
