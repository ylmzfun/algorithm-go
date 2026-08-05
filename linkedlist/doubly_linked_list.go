package linkedlist

import (
	"errors"
	"fmt"
	"strings"
)

// DNode 双向链表节点
type DNode struct {
	Data interface{} // 节点数据
	Prev *DNode      // 指向前一个节点的指针
	Next *DNode      // 指向后一个节点的指针
}

// DoublyLinkedList 双向链表实现
// 思路：每个节点维护前驱和后继两个指针，支持 O(1) 的双向遍历和两端操作
// 作用：提供高效的双向遍历能力，适合实现 LRU 缓存、双向队列等
// 业务场景：
// 1. LRU 缓存淘汰算法：快速移动节点到头部
// 2. 浏览器的前进/后退导航
// 3. 编辑器的撤销/重做功能
// 4. 音乐播放器的双向播放列表
// 5. 操作系统的页面置换算法
type DoublyLinkedList struct {
	head *DNode // 头哨兵节点
	tail *DNode // 尾哨兵节点
	size int    // 链表长度
}

// NewDoublyLinkedList 创建新的双向链表
func NewDoublyLinkedList() *DoublyLinkedList {
	dll := &DoublyLinkedList{
		head: &DNode{Data: nil, Prev: nil, Next: nil},
		tail: &DNode{Data: nil, Prev: nil, Next: nil},
		size: 0,
	}
	dll.head.Next = dll.tail
	dll.tail.Prev = dll.head
	return dll
}

// AddFirst 在头部添加元素
// 时间复杂度：O(1)
func (dll *DoublyLinkedList) AddFirst(data interface{}) {
	dll.addBetween(data, dll.head, dll.head.Next)
}

// AddLast 在尾部添加元素
// 时间复杂度：O(1)
func (dll *DoublyLinkedList) AddLast(data interface{}) {
	dll.addBetween(data, dll.tail.Prev, dll.tail)
}

// addBetween 在两个节点之间插入新节点
func (dll *DoublyLinkedList) addBetween(data interface{}, prev, next *DNode) {
	newNode := &DNode{
		Data: data,
		Prev: prev,
		Next: next,
	}
	prev.Next = newNode
	next.Prev = newNode
	dll.size++
}

// RemoveFirst 删除第一个元素
// 时间复杂度：O(1)
func (dll *DoublyLinkedList) RemoveFirst() (interface{}, error) {
	if dll.IsEmpty() {
		return nil, errors.New("list is empty")
	}
	return dll.remove(dll.head.Next), nil
}

// RemoveLast 删除最后一个元素
// 时间复杂度：O(1)
func (dll *DoublyLinkedList) RemoveLast() (interface{}, error) {
	if dll.IsEmpty() {
		return nil, errors.New("list is empty")
	}
	return dll.remove(dll.tail.Prev), nil
}

// remove 删除指定节点并返回其数据
func (dll *DoublyLinkedList) remove(node *DNode) interface{} {
	prev := node.Prev
	next := node.Next
	prev.Next = next
	next.Prev = prev
	dll.size--
	return node.Data
}

// Insert 在指定位置插入元素
// 时间复杂度：O(n)
func (dll *DoublyLinkedList) Insert(index int, data interface{}) error {
	if index < 0 || index > dll.size {
		return errors.New("index out of bounds")
	}

	if index == 0 {
		dll.AddFirst(data)
		return nil
	}
	if index == dll.size {
		dll.AddLast(data)
		return nil
	}

	// 从较近的一端开始遍历
	var current *DNode
	if index <= dll.size/2 {
		current = dll.head.Next
		for i := 0; i < index; i++ {
			current = current.Next
		}
	} else {
		current = dll.tail.Prev
		for i := dll.size - 1; i > index; i-- {
			current = current.Prev
		}
	}
	dll.addBetween(data, current.Prev, current)
	return nil
}

// Remove 删除指定位置的元素
// 时间复杂度：O(n)
func (dll *DoublyLinkedList) Remove(index int) (interface{}, error) {
	if index < 0 || index >= dll.size {
		return nil, errors.New("index out of bounds")
	}

	if index == 0 {
		return dll.RemoveFirst()
	}
	if index == dll.size-1 {
		return dll.RemoveLast()
	}

	var current *DNode
	if index <= dll.size/2 {
		current = dll.head.Next
		for i := 0; i < index; i++ {
			current = current.Next
		}
	} else {
		current = dll.tail.Prev
		for i := dll.size - 1; i > index; i-- {
			current = current.Prev
		}
	}
	return dll.remove(current), nil
}

