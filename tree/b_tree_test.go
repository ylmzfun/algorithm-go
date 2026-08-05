package tree

import (
	"testing"
)

func TestBTree_Insert(t *testing.T) {
	bt := NewBTree(3)
	values := []int{10, 20, 5, 6, 12, 30, 7, 17}
	for _, v := range values {
		bt.Insert(IntValue(v), v)
	}

	if bt.Size() != 8 {
		t.Errorf("expected size 8, got %d", bt.Size())
	}

	inorder := bt.InorderTraversal()
	for i := 1; i < len(inorder); i++ {
		if inorder[i-1].CompareTo(inorder[i]) > 0 {
			t.Error("inorder traversal is not sorted")
		}
	}
}

func TestBTree_Search(t *testing.T) {
	bt := NewBTree(3)
	bt.Insert(IntValue(10), "ten")
	bt.Insert(IntValue(20), "twenty")

	val, err := bt.Search(IntValue(10))
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if val != "ten" {
		t.Errorf("expected 'ten', got %v", val)
	}

	_, err = bt.Search(IntValue(99))
	if err == nil {
		t.Error("expected error for non-existing key")
	}
}

func TestBTree_Delete(t *testing.T) {
	bt := NewBTree(3)
	for _, v := range []int{10, 20, 5, 6, 12, 30, 7, 17} {
		bt.Insert(IntValue(v), v)
	}

	err := bt.Delete(IntValue(6))
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if bt.Size() != 7 {
		t.Errorf("expected size 7, got %d", bt.Size())
	}
	if bt.Contains(IntValue(6)) {
		t.Error("expected key 6 to be deleted")
	}
}

func TestBTree_MinMax(t *testing.T) {
	bt := NewBTree(3)
	for _, v := range []int{50, 30, 70, 20, 40, 60, 80} {
		bt.Insert(IntValue(v), v)
	}

	min, _ := bt.Min()
	if min.CompareTo(IntValue(20)) != 0 {
		t.Errorf("expected min=20, got %v", min)
	}
	max, _ := bt.Max()
	if max.CompareTo(IntValue(80)) != 0 {
		t.Errorf("expected max=80, got %v", max)
	}
}

func TestBTree_LargeDataset(t *testing.T) {
	bt := NewBTree(5)
	for i := 0; i < 1000; i++ {
		bt.Insert(IntValue(i), i)
	}

	if bt.Size() != 1000 {
		t.Errorf("expected size 1000, got %d", bt.Size())
	}

	inorder := bt.InorderTraversal()
	for i := 0; i < 1000; i++ {
		if inorder[i].CompareTo(IntValue(i)) != 0 {
			t.Errorf("expected inorder[%d]=%d, got %v", i, i, inorder[i])
		}
	}
}
