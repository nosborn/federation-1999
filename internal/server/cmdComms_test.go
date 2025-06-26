package server

import (
	"testing"

	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/server/database"
	"github.com/nosborn/federation-1999/internal/testutil"
	"github.com/nosborn/federation-1999/internal/text"
)

func TestCmdComms(t *testing.T) {
	tests := []struct {
		name        string
		player      *Player
		expectedMsg string
	}{
		{
			name: "no comms unit",
			player: func() *Player {
				p := createTestPlayerWithSession(666000, "TestActor", model.RankTrader)
				p.Flags0 &^= model.PL0_COMM_UNIT
				return p
			}(),
			expectedMsg: text.Msg(text.NoCommUnit),
		},
		{
			name: "comms on",
			player: func() *Player {
				p := createTestPlayerWithSession(666000, "TestActor", model.RankTrader)
				p.flags2 &^= PL2_COMMS_OFF
				return p
			}(),
			expectedMsg: text.Msg(text.CommsAreOn),
		},
		{
			name: "comms off",
			player: func() *Player {
				p := createTestPlayerWithSession(666000, "TestActor", model.RankTrader)
				p.flags2 |= PL2_COMMS_OFF
				return p
			}(),
			expectedMsg: text.Msg(text.CommsAreOff),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			saver := &testutil.MockPlayerSaver[*Player]{}
			tt.player.saveFunc = saver.Save

			tt.player.CmdComms()

			assertMessageEquals(t, tt.player, tt.expectedMsg)
			testutil.AssertPlayerNotSaved(t, saver)
		})
	}
}

func TestCmdCommsOff(t *testing.T) {
	tests := []struct {
		name           string
		player         *Player
		expectedMsg    string
		expectedFlags2 uint32
		shouldSave     bool
	}{
		{
			name: "no comms unit",
			player: func() *Player {
				p := createTestPlayerWithSession(666000, "TestActor", model.RankTrader)
				p.Flags0 &^= model.PL0_COMM_UNIT
				return p
			}(),
			expectedMsg: text.Msg(text.NoCommUnit),
			shouldSave:  false,
		},
		{
			name: "comms already off",
			player: func() *Player {
				p := createTestPlayerWithSession(666000, "TestActor", model.RankTrader)
				p.flags2 |= PL2_COMMS_OFF
				return p
			}(),
			expectedMsg:    text.Msg(text.CommsAlreadyOff),
			expectedFlags2: PL2_COMMS_OFF,
			shouldSave:     false,
		},
		{
			name: "comms now off",
			player: func() *Player {
				p := createTestPlayerWithSession(666000, "TestActor", model.RankTrader)
				p.flags2 &^= PL2_COMMS_OFF
				return p
			}(),
			expectedMsg:    text.Msg(text.CommsNowOff),
			expectedFlags2: PL2_COMMS_OFF,
			shouldSave:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			saver := &testutil.MockPlayerSaver[*Player]{}
			tt.player.saveFunc = saver.Save

			tt.player.CmdCommsOff()

			assertMessageContains(t, tt.player, tt.expectedMsg)
			if tt.expectedFlags2 != 0 {
				testutil.AssertFlagSet(t, tt.player.flags2, tt.expectedFlags2)
			}
			if tt.shouldSave {
				testutil.AssertPlayerSave(t, saver, database.SaveNow)
			} else {
				testutil.AssertPlayerNotSaved(t, saver)
			}
		})
	}
}

func TestCmdCommsOn(t *testing.T) {
	tests := []struct {
		name              string
		player            *Player
		expectedMsg       string
		notExpectedFlags2 uint32
		shouldSave        bool
	}{
		{
			name: "no comms unit",
			player: func() *Player {
				p := createTestPlayerWithSession(666000, "TestActor", model.RankTrader)
				p.Flags0 &^= model.PL0_COMM_UNIT
				return p
			}(),
			expectedMsg: text.Msg(text.NoCommUnit),
			shouldSave:  false,
		},
		{
			name: "comms already on",
			player: func() *Player {
				p := createTestPlayerWithSession(666000, "TestActor", model.RankTrader)
				p.flags2 &^= PL2_COMMS_OFF
				return p
			}(),
			expectedMsg:       text.Msg(text.CommsNowOn),
			notExpectedFlags2: PL2_COMMS_OFF,
			shouldSave:        false,
		},
		{
			name: "comms now on",
			player: func() *Player {
				p := createTestPlayerWithSession(666000, "TestActor", model.RankTrader)
				p.flags2 |= PL2_COMMS_OFF
				return p
			}(),
			expectedMsg:       text.Msg(text.CommsNowOn),
			notExpectedFlags2: PL2_COMMS_OFF,
			shouldSave:        false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			saver := &testutil.MockPlayerSaver[*Player]{}
			tt.player.saveFunc = saver.Save

			tt.player.CmdCommsOn()

			assertMessageEquals(t, tt.player, tt.expectedMsg)
			if tt.notExpectedFlags2 != 0 {
				testutil.AssertFlagNotSet(t, tt.player.flags2, tt.notExpectedFlags2)
			}
			if tt.shouldSave {
				testutil.AssertPlayerSave(t, saver, database.SaveNow)
			} else {
				testutil.AssertPlayerNotSaved(t, saver)
			}
		})
	}
}
