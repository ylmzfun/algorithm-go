package queue

import (
	"errors"
	"fmt"
	"strings"
)

// Queue 队列的实现（基于切片的循环队列）
// 思路：使用切片作为底层存储，通过front和rear指针实现循环队列，遵循FIFO（先进先出）原则
// 作用：提供先进先出的数据访问模式，支持高效的入队和出队操作
// 业务场景：
// 1. 任务调度：操作系统的进程调度，打印队列
// 2. 广度优先搜索（BFS）：图和树的层序遍历
// 3. 缓冲区：网络数据包处理，IO缓冲
// 4. 消息队列：异步消息处理系统
// 5. 客服系统：客户排队等待服务
// 6. 游戏开发：事件处理队列
type Queue struct {
	data     []interface{} // 存储队列元素的切片
	front    int           // 队头指针
	rear     int           // 队尾指针
	size     int           // 当前元素数量
	capacity int           // 队列容量
}

// NewQueue 创建新的队列
// initialCapacity: 初始容量，如果为0则使用默认值10
func NewQueue(initialCapacity int) *Queue {
	if initialCapacity <= 0 {
		initialCapacity = 10
	}
	return &Queue{
		data:     make([]interface{}, initialCapacity),
		front:    0,
		rear:     0,
		size:     0,
		capacity: initialCapacity,
	}
}

// Enqueue 入队（在队尾添加元素）
// 时间复杂度：平均O(1)，最坏O(n)（需要扩容时）
func (q *Queue) Enqueue(element interface{}) {
	// 检查是否需要扩容
	if q.size >= q.capacity {
		q.resize()
	}
	
	q.data[q.rear] = element
	q.rear = (q.rear + 1) % q.capacity
	q.size++
}

// Dequeue 出队（从队头移除元素）
// 时间复杂度：O(1)
func (q *Queue) Dequeue() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}
	
	element := q.data[q.front]
	q.data[q.front] = nil // 避免内存泄漏
	q.front = (q.front + 1) % q.capacity
	q.size--
	
	return element, nil
}

// Front 查看队头元素但不移除
// 时间复杂度：O(1)
func (q *Queue) Front() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}
	
	return q.data[q.front], nil
}

// Rear 查看队尾元素但不移除
// 时间复杂度：O(1)
func (q *Queue) Rear() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}
	
	// 计算队尾元素的实际位置
	rearIndex := (q.rear - 1 + q.capacity) % q.capacity
	return q.data[rearIndex], nil
}

// IsEmpty 判断队列是否为空
// 时间复杂度：O(1)
func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

// IsFull 判断队列是否已满
// 时间复杂度：O(1)
func (q *Queue) IsFull() bool {
	return q.size == q.capacity
}

// Size 返回队列的大小
// 时间复杂度：O(1)
func (q *Queue) Size() int {
	return q.size
}

// Capacity 返回当前容量
func (q *Queue) Capacity() int {
	return q.capacity
}

// Clear 清空队列
// 时间复杂度：O(n)
func (q *Queue) Clear() {
	for i := 0; i < q.capacity; i++ {
		q.data[i] = nil
	}
	q.front = 0
	q.rear = 0
	q.size = 0
}

// Contains 检查队列是否包含指定元素
// 时间复杂度：O(n)
func (q *Queue) Contains(element interface{}) bool {
	for i := 0; i < q.size; i++ {
		index := (q.front + i) % q.capacity
		if q.data[index] == element {
			return true
		}
	}
	return false
}

// ToSlice 转换为切片（从队头到队尾）
// 时间复杂度：O(n)
func (q *Queue) ToSlice() []interface{} {
	result := make([]interface{}, q.size)
	for i := 0; i < q.size; i++ {
		index := (q.front + i) % q.capacity
		result[i] = q.data[index]
	}
	return result
}

