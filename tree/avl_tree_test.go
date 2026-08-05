package tree

import (
	"testing"
)

func TestAVLTree_Insert(t *testing.T) {
	avl := NewAVLTree()
	values := []int{10, 20, 30, 40, 50, 25}
	for _, v := range values {
		avl.Insert(IntValue(v), v)
	}

	if avl.Size() != 6 {
		t.Errorf("expected size 6, got %d", avl.Size())
	}
	if !avl.IsBalanced() {
		t.Error("expected tree to be balanced")
	}
}

func TestAVLTree_Search(t *testing.T) {
	avl := NewAVLTree()
	avl.Insert(IntValue(10), "ten")
	avl.Insert(IntValue(20), "twenty")

	val, err := avl.Search(IntValue(10))
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if val != "ten" {
		t.Errorf("expected 'ten', got %v", val)
	}

	_, err = avl.Search(IntValue(99))
	if err == nil {
		t.Error("expected error for non-existing key")
	}
}

func TestAVLTree_Delete(t *testing.T) {
	avl := NewAVLTree()
	for _, v := range []int{10, 20, 30, 40, 50} {
		avl.Insert(IntValue(v), v)
	}

	err := avl.Delete(IntValue(30))
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if avl.Size() != 4 {
		t.Errorf("expected size 4, got %d", avl.Size())
	}
	if !avl.IsBalanced() {
		t.Error("expected tree to be balanced after deletion")
	}
	if avl.Contains(IntValue(30)) {
		t.Error("expected key 30 to be deleted")
	}
}

func TestAVLTree_MinMax(t *testing.T) {
	avl := NewAVLTree()
	for _, v := range []int{50, 30, 70, 20, 40, 60, 80} {
		avl.Insert(IntValue(v), v)
	}

	min, _ := avl.Min()
	if min.CompareTo(IntValue(20)) != 0 {
		t.Errorf("expected min=20, got %v", min)
	}
	max, _ := avl.Max()
	if max.CompareTo(IntValue(80)) != 0 {
		t.Errorf("expected max=80, got %v", max)
	}
}

func TestAVLTree_BalanceAfterSequentialInsert(t *testing.T) {
	avl := NewAVLTree()
	// 顺序插入会触发多次旋转
	for i := 0; i < 100; i++ {
		avl.Insert(IntValue(i), i)
		if !avl.IsBalanced() {
			t.Errorf("tree unbalanced after inserting %d", i)
		}
	}
	if avl.Height() > 10 {
		t.Errorf("height too large for balanced tree: %d", avl.Height())
	}
}

func TestAVLTree_Empty(t *testing.T) {
	avl := NewAVLTree()
	if !avl.IsEmpty() {
		t.Error("expected empty tree")
	}
	_, err := avl.Min()
	if err == nil {
		t.Error("expected error for Min on empty tree")
	}
}
