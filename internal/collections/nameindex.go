package collections

import "github.com/nosborn/federation-1999/internal/text"

type NameIndex[T any] struct {
	items map[string]T
}

func NewNameIndex[T any]() *NameIndex[T] {
	return &NameIndex[T]{
		items: make(map[string]T),
	}
}

func (ni *NameIndex[T]) Insert(key string, value T) {
	lowerKey := text.ToLowerString(key)
	if _, exists := ni.items[lowerKey]; exists {
		panic("key already exists: " + key)
	}
	ni.items[lowerKey] = value
}

func (ni *NameIndex[T]) Remove(key string) {
	delete(ni.items, text.ToLowerString(key))
}

func (ni *NameIndex[T]) Find(key string) (T, bool) {
	value, exists := ni.items[text.ToLowerString(key)]
	return value, exists
}
