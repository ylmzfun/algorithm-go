package heap

import (
	"errors"
	"fmt"
	"strings"
)

// Heap 堆数据结构实现
// 思路：使用数组表示完全二叉树，通过比较函数支持最大堆和最小堆
// 作用：高效地维护集合中的最大值或最小值，支持O(log n)的插入和删除
// 业务场景：
// 1. 优先级队列：任务调度，事件处理
// 2. 排序算法：堆排序
// 3. 图算法：Dijkstra最短路径，Prim最小生成树
// 4. 数据流处理：Top K问题，中位数维护
// 5. 内存管理：垃圾回收器的优先级管理
// 6. 操作系统：进程调度
// 7. 搜索算法：A*算法的开放列表
type Heap struct {
	data    []interface{}           // 存储堆元素的数组
	size    int                     // 当前堆的大小
	compare func(a, b interface{}) int // 比较函数：返回负数表示a<b，0表示a==b，正数表示a>b
}

// NewMaxHeap 创建最大堆
func NewMaxHeap(compare func(a, b interface{}) int) *Heap {
	return &Heap{
		data:    make([]interface{}, 0),
		size:    0,
		compare: func(a, b interface{}) int { return -compare(a, b) }, // 反转比较函数实现最大堆
	}
}

// NewMinHeap 创建最小堆
func NewMinHeap(compare func(a, b interface{}) int) *Heap {
	return &Heap{
		data:    make([]interface{}, 0),
		size:    0,
		compare: compare,
	}
}

// NewIntMaxHeap 创建整数最大堆
func NewIntMaxHeap() *Heap {
	return NewMaxHeap(func(a, b interface{}) int {
		ia, ib := a.(int), b.(int)
		if ia < ib {
			return -1
		} else if ia > ib {
			return 1
		}
		return 0
	})
}

// NewIntMinHeap 创建整数最小堆
func NewIntMinHeap() *Heap {
	return NewMinHeap(func(a, b interface{}) int {
		ia, ib := a.(int), b.(int)
		if ia < ib {
			return -1
		} else if ia > ib {
			return 1
		}
		return 0
	})
}

// parent 获取父节点索引
func (h *Heap) parent(index int) int {
	return (index - 1) / 2
}

// leftChild 获取左子节点索引
func (h *Heap) leftChild(index int) int {
	return 2*index + 1
}

// rightChild 获取右子节点索引
func (h *Heap) rightChild(index int) int {
	return 2*index + 2
}

// hasParent 检查是否有父节点
func (h *Heap) hasParent(index int) bool {
	return h.parent(index) >= 0
}

// hasLeftChild 检查是否有左子节点
func (h *Heap) hasLeftChild(index int) bool {
	return h.leftChild(index) < h.size
}

// hasRightChild 检查是否有右子节点
func (h *Heap) hasRightChild(index int) bool {
	return h.rightChild(index) < h.size
}

// swap 交换两个元素
func (h *Heap) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

// Insert 插入元素
// 时间复杂度：O(log n)
func (h *Heap) Insert(element interface{}) {
	// 在数组末尾添加元素
	h.data = append(h.data, element)
	h.size++
	
	// 向上调整维护堆性质
	h.heapifyUp(h.size - 1)
}

// heapifyUp 向上调整堆
func (h *Heap) heapifyUp(index int) {
	// 如果有父节点且当前节点比父节点优先级高，则交换
	for h.hasParent(index) && h.compare(h.data[index], h.data[h.parent(index)]) < 0 {
		h.swap(index, h.parent(index))
		index = h.parent(index)
	}
}

// ExtractTop 提取堆顶元素（最大堆的最大值或最小堆的最小值）
// 时间复杂度：O(log n)
func (h *Heap) ExtractTop() (interface{}, error) {
	if h.IsEmpty() {
		return nil, errors.New("heap is empty")
	}
	
	// 保存堆顶元素
	top := h.data[0]
	
	// 将最后一个元素移到堆顶
	h.data[0] = h.data[h.size-1]
	h.size--
	h.data = h.data[:h.size]
	
	// 向下调整维护堆性质
	if h.size > 0 {
		h.heapifyDown(0)
	}
	
	return top, nil
}

// heapifyDown 向下调整堆
func (h *Heap) heapifyDown(index int) {
	for h.hasLeftChild(index) {
		// 找到优先级最高的子节点
		smallerChildIndex := h.leftChild(index)
		if h.hasRightChild(index) && h.compare(h.data[h.rightChild(index)], h.data[h.leftChild(index)]) < 0 {
			smallerChildIndex = h.rightChild(index)
		}
		
		// 如果当前节点优先级已经比子节点高，则停止
		if h.compare(h.data[index], h.data[smallerChildIndex]) <= 0 {
			break
		}
		
		// 交换并继续向下调整
		h.swap(index, smallerChildIndex)
		index = smallerChildIndex
	}
}

