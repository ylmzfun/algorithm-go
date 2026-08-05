package linkedlist

import (
	"testing"
)

func TestNewSinglyLinkedList(t *testing.T) {
	sll := NewSinglyLinkedList()
	
	if sll.Size() != 0 {
		t.Errorf("Expected size 0, got %d", sll.Size())
	}
	
	if !sll.IsEmpty() {
		t.Error("New list should be empty")
	}
	
	if sll.head != nil {
		t.Error("Head should be nil for new list")
	}
	
	if sll.tail != nil {
		t.Error("Tail should be nil for new list")
	}
}

func TestAddFirst(t *testing.T) {
	sll := NewSinglyLinkedList()
	
	sll.AddFirst(1)
	sll.AddFirst(2)
	sll.AddFirst(3)
	
	if sll.Size() != 3 {
		t.Errorf("Expected size 3, got %d", sll.Size())
	}
	
	// 验证顺序：3 -> 2 -> 1
	val, _ := sll.Get(0)
	if val != 3 {
		t.Errorf("Expected first element 3, got %v", val)
	}
	
	val, _ = sll.Get(1)
	if val != 2 {
		t.Errorf("Expected second element 2, got %v", val)
	}
	
	val, _ = sll.Get(2)
	if val != 1 {
		t.Errorf("Expected third element 1, got %v", val)
	}
}

func TestAddLast(t *testing.T) {
	sll := NewSinglyLinkedList()
	
	sll.AddLast(1)
	sll.AddLast(2)
	sll.AddLast(3)
	
	if sll.Size() != 3 {
		t.Errorf("Expected size 3, got %d", sll.Size())
	}
	
	// 验证顺序：1 -> 2 -> 3
	val, _ := sll.Get(0)
	if val != 1 {
		t.Errorf("Expected first element 1, got %v", val)
	}
	
	val, _ = sll.Get(2)
	if val != 3 {
		t.Errorf("Expected last element 3, got %v", val)
	}
	
	// 验证尾指针
	lastVal, _ := sll.GetLast()
	if lastVal != 3 {
		t.Errorf("Expected tail element 3, got %v", lastVal)
	}
}

func TestInsert(t *testing.T) {
	sll := NewSinglyLinkedList()
	sll.Add(1)
	sll.Add(3)
	sll.Add(5)
	
	// 在中间插入
	err := sll.Insert(1, 2)
	if err != nil {
		t.Errorf("Insert failed: %v", err)
	}
	
	if sll.Size() != 4 {
		t.Errorf("Expected size 4, got %d", sll.Size())
	}
	
	// 验证插入结果：1 -> 2 -> 3 -> 5
	expected := []interface{}{1, 2, 3, 5}
	for i, exp := range expected {
		val, _ := sll.Get(i)
		if val != exp {
			t.Errorf("At index %d, expected %v, got %v", i, exp, val)
		}
	}
	
	// 在开头插入
	err = sll.Insert(0, 0)
	if err != nil {
		t.Errorf("Insert at beginning failed: %v", err)
	}
	
	first, _ := sll.GetFirst()
	if first != 0 {
		t.Errorf("Expected first element 0, got %v", first)
	}
	
	// 在末尾插入
	err = sll.Insert(sll.Size(), 6)
	if err != nil {
		t.Errorf("Insert at end failed: %v", err)
	}
	
	last, _ := sll.GetLast()
	if last != 6 {
		t.Errorf("Expected last element 6, got %v", last)
	}
	
	// 测试越界
	err = sll.Insert(-1, 0)
	if err == nil {
		t.Error("Expected error for negative index")
	}
	
	err = sll.Insert(sll.Size()+1, 0)
	if err == nil {
		t.Error("Expected error for index out of bounds")
	}
}

func TestRemoveFirst(t *testing.T) {
	sll := NewSinglyLinkedList()
	
	// 空链表删除
	_, err := sll.RemoveFirst()
	if err == nil {
		t.Error("Expected error for empty list")
	}
	
	// 添加元素后删除
	sll.Add(1)
	sll.Add(2)
	sll.Add(3)
	
	val, err := sll.RemoveFirst()
	if err != nil {
		t.Errorf("RemoveFirst failed: %v", err)
	}
	if val != 1 {
		t.Errorf("Expected removed value 1, got %v", val)
	}
	
	if sll.Size() != 2 {
		t.Errorf("Expected size 2, got %d", sll.Size())
	}
	
	first, _ := sll.GetFirst()
	if first != 2 {
		t.Errorf("Expected new first element 2, got %v", first)
	}
}

