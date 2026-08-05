package queue

import (
	"testing"
)

func TestNewQueue(t *testing.T) {
	// 测试默认容量
	q1 := NewQueue(0)
	if q1.Capacity() != 10 {
		t.Errorf("Expected capacity 10, got %d", q1.Capacity())
	}
	
	// 测试指定容量
	q2 := NewQueue(20)
	if q2.Capacity() != 20 {
		t.Errorf("Expected capacity 20, got %d", q2.Capacity())
	}
	
	if q2.Size() != 0 {
		t.Errorf("Expected size 0, got %d", q2.Size())
	}
	
	if !q2.IsEmpty() {
		t.Error("New queue should be empty")
	}
	
	if q2.IsFull() {
		t.Error("New queue should not be full")
	}
}

func TestEnqueueDequeue(t *testing.T) {
	q := NewQueue(5)
	
	// 测试入队
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	
	if q.Size() != 3 {
		t.Errorf("Expected size 3, got %d", q.Size())
	}
	
	if q.IsEmpty() {
		t.Error("Queue should not be empty")
	}
	
	// 测试出队（FIFO顺序）
	val, err := q.Dequeue()
	if err != nil {
		t.Errorf("Dequeue failed: %v", err)
	}
	if val != 1 {
		t.Errorf("Expected 1, got %v", val)
	}
	
	val, err = q.Dequeue()
	if err != nil {
		t.Errorf("Dequeue failed: %v", err)
	}
	if val != 2 {
		t.Errorf("Expected 2, got %v", val)
	}
	
	if q.Size() != 1 {
		t.Errorf("Expected size 1, got %d", q.Size())
	}
	
	val, err = q.Dequeue()
	if err != nil {
		t.Errorf("Dequeue failed: %v", err)
	}
	if val != 3 {
		t.Errorf("Expected 3, got %v", val)
	}
	
	if !q.IsEmpty() {
		t.Error("Queue should be empty after dequeuing all elements")
	}
	
	// 测试空队列出队
	_, err = q.Dequeue()
	if err == nil {
		t.Error("Expected error when dequeuing from empty queue")
	}
}

func TestFrontRear(t *testing.T) {
	q := NewQueue(5)
	
	// 空队列查看
	_, err := q.Front()
	if err == nil {
		t.Error("Expected error when checking front of empty queue")
	}
	
	_, err = q.Rear()
	if err == nil {
		t.Error("Expected error when checking rear of empty queue")
	}
	
	// 入队后查看
	q.Enqueue("first")
	q.Enqueue("second")
	q.Enqueue("third")
	
	front, err := q.Front()
	if err != nil {
		t.Errorf("Front failed: %v", err)
	}
	if front != "first" {
		t.Errorf("Expected 'first', got %v", front)
	}
	
	rear, err := q.Rear()
	if err != nil {
		t.Errorf("Rear failed: %v", err)
	}
	if rear != "third" {
		t.Errorf("Expected 'third', got %v", rear)
	}
	
	// 确保Front和Rear不改变队列大小
	if q.Size() != 3 {
		t.Errorf("Expected size 3 after front/rear operations, got %d", q.Size())
	}
}

func TestCircularBehavior(t *testing.T) {
	q := NewQueue(3)
	
	// 填满队列
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	
	if !q.IsFull() {
		t.Error("Queue should be full")
	}
	
	// 出队一个元素
	val, _ := q.Dequeue()
	if val != 1 {
		t.Errorf("Expected 1, got %v", val)
	}
	
	// 再入队一个元素（测试循环特性）
	q.Enqueue(4)
	
	if q.Size() != 3 {
		t.Errorf("Expected size 3, got %d", q.Size())
	}
	
	// 验证顺序：2, 3, 4
	expected := []interface{}{2, 3, 4}
	for _, exp := range expected {
		val, _ := q.Dequeue()
		if val != exp {
			t.Errorf("Expected %v, got %v", exp, val)
		}
	}
}

