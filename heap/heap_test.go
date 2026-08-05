package heap

import (
	"reflect"
	"testing"
)

// TestNewIntMaxHeap 测试整数最大堆创建
func TestNewIntMaxHeap(t *testing.T) {
	heap := NewIntMaxHeap()
	if !heap.IsEmpty() {
		t.Error("Expected heap to be empty")
	}
	if heap.Size() != 0 {
		t.Errorf("Expected size 0, got %d", heap.Size())
	}
}

// TestNewIntMinHeap 测试整数最小堆创建
func TestNewIntMinHeap(t *testing.T) {
	heap := NewIntMinHeap()
	if !heap.IsEmpty() {
		t.Error("Expected heap to be empty")
	}
	if heap.Size() != 0 {
		t.Errorf("Expected size 0, got %d", heap.Size())
	}
}

// TestMaxHeapInsertAndExtract 测试最大堆插入和提取
func TestMaxHeapInsertAndExtract(t *testing.T) {
	heap := NewIntMaxHeap()

	// 插入元素
	elements := []int{3, 1, 4, 1, 5, 9, 2, 6}
	for _, elem := range elements {
		heap.Insert(elem)
	}

	if heap.Size() != len(elements) {
		t.Errorf("Expected size %d, got %d", len(elements), heap.Size())
	}

	// 验证堆性质
	if !heap.IsValidHeap() {
		t.Error("Heap property violated")
	}

	// 提取元素应该按降序排列
	expected := []int{9, 6, 5, 4, 3, 2, 1, 1}
	for i, expectedVal := range expected {
		top, err := heap.ExtractTop()
		if err != nil {
			t.Errorf("Unexpected error at index %d: %v", i, err)
		}
		if top.(int) != expectedVal {
			t.Errorf("Expected %d at index %d, got %d", expectedVal, i, top.(int))
		}
	}

	if !heap.IsEmpty() {
		t.Error("Expected heap to be empty after extracting all elements")
	}
}

// TestMinHeapInsertAndExtract 测试最小堆插入和提取
func TestMinHeapInsertAndExtract(t *testing.T) {
	heap := NewIntMinHeap()

	// 插入元素
	elements := []int{3, 1, 4, 1, 5, 9, 2, 6}
	for _, elem := range elements {
		heap.Insert(elem)
	}

	if heap.Size() != len(elements) {
		t.Errorf("Expected size %d, got %d", len(elements), heap.Size())
	}

	// 验证堆性质
	if !heap.IsValidHeap() {
		t.Error("Heap property violated")
	}

	// 提取元素应该按升序排列
	expected := []int{1, 1, 2, 3, 4, 5, 6, 9}
	for i, expectedVal := range expected {
		top, err := heap.ExtractTop()
		if err != nil {
			t.Errorf("Unexpected error at index %d: %v", i, err)
		}
		if top.(int) != expectedVal {
			t.Errorf("Expected %d at index %d, got %d", expectedVal, i, top.(int))
		}
	}

	if !heap.IsEmpty() {
		t.Error("Expected heap to be empty after extracting all elements")
	}
}

// TestPeek 测试查看堆顶元素
func TestPeek(t *testing.T) {
	heap := NewIntMaxHeap()

	// 空堆测试
	_, err := heap.Peek()
	if err == nil {
		t.Error("Expected error for empty heap")
	}

	// 插入元素并测试
	heap.Insert(5)
	heap.Insert(3)
	heap.Insert(8)
	heap.Insert(1)

	top, err := heap.Peek()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if top.(int) != 8 {
		t.Errorf("Expected 8, got %d", top.(int))
	}

	// 验证Peek不会改变堆的大小
	if heap.Size() != 4 {
		t.Errorf("Expected size 4, got %d", heap.Size())
	}
}

// TestBuildHeap 测试从数组构建堆
func TestBuildHeap(t *testing.T) {
	heap := NewIntMinHeap()

	elements := []interface{}{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}
	heap.BuildHeap(elements)

	if heap.Size() != len(elements) {
		t.Errorf("Expected size %d, got %d", len(elements), heap.Size())
	}

	// 验证堆性质
	if !heap.IsValidHeap() {
		t.Error("Heap property violated after BuildHeap")
	}

	// 验证堆顶是最小元素
	top, err := heap.Peek()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if top.(int) != 1 {
		t.Errorf("Expected 1, got %d", top.(int))
	}
}

// TestContains 测试包含检查
func TestContains(t *testing.T) {
	heap := NewIntMinHeap()

	elements := []int{3, 1, 4, 1, 5, 9, 2, 6}
	for _, elem := range elements {
		heap.Insert(elem)
	}

	// 测试存在的元素
	for _, elem := range elements {
		if !heap.Contains(elem) {
			t.Errorf("Expected heap to contain %d", elem)
		}
	}

	// 测试不存在的元素
	if heap.Contains(10) {
		t.Error("Expected heap not to contain 10")
	}
}

