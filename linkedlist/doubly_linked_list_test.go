package linkedlist

import (
	"testing"
)

func TestDoublyLinkedList_AddFirst(t *testing.T) {
	dll := NewDoublyLinkedList()
	dll.AddFirst(3)
	dll.AddFirst(2)
	dll.AddFirst(1)

	if dll.Size() != 3 {
		t.Errorf("expected size 3, got %d", dll.Size())
	}
	first, _ := dll.GetFirst()
	if first != 1 {
		t.Errorf("expected first=1, got %v", first)
	}
	last, _ := dll.GetLast()
	if last != 3 {
		t.Errorf("expected last=3, got %v", last)
	}
}

func TestDoublyLinkedList_AddLast(t *testing.T) {
	dll := NewDoublyLinkedList()
	dll.AddLast(1)
	dll.AddLast(2)
	dll.AddLast(3)

	if dll.Size() != 3 {
		t.Errorf("expected size 3, got %d", dll.Size())
	}
	first, _ := dll.GetFirst()
	if first != 1 {
		t.Errorf("expected first=1, got %v", first)
	}
}

func TestDoublyLinkedList_RemoveFirst(t *testing.T) {
	dll := NewDoublyLinkedList()
	dll.AddLast(1)
	dll.AddLast(2)
	dll.AddLast(3)

	val, err := dll.RemoveFirst()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if val != 1 {
		t.Errorf("expected 1, got %v", val)
	}
	if dll.Size() != 2 {
		t.Errorf("expected size 2, got %d", dll.Size())
	}
}

func TestDoublyLinkedList_RemoveLast(t *testing.T) {
	dll := NewDoublyLinkedList()
	dll.AddLast(1)
	dll.AddLast(2)
	dll.AddLast(3)

	val, err := dll.RemoveLast()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if val != 3 {
		t.Errorf("expected 3, got %v", val)
	}
}

func TestDoublyLinkedList_Insert(t *testing.T) {
	dll := NewDoublyLinkedList()
	dll.AddLast(1)
	dll.AddLast(3)
	dll.Insert(1, 2)

	if dll.Size() != 3 {
		t.Errorf("expected size 3, got %d", dll.Size())
	}
	val, _ := dll.Get(1)
	if val != 2 {
		t.Errorf("expected 2 at index 1, got %v", val)
	}
}

func TestDoublyLinkedList_RemoveElement(t *testing.T) {
	dll := NewDoublyLinkedList()
	dll.AddLast(1)
	dll.AddLast(2)
	dll.AddLast(3)

	if !dll.RemoveElement(2) {
		t.Error("expected RemoveElement to return true")
	}
	if dll.Size() != 2 {
		t.Errorf("expected size 2, got %d", dll.Size())
	}
	if dll.RemoveElement(99) {
		t.Error("expected RemoveElement to return false for non-existing element")
	}
}

func TestDoublyLinkedList_Contains(t *testing.T) {
	dll := NewDoublyLinkedList()
	dll.AddLast(1)
	dll.AddLast(2)

	if !dll.Contains(1) {
		t.Error("expected Contains(1) to be true")
	}
	if dll.Contains(3) {
		t.Error("expected Contains(3) to be false")
	}
}

func TestDoublyLinkedList_ToSliceReverse(t *testing.T) {
	dll := NewDoublyLinkedList()
	dll.AddLast(1)
	dll.AddLast(2)
	dll.AddLast(3)

	reversed := dll.ToSliceReverse()
	if len(reversed) != 3 || reversed[0] != 3 || reversed[1] != 2 || reversed[2] != 1 {
		t.Errorf("expected [3,2,1], got %v", reversed)
	}
}

func TestDoublyLinkedList_Empty(t *testing.T) {
	dll := NewDoublyLinkedList()
	if !dll.IsEmpty() {
		t.Error("expected empty list")
	}
	_, err := dll.RemoveFirst()
	if err == nil {
		t.Error("expected error on RemoveFirst from empty list")
	}
	_, err = dll.RemoveLast()
	if err == nil {
		t.Error("expected error on RemoveLast from empty list")
	}
}

func TestDoublyLinkedList_Clear(t *testing.T) {
	dll := NewDoublyLinkedList()
	dll.AddLast(1)
	dll.AddLast(2)
	dll.Clear()

	if !dll.IsEmpty() {
		t.Error("expected empty list after Clear")
	}
	if dll.Size() != 0 {
		t.Errorf("expected size 0, got %d", dll.Size())
	}
}