func TestResize(t *testing.T) {
	q := NewQueue(2)
	
	// 填满初始容量
	q.Enqueue(1)
	q.Enqueue(2)
	
	if q.Capacity() != 2 {
		t.Errorf("Expected capacity 2, got %d", q.Capacity())
	}
	
	// 触发扩容
	q.Enqueue(3)
	
	if q.Capacity() != 4 {
		t.Errorf("Expected capacity 4 after resize, got %d", q.Capacity())
	}
	
	if q.Size() != 3 {
		t.Errorf("Expected size 3, got %d", q.Size())
	}
	
	// 验证元素顺序保持正确
	expected := []interface{}{1, 2, 3}
	for _, exp := range expected {
		val, _ := q.Dequeue()
		if val != exp {
			t.Errorf("Expected %v, got %v", exp, val)
		}
	}
}

func TestClear(t *testing.T) {
	q := NewQueue(5)
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	
	q.Clear()
	
	if q.Size() != 0 {
		t.Errorf("Expected size 0 after clear, got %d", q.Size())
	}
	
	if !q.IsEmpty() {
		t.Error("Queue should be empty after clear")
	}
	
	// 清空后应该能正常使用
	q.Enqueue("new")
	front, _ := q.Front()
	if front != "new" {
		t.Errorf("Expected 'new', got %v", front)
	}
}

func TestContains(t *testing.T) {
	q := NewQueue(5)
	q.Enqueue("apple")
	q.Enqueue("banana")
	q.Enqueue("cherry")
	
	if !q.Contains("banana") {
		t.Error("Expected to contain 'banana'")
	}
	
	if q.Contains("orange") {
		t.Error("Should not contain 'orange'")
	}
	
	// 出队元素后检查
	q.Dequeue() // 移除 "apple"
	if q.Contains("apple") {
		t.Error("Should not contain 'apple' after dequeuing")
	}
	
	if !q.Contains("banana") {
		t.Error("Should still contain 'banana'")
	}
}

func TestToSlice(t *testing.T) {
	q := NewQueue(5)
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	
	slice := q.ToSlice()
	expected := []interface{}{1, 2, 3}
	
	if len(slice) != 3 {
		t.Errorf("Expected slice length 3, got %d", len(slice))
	}
	
	for i, exp := range expected {
		if slice[i] != exp {
			t.Errorf("At index %d, expected %v, got %v", i, exp, slice[i])
		}
	}
	
	// 测试循环队列的ToSlice
	q.Dequeue() // 移除1
	q.Enqueue(4) // 添加4
	
	slice = q.ToSlice()
	expected = []interface{}{2, 3, 4}
	
	for i, exp := range expected {
		if slice[i] != exp {
			t.Errorf("At index %d, expected %v, got %v", i, exp, slice[i])
		}
	}
}

func TestString(t *testing.T) {
	q := NewQueue(5)
	
	// 空队列
	str := q.String()
	if str != "Queue{}" {
		t.Errorf("Expected empty queue string, got %s", str)
	}
	
	// 非空队列
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	
	str = q.String()
	if str == "" {
		t.Error("String representation should not be empty")
	}
	t.Logf("String representation: %s", str)
}

// 测试基于链表的队列
func TestQueueWithLinkedList(t *testing.T) {
	q := NewQueueWithLinkedList()
	
	if !q.IsEmpty() {
		t.Error("New queue should be empty")
	}
	
	if q.Size() != 0 {
		t.Errorf("Expected size 0, got %d", q.Size())
	}
	
	// 测试入队和出队
	q.Enqueue("first")
	q.Enqueue("second")
	q.Enqueue("third")
	
	if q.Size() != 3 {
		t.Errorf("Expected size 3, got %d", q.Size())
	}
	
	// 测试Front和Rear
	front, err := q.Front()
	if err != nil {
		t.Errorf("Front failed: %v", err)
	}
	if front != "first" {
		t.Errorf("Expected 'first', got %v", front)
	}
	
	rear, err := q.Rear()
	if err != nil {
		t.Errorf("Rear failed: %v", err)
	}
	if rear != "third" {
		t.Errorf("Expected 'third', got %v", rear)
	}
	
	// 测试Dequeue
	val, err := q.Dequeue()
	if err != nil {
		t.Errorf("Dequeue failed: %v", err)
	}
	if val != "first" {
		t.Errorf("Expected 'first', got %v", val)
	}
	
	if q.Size() != 2 {
		t.Errorf("Expected size 2 after dequeue, got %d", q.Size())
	}
	
	// 清空测试
	q.Clear()
	if !q.IsEmpty() {
		t.Error("Queue should be empty after clear")
	}
	
	// 空队列操作测试
	_, err = q.Dequeue()
	if err == nil {
		t.Error("Expected error when dequeuing from empty queue")
	}
	
	_, err = q.Front()
	if err == nil {
		t.Error("Expected error when checking front of empty queue")
	}
	
	_, err = q.Rear()
	if err == nil {
		t.Error("Expected error when checking rear of empty queue")
	}
}

