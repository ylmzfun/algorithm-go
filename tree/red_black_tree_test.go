package tree

import (
	"testing"
)

func TestRedBlackTree_Insert(t *testing.T) {
	rbt := NewRedBlackTree()
	values := []int{10, 20, 30, 15, 25, 5, 1}
	for _, v := range values {
		rbt.Insert(IntValue(v), v)
	}

	if rbt.Size() != 7 {
		t.Errorf("expected size 7, got %d", rbt.Size())
	}
	if !rbt.IsValid() {
		t.Error("expected valid red-black tree")
	}
}

func TestRedBlackTree_Search(t *testing.T) {
	rbt := NewRedBlackTree()
	rbt.Insert(IntValue(10), "ten")
	rbt.Insert(IntValue(20), "twenty")

	val, err := rbt.Search(IntValue(10))
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if val != "ten" {
		t.Errorf("expected 'ten', got %v", val)
	}

	_, err = rbt.Search(IntValue(99))
	if err == nil {
		t.Error("expected error for non-existing key")
	}
}

func TestRedBlackTree_Delete(t *testing.T) {
	rbt := NewRedBlackTree()
	for _, v := range []int{10, 20, 30, 15, 25} {
		rbt.Insert(IntValue(v), v)
	}

	err := rbt.Delete(IntValue(15))
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if rbt.Size() != 4 {
		t.Errorf("expected size 4, got %d", rbt.Size())
	}
	if !rbt.IsValid() {
		t.Error("expected valid red-black tree after deletion")
	}
	if rbt.Contains(IntValue(15)) {
		t.Error("expected key 15 to be deleted")
	}
}

func TestRedBlackTree_MinMax(t *testing.T) {
	rbt := NewRedBlackTree()
	for _, v := range []int{50, 30, 70, 20, 40, 60, 80} {
		rbt.Insert(IntValue(v), v)
	}

	min, _ := rbt.Min()
	if min.CompareTo(IntValue(20)) != 0 {
		t.Errorf("expected min=20, got %v", min)
	}
	max, _ := rbt.Max()
	if max.CompareTo(IntValue(80)) != 0 {
		t.Errorf("expected max=80, got %v", max)
	}
}

func TestRedBlackTree_SequentialInsert(t *testing.T) {
	rbt := NewRedBlackTree()
	for i := 0; i < 100; i++ {
		rbt.Insert(IntValue(i), i)
	}
	if !rbt.IsValid() {
		t.Error("expected valid red-black tree after sequential inserts")
	}
	// 所有键存在且有序
	traversal := rbt.InorderTraversal()
	for i := 0; i < 100; i++ {
		if traversal[i].CompareTo(IntValue(i)) != 0 {
			t.Errorf("expected inorder[%d]=%d, got %v", i, i, traversal[i])
		}
	}
}

func TestRedBlackTree_Empty(t *testing.T) {
	rbt := NewRedBlackTree()
	if !rbt.IsEmpty() {
		t.Error("expected empty tree")
	}
	_, err := rbt.Min()
	if err == nil {
		t.Error("expected error for Min on empty tree")
	}
}