// Peek 查看堆顶元素但不删除
// 时间复杂度：O(1)
func (h *Heap) Peek() (interface{}, error) {
	if h.IsEmpty() {
		return nil, errors.New("heap is empty")
	}
	return h.data[0], nil
}

// Size 返回堆的大小
func (h *Heap) Size() int {
	return h.size
}

// IsEmpty 判断堆是否为空
func (h *Heap) IsEmpty() bool {
	return h.size == 0
}

// Clear 清空堆
func (h *Heap) Clear() {
	h.data = h.data[:0]
	h.size = 0
}

// ToSlice 转换为切片（不保证顺序）
func (h *Heap) ToSlice() []interface{} {
	result := make([]interface{}, h.size)
	copy(result, h.data[:h.size])
	return result
}

// ToSortedSlice 转换为有序切片（会清空原堆）
func (h *Heap) ToSortedSlice() []interface{} {
	result := make([]interface{}, 0, h.size)
	
	for !h.IsEmpty() {
		top, _ := h.ExtractTop()
		result = append(result, top)
	}
	
	return result
}

// BuildHeap 从切片构建堆
// 时间复杂度：O(n)
func (h *Heap) BuildHeap(elements []interface{}) {
	h.data = make([]interface{}, len(elements))
	copy(h.data, elements)
	h.size = len(elements)
	
	// 从最后一个非叶子节点开始向下调整
	for i := h.parent(h.size - 1); i >= 0; i-- {
		h.heapifyDown(i)
	}
}

// Contains 检查是否包含指定元素
// 时间复杂度：O(n)
func (h *Heap) Contains(element interface{}) bool {
	for i := 0; i < h.size; i++ {
		if h.compare(h.data[i], element) == 0 {
			return true
		}
	}
	return false
}

// Remove 删除指定元素的第一个出现
// 时间复杂度：O(n)
func (h *Heap) Remove(element interface{}) error {
	// 查找元素
	index := -1
	for i := 0; i < h.size; i++ {
		if h.compare(h.data[i], element) == 0 {
			index = i
			break
		}
	}
	
	if index == -1 {
		return errors.New("element not found")
	}
	
	// 将最后一个元素移到要删除的位置
	h.data[index] = h.data[h.size-1]
	h.size--
	h.data = h.data[:h.size]
	
	// 调整堆性质
	if index < h.size {
		// 先尝试向上调整
		if h.hasParent(index) && h.compare(h.data[index], h.data[h.parent(index)]) < 0 {
			h.heapifyUp(index)
		} else {
			// 再尝试向下调整
			h.heapifyDown(index)
		}
	}
	
	return nil
}

// String 字符串表示
func (h *Heap) String() string {
	if h.IsEmpty() {
		return "Heap[]"
	}
	
	elements := make([]string, h.size)
	for i := 0; i < h.size; i++ {
		elements[i] = fmt.Sprintf("%v", h.data[i])
	}
	
	return fmt.Sprintf("Heap[%s]", strings.Join(elements, ", "))
}

// PrintTree 打印堆的树形结构
func (h *Heap) PrintTree() {
	if h.IsEmpty() {
		fmt.Println("Empty heap")
		return
	}
	
	h.printTreeHelper(0, 0)
}

// printTreeHelper 递归打印树形结构
func (h *Heap) printTreeHelper(index, depth int) {
	if index >= h.size {
		return
	}
	
	// 打印右子树
	if h.hasRightChild(index) {
		h.printTreeHelper(h.rightChild(index), depth+1)
	}
	
	// 打印当前节点
	for i := 0; i < depth; i++ {
		fmt.Print("    ")
	}
	fmt.Printf("%v\n", h.data[index])
	
	// 打印左子树
	if h.hasLeftChild(index) {
		h.printTreeHelper(h.leftChild(index), depth+1)
	}
}

// IsValidHeap 验证是否为有效的堆
func (h *Heap) IsValidHeap() bool {
	return h.isValidHeapHelper(0)
}

// isValidHeapHelper 递归验证堆性质
func (h *Heap) isValidHeapHelper(index int) bool {
	if index >= h.size {
		return true
	}
	
	// 检查左子节点
	if h.hasLeftChild(index) {
		if h.compare(h.data[index], h.data[h.leftChild(index)]) > 0 {
			return false
		}
		if !h.isValidHeapHelper(h.leftChild(index)) {
			return false
		}
	}
	
	// 检查右子节点
	if h.hasRightChild(index) {
		if h.compare(h.data[index], h.data[h.rightChild(index)]) > 0 {
			return false
		}
		if !h.isValidHeapHelper(h.rightChild(index)) {
			return false
		}
	}
	
	return true
}