// String 字符串表示
func (q *Queue) String() string {
	if q.IsEmpty() {
		return "Queue{}"
	}
	
	var elements []string
	for i := 0; i < q.size; i++ {
		index := (q.front + i) % q.capacity
		elements = append(elements, fmt.Sprintf("%v", q.data[index]))
	}
	
	return fmt.Sprintf("Queue{size: %d, elements: [%s] front->rear}", 
		q.size, strings.Join(elements, ", "))
}

// resize 扩容，容量翻倍
func (q *Queue) resize() {
	newCapacity := q.capacity * 2
	newData := make([]interface{}, newCapacity)
	
	// 将现有元素复制到新数组中，保持顺序
	for i := 0; i < q.size; i++ {
		index := (q.front + i) % q.capacity
		newData[i] = q.data[index]
	}
	
	q.data = newData
	q.front = 0
	q.rear = q.size
	q.capacity = newCapacity
}

// QueueWithLinkedList 基于链表实现的队列
// 优点：动态大小，不需要预分配内存，不会出现"假满"情况
// 缺点：每个元素需要额外的指针存储空间
type QueueWithLinkedList struct {
	front *QueueNode // 队头指针
	rear  *QueueNode // 队尾指针
	size  int        // 队列大小
}

// QueueNode 队列节点
type QueueNode struct {
	Data interface{} // 节点数据
	Next *QueueNode  // 指向下一个节点
}

// NewQueueWithLinkedList 创建基于链表的队列
func NewQueueWithLinkedList() *QueueWithLinkedList {
	return &QueueWithLinkedList{
		front: nil,
		rear:  nil,
		size:  0,
	}
}

// Enqueue 入队
// 时间复杂度：O(1)
func (q *QueueWithLinkedList) Enqueue(element interface{}) {
	newNode := &QueueNode{
		Data: element,
		Next: nil,
	}
	
	if q.rear == nil {
		// 空队列
		q.front = newNode
		q.rear = newNode
	} else {
		// 非空队列
		q.rear.Next = newNode
		q.rear = newNode
	}
	
	q.size++
}

// Dequeue 出队
// 时间复杂度：O(1)
func (q *QueueWithLinkedList) Dequeue() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}
	
	element := q.front.Data
	q.front = q.front.Next
	
	// 如果队列变空，更新rear指针
	if q.front == nil {
		q.rear = nil
	}
	
	q.size--
	return element, nil
}

// Front 查看队头元素
// 时间复杂度：O(1)
func (q *QueueWithLinkedList) Front() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}
	
	return q.front.Data, nil
}

// Rear 查看队尾元素
// 时间复杂度：O(1)
func (q *QueueWithLinkedList) Rear() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}
	
	return q.rear.Data, nil
}

// IsEmpty 判断是否为空
func (q *QueueWithLinkedList) IsEmpty() bool {
	return q.front == nil
}

// Size 返回大小
func (q *QueueWithLinkedList) Size() int {
	return q.size
}

// Clear 清空队列
func (q *QueueWithLinkedList) Clear() {
	q.front = nil
	q.rear = nil
	q.size = 0
}

// String 字符串表示
func (q *QueueWithLinkedList) String() string {
	if q.IsEmpty() {
		return "QueueWithLinkedList{}"
	}
	
	var elements []string
	current := q.front
	
	for current != nil {
		elements = append(elements, fmt.Sprintf("%v", current.Data))
		current = current.Next
	}
	
	return fmt.Sprintf("QueueWithLinkedList{size: %d, elements: [%s] front->rear}", 
		q.size, strings.Join(elements, ", "))
}

// Deque 双端队列（Double-ended Queue）
// 支持在两端进行插入和删除操作
type Deque struct {
	data     []interface{} // 存储数据的切片
	front    int           // 队头指针
	rear     int           // 队尾指针
	size     int           // 当前元素数量
	capacity int           // 容量
}

// NewDeque 创建新的双端队列
func NewDeque(initialCapacity int) *Deque {
	if initialCapacity <= 0 {
		initialCapacity = 10
	}
	return &Deque{
		data:     make([]interface{}, initialCapacity),
		front:    0,
		rear:     0,
		size:     0,
		capacity: initialCapacity,
	}
}

