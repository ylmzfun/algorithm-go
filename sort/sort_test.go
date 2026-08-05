package sort

import (
	"math/rand"
	"sort"
	"testing"
)

// 基准测试用的大数组
var testSizes = []int{0, 1, 5, 100}

func generateRandomArray(size int) []int {
	arr := make([]int, size)
	for i := range arr {
		arr[i] = rand.Intn(1000)
	}
	return arr
}

func generateSortedArray(size int) []int {
	arr := make([]int, size)
	for i := range arr {
		arr[i] = i
	}
	return arr
}

func generateReversedArray(size int) []int {
	arr := make([]int, size)
	for i := range arr {
		arr[i] = size - i - 1
	}
	return arr
}

func copyArr(arr []int) []int {
	cpy := make([]int, len(arr))
	copy(cpy, arr)
	return cpy
}

func isSorted(arr []int) bool {
	return sort.IntsAreSorted(arr)
}

// --- 冒泡排序测试 ---

func TestBubbleSort(t *testing.T) {
	for _, size := range testSizes {
		arr := generateRandomArray(size)
		BubbleSort(arr)
		if !isSorted(arr) {
			t.Errorf("BubbleSort failed for size %d", size)
		}
	}
}

func TestBubbleSort_Sorted(t *testing.T) {
	arr := generateSortedArray(100)
	BubbleSort(arr)
	if !isSorted(arr) {
		t.Error("BubbleSort failed for sorted array")
	}
}

// --- 选择排序测试 ---

func TestSelectionSort(t *testing.T) {
	for _, size := range testSizes {
		arr := generateRandomArray(size)
		SelectionSort(arr)
		if !isSorted(arr) {
			t.Errorf("SelectionSort failed for size %d", size)
		}
	}
}

// --- 插入排序测试 ---

func TestInsertionSort(t *testing.T) {
	for _, size := range testSizes {
		arr := generateRandomArray(size)
		InsertionSort(arr)
		if !isSorted(arr) {
			t.Errorf("InsertionSort failed for size %d", size)
		}
	}
}

// --- 希尔排序测试 ---

func TestShellSort(t *testing.T) {
	for _, size := range testSizes {
		arr := generateRandomArray(size)
		ShellSort(arr)
		if !isSorted(arr) {
			t.Errorf("ShellSort failed for size %d", size)
		}
	}
}

// --- 归并排序测试 ---

func TestMergeSort(t *testing.T) {
	for _, size := range testSizes {
		arr := generateRandomArray(size)
		MergeSort(arr)
		if !isSorted(arr) {
			t.Errorf("MergeSort failed for size %d", size)
		}
	}
}

// --- 快速排序测试 ---

func TestQuickSort(t *testing.T) {
	for _, size := range testSizes {
		arr := generateRandomArray(size)
		QuickSort(arr)
		if !isSorted(arr) {
			t.Errorf("QuickSort failed for size %d", size)
		}
	}
}

func TestQuickSort_Sorted(t *testing.T) {
	arr := generateSortedArray(100)
	QuickSort(arr)
	if !isSorted(arr) {
		t.Error("QuickSort failed for sorted array")
	}
}

func TestQuickSort_Reversed(t *testing.T) {
	arr := generateReversedArray(100)
	QuickSort(arr)
	if !isSorted(arr) {
		t.Error("QuickSort failed for reversed array")
	}
}

// --- 堆排序测试 ---

func TestHeapSort(t *testing.T) {
	for _, size := range testSizes {
		arr := generateRandomArray(size)
		HeapSort(arr)
		if !isSorted(arr) {
			t.Errorf("HeapSort failed for size %d", size)
		}
	}
}

// --- 计数排序测试 ---

func TestCountingSort(t *testing.T) {
	for _, size := range testSizes {
		arr := generateRandomArray(size)
		CountingSort(arr)
		if !isSorted(arr) {
			t.Errorf("CountingSort failed for size %d", size)
		}
	}
}

// --- 基数排序测试 ---

func TestRadixSort(t *testing.T) {
	for _, size := range testSizes {
		arr := generateRandomArray(size)
		RadixSort(arr)
		if !isSorted(arr) {
			t.Errorf("RadixSort failed for size %d", size)
		}
	}
}

// --- 桶排序测试 ---

func TestBucketSort(t *testing.T) {
	for _, size := range testSizes {
		arr := generateRandomArray(size)
		BucketSort(arr)
		if !isSorted(arr) {
			t.Errorf("BucketSort failed for size %d", size)
		}
	}
}

// --- 快速选择测试 ---

func TestFindKthSmallest(t *testing.T) {
	arr := []int{7, 10, 4, 3, 20, 15}
	result := FindKthSmallest(arr, 3)
	if result != 7 {
		t.Errorf("expected 3rd smallest = 7, got %d", result)
	}

	result = FindKthSmallest(arr, 1)
	if result != 3 {
		t.Errorf("expected 1st smallest = 3, got %d", result)
	}

	result = FindKthSmallest(arr, 6)
	if result != 20 {
		t.Errorf("expected 6th smallest = 20, got %d", result)
	}
}

// --- 综合正确性验证 ---

func TestAllSorts_LargeRandom(t *testing.T) {
	original := generateRandomArray(200)

	sorts := map[string]func([]int){
		"BubbleSort":    BubbleSort,
		"SelectionSort": SelectionSort,
		"InsertionSort": InsertionSort,
		"ShellSort":     ShellSort,
		"MergeSort":     MergeSort,
		"QuickSort":     QuickSort,
		"HeapSort":      HeapSort,
		"CountingSort":  CountingSort,
		"RadixSort":     RadixSort,
		"BucketSort":    BucketSort,
	}

	for name, sortFn := range sorts {
		arr := copyArr(original)
		sortFn(arr)
		if !isSorted(arr) {
			t.Errorf("%s failed to sort correctly", name)
		}
		if len(arr) != len(original) {
			t.Errorf("%s changed array length", name)
		}
	}
}

// --- IsSorted 测试 ---

func TestIsSorted(t *testing.T) {
	if !IsSorted([]int{1, 2, 3, 4, 5}) {
		t.Error("expected true for sorted array")
	}
	if IsSorted([]int{1, 3, 2, 4, 5}) {
		t.Error("expected false for unsorted array")
	}
	if !IsSorted([]int{}) {
		t.Error("expected true for empty array")
	}
	if !IsSorted([]int{42}) {
		t.Error("expected true for single element")
	}
}

// --- Reverse 测试 ---

func TestReverse(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	Reverse(arr)
	for i := 0; i < 5; i++ {
		if arr[i] != 5-i {
			t.Errorf("expected arr[%d]=%d, got %d", i, 5-i, arr[i])
		}
	}
}