// TestRemove 测试删除元素
func TestRemove(t *testing.T) {
	heap := NewIntMinHeap()

	elements := []int{3, 1, 4, 1, 5, 9, 2, 6}
	for _, elem := range elements {
		heap.Insert(elem)
	}

	initialSize := heap.Size()

	// 删除存在的元素
	err := heap.Remove(4)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if heap.Size() != initialSize-1 {
		t.Errorf("Expected size %d, got %d", initialSize-1, heap.Size())
	}
	if heap.Contains(4) {
		t.Error("Expected element 4 to be removed")
	}

	// 验证堆性质
	if !heap.IsValidHeap() {
		t.Error("Heap property violated after removal")
	}

	// 删除不存在的元素
	err = heap.Remove(10)
	if err == nil {
		t.Error("Expected error for removing non-existent element")
	}
}

// TestClear 测试清空堆
func TestClear(t *testing.T) {
	heap := NewIntMinHeap()

	elements := []int{3, 1, 4, 1, 5, 9, 2, 6}
	for _, elem := range elements {
		heap.Insert(elem)
	}

	heap.Clear()

	if !heap.IsEmpty() {
		t.Error("Expected heap to be empty after clear")
	}
	if heap.Size() != 0 {
		t.Errorf("Expected size 0, got %d", heap.Size())
	}
}

// TestToSlice 测试转换为切片
func TestToSlice(t *testing.T) {
	heap := NewIntMinHeap()

	elements := []int{3, 1, 4, 1, 5}
	for _, elem := range elements {
		heap.Insert(elem)
	}

	slice := heap.ToSlice()
	if len(slice) != len(elements) {
		t.Errorf("Expected slice length %d, got %d", len(elements), len(slice))
	}

	// 验证原堆未被修改
	if heap.Size() != len(elements) {
		t.Errorf("Expected heap size %d, got %d", len(elements), heap.Size())
	}
}

// TestToSortedSlice 测试转换为有序切片
func TestToSortedSlice(t *testing.T) {
	heap := NewIntMinHeap()

	elements := []int{3, 1, 4, 1, 5, 9, 2, 6}
	for _, elem := range elements {
		heap.Insert(elem)
	}

	sortedSlice := heap.ToSortedSlice()
	expected := []int{1, 1, 2, 3, 4, 5, 6, 9}

	if len(sortedSlice) != len(expected) {
		t.Errorf("Expected slice length %d, got %d", len(expected), len(sortedSlice))
	}

	for i, expectedVal := range expected {
		if sortedSlice[i].(int) != expectedVal {
			t.Errorf("Expected %d at index %d, got %d", expectedVal, i, sortedSlice[i].(int))
		}
	}

	// 验证原堆已被清空
	if !heap.IsEmpty() {
		t.Error("Expected heap to be empty after ToSortedSlice")
	}
}

// TestString 测试字符串表示
func TestString(t *testing.T) {
	heap := NewIntMinHeap()

	// 空堆
	str := heap.String()
	if str != "Heap[]" {
		t.Errorf("Expected 'Heap[]', got '%s'", str)
	}

	// 非空堆
	heap.Insert(1)
	heap.Insert(2)
	str = heap.String()
	if str != "Heap[1, 2]" {
		t.Errorf("Expected 'Heap[1, 2]', got '%s'", str)
	}
}

// TestPriorityQueue 测试优先级队列
func TestPriorityQueue(t *testing.T) {
	// 创建优先级队列（数字越小优先级越高）
	pq := NewPriorityQueue(func(a, b interface{}) int {
		ia, ib := a.(int), b.(int)
		if ia < ib {
			return -1
		} else if ia > ib {
			return 1
		}
		return 0
	})

	// 入队
	elements := []int{3, 1, 4, 1, 5, 9, 2, 6}
	for _, elem := range elements {
		pq.Enqueue(elem)
	}

	if pq.Size() != len(elements) {
		t.Errorf("Expected size %d, got %d", len(elements), pq.Size())
	}

	// 查看队首
	front, err := pq.Front()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if front.(int) != 1 {
		t.Errorf("Expected 1, got %d", front.(int))
	}

	// 出队应该按优先级顺序
	expected := []int{1, 1, 2, 3, 4, 5, 6, 9}
	for i, expectedVal := range expected {
		elem, err := pq.Dequeue()
		if err != nil {
			t.Errorf("Unexpected error at index %d: %v", i, err)
		}
		if elem.(int) != expectedVal {
			t.Errorf("Expected %d at index %d, got %d", expectedVal, i, elem.(int))
		}
	}

	if !pq.IsEmpty() {
		t.Error("Expected priority queue to be empty")
	}
}

