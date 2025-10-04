package collections

import "testing"

type NamedItem struct {
	Name string
	ID   int
}

type TestRecord struct {
	Label       string
	Description string
	Value       int
}

func TestNameIndex_InsertAndFind(t *testing.T) {
	ni := NewNameIndex[*NamedItem]()
	item := &NamedItem{Name: "Alice", ID: 123}

	ni.Insert(item.Name, item)

	found, exists := ni.Find(item.Name)
	if !exists {
		t.Error("Expected to find Alice")
	}
	if found != item {
		t.Error("Expected same pointer reference")
	}
	if found.Name != "Alice" || found.ID != 123 {
		t.Errorf("Expected NamedItem{Alice, 123}, got %+v", found)
	}
}

func TestNameIndex_CaseInsensitive(t *testing.T) {
	ni := NewNameIndex[*TestRecord]()
	record := &TestRecord{Label: "Entry", Description: "Test record", Value: 100}

	ni.Insert(record.Label, record)

	found, exists := ni.Find("entry")
	if !exists {
		t.Error("Expected to find entry (case insensitive)")
	}
	if found != record {
		t.Error("Expected same pointer reference")
	}

	found, exists = ni.Find("ENTRY")
	if !exists {
		t.Error("Expected to find ENTRY (case insensitive)")
	}
	if found != record {
		t.Error("Expected same pointer reference")
	}
}

func TestNameIndex_Remove(t *testing.T) {
	ni := NewNameIndex[*NamedItem]()
	item := &NamedItem{Name: "Bob", ID: 456}

	ni.Insert(item.Name, item)
	ni.Remove(item.Name)

	_, found := ni.Find(item.Name)
	if found {
		t.Error("Expected Bob to be removed")
	}
}

func TestNameIndex_RemoveCaseInsensitive(t *testing.T) {
	ni := NewNameIndex[*NamedItem]()
	item := &NamedItem{Name: "Charlie", ID: 789}

	ni.Insert(item.Name, item)
	ni.Remove("charlie")

	_, found := ni.Find(item.Name)
	if found {
		t.Error("Expected Charlie to be removed (case insensitive)")
	}
}

func TestNameIndex_FindNonExistent(t *testing.T) {
	ni := NewNameIndex[*NamedItem]()

	value, found := ni.Find("NonExistent")
	if found {
		t.Error("Expected not to find NonExistent")
	}
	if value != nil {
		t.Errorf("Expected nil for zero value, got %v", value)
	}
}

func TestNameIndex_DuplicateKeyPanic(t *testing.T) {
	ni := NewNameIndex[*NamedItem]()
	item1 := &NamedItem{Name: "Dave", ID: 111}
	item2 := &NamedItem{Name: "Dave", ID: 222}

	ni.Insert(item1.Name, item1)

	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic when inserting duplicate key")
		}
	}()

	ni.Insert(item2.Name, item2)
}

func TestNameIndex_DuplicateKeyCaseInsensitivePanic(t *testing.T) {
	ni := NewNameIndex[*NamedItem]()
	item1 := &NamedItem{Name: "Eve", ID: 333}
	item2 := &NamedItem{Name: "eve", ID: 444}

	ni.Insert(item1.Name, item1)

	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic when inserting duplicate key (case insensitive)")
		}
	}()

	ni.Insert(item2.Name, item2)
}
