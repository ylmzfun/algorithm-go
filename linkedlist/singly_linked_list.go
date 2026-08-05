package linkedlist

import (
	"errors"
	"fmt"
	"strings"
)

// Node 单向链表节点
type Node struct {
	Data interface{} // 节点数据
	Next *Node       // 指向下一个节点的指针
}

// SinglyLinkedList 单向链表实现
// 思路：通过节点之间的指针连接形成线性结构，每个节点只有一个指向下一个节点的指针
// 作用：提供动态的线性数据结构，插入和删除操作高效
// 业务场景：
// 1. 音乐播放器的播放列表（顺序播放）
// 2. 浏览器历史记录（前进功能）
// 3. 撤销操作的历史栈
// 4. 任务队列系统
// 5. 缓存淘汰策略（LRU的一部分）
type SinglyLinkedList struct {
	head *Node // 头节点
	tail *Node // 尾节点（可选，用于优化尾部插入）
	size int   // 链表长度
}

// NewSinglyLinkedList 创建新的单向链表
func NewSinglyLinkedList() *SinglyLinkedList {
	return &SinglyLinkedList{
		head: nil,
		tail: nil,
		size: 0,
	}
}

// AddFirst 在链表头部添加元素
// 时间复杂度：O(1)
func (sll *SinglyLinkedList) AddFirst(data interface{}) {
	newNode := &Node{Data: data, Next: sll.head}
	sll.head = newNode
	
	// 如果是第一个节点，同时设置为尾节点
	if sll.tail == nil {
		sll.tail = newNode
	}
	
	sll.size++
}

// AddLast 在链表尾部添加元素
// 时间复杂度：O(1)（有尾指针）
func (sll *SinglyLinkedList) AddLast(data interface{}) {
	newNode := &Node{Data: data, Next: nil}
	
	if sll.head == nil {
		// 空链表
		sll.head = newNode
		sll.tail = newNode
	} else {
		// 非空链表
		sll.tail.Next = newNode
		sll.tail = newNode
	}
	
	sll.size++
}

// Add 添加元素（默认添加到尾部）
func (sll *SinglyLinkedList) Add(data interface{}) {
	sll.AddLast(data)
}

// Insert 在指定位置插入元素
// 时间复杂度：O(n)
func (sll *SinglyLinkedList) Insert(index int, data interface{}) error {
	if index < 0 || index > sll.size {
		return errors.New("index out of bounds")
	}
	
	// 在头部插入
	if index == 0 {
		sll.AddFirst(data)
		return nil
	}
	
	// 在尾部插入
	if index == sll.size {
		sll.AddLast(data)
		return nil
	}
	
	// 在中间插入
	newNode := &Node{Data: data}
	current := sll.head
	
	// 找到插入位置的前一个节点
	for i := 0; i < index-1; i++ {
		current = current.Next
	}
	
	newNode.Next = current.Next
	current.Next = newNode
	sll.size++
	
	return nil
}

// RemoveFirst 删除第一个元素
// 时间复杂度：O(1)
func (sll *SinglyLinkedList) RemoveFirst() (interface{}, error) {
	if sll.head == nil {
		return nil, errors.New("list is empty")
	}
	
	data := sll.head.Data
	sll.head = sll.head.Next
	
	// 如果删除后链表为空，更新尾指针
	if sll.head == nil {
		sll.tail = nil
	}
	
	sll.size--
	return data, nil
}

// RemoveLast 删除最后一个元素
// 时间复杂度：O(n)（需要找到倒数第二个节点）
func (sll *SinglyLinkedList) RemoveLast() (interface{}, error) {
	if sll.head == nil {
		return nil, errors.New("list is empty")
	}
	
	// 只有一个节点
	if sll.head == sll.tail {
		data := sll.head.Data
		sll.head = nil
		sll.tail = nil
		sll.size--
		return data, nil
	}
	
	// 找到倒数第二个节点
	current := sll.head
	for current.Next != sll.tail {
		current = current.Next
	}
	
	data := sll.tail.Data
	current.Next = nil
	sll.tail = current
	sll.size--
	
	return data, nil
}

// Remove 删除指定位置的元素
// 时间复杂度：O(n)
func (sll *SinglyLinkedList) Remove(index int) (interface{}, error) {
	if index < 0 || index >= sll.size {
		return nil, errors.New("index out of bounds")
	}
	
	// 删除第一个元素
	if index == 0 {
		return sll.RemoveFirst()
	}
	
	// 删除最后一个元素
	if index == sll.size-1 {
		return sll.RemoveLast()
	}
	
	// 删除中间元素
	current := sll.head
	
	// 找到要删除节点的前一个节点
	for i := 0; i < index-1; i++ {
		current = current.Next
	}
	
	nodeToRemove := current.Next
	data := nodeToRemove.Data
	current.Next = nodeToRemove.Next
	sll.size--
	
	return data, nil
}

