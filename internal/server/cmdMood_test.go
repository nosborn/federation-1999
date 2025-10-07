package server

import (
	"testing"

	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/server/database"
	"github.com/nosborn/federation-1999/internal/testutil"
	"github.com/nosborn/federation-1999/internal/text"
	"github.com/stretchr/testify/require"
)

func TestCmdMood(t *testing.T) {
	tests := []struct {
		name         string
		player       *Player
		text         string
		expectedMsg  string
		expectedMood string
		shouldSave   bool
	}{
		{
			name:         "display mood when not set",
			player:       createTestPlayerWithSession(666000, "TestActor", model.RankTrader),
			text:         "",
			expectedMsg:  text.Msg(text.MoodNotSet),
			expectedMood: "",
			shouldSave:   false,
		},
		{
			name: "display mood when set",
			player: func() *Player {
				p := createTestPlayerWithSession(666000, "TestActor", model.RankTrader)
				p.mood = "is happy"
				return p
			}(),
			text:         "",
			expectedMsg:  text.Msg(text.Mood, "is happy"),
			expectedMood: "is happy",
			shouldSave:   false,
		},
		{
			name:         "set mood",
			player:       createTestPlayerWithSession(666000, "TestActor", model.RankTrader),
			text:         "is coding",
			expectedMsg:  text.Msg(text.MoodSet, "is coding", "TestActor"),
			expectedMood: "is coding",
			shouldSave:   true,
		},
		{
			name:         "mood too long",
			player:       createTestPlayerWithSession(666000, "TestActor", model.RankTrader),
			text:         "is writing a very long mood that will not fit",
			expectedMsg:  text.Msg(text.MoodTooLong, model.MOOD_SIZE-1),
			expectedMood: "",
			shouldSave:   false,
		},
		{
			name:         "mood with bad leader",
			player:       createTestPlayerWithSession(666000, "TestActor", model.RankTrader),
			text:         "/is hacking",
			expectedMsg:  text.Msg(text.MoodBadLeader),
			expectedMood: "",
			shouldSave:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			saver := &testutil.MockPlayerSaver[*Player]{}
			tt.player.saveFunc = saver.Save

			tt.player.CmdMood(tt.text)

			assertMessageEquals(t, tt.player, tt.expectedMsg)
			require.Equal(t, tt.expectedMood, tt.player.mood, "Mood mismatch")
			if tt.shouldSave {
				testutil.AssertPlayerSave(t, saver, database.SaveNow)
			} else {
				testutil.AssertPlayerNotSaved(t, saver)
			}
		})
	}
}

func TestCmdMoodOff(t *testing.T) {
	tests := []struct {
		name         string
		player       *Player
		expectedMsg  string
		expectedMood string
		shouldSave   bool
	}{
		{
			name:         "mood not set",
			player:       createTestPlayerWithSession(666000, "TestActor", model.RankTrader),
			expectedMsg:  text.Msg(text.MoodNotSet),
			expectedMood: "",
			shouldSave:   false,
		},
		{
			name: "clear mood",
			player: func() *Player {
				p := createTestPlayerWithSession(666000, "TestActor", model.RankTrader)
				p.mood = "is happy"
				return p
			}(),
			expectedMsg:  text.Msg(text.MoodCleared),
			expectedMood: "",
			shouldSave:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			saver := &testutil.MockPlayerSaver[*Player]{}
			tt.player.saveFunc = saver.Save

			tt.player.CmdMoodOff()

			assertMessageEquals(t, tt.player, tt.expectedMsg)
			require.Equal(t, tt.expectedMood, tt.player.mood, "Mood mismatch")
			if tt.shouldSave {
				testutil.AssertPlayerSave(t, saver, database.SaveNow)
			} else {
				testutil.AssertPlayerNotSaved(t, saver)
			}
		})
	}
}