// RemoveElement 删除第一个匹配的元素
// 时间复杂度：O(n)
func (dll *DoublyLinkedList) RemoveElement(data interface{}) bool {
	current := dll.head.Next
	for current != dll.tail {
		if current.Data == data {
			dll.remove(current)
			return true
		}
		current = current.Next
	}
	return false
}

// Get 获取指定位置的元素
// 时间复杂度：O(n)
func (dll *DoublyLinkedList) Get(index int) (interface{}, error) {
	if index < 0 || index >= dll.size {
		return nil, errors.New("index out of bounds")
	}

	var current *DNode
	if index <= dll.size/2 {
		current = dll.head.Next
		for i := 0; i < index; i++ {
			current = current.Next
		}
	} else {
		current = dll.tail.Prev
		for i := dll.size - 1; i > index; i-- {
			current = current.Prev
		}
	}
	return current.Data, nil
}

// GetFirst 获取第一个元素
func (dll *DoublyLinkedList) GetFirst() (interface{}, error) {
	if dll.IsEmpty() {
		return nil, errors.New("list is empty")
	}
	return dll.head.Next.Data, nil
}

// GetLast 获取最后一个元素
func (dll *DoublyLinkedList) GetLast() (interface{}, error) {
	if dll.IsEmpty() {
		return nil, errors.New("list is empty")
	}
	return dll.tail.Prev.Data, nil
}

// Set 设置指定位置的元素
func (dll *DoublyLinkedList) Set(index int, data interface{}) error {
	if index < 0 || index >= dll.size {
		return errors.New("index out of bounds")
	}

	var current *DNode
	if index <= dll.size/2 {
		current = dll.head.Next
		for i := 0; i < index; i++ {
			current = current.Next
		}
	} else {
		current = dll.tail.Prev
		for i := dll.size - 1; i > index; i-- {
			current = current.Prev
		}
	}
	current.Data = data
	return nil
}

// Contains 检查是否包含指定元素
func (dll *DoublyLinkedList) Contains(data interface{}) bool {
	current := dll.head.Next
	for current != dll.tail {
		if current.Data == data {
			return true
		}
		current = current.Next
	}
	return false
}

// IndexOf 查找元素的索引
func (dll *DoublyLinkedList) IndexOf(data interface{}) int {
	current := dll.head.Next
	for i := 0; current != dll.tail; i++ {
		if current.Data == data {
			return i
		}
		current = current.Next
	}
	return -1
}

// Size 返回链表大小
func (dll *DoublyLinkedList) Size() int {
	return dll.size
}

// IsEmpty 判断链表是否为空
func (dll *DoublyLinkedList) IsEmpty() bool {
	return dll.size == 0
}

// Clear 清空链表
func (dll *DoublyLinkedList) Clear() {
	dll.head.Next = dll.tail
	dll.tail.Prev = dll.head
	dll.size = 0
}

// ToSlice 转换为切片（正向）
func (dll *DoublyLinkedList) ToSlice() []interface{} {
	result := make([]interface{}, 0, dll.size)
	current := dll.head.Next
	for current != dll.tail {
		result = append(result, current.Data)
		current = current.Next
	}
	return result
}

// ToSliceReverse 转换为切片（反向）
func (dll *DoublyLinkedList) ToSliceReverse() []interface{} {
	result := make([]interface{}, 0, dll.size)
	current := dll.tail.Prev
	for current != dll.head {
		result = append(result, current.Data)
		current = current.Prev
	}
	return result
}

// String 字符串表示
func (dll *DoublyLinkedList) String() string {
	if dll.IsEmpty() {
		return "DoublyLinkedList{}"
	}

	var elements []string
	current := dll.head.Next
	for current != dll.tail {
		elements = append(elements, fmt.Sprintf("%v", current.Data))
		current = current.Next
	}

	return fmt.Sprintf("DoublyLinkedList{size: %d, elements: [%s]}",
		dll.size, strings.Join(elements, " <-> "))
}
