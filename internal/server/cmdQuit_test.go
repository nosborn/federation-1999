package server

// import (
// 	"testing"
//
// 	"github.com/nosborn/federation-1999/internal/model"
// 	"github.com/nosborn/federation-1999/internal/testutil"
// 	"github.com/nosborn/federation-1999/internal/text"
// )
//
// func TestCmdQuit(t *testing.T) {
// 	t.Run("displays quit message and calls session Quit", func(t *testing.T) {
// 		player := createTestPlayerWithSession(666000, "TestActor", model.RankTrader)
// 		saver := &testutil.MockPlayerSaver[*Player]{}
// 		player.saveFunc = saver.Save
//
// 		player.CmdQuit()
//
// 		// Check message output
// 		expectedMsg := text.Msg(text.Quit)
// 		assertMessageEquals(t, player, expectedMsg)
//
// 		// Check that session Quit was called
// 		testutil.AssertSessionQuit(t, player.Session)
//
// 		// CmdQuit doesn't save player data
// 		testutil.AssertPlayerNotSaved(t, saver)
// 	})
// }
