package stack

import (
	"errors"
	"fmt"
	"strings"
)

// Stack 栈的实现（基于切片）
// 思路：使用切片作为底层存储，遵循LIFO（后进先出）原则
// 作用：提供后进先出的数据访问模式，支持快速的压入和弹出操作
// 业务场景：
// 1. 函数调用栈：程序执行时的函数调用管理
// 2. 表达式求值：中缀表达式转后缀表达式，括号匹配
// 3. 浏览器历史：后退功能的实现
// 4. 撤销操作：编辑器的撤销/重做功能
// 5. 深度优先搜索（DFS）：图和树的遍历
// 6. 内存管理：局部变量的分配和释放
type Stack struct {
	data []interface{} // 存储栈元素的切片
	top  int           // 栈顶指针（指向栈顶元素的下一个位置）
}

// NewStack 创建新的栈
// initialCapacity: 初始容量，如果为0则使用默认值10
func NewStack(initialCapacity int) *Stack {
	if initialCapacity <= 0 {
		initialCapacity = 10
	}
	return &Stack{
		data: make([]interface{}, initialCapacity),
		top:  0,
	}
}

// Push 压入元素到栈顶
// 时间复杂度：平均O(1)，最坏O(n)（需要扩容时）
func (s *Stack) Push(element interface{}) {
	// 检查是否需要扩容
	if s.top >= len(s.data) {
		s.resize()
	}
	
	s.data[s.top] = element
	s.top++
}

// Pop 弹出栈顶元素
// 时间复杂度：O(1)
func (s *Stack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	
	s.top--
	element := s.data[s.top]
	s.data[s.top] = nil // 避免内存泄漏
	
	return element, nil
}

// Peek 查看栈顶元素但不弹出
// 时间复杂度：O(1)
func (s *Stack) Peek() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	
	return s.data[s.top-1], nil
}

// IsEmpty 判断栈是否为空
// 时间复杂度：O(1)
func (s *Stack) IsEmpty() bool {
	return s.top == 0
}

// Size 返回栈的大小
// 时间复杂度：O(1)
func (s *Stack) Size() int {
	return s.top
}

// Capacity 返回当前容量
func (s *Stack) Capacity() int {
	return len(s.data)
}

// Clear 清空栈
// 时间复杂度：O(n)
func (s *Stack) Clear() {
	for i := 0; i < s.top; i++ {
		s.data[i] = nil
	}
	s.top = 0
}

// Contains 检查栈是否包含指定元素
// 时间复杂度：O(n)
func (s *Stack) Contains(element interface{}) bool {
	for i := 0; i < s.top; i++ {
		if s.data[i] == element {
			return true
		}
	}
	return false
}

// ToSlice 转换为切片（从栈底到栈顶）
// 时间复杂度：O(n)
func (s *Stack) ToSlice() []interface{} {
	result := make([]interface{}, s.top)
	copy(result, s.data[:s.top])
	return result
}

// ToSliceReversed 转换为切片（从栈顶到栈底）
// 时间复杂度：O(n)
func (s *Stack) ToSliceReversed() []interface{} {
	result := make([]interface{}, s.top)
	for i := 0; i < s.top; i++ {
		result[i] = s.data[s.top-1-i]
	}
	return result
}

// String 字符串表示
func (s *Stack) String() string {
	if s.IsEmpty() {
		return "Stack{}"
	}
	
	var elements []string
	// 从栈底到栈顶显示
	for i := 0; i < s.top; i++ {
		elements = append(elements, fmt.Sprintf("%v", s.data[i]))
	}
	
	return fmt.Sprintf("Stack{size: %d, elements: [%s] <- top}", 
		s.top, strings.Join(elements, ", "))
}

// resize 扩容，容量翻倍
func (s *Stack) resize() {
	newCapacity := len(s.data) * 2
	newData := make([]interface{}, newCapacity)
	copy(newData, s.data)
	s.data = newData
}

