package collections

import (
	"slices"
	"testing"
)

// TestItem implements NamedThing for testing
type TestItem struct {
	name  string
	value int
}

func (t *TestItem) Name() string {
	return t.name
}

func TestOrderedCollection_Insert(t *testing.T) {
	oc := NewOrderedCollection[*TestItem]()

	item1 := &TestItem{name: "first", value: 1}
	item2 := &TestItem{name: "second", value: 2}
	item3 := &TestItem{name: "third", value: 3}

	// Test successful insertions
	if err := oc.Insert(item1); err != nil {
		t.Errorf("Insert should succeed: %v", err)
	}
	if err := oc.Insert(item2); err != nil {
		t.Errorf("Insert should succeed: %v", err)
	}
	if err := oc.Insert(item3); err != nil {
		t.Errorf("Insert should succeed: %v", err)
	}

	if oc.Len() != 3 {
		t.Errorf("Expected length 3, got %d", oc.Len())
	}

	// Test duplicate insertion fails
	if err := oc.Insert(item1); err == nil {
		t.Error("Insert should fail for duplicate item")
	}
}

func TestOrderedCollection_Remove(t *testing.T) {
	oc := NewOrderedCollection[*TestItem]()

	item1 := &TestItem{name: "first", value: 1}
	item2 := &TestItem{name: "second", value: 2}
	item3 := &TestItem{name: "third", value: 3}

	oc.Insert(item1)
	oc.Insert(item2)
	oc.Insert(item3)

	// Test successful removal
	if err := oc.Remove(item2); err != nil {
		t.Errorf("Remove should succeed: %v", err)
	}

	if oc.Len() != 2 {
		t.Errorf("Expected length 2, got %d", oc.Len())
	}

	// Test removal of non-existent item fails
	if err := oc.Remove(item2); err == nil {
		t.Error("Remove should fail for non-existent item")
	}

	// Test removal of remaining items
	if err := oc.Remove(item1); err != nil {
		t.Errorf("Remove should succeed: %v", err)
	}
	if err := oc.Remove(item3); err != nil {
		t.Errorf("Remove should succeed: %v", err)
	}

	if oc.Len() != 0 {
		t.Errorf("Expected length 0, got %d", oc.Len())
	}
}

func TestOrderedCollection_FIFOOrder(t *testing.T) {
	oc := NewOrderedCollection[*TestItem]()

	// Insert items in a specific order
	items := []*TestItem{
		{name: "first", value: 1},
		{name: "second", value: 2},
		{name: "third", value: 3},
		{name: "fourth", value: 4},
	}

	for _, item := range items {
		oc.Insert(item)
	}

	// Check that All() returns items in insertion order
	all := oc.All()
	if len(all) != len(items) {
		t.Errorf("Expected %d items, got %d", len(items), len(all))
	}

	for i, item := range all {
		if item != items[i] {
			t.Errorf("Expected item at index %d to be %v, got %v", i, items[i], item)
		}
	}

	// Test that removing middle item preserves order
	oc.Remove(items[1]) // Remove "second"

	expected := []*TestItem{items[0], items[2], items[3]}
	all = oc.All()

	if len(all) != len(expected) {
		t.Errorf("Expected %d items after removal, got %d", len(expected), len(all))
	}

	for i, item := range all {
		if item != expected[i] {
			t.Errorf("Expected item at index %d to be %v, got %v", i, expected[i], item)
		}
	}
}

func TestOrderedCollection_FindByName(t *testing.T) {
	oc := NewOrderedCollection[*TestItem]()

	item1 := &TestItem{name: "Alpha", value: 1}
	item2 := &TestItem{name: "Beta", value: 2}
	item3 := &TestItem{name: "Gamma", value: 3}

	oc.Insert(item1)
	oc.Insert(item2)
	oc.Insert(item3)

	// Test case-insensitive search
	found, ok := oc.FindByName("alpha")
	if !ok {
		t.Error("Should find item with case-insensitive search")
	}
	if found != item1 {
		t.Errorf("Expected to find item1, got %v", found)
	}

	found, ok = oc.FindByName("BETA")
	if !ok {
		t.Error("Should find item with case-insensitive search")
	}
	if found != item2 {
		t.Errorf("Expected to find item2, got %v", found)
	}

	found, ok = oc.FindByName("gamma")
	if !ok {
		t.Error("Should find item with case-insensitive search")
	}
	if found != item3 {
		t.Errorf("Expected to find item3, got %v", found)
	}

	// Test non-existent name
	_, ok = oc.FindByName("nonexistent")
	if ok {
		t.Error("Should not find non-existent item")
	}
}

