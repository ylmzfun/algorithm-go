package advanced

import (
	"strings"
	"testing"
)

func TestMax(t *testing.T) {
	if got := Max(3, 5); got != 5 {
		t.Errorf("Max(3,5): expected 5, got %d", got)
	}
	if got := Max(3, 3); got != 3 {
		t.Errorf("Max(3,3): expected 3, got %d", got)
	}
	if got := Max(1.5, 1.2); got != 1.5 {
		t.Errorf("Max(1.5,1.2): expected 1.5, got %v", got)
	}
	if got := Max("apple", "banana"); got != "banana" {
		t.Errorf("Max(apple,banana): expected banana, got %q", got)
	}
}

func TestMin(t *testing.T) {
	if got := Min(5, 3); got != 3 {
		t.Errorf("Min(5,3): expected 3, got %d", got)
	}
	if got := Min("banana", "apple"); got != "apple" {
		t.Errorf("Min(banana,apple): expected apple, got %q", got)
	}
}

func TestFilter(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6}
	even := Filter(nums, func(v int) bool { return v%2 == 0 })
	expected := []int{2, 4, 6}
	if len(even) != len(expected) {
		t.Fatalf("Expected %v, got %v", expected, even)
	}
	for i := range expected {
		if even[i] != expected[i] {
			t.Errorf("At index %d, expected %d, got %d", i, expected[i], even[i])
		}
	}

	// 空切片
	if got := Filter([]int{}, func(v int) bool { return true }); len(got) != 0 {
		t.Errorf("Expected empty result, got %v", got)
	}

	// 全部过滤掉
	if got := Filter(nums, func(v int) bool { return false }); len(got) != 0 {
		t.Errorf("Expected empty result when all filtered out, got %v", got)
	}
}

func TestMap(t *testing.T) {
	nums := []int{1, 2, 3}
	strs := Map(nums, func(v int) string { return strings.Repeat("x", v) })
	expected := []string{"x", "xx", "xxx"}
	if len(strs) != len(expected) {
		t.Fatalf("Expected %v, got %v", expected, strs)
	}
	for i := range expected {
		if strs[i] != expected[i] {
			t.Errorf("At index %d, expected %q, got %q", i, expected[i], strs[i])
		}
	}
}

func TestContains(t *testing.T) {
	nums := []int{1, 2, 3}
	if !Contains(nums, 2) {
		t.Error("Expected nums to contain 2")
	}
	if Contains(nums, 9) {
		t.Error("Expected nums not to contain 9")
	}
	if !Contains([]string{"a", "b"}, "a") {
		t.Error("Expected string slice to contain 'a'")
	}
	if Contains([]string{}, "x") {
		t.Error("Empty slice should not contain anything")
	}
}

func TestStack(t *testing.T) {
	s := NewStack[int]()
	if !s.IsEmpty() {
		t.Error("New stack should be empty")
	}
	if s.Size() != 0 {
		t.Errorf("Expected size 0, got %d", s.Size())
	}

	// 空栈 Pop/Peek：返回零值和 false
	if v, ok := s.Pop(); ok {
		t.Errorf("Pop on empty stack should return false, got %v", v)
	}
	if v, ok := s.Peek(); ok {
		t.Errorf("Peek on empty stack should return false, got %v", v)
	}

	// 压入与弹出（LIFO 顺序）
	s.Push(1)
	s.Push(2)
	s.Push(3)
	if s.Size() != 3 {
		t.Errorf("Expected size 3, got %d", s.Size())
	}
	if v, ok := s.Peek(); !ok || v != 3 {
		t.Errorf("Peek: expected 3, got %d (ok=%v)", v, ok)
	}

	expected := []int{3, 2, 1}
	for i, exp := range expected {
		v, ok := s.Pop()
		if !ok || v != exp {
			t.Errorf("Pop %d: expected %d, got %d (ok=%v)", i, exp, v, ok)
		}
	}
	if !s.IsEmpty() {
		t.Error("Stack should be empty after popping all elements")
	}

	// 泛型栈支持不同类型
	strStack := NewStack[string]()
	strStack.Push("hello")
	if v, _ := strStack.Pop(); v != "hello" {
		t.Errorf("Expected 'hello', got %q", v)
	}
}
