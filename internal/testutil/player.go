package testutil

import "github.com/nosborn/federation-1999/internal/server/database"

type MockPlayerSaver[T any] struct {
	Player    T
	WasCalled bool
	When      database.SaveWhen
}

func (m *MockPlayerSaver[T]) Save(p T, when database.SaveWhen) {
	m.Player = p
	m.WasCalled = true
	m.When = when
}

type MockPlayer interface {
	// TODO
}
