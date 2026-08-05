package array

import (
	"testing"
)

func TestNewDynamicArray(t *testing.T) {
	// 测试默认容量
	da1 := NewDynamicArray(0)
	if da1.Capacity() != 10 {
		t.Errorf("Expected capacity 10, got %d", da1.Capacity())
	}
	
	// 测试指定容量
	da2 := NewDynamicArray(20)
	if da2.Capacity() != 20 {
		t.Errorf("Expected capacity 20, got %d", da2.Capacity())
	}
	
	if da2.Size() != 0 {
		t.Errorf("Expected size 0, got %d", da2.Size())
	}
}

func TestAdd(t *testing.T) {
	da := NewDynamicArray(2)
	
	// 添加元素
	da.Add(1)
	da.Add(2)
	da.Add(3) // 触发扩容
	
	if da.Size() != 3 {
		t.Errorf("Expected size 3, got %d", da.Size())
	}
	
	if da.Capacity() != 4 { // 容量应该翻倍
		t.Errorf("Expected capacity 4, got %d", da.Capacity())
	}
	
	// 验证元素
	val, _ := da.Get(0)
	if val != 1 {
		t.Errorf("Expected 1, got %v", val)
	}
	
	val, _ = da.Get(2)
	if val != 3 {
		t.Errorf("Expected 3, got %v", val)
	}
}

func TestInsert(t *testing.T) {
	da := NewDynamicArray(5)
	da.Add(1)
	da.Add(3)
	da.Add(5)
	
	// 在中间插入
	err := da.Insert(1, 2)
	if err != nil {
		t.Errorf("Insert failed: %v", err)
	}
	
	if da.Size() != 4 {
		t.Errorf("Expected size 4, got %d", da.Size())
	}
	
	// 验证插入结果
	expected := []interface{}{1, 2, 3, 5}
	for i, exp := range expected {
		val, _ := da.Get(i)
		if val != exp {
			t.Errorf("At index %d, expected %v, got %v", i, exp, val)
		}
	}
	
	// 测试边界插入
	err = da.Insert(0, 0) // 开头插入
	if err != nil {
		t.Errorf("Insert at beginning failed: %v", err)
	}
	
	err = da.Insert(da.Size(), 6) // 末尾插入
	if err != nil {
		t.Errorf("Insert at end failed: %v", err)
	}
	
	// 测试越界
	err = da.Insert(-1, 0)
	if err == nil {
		t.Error("Expected error for negative index")
	}
	
	err = da.Insert(da.Size()+1, 0)
	if err == nil {
		t.Error("Expected error for index out of bounds")
	}
}

func TestRemove(t *testing.T) {
	da := NewDynamicArray(5)
	da.Add(1)
	da.Add(2)
	da.Add(3)
	da.Add(4)
	
	// 删除中间元素
	val, err := da.Remove(1)
	if err != nil {
		t.Errorf("Remove failed: %v", err)
	}
	if val != 2 {
		t.Errorf("Expected removed value 2, got %v", val)
	}
	
	if da.Size() != 3 {
		t.Errorf("Expected size 3, got %d", da.Size())
	}
	
	// 验证删除后的数组
	expected := []interface{}{1, 3, 4}
	for i, exp := range expected {
		val, _ := da.Get(i)
		if val != exp {
			t.Errorf("At index %d, expected %v, got %v", i, exp, val)
		}
	}
	
	// 测试越界删除
	_, err = da.Remove(-1)
	if err == nil {
		t.Error("Expected error for negative index")
	}
	
	_, err = da.Remove(da.Size())
	if err == nil {
		t.Error("Expected error for index out of bounds")
	}
}

func TestRemoveLast(t *testing.T) {
	da := NewDynamicArray(5)
	
	// 空数组删除
	_, err := da.RemoveLast()
	if err == nil {
		t.Error("Expected error for empty array")
	}
	
	// 添加元素后删除
	da.Add(1)
	da.Add(2)
	da.Add(3)
	
	val, err := da.RemoveLast()
	if err != nil {
		t.Errorf("RemoveLast failed: %v", err)
	}
	if val != 3 {
		t.Errorf("Expected removed value 3, got %v", val)
	}
	
	if da.Size() != 2 {
		t.Errorf("Expected size 2, got %d", da.Size())
	}
}

func TestGetSet(t *testing.T) {
	da := NewDynamicArray(5)
	da.Add(1)
	da.Add(2)
	da.Add(3)
	
	// 测试Get
	val, err := da.Get(1)
	if err != nil {
		t.Errorf("Get failed: %v", err)
	}
	if val != 2 {
		t.Errorf("Expected 2, got %v", val)
	}
	
	// 测试Set
	err = da.Set(1, 20)
	if err != nil {
		t.Errorf("Set failed: %v", err)
	}
	
	val, _ = da.Get(1)
	if val != 20 {
		t.Errorf("Expected 20, got %v", val)
	}
	
	// 测试越界
	_, err = da.Get(-1)
	if err == nil {
		t.Error("Expected error for negative index")
	}
	
	err = da.Set(da.Size(), 0)
	if err == nil {
		t.Error("Expected error for index out of bounds")
	}
}

func TestContainsIndexOf(t *testing.T) {
	da := NewDynamicArray(5)
	da.Add("apple")
	da.Add("banana")
	da.Add("cherry")
	
	// 测试Contains
	if !da.Contains("banana") {
		t.Error("Expected to contain 'banana'")
	}
	
	if da.Contains("orange") {
		t.Error("Should not contain 'orange'")
	}
	
	// 测试IndexOf
	index := da.IndexOf("cherry")
	if index != 2 {
		t.Errorf("Expected index 2, got %d", index)
	}
	
	index = da.IndexOf("orange")
	if index != -1 {
		t.Errorf("Expected index -1, got %d", index)
	}
}

func TestClear(t *testing.T) {
	da := NewDynamicArray(5)
	da.Add(1)
	da.Add(2)
	da.Add(3)
	
	da.Clear()
	
	if da.Size() != 0 {
		t.Errorf("Expected size 0 after clear, got %d", da.Size())
	}
	
	if !da.IsEmpty() {
		t.Error("Array should be empty after clear")
	}
}

func TestToSlice(t *testing.T) {
	da := NewDynamicArray(5)
	da.Add(1)
	da.Add(2)
	da.Add(3)
	
	slice := da.ToSlice()
	
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

func TestString(t *testing.T) {
	da := NewDynamicArray(5)
	da.Add(1)
	da.Add(2)
	
	str := da.String()
	if str == "" {
		t.Error("String representation should not be empty")
	}
	t.Logf("String representation: %s", str)
}

// 性能测试
func BenchmarkAdd(b *testing.B) {
	da := NewDynamicArray(1)
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		da.Add(i)
	}
}

func BenchmarkGet(b *testing.B) {
	da := NewDynamicArray(1000)
	for i := 0; i < 1000; i++ {
		da.Add(i)
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		da.Get(i % 1000)
	}
}