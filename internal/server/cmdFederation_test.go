package server

import (
	"testing"

	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/testutil"
	"github.com/nosborn/federation-1999/internal/text"
	"github.com/nosborn/federation-1999/internal/version"
)

func TestCmdFederation(t *testing.T) {
	t.Run("displays federation version", func(t *testing.T) {
		saver := &testutil.MockPlayerSaver[*Player]{}
		player := createTestPlayerWithSession(666000, "TestActor", model.RankTrader)
		player.saveFunc = saver.Save

		player.CmdFederation()

		expectedMsg := text.Msg(text.Federation, version.String())
		assertMessageEquals(t, player, expectedMsg)
		testutil.AssertPlayerNotSaved(t, saver)
	})
}
