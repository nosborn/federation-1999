package collections

import (
	"errors"
	"iter"
	"strings"
)

// NamedThing defines the interface for objects that can be found by name. This
// follows Federation's pattern for case-insensitive name lookup.
type NamedThing interface {
	Name() string
}

// OrderedCollection is a generic collection that maintains FIFO insertion
// order. It provides operations for insert (error if already present), remove
// (error if not found), and find by case-insensitive name for types that
// implement NamedThing.
type OrderedCollection[T any] struct {
	items []T
}

// NewOrderedCollection creates a new empty ordered collection.
func NewOrderedCollection[T any]() *OrderedCollection[T] {
	return &OrderedCollection[T]{
		items: make([]T, 0),
	}
}

// NewOrderedCollectionWith creates a new ordered collection with initial items.
func NewOrderedCollectionWith[T any](items ...T) *OrderedCollection[T] {
	return &OrderedCollection[T]{
		items: append([]T(nil), items...),
	}
}

// Insert adds an item to the collection. Returns an error if the item is
// already present. For pointer types, comparison is done by pointer equality.
func (oc *OrderedCollection[T]) Insert(item T) error {
	for _, existing := range oc.items {
		if any(existing) == any(item) {
			return errors.New("item already exists in collection")
		}
	}
	oc.items = append(oc.items, item)
	return nil
}

// Remove removes an item from the collection. Returns an error if the item is
// not found. For pointer types, comparison is done by pointer equality.
func (oc *OrderedCollection[T]) Remove(item T) error {
	for i, existing := range oc.items {
		if any(existing) == any(item) {
			oc.items = append(oc.items[:i], oc.items[i+1:]...)
			return nil
		}
	}
	return errors.New("item not found in collection")
}

// All returns a copy of all items in the collection, preserving insertion
// order.
func (oc *OrderedCollection[T]) All() []T {
	result := make([]T, len(oc.items))
	copy(result, oc.items)
	return result
}

// FindByName finds an item by case-insensitive name comparison. This only
// works for types that implement NamedThing.
func (oc *OrderedCollection[T]) FindByName(name string) (T, bool) {
	for _, item := range oc.items {
		if namedItem, ok := any(item).(NamedThing); ok {
			if strings.EqualFold(namedItem.Name(), name) {
				return item, true
			}
		}
	}
	var zero T
	return zero, false
}

// Len returns the number of items in the collection.
func (oc *OrderedCollection[T]) Len() int {
	return len(oc.items)
}

// Values returns an iterator for the collection (Go 1.23+).
func (oc *OrderedCollection[T]) Values() iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, item := range oc.items {
			if !yield(item) {
				return
			}
		}
	}
}

// Enumerate returns an iterator with index and value (Go 1.23+).
func (oc *OrderedCollection[T]) Enumerate() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for i, item := range oc.items {
			if !yield(i, item) {
				return
			}
		}
	}
}