// AddFront 在队头添加元素
// 时间复杂度：平均O(1)
func (dq *Deque) AddFront(element interface{}) {
	if dq.size >= dq.capacity {
		dq.resize()
	}
	
	dq.front = (dq.front - 1 + dq.capacity) % dq.capacity
	dq.data[dq.front] = element
	dq.size++
}

// AddRear 在队尾添加元素
// 时间复杂度：平均O(1)
func (dq *Deque) AddRear(element interface{}) {
	if dq.size >= dq.capacity {
		dq.resize()
	}
	
	dq.data[dq.rear] = element
	dq.rear = (dq.rear + 1) % dq.capacity
	dq.size++
}

// RemoveFront 从队头移除元素
// 时间复杂度：O(1)
func (dq *Deque) RemoveFront() (interface{}, error) {
	if dq.IsEmpty() {
		return nil, errors.New("deque is empty")
	}
	
	element := dq.data[dq.front]
	dq.data[dq.front] = nil
	dq.front = (dq.front + 1) % dq.capacity
	dq.size--
	
	return element, nil
}

// RemoveRear 从队尾移除元素
// 时间复杂度：O(1)
func (dq *Deque) RemoveRear() (interface{}, error) {
	if dq.IsEmpty() {
		return nil, errors.New("deque is empty")
	}
	
	dq.rear = (dq.rear - 1 + dq.capacity) % dq.capacity
	element := dq.data[dq.rear]
	dq.data[dq.rear] = nil
	dq.size--
	
	return element, nil
}

// Front 查看队头元素
func (dq *Deque) Front() (interface{}, error) {
	if dq.IsEmpty() {
		return nil, errors.New("deque is empty")
	}
	return dq.data[dq.front], nil
}

// Rear 查看队尾元素
func (dq *Deque) Rear() (interface{}, error) {
	if dq.IsEmpty() {
		return nil, errors.New("deque is empty")
	}
	rearIndex := (dq.rear - 1 + dq.capacity) % dq.capacity
	return dq.data[rearIndex], nil
}

// IsEmpty 判断是否为空
func (dq *Deque) IsEmpty() bool {
	return dq.size == 0
}

// Size 返回大小
func (dq *Deque) Size() int {
	return dq.size
}

// Clear 清空双端队列
func (dq *Deque) Clear() {
	for i := 0; i < dq.capacity; i++ {
		dq.data[i] = nil
	}
	dq.front = 0
	dq.rear = 0
	dq.size = 0
}

// resize 扩容
func (dq *Deque) resize() {
	newCapacity := dq.capacity * 2
	newData := make([]interface{}, newCapacity)
	
	for i := 0; i < dq.size; i++ {
		index := (dq.front + i) % dq.capacity
		newData[i] = dq.data[index]
	}
	
	dq.data = newData
	dq.front = 0
	dq.rear = dq.size
	dq.capacity = newCapacity
}

// String 字符串表示
func (dq *Deque) String() string {
	if dq.IsEmpty() {
		return "Deque{}"
	}
	
	var elements []string
	for i := 0; i < dq.size; i++ {
		index := (dq.front + i) % dq.capacity
		elements = append(elements, fmt.Sprintf("%v", dq.data[index]))
	}
	
	return fmt.Sprintf("Deque{size: %d, elements: [%s] front<->rear}", 
		dq.size, strings.Join(elements, ", "))
}

// 业务应用示例：
// 1. 操作系统：进程调度队列，就绪队列
// 2. 网络编程：数据包缓冲区，请求队列
// 3. 游戏开发：事件队列，动画队列
// 4. 消息系统：消息队列，任务队列
// 5. 算法：广度优先搜索，层序遍历
// 6. 缓存系统：LRU缓存的实现
// 7. 打印系统：打印任务队列
// 8. 客服系统：客户等待队列