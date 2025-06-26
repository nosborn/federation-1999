package testutil

import (
	"testing"

	"github.com/nosborn/federation-1999/internal/server/database"
	"github.com/stretchr/testify/require"
)

func AssertFlagNotSet(t *testing.T, flags uint32, unexpectedFlag uint32) {
	t.Helper()
	require.Zero(t, flags&unexpectedFlag, "Expected flag %d to not be set, but it was", unexpectedFlag)
}

func AssertFlagSet(t *testing.T, flags uint32, expectedFlag uint32) {
	t.Helper()
	require.NotZero(t, flags&expectedFlag, "Expected flag %d to be set, but it was not", expectedFlag)
}

func AssertPlayerNotSaved[T any](t *testing.T, saver *MockPlayerSaver[T]) {
	t.Helper()
	require.False(t, saver.WasCalled, "Expected player not to be saved, but it was")
}

func AssertPlayerSave[T any](t *testing.T, saver *MockPlayerSaver[T], when database.SaveWhen) {
	t.Helper()
	require.True(t, saver.WasCalled, "Expected player to be saved, but it was not")
	require.Equal(t, when, saver.When, "Player was saved with unexpected 'when' value")
}

// func AssertSessionQuit(t *testing.T, session *MockSession) {
// 	t.Helper()
// 	require.True(t, session.QuitCalled, "Expected session Quit() to be called, but it was not")
// }
