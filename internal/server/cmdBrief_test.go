package server

import (
	"testing"

	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/server/database"
	"github.com/nosborn/federation-1999/internal/testutil"
	"github.com/nosborn/federation-1999/internal/text"
)

func TestCmdBrief(t *testing.T) {
	tests := []struct {
		name           string
		player         *Player
		expectedMsg    string
		expectedFlags0 uint32
		shouldSave     bool
	}{
		{
			name: "brief mode not set",
			player: func() *Player {
				p := createTestPlayerWithSession(666000, "TestActor", model.RankTrader)
				p.Flags0 &^= model.PL0_BRIEF
				return p
			}(),
			expectedMsg:    text.Msg(text.BriefOK),
			expectedFlags0: model.PL0_BRIEF,
			shouldSave:     true,
		},
		{
			name: "brief mode already set",
			player: func() *Player {
				p := createTestPlayerWithSession(666000, "TestActor", model.RankTrader)
				p.Flags0 |= model.PL0_BRIEF
				return p
			}(),
			expectedMsg:    text.Msg(text.BriefOK),
			expectedFlags0: model.PL0_BRIEF,
			shouldSave:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			saver := &testutil.MockPlayerSaver[*Player]{}
			tt.player.saveFunc = saver.Save

			tt.player.CmdBrief()

			assertMessageEquals(t, tt.player, tt.expectedMsg)
			testutil.AssertFlagSet(t, tt.player.Flags0, tt.expectedFlags0)
			if tt.shouldSave {
				testutil.AssertPlayerSave(t, saver, database.SaveNow)
			} else {
				testutil.AssertPlayerNotSaved(t, saver)
			}
		})
	}
}
