package server

import (
	"testing"

	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/server/database"
	"github.com/nosborn/federation-1999/internal/testutil"
	"github.com/nosborn/federation-1999/internal/text"
)

func TestCmdJobs(t *testing.T) {
	tests := []struct {
		name        string
		player      *Player
		expectedMsg string
	}{
		{
			name: "jobs off",
			player: func() *Player {
				p := createTestPlayerWithSession(666000, "TestActor", model.RankTrader)
				p.Flags0 &^= model.PL0_JOB
				return p
			}(),
			expectedMsg: text.Msg(text.JobsAreOff),
		},
		{
			name: "jobs on",
			player: func() *Player {
				p := createTestPlayerWithSession(666000, "TestActor", model.RankTrader)
				p.Flags0 |= model.PL0_JOB
				return p
			}(),
			expectedMsg: text.Msg(text.JobsAreOn),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			saver := &testutil.MockPlayerSaver[*Player]{}
			tt.player.saveFunc = saver.Save

			tt.player.CmdJobs()

			assertMessageEquals(t, tt.player, tt.expectedMsg)
			testutil.AssertPlayerNotSaved(t, saver)
		})
	}
}

func TestCmdJobsOff(t *testing.T) {
	tests := []struct {
		name              string
		player            *Player
		expectedMsg       string
		notExpectedFlags0 uint32
		shouldSave        bool
	}{
		{
			name: "jobs already off",
			player: func() *Player {
				p := createTestPlayerWithSession(666000, "TestActor", model.RankTrader)
				p.Flags0 &^= model.PL0_JOB
				return p
			}(),
			expectedMsg:       text.Msg(text.JobsNowOff),
			notExpectedFlags0: model.PL0_JOB,
			shouldSave:        false,
		},
		{
			name: "jobs now off",
			player: func() *Player {
				p := createTestPlayerWithSession(666000, "TestActor", model.RankTrader)
				p.Flags0 |= model.PL0_JOB
				return p
			}(),
			expectedMsg:       text.Msg(text.JobsNowOff),
			notExpectedFlags0: model.PL0_JOB,
			shouldSave:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			saver := &testutil.MockPlayerSaver[*Player]{}
			tt.player.saveFunc = saver.Save

			tt.player.CmdJobsOff()

			assertMessageEquals(t, tt.player, tt.expectedMsg)
			testutil.AssertFlagNotSet(t, tt.player.Flags0, tt.notExpectedFlags0)
			if tt.shouldSave {
				testutil.AssertPlayerSave(t, saver, database.SaveNow)
			} else {
				testutil.AssertPlayerNotSaved(t, saver)
			}
		})
	}
}

func TestCmdJobsOn(t *testing.T) {
	tests := []struct {
		name           string
		player         *Player
		expectedMsg    string
		expectedFlags0 uint32
		shouldSave     bool
	}{
		{
			name: "jobs already on",
			player: func() *Player {
				p := createTestPlayerWithSession(666000, "TestActor", model.RankTrader)
				p.Flags0 |= model.PL0_JOB
				return p
			}(),
			expectedMsg:    text.Msg(text.JobsNowOn),
			expectedFlags0: model.PL0_JOB,
			shouldSave:     false,
		},
		{
			name: "jobs now on",
			player: func() *Player {
				p := createTestPlayerWithSession(666000, "TestActor", model.RankTrader)
				p.Flags0 &^= model.PL0_JOB
				return p
			}(),
			expectedMsg:    text.Msg(text.JobsNowOn),
			expectedFlags0: model.PL0_JOB,
			shouldSave:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			saver := &testutil.MockPlayerSaver[*Player]{}
			tt.player.saveFunc = saver.Save

			tt.player.CmdJobsOn()

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