// TestHeapSort 测试堆排序
func TestHeapSort(t *testing.T) {
	arr := []interface{}{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	expected := []interface{}{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9}

	HeapSort(arr, func(a, b interface{}) int {
		ia, ib := a.(int), b.(int)
		if ia < ib {
			return -1
		} else if ia > ib {
			return 1
		}
		return 0
	})

	if !reflect.DeepEqual(arr, expected) {
		t.Errorf("Expected %v, got %v", expected, arr)
	}
}

// TestFindKthLargest 测试查找第K大元素
func TestFindKthLargest(t *testing.T) {
	arr := []interface{}{3, 2, 1, 5, 6, 4}

	// 测试第2大元素
	kthLargest, err := FindKthLargest(arr, 2, func(a, b interface{}) int {
		ia, ib := a.(int), b.(int)
		if ia < ib {
			return -1
		} else if ia > ib {
			return 1
		}
		return 0
	})

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if kthLargest.(int) != 5 {
		t.Errorf("Expected 5, got %d", kthLargest.(int))
	}

	// 测试无效的k
	_, err = FindKthLargest(arr, 0, func(a, b interface{}) int { return 0 })
	if err == nil {
		t.Error("Expected error for k=0")
	}

	_, err = FindKthLargest(arr, 7, func(a, b interface{}) int { return 0 })
	if err == nil {
		t.Error("Expected error for k > array length")
	}
}

// TestMergeKSortedArrays 测试合并K个有序数组
func TestMergeKSortedArrays(t *testing.T) {
	arrays := [][]interface{}{
		{1, 4, 5},
		{1, 3, 4},
		{2, 6},
	}

	result := MergeKSortedArrays(arrays, func(a, b interface{}) int {
		ia, ib := a.(int), b.(int)
		if ia < ib {
			return -1
		} else if ia > ib {
			return 1
		}
		return 0
	})

	expected := []interface{}{1, 1, 2, 3, 4, 4, 5, 6}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	// 测试空数组
	emptyResult := MergeKSortedArrays([][]interface{}{}, func(a, b interface{}) int { return 0 })
	if len(emptyResult) != 0 {
		t.Errorf("Expected empty result, got %v", emptyResult)
	}
}

// TestEmptyHeapOperations 测试空堆操作
func TestEmptyHeapOperations(t *testing.T) {
	heap := NewIntMinHeap()

	// 测试空堆的ExtractTop
	_, err := heap.ExtractTop()
	if err == nil {
		t.Error("Expected error for ExtractTop on empty heap")
	}

	// 测试空堆的Peek
	_, err = heap.Peek()
	if err == nil {
		t.Error("Expected error for Peek on empty heap")
	}

	// 测试空堆的Remove
	err = heap.Remove(1)
	if err == nil {
		t.Error("Expected error for Remove on empty heap")
	}
}

// TestCustomComparator 测试自定义比较器
func TestCustomComparator(t *testing.T) {
	// 创建字符串长度比较的最小堆
	heap := NewMinHeap(func(a, b interface{}) int {
		sa, sb := a.(string), b.(string)
		if len(sa) < len(sb) {
			return -1
		} else if len(sa) > len(sb) {
			return 1
		}
		return 0
	})

	strings := []string{"hello", "hi", "world", "a", "test"}
	for _, str := range strings {
		heap.Insert(str)
	}

	// 应该按字符串长度升序提取
	expectedLengths := []int{1, 2, 4, 5, 5} // "a", "hi", "test", "hello", "world"
	for i, expectedLen := range expectedLengths {
		top, err := heap.ExtractTop()
		if err != nil {
			t.Errorf("Unexpected error at index %d: %v", i, err)
		}
		if len(top.(string)) != expectedLen {
			t.Errorf("Expected string length %d at index %d, got %d (%s)", 
				expectedLen, i, len(top.(string)), top.(string))
		}
	}
}

// BenchmarkHeapInsert 基准测试：插入操作
func BenchmarkHeapInsert(b *testing.B) {
	heap := NewIntMinHeap()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		heap.Insert(i)
	}
}

// BenchmarkHeapExtractTop 基准测试：提取堆顶
func BenchmarkHeapExtractTop(b *testing.B) {
	heap := NewIntMinHeap()

	// 预先插入数据
	for i := 0; i < b.N; i++ {
		heap.Insert(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		heap.ExtractTop()
	}
}

// BenchmarkHeapBuildHeap 基准测试：构建堆
func BenchmarkHeapBuildHeap(b *testing.B) {
	elements := make([]interface{}, 10000)
	for i := 0; i < 10000; i++ {
		elements[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		heap := NewIntMinHeap()
		heap.BuildHeap(elements)
	}
}

// BenchmarkHeapSort 基准测试：堆排序
func BenchmarkHeapSort(b *testing.B) {
	original := make([]interface{}, 1000)
	for i := 0; i < 1000; i++ {
		original[i] = 1000 - i // 逆序数组
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		arr := make([]interface{}, len(original))
		copy(arr, original)
		HeapSort(arr, func(a, b interface{}) int {
			ia, ib := a.(int), b.(int)
			if ia < ib {
				return -1
			} else if ia > ib {
				return 1
			}
			return 0
		})
	}
}

// BenchmarkPriorityQueue 基准测试：优先级队列
func BenchmarkPriorityQueue(b *testing.B) {
	pq := NewPriorityQueue(func(a, b interface{}) int {
		ia, ib := a.(int), b.(int)
		if ia < ib {
			return -1
		} else if ia > ib {
			return 1
		}
		return 0
	})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pq.Enqueue(i)
		if i%2 == 0 && !pq.IsEmpty() {
			pq.Dequeue()
		}
	}
}