// 测试双端队列
func TestDeque(t *testing.T) {
	dq := NewDeque(5)
	
	if !dq.IsEmpty() {
		t.Error("New deque should be empty")
	}
	
	// 测试在两端添加元素
	dq.AddRear(1)   // [1]
	dq.AddFront(2)  // [2, 1]
	dq.AddRear(3)   // [2, 1, 3]
	dq.AddFront(4)  // [4, 2, 1, 3]
	
	if dq.Size() != 4 {
		t.Errorf("Expected size 4, got %d", dq.Size())
	}
	
	// 测试查看两端元素
	front, err := dq.Front()
	if err != nil {
		t.Errorf("Front failed: %v", err)
	}
	if front != 4 {
		t.Errorf("Expected front 4, got %v", front)
	}
	
	rear, err := dq.Rear()
	if err != nil {
		t.Errorf("Rear failed: %v", err)
	}
	if rear != 3 {
		t.Errorf("Expected rear 3, got %v", rear)
	}
	
	// 测试从两端移除元素
	val, err := dq.RemoveFront() // 移除4，剩余[2, 1, 3]
	if err != nil {
		t.Errorf("RemoveFront failed: %v", err)
	}
	if val != 4 {
		t.Errorf("Expected removed front 4, got %v", val)
	}
	
	val, err = dq.RemoveRear() // 移除3，剩余[2, 1]
	if err != nil {
		t.Errorf("RemoveRear failed: %v", err)
	}
	if val != 3 {
		t.Errorf("Expected removed rear 3, got %v", val)
	}
	
	if dq.Size() != 2 {
		t.Errorf("Expected size 2, got %d", dq.Size())
	}
	
	// 验证剩余元素
	front, _ = dq.Front()
	if front != 2 {
		t.Errorf("Expected front 2, got %v", front)
	}
	
	rear, _ = dq.Rear()
	if rear != 1 {
		t.Errorf("Expected rear 1, got %v", rear)
	}
	
	// 清空测试
	dq.Clear()
	if !dq.IsEmpty() {
		t.Error("Deque should be empty after clear")
	}
	
	// 空双端队列操作测试
	_, err = dq.RemoveFront()
	if err == nil {
		t.Error("Expected error when removing from empty deque")
	}
	
	_, err = dq.RemoveRear()
	if err == nil {
		t.Error("Expected error when removing from empty deque")
	}
}

func TestDequeResize(t *testing.T) {
	dq := NewDeque(2)
	
	// 填满初始容量
	dq.AddRear(1)
	dq.AddRear(2)
	
	// 触发扩容
	dq.AddRear(3)
	
	if dq.Size() != 3 {
		t.Errorf("Expected size 3, got %d", dq.Size())
	}
	
	// 验证元素顺序
	expected := []interface{}{1, 2, 3}
	for _, exp := range expected {
		val, _ := dq.RemoveFront()
		if val != exp {
			t.Errorf("Expected %v, got %v", exp, val)
		}
	}
}

// 性能测试
func BenchmarkEnqueue(b *testing.B) {
	q := NewQueue(1)
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		q.Enqueue(i)
	}
}

func BenchmarkDequeue(b *testing.B) {
	q := NewQueue(b.N)
	for i := 0; i < b.N; i++ {
		q.Enqueue(i)
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Dequeue()
	}
}

func BenchmarkQueueWithLinkedListEnqueue(b *testing.B) {
	q := NewQueueWithLinkedList()
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		q.Enqueue(i)
	}
}

func BenchmarkQueueWithLinkedListDequeue(b *testing.B) {
	q := NewQueueWithLinkedList()
	for i := 0; i < b.N; i++ {
		q.Enqueue(i)
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Dequeue()
	}
}

func BenchmarkDequeAddRear(b *testing.B) {
	dq := NewDeque(1)
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		dq.AddRear(i)
	}
}

func BenchmarkDequeAddFront(b *testing.B) {
	dq := NewDeque(1)
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		dq.AddFront(i)
	}
}