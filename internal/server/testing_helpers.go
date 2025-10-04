// Claude:
// - All assertX functions must be sorted alphabetically by name.

package server

import (
	"strings"
	"testing"

	"github.com/nosborn/federation-1999/internal/collections"
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/testutil"
	"github.com/nosborn/federation-1999/pkg/ibgames"
	"github.com/stretchr/testify/require"
)

func testResetGameWorld() {
	allDuchies = collections.NewOrderedCollection[*Duchy]()
	duchyIndex = collections.NewNameIndex[*Duchy]()
	allSystems = collections.NewOrderedCollection[Systemer]()
	systemIndex = collections.NewNameIndex[Systemer]()
	planetIndex = collections.NewNameIndex[*Planet]()
	Players = make(map[string]*Player)
}

func testInitializeGameWorld() {
	testResetGameWorld()
	NewLoader(func() {}) // Initialize loader to avoid nil pointer crash
	SolDuchy = NewSolDuchy()
	SolSystem = NewSolSystem(SolDuchy)
	arenaSystem = NewArenaSystem(SolDuchy)
	snarkSystem = NewSnarkSystem(SolDuchy)
}

func createTestPlayer(uid ibgames.AccountID, name string, rank model.Rank) *Player {
	if Players == nil {
		Players = make(map[string]*Player)
	}

	player := &Player{
		name:       name,
		uid:        uid,
		rank:       rank,
		MsgOut:     strings.Builder{},
		saveFunc:   (&testutil.MockPlayerSaver[*Player]{}).Save,
		Flags0:     model.PL0_COMM_UNIT,
		curSysName: "Sol",
		curSys:     SolSystem,
	}
	Players[name] = player
	return player
}

func createTestPlayerWithSession(uid ibgames.AccountID, name string, rank model.Rank) *Player {
	player := createTestPlayer(uid, name, rank)
	player.session = &Session{}
	return player
}

func assertMessageContains(t *testing.T, player *Player, expected string) {
	t.Helper()
	actual := player.MsgOut.String()
	require.Contains(t, actual, expected, "Message output should contain expected text")
}

func assertMessageEquals(t *testing.T, player *Player, expected string) {
	t.Helper()
	actual := player.MsgOut.String()
	require.Equal(t, expected, actual, "Message output mismatch")
}

// func assertPlayerStatDecrease(t *testing.T, actual PlayerStat, expectedCur, expectedMax int, statName string) {
// 	t.Helper()
// 	require.Equal(t, expectedCur, actual.Cur, "%s current should be %d", statName, expectedCur)
// 	require.Equal(t, expectedMax, actual.Max, "%s max should be %d", statName, expectedMax)
// 	require.LessOrEqual(t, actual.Cur, actual.Max, "%s current should not exceed max", statName)
// 	require.GreaterOrEqual(t, actual.Max, 1, "%s max should not be below 1", statName)
// }

// func assertPlayerStat(t *testing.T, actual PlayerStat, expectedCur, expectedMax int, statName string) {
// 	t.Helper()
// 	require.Equal(t, expectedCur, actual.Cur, "%s current should be %d", statName, expectedCur)
// 	require.Equal(t, expectedMax, actual.Max, "%s max should be %d", statName, expectedMax)
// 	require.LessOrEqual(t, actual.Cur, actual.Max, "%s current should not exceed max", statName)
// 	require.LessOrEqual(t, actual.Max, 120, "%s max should not exceed 120", statName)
// }

// func assertPlayerStatIncrease(t *testing.T, actual PlayerStat, expectedCur, expectedMax int, statName string) {
// 	t.Helper()
// 	require.Equal(t, expectedCur, actual.Cur, "%s current should be %d", statName, expectedCur)
// 	require.Equal(t, expectedMax, actual.Max, "%s max should be %d", statName, expectedMax)
// 	require.LessOrEqual(t, actual.Cur, actual.Max, "%s current should not exceed max", statName)
// 	require.LessOrEqual(t, actual.Max, 120, "%s max should not exceed 120", statName)
// }
