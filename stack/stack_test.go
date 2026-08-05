package stack

import (
	"testing"
)

func TestNewStack(t *testing.T) {
	// 测试默认容量
	s1 := NewStack(0)
	if s1.Capacity() != 10 {
		t.Errorf("Expected capacity 10, got %d", s1.Capacity())
	}
	
	// 测试指定容量
	s2 := NewStack(20)
	if s2.Capacity() != 20 {
		t.Errorf("Expected capacity 20, got %d", s2.Capacity())
	}
	
	if s2.Size() != 0 {
		t.Errorf("Expected size 0, got %d", s2.Size())
	}
	
	if !s2.IsEmpty() {
		t.Error("New stack should be empty")
	}
}

func TestPushPop(t *testing.T) {
	s := NewStack(5)
	
	// 测试压入
	s.Push(1)
	s.Push(2)
	s.Push(3)
	
	if s.Size() != 3 {
		t.Errorf("Expected size 3, got %d", s.Size())
	}
	
	if s.IsEmpty() {
		t.Error("Stack should not be empty")
	}
	
	// 测试弹出（LIFO顺序）
	val, err := s.Pop()
	if err != nil {
		t.Errorf("Pop failed: %v", err)
	}
	if val != 3 {
		t.Errorf("Expected 3, got %v", val)
	}
	
	val, err = s.Pop()
	if err != nil {
		t.Errorf("Pop failed: %v", err)
	}
	if val != 2 {
		t.Errorf("Expected 2, got %v", val)
	}
	
	if s.Size() != 1 {
		t.Errorf("Expected size 1, got %d", s.Size())
	}
	
	val, err = s.Pop()
	if err != nil {
		t.Errorf("Pop failed: %v", err)
	}
	if val != 1 {
		t.Errorf("Expected 1, got %v", val)
	}
	
	if !s.IsEmpty() {
		t.Error("Stack should be empty after popping all elements")
	}
	
	// 测试空栈弹出
	_, err = s.Pop()
	if err == nil {
		t.Error("Expected error when popping from empty stack")
	}
}

func TestPeek(t *testing.T) {
	s := NewStack(5)
	
	// 空栈查看
	_, err := s.Peek()
	if err == nil {
		t.Error("Expected error when peeking empty stack")
	}
	
	// 压入元素后查看
	s.Push("hello")
	s.Push("world")
	
	val, err := s.Peek()
	if err != nil {
		t.Errorf("Peek failed: %v", err)
	}
	if val != "world" {
		t.Errorf("Expected 'world', got %v", val)
	}
	
	// 确保Peek不改变栈的大小
	if s.Size() != 2 {
		t.Errorf("Expected size 2 after peek, got %d", s.Size())
	}
	
	// 再次查看应该返回相同的值
	val2, _ := s.Peek()
	if val2 != val {
		t.Errorf("Peek should return same value, got %v and %v", val, val2)
	}
}

func TestResize(t *testing.T) {
	s := NewStack(2)
	
	// 填满初始容量
	s.Push(1)
	s.Push(2)
	
	if s.Capacity() != 2 {
		t.Errorf("Expected capacity 2, got %d", s.Capacity())
	}
	
	// 触发扩容
	s.Push(3)
	
	if s.Capacity() != 4 {
		t.Errorf("Expected capacity 4 after resize, got %d", s.Capacity())
	}
	
	if s.Size() != 3 {
		t.Errorf("Expected size 3, got %d", s.Size())
	}
	
	// 验证元素顺序
	val, _ := s.Pop()
	if val != 3 {
		t.Errorf("Expected 3, got %v", val)
	}
}

func TestClear(t *testing.T) {
	s := NewStack(5)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	
	s.Clear()
	
	if s.Size() != 0 {
		t.Errorf("Expected size 0 after clear, got %d", s.Size())
	}
	
	if !s.IsEmpty() {
		t.Error("Stack should be empty after clear")
	}
	
	// 清空后应该能正常使用
	s.Push("new")
	val, _ := s.Peek()
	if val != "new" {
		t.Errorf("Expected 'new', got %v", val)
	}
}

func TestContains(t *testing.T) {
	s := NewStack(5)
	s.Push("apple")
	s.Push("banana")
	s.Push("cherry")
	
	if !s.Contains("banana") {
		t.Error("Expected to contain 'banana'")
	}
	
	if s.Contains("orange") {
		t.Error("Should not contain 'orange'")
	}
	
	// 弹出元素后检查
	s.Pop() // 移除 "cherry"
	if s.Contains("cherry") {
		t.Error("Should not contain 'cherry' after popping")
	}
	
	if !s.Contains("banana") {
		t.Error("Should still contain 'banana'")
	}
}