// StackWithLinkedList 基于链表实现的栈
// 优点：动态大小，不需要预分配内存
// 缺点：每个元素需要额外的指针存储空间
type StackWithLinkedList struct {
	top  *StackNode // 栈顶指针
	size int        // 栈大小
}

// StackNode 栈节点
type StackNode struct {
	Data interface{} // 节点数据
	Next *StackNode  // 指向下一个节点
}

// NewStackWithLinkedList 创建基于链表的栈
func NewStackWithLinkedList() *StackWithLinkedList {
	return &StackWithLinkedList{
		top:  nil,
		size: 0,
	}
}

// Push 压入元素
// 时间复杂度：O(1)
func (s *StackWithLinkedList) Push(element interface{}) {
	newNode := &StackNode{
		Data: element,
		Next: s.top,
	}
	s.top = newNode
	s.size++
}

// Pop 弹出元素
// 时间复杂度：O(1)
func (s *StackWithLinkedList) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	
	element := s.top.Data
	s.top = s.top.Next
	s.size--
	
	return element, nil
}

// Peek 查看栈顶元素
// 时间复杂度：O(1)
func (s *StackWithLinkedList) Peek() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	
	return s.top.Data, nil
}

// IsEmpty 判断是否为空
func (s *StackWithLinkedList) IsEmpty() bool {
	return s.top == nil
}

// Size 返回大小
func (s *StackWithLinkedList) Size() int {
	return s.size
}

// Clear 清空栈
func (s *StackWithLinkedList) Clear() {
	s.top = nil
	s.size = 0
}

// String 字符串表示
func (s *StackWithLinkedList) String() string {
	if s.IsEmpty() {
		return "StackWithLinkedList{}"
	}
	
	var elements []string
	current := s.top
	
	// 从栈顶到栈底遍历
	for current != nil {
		elements = append(elements, fmt.Sprintf("%v", current.Data))
		current = current.Next
	}
	
	return fmt.Sprintf("StackWithLinkedList{size: %d, elements: [%s] <- top}", 
		s.size, strings.Join(elements, ", "))
}

// 实用工具函数

// IsValidParentheses 检查括号是否匹配（栈的经典应用）
func IsValidParentheses(s string) bool {
	stack := NewStack(len(s))
	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	
	for _, char := range s {
		switch char {
		case '(', '[', '{':
			stack.Push(char)
		case ')', ']', '}':
			if stack.IsEmpty() {
				return false
			}
			top, _ := stack.Pop()
			if top != pairs[char] {
				return false
			}
		}
	}
	
	return stack.IsEmpty()
}

// EvaluatePostfix 计算后缀表达式（栈的经典应用）
func EvaluatePostfix(expression []string) (int, error) {
	stack := NewStack(len(expression))
	
	for _, token := range expression {
		switch token {
		case "+", "-", "*", "/":
			if stack.Size() < 2 {
				return 0, errors.New("invalid expression")
			}
			
			b, _ := stack.Pop()
			a, _ := stack.Pop()
			
			var result int
			switch token {
			case "+":
				result = a.(int) + b.(int)
			case "-":
				result = a.(int) - b.(int)
			case "*":
				result = a.(int) * b.(int)
			case "/":
				if b.(int) == 0 {
					return 0, errors.New("division by zero")
				}
				result = a.(int) / b.(int)
			}
			
			stack.Push(result)
		default:
			// 假设是数字
			var num int
			if _, err := fmt.Sscanf(token, "%d", &num); err != nil {
				return 0, fmt.Errorf("invalid token: %s", token)
			}
			stack.Push(num)
		}
	}
	
	if stack.Size() != 1 {
		return 0, errors.New("invalid expression")
	}
	
	result, _ := stack.Pop()
	return result.(int), nil
}

// 业务应用示例：
// 1. 编译器：语法分析，表达式求值
// 2. 浏览器：页面历史管理（后退按钮）
// 3. 编辑器：撤销/重做操作栈
// 4. 游戏开发：状态管理，场景切换
// 5. 算法：深度优先搜索，回溯算法
// 6. 系统调用：函数调用栈管理
// 7. 计算器：表达式计算和括号匹配