// PriorityQueue 优先级队列（基于堆实现）
type PriorityQueue struct {
	heap *Heap
}

// NewPriorityQueue 创建优先级队列
func NewPriorityQueue(compare func(a, b interface{}) int) *PriorityQueue {
	return &PriorityQueue{
		heap: NewMinHeap(compare), // 使用最小堆，优先级高的元素在堆顶
	}
}

// Enqueue 入队
func (pq *PriorityQueue) Enqueue(element interface{}) {
	pq.heap.Insert(element)
}

// Dequeue 出队
func (pq *PriorityQueue) Dequeue() (interface{}, error) {
	return pq.heap.ExtractTop()
}

// Front 查看队首元素
func (pq *PriorityQueue) Front() (interface{}, error) {
	return pq.heap.Peek()
}

// Size 返回队列大小
func (pq *PriorityQueue) Size() int {
	return pq.heap.Size()
}

// IsEmpty 判断是否为空
func (pq *PriorityQueue) IsEmpty() bool {
	return pq.heap.IsEmpty()
}

// Clear 清空队列
func (pq *PriorityQueue) Clear() {
	pq.heap.Clear()
}

// String 字符串表示
func (pq *PriorityQueue) String() string {
	return fmt.Sprintf("PriorityQueue%s", pq.heap.String()[4:]) // 去掉"Heap"前缀
}

// HeapSort 堆排序函数
// 时间复杂度：O(n log n)
// 空间复杂度：O(1)
func HeapSort(arr []interface{}, compare func(a, b interface{}) int) {
	n := len(arr)
	if n <= 1 {
		return
	}
	
	// 构建最大堆
	heap := NewMaxHeap(compare)
	heap.data = arr
	heap.size = n
	
	// 从最后一个非叶子节点开始向下调整
	for i := heap.parent(n - 1); i >= 0; i-- {
		heap.heapifyDown(i)
	}
	
	// 依次将堆顶元素与末尾元素交换，并调整堆
	for i := n - 1; i > 0; i-- {
		heap.swap(0, i)
		heap.size--
		heap.heapifyDown(0)
	}
}

// FindKthLargest 查找第K大的元素
// 时间复杂度：O(n log k)
func FindKthLargest(arr []interface{}, k int, compare func(a, b interface{}) int) (interface{}, error) {
	if k <= 0 || k > len(arr) {
		return nil, errors.New("invalid k")
	}
	
	// 使用最小堆维护前K大的元素
	heap := NewMinHeap(compare)
	
	for _, element := range arr {
		if heap.Size() < k {
			heap.Insert(element)
		} else {
			top, _ := heap.Peek()
			if compare(element, top) > 0 {
				heap.ExtractTop()
				heap.Insert(element)
			}
		}
	}
	
	return heap.Peek()
}

// MergeKSortedArrays 合并K个有序数组
// 时间复杂度：O(n log k)，其中n是所有元素的总数
func MergeKSortedArrays(arrays [][]interface{}, compare func(a, b interface{}) int) []interface{} {
	if len(arrays) == 0 {
		return []interface{}{}
	}
	
	// 定义堆元素结构
	type heapElement struct {
		value      interface{}
		arrayIndex int
		elementIndex int
	}
	
	// 创建最小堆
	heap := NewMinHeap(func(a, b interface{}) int {
		ea, eb := a.(heapElement), b.(heapElement)
		return compare(ea.value, eb.value)
	})
	
	// 将每个数组的第一个元素加入堆
	for i, arr := range arrays {
		if len(arr) > 0 {
			heap.Insert(heapElement{
				value:        arr[0],
				arrayIndex:   i,
				elementIndex: 0,
			})
		}
	}
	
	result := make([]interface{}, 0)
	
	// 依次取出堆顶元素，并将对应数组的下一个元素加入堆
	for !heap.IsEmpty() {
		top, _ := heap.ExtractTop()
		element := top.(heapElement)
		
		result = append(result, element.value)
		
		// 如果对应数组还有下一个元素，加入堆
		if element.elementIndex+1 < len(arrays[element.arrayIndex]) {
			heap.Insert(heapElement{
				value:        arrays[element.arrayIndex][element.elementIndex+1],
				arrayIndex:   element.arrayIndex,
				elementIndex: element.elementIndex + 1,
			})
		}
	}
	
	return result
}

// 业务应用示例：
// 1. 任务调度系统：根据优先级调度任务
// 2. 事件处理：按时间戳处理事件
// 3. 路径规划：Dijkstra算法中的最短路径
// 4. 数据流处理：实时维护Top K热门商品
// 5. 内存管理：LRU缓存的实现
// 6. 游戏开发：AI决策树的评估
// 7. 网络协议：TCP拥塞控制
// 8. 数据库：查询优化器的成本估算