func TestRemoveLast(t *testing.T) {
	sll := NewSinglyLinkedList()
	
	// 空链表删除
	_, err := sll.RemoveLast()
	if err == nil {
		t.Error("Expected error for empty list")
	}
	
	// 只有一个元素
	sll.Add(1)
	val, err := sll.RemoveLast()
	if err != nil {
		t.Errorf("RemoveLast failed: %v", err)
	}
	if val != 1 {
		t.Errorf("Expected removed value 1, got %v", val)
	}
	
	if sll.Size() != 0 {
		t.Errorf("Expected size 0, got %d", sll.Size())
	}
	
	// 多个元素
	sll.Add(1)
	sll.Add(2)
	sll.Add(3)
	
	val, err = sll.RemoveLast()
	if err != nil {
		t.Errorf("RemoveLast failed: %v", err)
	}
	if val != 3 {
		t.Errorf("Expected removed value 3, got %v", val)
	}
	
	last, _ := sll.GetLast()
	if last != 2 {
		t.Errorf("Expected new last element 2, got %v", last)
	}
}

func TestRemove(t *testing.T) {
	sll := NewSinglyLinkedList()
	sll.Add(1)
	sll.Add(2)
	sll.Add(3)
	sll.Add(4)
	
	// 删除中间元素
	val, err := sll.Remove(1)
	if err != nil {
		t.Errorf("Remove failed: %v", err)
	}
	if val != 2 {
		t.Errorf("Expected removed value 2, got %v", val)
	}
	
	if sll.Size() != 3 {
		t.Errorf("Expected size 3, got %d", sll.Size())
	}
	
	// 验证删除后的链表：1 -> 3 -> 4
	expected := []interface{}{1, 3, 4}
	for i, exp := range expected {
		val, _ := sll.Get(i)
		if val != exp {
			t.Errorf("At index %d, expected %v, got %v", i, exp, val)
		}
	}
	
	// 测试越界
	_, err = sll.Remove(-1)
	if err == nil {
		t.Error("Expected error for negative index")
	}
	
	_, err = sll.Remove(sll.Size())
	if err == nil {
		t.Error("Expected error for index out of bounds")
	}
}

func TestRemoveElement(t *testing.T) {
	sll := NewSinglyLinkedList()
	sll.Add("apple")
	sll.Add("banana")
	sll.Add("cherry")
	sll.Add("banana")
	
	// 删除存在的元素（删除第一个匹配的）
	removed := sll.RemoveElement("banana")
	if !removed {
		t.Error("Expected to remove 'banana'")
	}
	
	if sll.Size() != 3 {
		t.Errorf("Expected size 3, got %d", sll.Size())
	}
	
	// 验证还有一个banana
	if !sll.Contains("banana") {
		t.Error("Should still contain 'banana'")
	}
	
	// 删除不存在的元素
	removed = sll.RemoveElement("orange")
	if removed {
		t.Error("Should not remove non-existent element")
	}
	
	// 删除第一个元素
	removed = sll.RemoveElement("apple")
	if !removed {
		t.Error("Expected to remove 'apple'")
	}
	
	first, _ := sll.GetFirst()
	if first == "apple" {
		t.Error("'apple' should be removed from first position")
	}
}

func TestGetters(t *testing.T) {
	sll := NewSinglyLinkedList()
	
	// 空链表测试
	_, err := sll.GetFirst()
	if err == nil {
		t.Error("Expected error for empty list")
	}
	
	_, err = sll.GetLast()
	if err == nil {
		t.Error("Expected error for empty list")
	}
	
	_, err = sll.Get(0)
	if err == nil {
		t.Error("Expected error for empty list")
	}
	
	// 添加元素
	sll.Add("first")
	sll.Add("middle")
	sll.Add("last")
	
	// 测试GetFirst
	first, err := sll.GetFirst()
	if err != nil {
		t.Errorf("GetFirst failed: %v", err)
	}
	if first != "first" {
		t.Errorf("Expected 'first', got %v", first)
	}
	
	// 测试GetLast
	last, err := sll.GetLast()
	if err != nil {
		t.Errorf("GetLast failed: %v", err)
	}
	if last != "last" {
		t.Errorf("Expected 'last', got %v", last)
	}
	
	// 测试Get
	middle, err := sll.Get(1)
	if err != nil {
		t.Errorf("Get failed: %v", err)
	}
	if middle != "middle" {
		t.Errorf("Expected 'middle', got %v", middle)
	}
}

