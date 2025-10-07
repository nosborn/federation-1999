package server

import (
	"testing"

	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/server/sol"
	"github.com/nosborn/federation-1999/internal/testutil"
	"github.com/nosborn/federation-1999/internal/text"
)

func TestCmdQuickStatus(t *testing.T) {
	tests := []struct {
		name        string
		player      *Player
		expectedMsg string
	}{
		{
			name: "no spaceship",
			player: func() *Player {
				p := createTestPlayerWithSession(666000, "TestActor", model.RankGroundHog)
				return p
			}(),
			expectedMsg: text.Msg(text.NoSpaceship),
		},
		{
			name: "with spaceship, no ammo or missiles",
			player: func() *Player {
				p := createTestPlayerWithSession(666000, "TestActor", model.RankCaptain)
				p.ShipLoc = sol.EarthLandingArea
				p.ShipKit = Equipment{
					Tonnage:     100,
					CurHull:     50,
					MaxHull:     100,
					CurShield:   25,
					MaxShield:   50,
					CurEngine:   10,
					MaxEngine:   20,
					CurComputer: 5,
					MaxComputer: 10,
					CurFuel:     100,
					MaxFuel:     200,
				}
				return p
			}(),
			expectedMsg: "Stats: H:50/100 S:25/50 E:10/20 C:5/10 F:100/200\n",
		},
		{
			name: "with spaceship, with ammo, no missiles",
			player: func() *Player {
				p := createTestPlayerWithSession(666000, "TestActor", model.RankCaptain)
				p.ShipLoc = sol.EarthLandingArea
				p.ShipKit = Equipment{
					Tonnage:     100,
					CurHull:     50,
					MaxHull:     100,
					CurShield:   25,
					MaxShield:   50,
					CurEngine:   10,
					MaxEngine:   20,
					CurComputer: 5,
					MaxComputer: 10,
					CurFuel:     100,
					MaxFuel:     200,
				}
				p.Ammo = 5
				return p
			}(),
			expectedMsg: "Stats: H:50/100 S:25/50 E:10/20 C:5/10 F:100/200 A:5\n",
		},
		{
			name: "with spaceship, no ammo, with missiles",
			player: func() *Player {
				p := createTestPlayerWithSession(666000, "TestActor", model.RankCaptain)
				p.ShipLoc = sol.EarthLandingArea
				p.ShipKit = Equipment{
					Tonnage:     100,
					CurHull:     50,
					MaxHull:     100,
					CurShield:   25,
					MaxShield:   50,
					CurEngine:   10,
					MaxEngine:   20,
					CurComputer: 5,
					MaxComputer: 10,
					CurFuel:     100,
					MaxFuel:     200,
				}
				p.Missiles = 10
				return p
			}(),
			expectedMsg: "Stats: H:50/100 S:25/50 E:10/20 C:5/10 F:100/200 M:10\n",
		},
		{
			name: "with spaceship, with ammo and missiles",
			player: func() *Player {
				p := createTestPlayerWithSession(666000, "TestActor", model.RankCaptain)
				p.ShipLoc = sol.EarthLandingArea
				p.ShipKit = Equipment{
					Tonnage:     100,
					CurHull:     50,
					MaxHull:     100,
					CurShield:   25,
					MaxShield:   50,
					CurEngine:   10,
					MaxEngine:   20,
					CurComputer: 5,
					MaxComputer: 10,
					CurFuel:     100,
					MaxFuel:     200,
				}
				p.Ammo = 5
				p.Missiles = 10
				return p
			}(),
			expectedMsg: "Stats: H:50/100 S:25/50 E:10/20 C:5/10 F:100/200 A:5 M:10\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			saver := &testutil.MockPlayerSaver[*Player]{}
			tt.player.saveFunc = saver.Save

			tt.player.CmdQuickStatus()

			assertMessageEquals(t, tt.player, tt.expectedMsg)
			testutil.AssertPlayerNotSaved(t, saver)
		})
	}
}