func TestToSlice(t *testing.T) {
	s := NewStack(5)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	
	// 测试正序切片（栈底到栈顶）
	slice := s.ToSlice()
	expected := []interface{}{1, 2, 3}
	
	if len(slice) != 3 {
		t.Errorf("Expected slice length 3, got %d", len(slice))
	}
	
	for i, exp := range expected {
		if slice[i] != exp {
			t.Errorf("At index %d, expected %v, got %v", i, exp, slice[i])
		}
	}
	
	// 测试逆序切片（栈顶到栈底）
	reversedSlice := s.ToSliceReversed()
	expectedReversed := []interface{}{3, 2, 1}
	
	for i, exp := range expectedReversed {
		if reversedSlice[i] != exp {
			t.Errorf("At index %d, expected %v, got %v", i, exp, reversedSlice[i])
		}
	}
}

func TestString(t *testing.T) {
	s := NewStack(5)
	
	// 空栈
	str := s.String()
	if str != "Stack{}" {
		t.Errorf("Expected empty stack string, got %s", str)
	}
	
	// 非空栈
	s.Push(1)
	s.Push(2)
	s.Push(3)
	
	str = s.String()
	if str == "" {
		t.Error("String representation should not be empty")
	}
	t.Logf("String representation: %s", str)
}

// 测试基于链表的栈
func TestStackWithLinkedList(t *testing.T) {
	s := NewStackWithLinkedList()
	
	if !s.IsEmpty() {
		t.Error("New stack should be empty")
	}
	
	if s.Size() != 0 {
		t.Errorf("Expected size 0, got %d", s.Size())
	}
	
	// 测试压入和弹出
	s.Push("first")
	s.Push("second")
	s.Push("third")
	
	if s.Size() != 3 {
		t.Errorf("Expected size 3, got %d", s.Size())
	}
	
	// 测试Peek
	val, err := s.Peek()
	if err != nil {
		t.Errorf("Peek failed: %v", err)
	}
	if val != "third" {
		t.Errorf("Expected 'third', got %v", val)
	}
	
	// 测试Pop
	val, err = s.Pop()
	if err != nil {
		t.Errorf("Pop failed: %v", err)
	}
	if val != "third" {
		t.Errorf("Expected 'third', got %v", val)
	}
	
	if s.Size() != 2 {
		t.Errorf("Expected size 2 after pop, got %d", s.Size())
	}
	
	// 清空测试
	s.Clear()
	if !s.IsEmpty() {
		t.Error("Stack should be empty after clear")
	}
	
	// 空栈操作测试
	_, err = s.Pop()
	if err == nil {
		t.Error("Expected error when popping from empty stack")
	}
	
	_, err = s.Peek()
	if err == nil {
		t.Error("Expected error when peeking empty stack")
	}
}

// 测试括号匹配
func TestIsValidParentheses(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"([{}])", true},
		{"((()))", true},
		{"", true},
		{"(", false},
		{")", false},
		{"([)]", false},
		{"(((", false},
		{")))", false},
		{"([{})])", false},
	}
	
	for _, tc := range testCases {
		result := IsValidParentheses(tc.input)
		if result != tc.expected {
			t.Errorf("IsValidParentheses(%q) = %v, expected %v", 
				tc.input, result, tc.expected)
		}
	}
}

// 测试后缀表达式求值
func TestEvaluatePostfix(t *testing.T) {
	testCases := []struct {
		input    []string
		expected int
		hasError bool
	}{
		{[]string{"2", "1", "+", "3", "*"}, 9, false}, // (2+1)*3 = 9
		{[]string{"4", "13", "5", "/", "+"}, 6, false}, // 4+(13/5) = 4+2 = 6
		{[]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}, 22, false},
		{[]string{"3", "4", "+"}, 7, false},
		{[]string{"3", "4", "-"}, -1, false},
		{[]string{"3", "4", "*"}, 12, false},
		{[]string{"8", "2", "/"}, 4, false},
		{[]string{"3", "+"}, 0, true}, // 操作数不足
		{[]string{"3", "0", "/"}, 0, true}, // 除零错误
		{[]string{"abc"}, 0, true}, // 无效token
	}
	
	for i, tc := range testCases {
		result, err := EvaluatePostfix(tc.input)
		
		if tc.hasError {
			if err == nil {
				t.Errorf("Test case %d: expected error but got none", i)
			}
		} else {
			if err != nil {
				t.Errorf("Test case %d: unexpected error: %v", i, err)
			}
			if result != tc.expected {
				t.Errorf("Test case %d: expected %d, got %d", i, tc.expected, result)
			}
		}
	}
}

// 性能测试
func BenchmarkPush(b *testing.B) {
	s := NewStack(1)
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		s.Push(i)
	}
}

func BenchmarkPop(b *testing.B) {
	s := NewStack(b.N)
	for i := 0; i < b.N; i++ {
		s.Push(i)
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Pop()
	}
}

func BenchmarkPeek(b *testing.B) {
	s := NewStack(1000)
	for i := 0; i < 1000; i++ {
		s.Push(i)
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Peek()
	}
}

func BenchmarkStackWithLinkedListPush(b *testing.B) {
	s := NewStackWithLinkedList()
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		s.Push(i)
	}
}

func BenchmarkStackWithLinkedListPop(b *testing.B) {
	s := NewStackWithLinkedList()
	for i := 0; i < b.N; i++ {
		s.Push(i)
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Pop()
	}
}