// RemoveElement 删除第一个匹配的元素
// 时间复杂度：O(n)
func (sll *SinglyLinkedList) RemoveElement(data interface{}) bool {
	if sll.head == nil {
		return false
	}
	
	// 如果要删除的是第一个元素
	if sll.head.Data == data {
		sll.RemoveFirst()
		return true
	}
	
	// 查找要删除的元素
	current := sll.head
	for current.Next != nil {
		if current.Next.Data == data {
			nodeToRemove := current.Next
			current.Next = nodeToRemove.Next
			
			// 如果删除的是尾节点，更新尾指针
			if nodeToRemove == sll.tail {
				sll.tail = current
			}
			
			sll.size--
			return true
		}
		current = current.Next
	}
	
	return false
}

// Get 获取指定位置的元素
// 时间复杂度：O(n)
func (sll *SinglyLinkedList) Get(index int) (interface{}, error) {
	if index < 0 || index >= sll.size {
		return nil, errors.New("index out of bounds")
	}
	
	current := sll.head
	for i := 0; i < index; i++ {
		current = current.Next
	}
	
	return current.Data, nil
}

// GetFirst 获取第一个元素
// 时间复杂度：O(1)
func (sll *SinglyLinkedList) GetFirst() (interface{}, error) {
	if sll.head == nil {
		return nil, errors.New("list is empty")
	}
	return sll.head.Data, nil
}

// GetLast 获取最后一个元素
// 时间复杂度：O(1)（有尾指针）
func (sll *SinglyLinkedList) GetLast() (interface{}, error) {
	if sll.tail == nil {
		return nil, errors.New("list is empty")
	}
	return sll.tail.Data, nil
}

// Set 设置指定位置的元素
// 时间复杂度：O(n)
func (sll *SinglyLinkedList) Set(index int, data interface{}) error {
	if index < 0 || index >= sll.size {
		return errors.New("index out of bounds")
	}
	
	current := sll.head
	for i := 0; i < index; i++ {
		current = current.Next
	}
	
	current.Data = data
	return nil
}

// Contains 检查是否包含指定元素
// 时间复杂度：O(n)
func (sll *SinglyLinkedList) Contains(data interface{}) bool {
	current := sll.head
	for current != nil {
		if current.Data == data {
			return true
		}
		current = current.Next
	}
	return false
}

// IndexOf 查找元素的索引
// 时间复杂度：O(n)
func (sll *SinglyLinkedList) IndexOf(data interface{}) int {
	current := sll.head
	for i := 0; current != nil; i++ {
		if current.Data == data {
			return i
		}
		current = current.Next
	}
	return -1
}

// Size 返回链表大小
func (sll *SinglyLinkedList) Size() int {
	return sll.size
}

// IsEmpty 判断链表是否为空
func (sll *SinglyLinkedList) IsEmpty() bool {
	return sll.size == 0
}

// Clear 清空链表
func (sll *SinglyLinkedList) Clear() {
	sll.head = nil
	sll.tail = nil
	sll.size = 0
}

// ToSlice 转换为切片
func (sll *SinglyLinkedList) ToSlice() []interface{} {
	result := make([]interface{}, 0, sll.size)
	current := sll.head
	
	for current != nil {
		result = append(result, current.Data)
		current = current.Next
	}
	
	return result
}

// Reverse 反转链表
// 时间复杂度：O(n)
func (sll *SinglyLinkedList) Reverse() {
	if sll.head == nil || sll.head.Next == nil {
		return
	}
	
	var prev *Node
	current := sll.head
	sll.tail = sll.head // 原来的头变成尾
	
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	
	sll.head = prev
}

// String 字符串表示
func (sll *SinglyLinkedList) String() string {
	if sll.head == nil {
		return "SinglyLinkedList{}"
	}
	
	var elements []string
	current := sll.head
	
	for current != nil {
		elements = append(elements, fmt.Sprintf("%v", current.Data))
		current = current.Next
	}
	
	return fmt.Sprintf("SinglyLinkedList{size: %d, elements: [%s]}", 
		sll.size, strings.Join(elements, " -> "))
}

// 业务应用示例：
// 1. 音乐播放器：歌曲列表的顺序播放
// 2. 浏览器历史：页面访问历史记录
// 3. 任务调度：按顺序执行的任务队列
// 4. 消息传递：消息的顺序处理
// 5. 文件系统：目录结构的遍历
// 6. 游戏开发：角色移动路径记录
// 7. 数据流处理：流式数据的顺序处理