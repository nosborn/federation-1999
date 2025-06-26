package server

import (
	"log"
	"testing"

	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/server/database"
	"github.com/nosborn/federation-1999/internal/testutil"
	"github.com/nosborn/federation-1999/internal/text"
	"github.com/stretchr/testify/require"
)

func TestCmdEat(t *testing.T) {
	log.Printf("SolDuchy = %#v", SolDuchy)

	tests := []struct {
		name           string
		setupPlayer    func() *Player
		objectName     model.Name
		expectedMsg    string
		expectedStaCur int32
		checkInventory bool
		expectInInv    bool
		shouldSave     bool
	}{
		{
			name: "eat edible object increases stamina",
			setupPlayer: func() *Player {
				p := createTestPlayerWithSession(666000, "TestActor", model.RankTrader)
				p.Sta = PlayerStat{Cur: 10, Max: 30}
				obj := &Object{name: "cookie", Flags: model.OfEdible}
				p.AddToInventory(obj)
				return p
			},
			objectName:     model.Name{The: true, Words: 1, Text: "cookie"},
			expectedMsg:    text.Msg(text.EatOK),
			expectedStaCur: 15,
			checkInventory: true,
			expectInInv:    false,
			shouldSave:     true,
		},
		{
			name: "eat edible object caps stamina at max",
			setupPlayer: func() *Player {
				p := createTestPlayerWithSession(666001, "TestActor", model.RankTrader)
				p.Sta = PlayerStat{Cur: 28, Max: 30}
				obj := &Object{name: "sandwich", Flags: model.OfEdible}
				p.AddToInventory(obj)
				return p
			},
			objectName:     model.Name{The: true, Words: 1, Text: "sandwich"},
			expectedMsg:    text.Msg(text.EatOK),
			expectedStaCur: 30,
			checkInventory: true,
			expectInInv:    false,
			shouldSave:     true,
		},
		{
			name: "eat object not in inventory",
			setupPlayer: func() *Player {
				p := createTestPlayerWithSession(666002, "TestActor", model.RankTrader)
				p.Sta = PlayerStat{Cur: 10, Max: 30}
				return p
			},
			objectName:     model.Name{The: true, Words: 1, Text: "cookie"},
			expectedMsg:    text.Msg(text.EatNotCarried),
			expectedStaCur: 10,
			checkInventory: false,
			expectInInv:    false,
			shouldSave:     false,
		},
		{
			name: "eat non-edible object",
			setupPlayer: func() *Player {
				p := createTestPlayerWithSession(666003, "TestActor", model.RankTrader)
				p.Sta = PlayerStat{Cur: 10, Max: 30}
				obj := &Object{name: "calendar"}
				p.AddToInventory(obj)
				return p
			},
			objectName:     model.Name{The: true, Words: 1, Text: "calendar"},
			expectedMsg:    text.Msg(text.EatNotEdible),
			expectedStaCur: 10,
			checkInventory: true,
			expectInInv:    true,
			shouldSave:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			player := tt.setupPlayer()
			saver := &testutil.MockPlayerSaver[*Player]{}
			player.saveFunc = saver.Save
			initialInvCount := len(player.inventory)

			player.CmdEat(tt.objectName)

			assertMessageEquals(t, player, tt.expectedMsg)
			require.Equal(t, tt.expectedStaCur, player.Sta.Cur, "Stamina should be updated correctly")
			if tt.shouldSave {
				testutil.AssertPlayerSave(t, saver, database.SaveNow)
			} else {
				testutil.AssertPlayerNotSaved(t, saver)
			}

			if tt.checkInventory {
				if tt.expectInInv {
					require.Len(t, player.inventory, initialInvCount, "Object should remain in inventory")
				} else {
					require.Len(t, player.inventory, initialInvCount-1, "Object should be removed from inventory")

					foundObj, _ := player.FindInventoryName(tt.objectName)
					require.Nil(t, foundObj, "Object should no longer be in inventory")
				}
			}
		})
	}
}
