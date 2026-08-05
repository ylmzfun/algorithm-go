package search

import (
	"sort"
	"testing"
)

func TestLinearSearch(t *testing.T) {
	arr := []int{1, 3, 5, 7, 9}
	if LinearSearch(arr, 5) != 2 {
		t.Error("LinearSearch failed for existing element")
	}
	if LinearSearch(arr, 4) != -1 {
		t.Error("LinearSearch failed for non-existing element")
	}
	if LinearSearch([]int{}, 1) != -1 {
		t.Error("LinearSearch failed for empty array")
	}
}

func TestLinearSearchAll(t *testing.T) {
	arr := []int{1, 3, 1, 5, 1}
	indices := LinearSearchAll(arr, 1)
	if len(indices) != 3 || indices[0] != 0 || indices[1] != 2 || indices[2] != 4 {
		t.Errorf("LinearSearchAll returned wrong indices: %v", indices)
	}
}

func TestBinarySearch(t *testing.T) {
	arr := []int{1, 3, 5, 7, 9}
	if BinarySearch(arr, 5) != 2 {
		t.Error("BinarySearch failed for existing element")
	}
	if BinarySearch(arr, 6) != -1 {
		t.Error("BinarySearch failed for non-existing element")
	}
	if BinarySearch(arr, 1) != 0 {
		t.Error("BinarySearch failed for first element")
	}
	if BinarySearch(arr, 9) != 4 {
		t.Error("BinarySearch failed for last element")
	}
}

func TestBinarySearchRecursive(t *testing.T) {
	arr := []int{1, 3, 5, 7, 9}
	if BinarySearchRecursive(arr, 5, 0, len(arr)-1) != 2 {
		t.Error("BinarySearchRecursive failed")
	}
	if BinarySearchRecursive(arr, 6, 0, len(arr)-1) != -1 {
		t.Error("BinarySearchRecursive failed for non-existing")
	}
}

func TestBinarySearchFirstLast(t *testing.T) {
	arr := []int{1, 2, 2, 2, 3, 4, 5}

	first := BinarySearchFirst(arr, 2)
	if first != 1 {
		t.Errorf("expected first index of 2 = 1, got %d", first)
	}
	last := BinarySearchLast(arr, 2)
	if last != 3 {
		t.Errorf("expected last index of 2 = 3, got %d", last)
	}
	if BinarySearchFirst(arr, 6) != -1 {
		t.Error("expected -1 for non-existing element")
	}
}

func TestLowerUpperBound(t *testing.T) {
	arr := []int{1, 2, 4, 4, 5, 7}

	if LowerBound(arr, 4) != 2 {
		t.Errorf("expected LowerBound(4) = 2, got %d", LowerBound(arr, 4))
	}
	if LowerBound(arr, 3) != 2 {
		t.Errorf("expected LowerBound(3) = 2, got %d", LowerBound(arr, 3))
	}
	if UpperBound(arr, 4) != 4 {
		t.Errorf("expected UpperBound(4) = 4, got %d", UpperBound(arr, 4))
	}
}

func TestJumpSearch(t *testing.T) {
	arr := []int{1, 3, 5, 7, 9}
	if JumpSearch(arr, 5) != 2 {
		t.Error("JumpSearch failed")
	}
	if JumpSearch(arr, 6) != -1 {
		t.Error("JumpSearch failed for non-existing")
	}
}

func TestInterpolationSearch(t *testing.T) {
	arr := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	if InterpolationSearch(arr, 30) != 2 {
		t.Error("InterpolationSearch failed")
	}
	if InterpolationSearch(arr, 55) != -1 {
		t.Error("InterpolationSearch failed for non-existing")
	}
}

func TestExponentialSearch(t *testing.T) {
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	if ExponentialSearch(arr, 9) != 4 {
		t.Error("ExponentialSearch failed")
	}
	if ExponentialSearch(arr, 10) != -1 {
		t.Error("ExponentialSearch failed for non-existing")
	}
	if ExponentialSearch(arr, 1) != 0 {
		t.Error("ExponentialSearch failed for first element")
	}
}

func TestFibonacciSearch(t *testing.T) {
	arr := []int{10, 22, 35, 40, 45, 50, 80, 82, 85, 90, 100}
	if FibonacciSearch(arr, 85) != 8 {
		t.Error("FibonacciSearch failed")
	}
	if FibonacciSearch(arr, 83) != -1 {
		t.Error("FibonacciSearch failed for non-existing")
	}
}

func TestTernarySearch(t *testing.T) {
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15}
	if TernarySearch(arr, 7) != 3 {
		t.Error("TernarySearch failed")
	}
	if TernarySearch(arr, 8) != -1 {
		t.Error("TernarySearch failed for non-existing")
	}
}

func TestAllSearches(t *testing.T) {
	arr := make([]int, 1000)
	for i := range arr {
		arr[i] = i * 2
	}

	searches := map[string]func([]int, int) int{
		"LinearSearch":       LinearSearch,
		"BinarySearch":       BinarySearch,
		"JumpSearch":         JumpSearch,
		"InterpolationSearch": InterpolationSearch,
		"ExponentialSearch":  ExponentialSearch,
		"FibonacciSearch":    FibonacciSearch,
		"TernarySearch":      TernarySearch,
	}

	for name, searchFn := range searches {
		idx := searchFn(arr, 500)
		if idx != 250 {
			t.Errorf("%s: expected index 250 for 500, got %d", name, idx)
		}
		idx = searchFn(arr, 501)
		if idx != -1 {
			t.Errorf("%s: expected -1 for 501, got %d", name, idx)
		}
	}

	// Verify array still sorted (no side effects)
	if !sort.IntsAreSorted(arr) {
		t.Error("search functions modified the array")
	}
}