func TestOrderedCollection_Values(t *testing.T) {
	oc := NewOrderedCollection[*TestItem]()

	items := []*TestItem{
		{name: "first", value: 1},
		{name: "second", value: 2},
		{name: "third", value: 3},
	}

	for _, item := range items {
		oc.Insert(item)
	}

	// Test Values() iterator
	var collected []*TestItem
	for item := range oc.Values() {
		collected = append(collected, item)
	}

	if len(collected) != len(items) {
		t.Errorf("Expected %d items from Values(), got %d", len(items), len(collected))
	}

	for i, item := range collected {
		if item != items[i] {
			t.Errorf("Expected item at index %d to be %v, got %v", i, items[i], item)
		}
	}
}

func TestOrderedCollection_Enumerate(t *testing.T) {
	oc := NewOrderedCollection[*TestItem]()

	items := []*TestItem{
		{name: "first", value: 1},
		{name: "second", value: 2},
		{name: "third", value: 3},
	}

	for _, item := range items {
		oc.Insert(item)
	}

	// Test Enumerate() iterator
	var collectedIndices []int
	var collectedItems []*TestItem
	for i, item := range oc.Enumerate() {
		collectedIndices = append(collectedIndices, i)
		collectedItems = append(collectedItems, item)
	}

	expectedIndices := []int{0, 1, 2}
	if !slices.Equal(collectedIndices, expectedIndices) {
		t.Errorf("Expected indices %v, got %v", expectedIndices, collectedIndices)
	}

	if len(collectedItems) != len(items) {
		t.Errorf("Expected %d items from Enumerate(), got %d", len(items), len(collectedItems))
	}

	for i, item := range collectedItems {
		if item != items[i] {
			t.Errorf("Expected item at index %d to be %v, got %v", i, items[i], item)
		}
	}
}

func TestOrderedCollection_All_IsCopy(t *testing.T) {
	oc := NewOrderedCollection[*TestItem]()

	item1 := &TestItem{name: "test", value: 1}
	oc.Insert(item1)

	// Get a copy via All()
	all := oc.All()

	// Modify the returned slice
	all[0] = &TestItem{name: "modified", value: 999}

	// Original collection should be unchanged
	original := oc.All()
	if original[0].name != "test" || original[0].value != 1 {
		t.Error("All() should return a copy that doesn't affect the original")
	}
}

func TestOrderedCollection_EmptyCollection(t *testing.T) {
	oc := NewOrderedCollection[*TestItem]()

	// Test empty collection
	if oc.Len() != 0 {
		t.Errorf("Expected length 0 for empty collection, got %d", oc.Len())
	}

	all := oc.All()
	if len(all) != 0 {
		t.Errorf("Expected empty slice from All(), got %d items", len(all))
	}

	// Test remove from empty collection
	item := &TestItem{name: "test", value: 1}
	if err := oc.Remove(item); err == nil {
		t.Error("Remove should fail on empty collection")
	}

	// Test find in empty collection
	_, ok := oc.FindByName("test")
	if ok {
		t.Error("FindByName should fail on empty collection")
	}

	// Test iterators on empty collection
	count := 0
	for range oc.Values() {
		count++
	}
	if count != 0 {
		t.Errorf("Expected 0 iterations for empty collection, got %d", count)
	}

	count = 0
	for range oc.Enumerate() {
		count++
	}
	if count != 0 {
		t.Errorf("Expected 0 iterations for empty collection, got %d", count)
	}
}

func TestOrderedCollection_NonNamedThing(t *testing.T) {
	// Test with types that don't implement NamedThing
	oc := NewOrderedCollection[string]()

	oc.Insert("hello")
	oc.Insert("world")

	// FindByName should return zero value and false for non-NamedThing types
	result, ok := oc.FindByName("hello")
	if ok {
		t.Error("FindByName should return false for non-NamedThing types")
	}
	if result != "" {
		t.Errorf("Expected zero value for non-NamedThing types, got %v", result)
	}
}
