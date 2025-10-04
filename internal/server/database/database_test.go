package database

import (
	"testing"
	"unsafe"

	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/text"
)

func TestDBPersonaSize(t *testing.T) {
	personaSize := unsafe.Sizeof(model.DBPersona{})
	t.Logf("DBPersona size: %d bytes", personaSize)
}

func TestBinaryReadWrite(t *testing.T) {
	t.Run("PP Data (Regular Player)", func(t *testing.T) {
		// Create a DBPersona with PP data (rank < model.RankSquire)
		var dbPersona model.DBPersona

		// Set basic fields
		copy(dbPersona.Name[:], "PPTest")
		dbPersona.ID = 666000
		dbPersona.Sex = byte(model.SexFemale)
		dbPersona.Rank = uint32(model.RankTrader) // Below model.RankSquire
		dbPersona.Balance = 75000
		copy(dbPersona.StarSystem[:], "Sol")
		dbPersona.LocNo = 696 /*MeetingPoint*/

		// Set PP data
		ppData := &model.DBPPData{
			GMLocation: 123,
			Company: model.DBCompany{
				Balance: 50000,
				Capital: 25000,
			},
		}
		copy(ppData.Company.Name[:], "TestCorp")
		dbPersona.PP = ppData

		// Test binary pack/unpack round-trip
		page := packDBPersona(&dbPersona)

		// Unpack from page
		readPersona := unpackDBPersona(page)

		// Verify basic data preserved
		if text.CStringToString(readPersona.Name[:]) != "PPTest" {
			t.Errorf("Binary round-trip name mismatch: got %s, want %s", text.CStringToString(readPersona.Name[:]), "PPTest")
		}

		if readPersona.ID != 666000 {
			t.Errorf("Binary round-trip ID mismatch: got %d, want %d", readPersona.ID, 666000)
		}

		// Verify PP data preserved
		if readPersona.PP == nil {
			t.Fatal("PP data was nil after round-trip")
		}

		if readPersona.PP.GMLocation != 123 {
			t.Errorf("PP GMLocation mismatch: got %d, want %d", readPersona.PP.GMLocation, 123)
		}

		if readPersona.PP.Company.Balance != 50000 {
			t.Errorf("PP Company Balance mismatch: got %d, want %d", readPersona.PP.Company.Balance, 50000)
		}

		if text.CStringToString(readPersona.PP.Company.Name[:]) != "TestCorp" {
			t.Errorf("PP Company Name mismatch: got %s, want %s", text.CStringToString(readPersona.PP.Company.Name[:]), "TestCorp")
		}
	})

	t.Run("RP Data (Planet Owner)", func(t *testing.T) {
		// Create a DBPersona with RP data (RankSquire <= rank <= model.RankDuke)
		var dbPersona model.DBPersona

		// Set basic fields
		copy(dbPersona.Name[:], "RPTest")
		dbPersona.ID = 666001
		dbPersona.Sex = byte(model.SexMale)
		dbPersona.Rank = uint32(model.RankBaron) // Between model.RankSquire and model.RankDuke
		dbPersona.Balance = 125000
		copy(dbPersona.StarSystem[:], "Arena")
		dbPersona.LocNo = 500

		// Set RP data
		rpData := &model.DBRPData{
			Planet: model.DBPlanet{
				Population: 1500,
				Tax:        15,
				Balance:    80000,
			},
			Facilities: [7]int32{85, 60, 45, 0, 0, 0, 0},
		}
		copy(rpData.Fief[:], "TestPlanet")
		copy(rpData.Planet.Duchy[:], "TestDuchy")
		dbPersona.RP = rpData

		// Test binary pack/unpack round-trip
		page := packDBPersona(&dbPersona)

		// Unpack from page
		readPersona := unpackDBPersona(page)

		// Verify basic data preserved
		if text.CStringToString(readPersona.Name[:]) != "RPTest" {
			t.Errorf("Binary round-trip name mismatch: got %s, want %s", text.CStringToString(readPersona.Name[:]), "RPTest")
		}

		if readPersona.ID != 666001 {
			t.Errorf("Binary round-trip ID mismatch: got %d, want %d", readPersona.ID, 666001)
		}

		// Verify RP data preserved
		if readPersona.RP == nil {
			t.Fatal("RP data was nil after round-trip")
		}

		if text.CStringToString(readPersona.RP.Fief[:]) != "TestPlanet" {
			t.Errorf("RP Fief mismatch: got %s, want %s", text.CStringToString(readPersona.RP.Fief[:]), "TestPlanet")
		}

		if readPersona.RP.Planet.Population != 1500 {
			t.Errorf("RP Planet Population mismatch: got %d, want %d", readPersona.RP.Planet.Population, 1500)
		}

		if readPersona.RP.Planet.Tax != 15 {
			t.Errorf("RP Planet Tax mismatch: got %d, want %d", readPersona.RP.Planet.Tax, 15)
		}

		if readPersona.RP.Facilities[0] != 85 || readPersona.RP.Facilities[1] != 60 {
			t.Errorf("RP Facilities mismatch: got [%d, %d], want [85, 60]", readPersona.RP.Facilities[0], readPersona.RP.Facilities[1])
		}
	})
}
