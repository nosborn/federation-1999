package server

import (
	"testing"

	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/server/database"
	"github.com/nosborn/federation-1999/internal/server/sol"
	"github.com/nosborn/federation-1999/internal/testutil"
	"github.com/nosborn/federation-1999/internal/text"
)

func TestCmdSalute(t *testing.T) {
	tests := []struct {
		name              string
		player            *Player
		expectedMsg       string
		expectedFlags1    uint32
		notExpectedFlags1 uint32
		shouldSave        bool
	}{
		{
			name: "salute in wrong location",
			player: func() *Player {
				p := createTestPlayerWithSession(666000, "TestActor", model.RankTrader)
				p.curSys = &System{name: "Sol"}
				p.LocNo = sol.EarthLandingArea
				return p
			}(),
			expectedMsg:       text.Msg(text.Salute),
			notExpectedFlags1: model.PL1_MI6,
			shouldSave:        false,
		},
		{
			name: "salute in correct location with ID card",
			player: func() *Player {
				p := createTestPlayerWithSession(666000, "TestActor", model.RankTrader)
				p.curSys = &System{name: "Sol"}
				p.LocNo = sol.ControlRoom3
				p.Flags1 |= model.PL1_MI6
				return p
			}(),
			expectedMsg:    text.Msg(text.Salute),
			expectedFlags1: model.PL1_MI6,
			shouldSave:     false,
		},
		{
			name: "salute in correct location without ID card",
			player: func() *Player {
				p := createTestPlayerWithSession(666000, "TestActor", model.RankTrader)
				p.curSys = &System{name: "Sol"}
				p.LocNo = sol.ControlRoom3
				p.Flags1 &^= model.PL1_MI6
				return p
			}(),
			expectedMsg:    text.Msg(text.SaluteKatov),
			expectedFlags1: model.PL1_MI6,
			shouldSave:     true,
		},
		{
			name: "salute in correct location with MI6 offered",
			player: func() *Player {
				p := createTestPlayerWithSession(666000, "TestActor", model.RankTrader)
				p.curSys = &System{name: "Sol"}
				p.LocNo = sol.ControlRoom3
				p.Flags1 |= model.PL1_MI6_OFFERED
				p.Flags1 &^= model.PL1_MI6
				return p
			}(),
			expectedMsg:       text.Msg(text.SaluteKatov),
			expectedFlags1:    model.PL1_MI6,
			notExpectedFlags1: model.PL1_MI6_OFFERED,
			shouldSave:        true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			saver := &testutil.MockPlayerSaver[*Player]{}
			tt.player.saveFunc = saver.Save

			tt.player.CmdSalute()

			assertMessageEquals(t, tt.player, tt.expectedMsg)
			if tt.expectedFlags1 != 0 {
				testutil.AssertFlagSet(t, tt.player.Flags1, tt.expectedFlags1)
			}
			if tt.notExpectedFlags1 != 0 {
				testutil.AssertFlagNotSet(t, tt.player.Flags1, tt.notExpectedFlags1)
			}
			if tt.shouldSave {
				testutil.AssertPlayerSave(t, saver, database.SaveNow)
			} else {
				testutil.AssertPlayerNotSaved(t, saver)
			}
		})
	}
}