func TestSet(t *testing.T) {
	sll := NewSinglyLinkedList()
	sll.Add(1)
	sll.Add(2)
	sll.Add(3)
	
	// 设置中间元素
	err := sll.Set(1, 20)
	if err != nil {
		t.Errorf("Set failed: %v", err)
	}
	
	val, _ := sll.Get(1)
	if val != 20 {
		t.Errorf("Expected 20, got %v", val)
	}
	
	// 测试越界
	err = sll.Set(-1, 0)
	if err == nil {
		t.Error("Expected error for negative index")
	}
	
	err = sll.Set(sll.Size(), 0)
	if err == nil {
		t.Error("Expected error for index out of bounds")
	}
}

func TestContainsIndexOf(t *testing.T) {
	sll := NewSinglyLinkedList()
	sll.Add("apple")
	sll.Add("banana")
	sll.Add("cherry")
	sll.Add("banana")
	
	// 测试Contains
	if !sll.Contains("banana") {
		t.Error("Expected to contain 'banana'")
	}
	
	if sll.Contains("orange") {
		t.Error("Should not contain 'orange'")
	}
	
	// 测试IndexOf（返回第一个匹配的索引）
	index := sll.IndexOf("banana")
	if index != 1 {
		t.Errorf("Expected index 1, got %d", index)
	}
	
	index = sll.IndexOf("orange")
	if index != -1 {
		t.Errorf("Expected index -1, got %d", index)
	}
}

func TestClear(t *testing.T) {
	sll := NewSinglyLinkedList()
	sll.Add(1)
	sll.Add(2)
	sll.Add(3)
	
	sll.Clear()
	
	if sll.Size() != 0 {
		t.Errorf("Expected size 0 after clear, got %d", sll.Size())
	}
	
	if !sll.IsEmpty() {
		t.Error("List should be empty after clear")
	}
	
	if sll.head != nil {
		t.Error("Head should be nil after clear")
	}
	
	if sll.tail != nil {
		t.Error("Tail should be nil after clear")
	}
}

func TestToSlice(t *testing.T) {
	sll := NewSinglyLinkedList()
	sll.Add(1)
	sll.Add(2)
	sll.Add(3)
	
	slice := sll.ToSlice()
	
	if len(slice) != 3 {
		t.Errorf("Expected slice length 3, got %d", len(slice))
	}
	
	expected := []interface{}{1, 2, 3}
	for i, exp := range expected {
		if slice[i] != exp {
			t.Errorf("At index %d, expected %v, got %v", i, exp, slice[i])
		}
	}
}

func TestReverse(t *testing.T) {
	sll := NewSinglyLinkedList()
	
	// 空链表反转
	sll.Reverse()
	if sll.Size() != 0 {
		t.Error("Empty list should remain empty after reverse")
	}
	
	// 单个元素反转
	sll.Add(1)
	sll.Reverse()
	if sll.Size() != 1 {
		t.Error("Single element list should remain size 1 after reverse")
	}
	val, _ := sll.Get(0)
	if val != 1 {
		t.Error("Single element should remain unchanged after reverse")
	}
	
	// 多个元素反转
	sll.Clear()
	sll.Add(1)
	sll.Add(2)
	sll.Add(3)
	sll.Add(4)
	
	sll.Reverse()
	
	// 验证反转结果：4 -> 3 -> 2 -> 1
	expected := []interface{}{4, 3, 2, 1}
	for i, exp := range expected {
		val, _ := sll.Get(i)
		if val != exp {
			t.Errorf("At index %d, expected %v, got %v", i, exp, val)
		}
	}
	
	// 验证头尾指针
	first, _ := sll.GetFirst()
	if first != 4 {
		t.Errorf("Expected first element 4, got %v", first)
	}
	
	last, _ := sll.GetLast()
	if last != 1 {
		t.Errorf("Expected last element 1, got %v", last)
	}
}

func TestString(t *testing.T) {
	sll := NewSinglyLinkedList()
	
	// 空链表
	str := sll.String()
	if str != "SinglyLinkedList{}" {
		t.Errorf("Expected empty list string, got %s", str)
	}
	
	// 非空链表
	sll.Add(1)
	sll.Add(2)
	sll.Add(3)
	
	str = sll.String()
	if str == "" {
		t.Error("String representation should not be empty")
	}
	t.Logf("String representation: %s", str)
}

// 性能测试
func BenchmarkAddFirst(b *testing.B) {
	sll := NewSinglyLinkedList()
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		sll.AddFirst(i)
	}
}

func BenchmarkAddLast(b *testing.B) {
	sll := NewSinglyLinkedList()
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		sll.AddLast(i)
	}
}

func BenchmarkGet(b *testing.B) {
	sll := NewSinglyLinkedList()
	for i := 0; i < 1000; i++ {
		sll.Add(i)
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sll.Get(i % 1000)
	}
}