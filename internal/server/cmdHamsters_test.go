package server

import (
	"testing"

	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/testutil"
	"github.com/nosborn/federation-1999/internal/text"
)

func TestCmdHamsters(t *testing.T) {
	t.Run("displays hamsters message", func(t *testing.T) {
		saver := &testutil.MockPlayerSaver[*Player]{}
		player := createTestPlayerWithSession(666000, "TestActor", model.RankTrader)
		player.saveFunc = saver.Save

		player.CmdHamsters()

		expectedMsg := text.Msg(text.Hamsters)
		assertMessageEquals(t, player, expectedMsg)
		testutil.AssertPlayerNotSaved(t, saver)
	})
}
