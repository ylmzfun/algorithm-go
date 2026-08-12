package hash

import (
	"fmt"
	"strings"
	"testing"
)

// TestNewHashTable 测试哈希表创建
func TestNewHashTable(t *testing.T) {
	// 测试默认构造函数
	ht := NewDefaultHashTable()
	if ht.Size() != 0 {
		t.Errorf("Expected size 0, got %d", ht.Size())
	}
	if ht.Capacity() != 16 {
		t.Errorf("Expected capacity 16, got %d", ht.Capacity())
	}
	if !ht.IsEmpty() {
		t.Error("Expected hash table to be empty")
	}

	// 测试自定义参数构造函数
	ht2 := NewHashTable(32, 0.8)
	if ht2.Capacity() != 32 {
		t.Errorf("Expected capacity 32, got %d", ht2.Capacity())
	}
	if ht2.LoadFactor() != 0.0 {
		t.Errorf("Expected load factor 0.0, got %f", ht2.LoadFactor())
	}

	// 测试无效参数
	ht3 := NewHashTable(-1, -1)
	if ht3.Capacity() != 16 {
		t.Errorf("Expected default capacity 16, got %d", ht3.Capacity())
	}
}

// TestPutAndGet 测试插入和获取
func TestPutAndGet(t *testing.T) {
	ht := NewDefaultHashTable()

	// 测试插入和获取
	ht.Put("key1", "value1")
	ht.Put("key2", 42)
	ht.Put("key3", []int{1, 2, 3})

	if ht.Size() != 3 {
		t.Errorf("Expected size 3, got %d", ht.Size())
	}

	// 测试获取存在的键
	value1, err := ht.Get("key1")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if value1 != "value1" {
		t.Errorf("Expected 'value1', got %v", value1)
	}

	value2, err := ht.Get("key2")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if value2 != 42 {
		t.Errorf("Expected 42, got %v", value2)
	}

	// 测试获取不存在的键
	_, err = ht.Get("nonexistent")
	if err == nil {
		t.Error("Expected error for nonexistent key")
	}

	// 测试更新现有键
	ht.Put("key1", "updated_value1")
	if ht.Size() != 3 {
		t.Errorf("Expected size to remain 3, got %d", ht.Size())
	}

	updatedValue, err := ht.Get("key1")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if updatedValue != "updated_value1" {
		t.Errorf("Expected 'updated_value1', got %v", updatedValue)
	}
}

// TestContains 测试包含检查
func TestContains(t *testing.T) {
	ht := NewDefaultHashTable()

	ht.Put("key1", "value1")
	ht.Put("key2", "value2")

	if !ht.Contains("key1") {
		t.Error("Expected hash table to contain 'key1'")
	}
	if !ht.Contains("key2") {
		t.Error("Expected hash table to contain 'key2'")
	}
	if ht.Contains("key3") {
		t.Error("Expected hash table not to contain 'key3'")
	}
}

// TestRemove 测试删除
func TestRemove(t *testing.T) {
	ht := NewDefaultHashTable()

	ht.Put("key1", "value1")
	ht.Put("key2", "value2")
	ht.Put("key3", "value3")

	// 测试删除存在的键
	removedValue, err := ht.Remove("key2")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if removedValue != "value2" {
		t.Errorf("Expected 'value2', got %v", removedValue)
	}
	if ht.Size() != 2 {
		t.Errorf("Expected size 2, got %d", ht.Size())
	}
	if ht.Contains("key2") {
		t.Error("Expected 'key2' to be removed")
	}

	// 测试删除不存在的键
	_, err = ht.Remove("nonexistent")
	if err == nil {
		t.Error("Expected error for removing nonexistent key")
	}

	// 测试删除头节点
	_, err = ht.Remove("key1")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if ht.Size() != 1 {
		t.Errorf("Expected size 1, got %d", ht.Size())
	}
}

// TestKeysValuesEntries 测试获取键、值、条目
func TestKeysValuesEntries(t *testing.T) {
	ht := NewDefaultHashTable()

	ht.Put("key1", "value1")
	ht.Put("key2", "value2")
	ht.Put("key3", "value3")

	// 测试获取所有键
	keys := ht.Keys()
	if len(keys) != 3 {
		t.Errorf("Expected 3 keys, got %d", len(keys))
	}

	// 验证所有键都存在
	keyMap := make(map[string]bool)
	for _, key := range keys {
		keyMap[key] = true
	}
	if !keyMap["key1"] || !keyMap["key2"] || !keyMap["key3"] {
		t.Error("Missing expected keys")
	}

	// 测试获取所有值
	values := ht.Values()
	if len(values) != 3 {
		t.Errorf("Expected 3 values, got %d", len(values))
	}

	// 测试获取所有条目
	entries := ht.Entries()
	if len(entries) != 3 {
		t.Errorf("Expected 3 entries, got %d", len(entries))
	}

	// 验证条目的完整性
	entryMap := make(map[string]interface{})
	for _, entry := range entries {
		entryMap[entry.Key] = entry.Value
	}
	if entryMap["key1"] != "value1" || entryMap["key2"] != "value2" || entryMap["key3"] != "value3" {
		t.Error("Entry values don't match expected values")
	}
}

// TestClear 测试清空
func TestClear(t *testing.T) {
	ht := NewDefaultHashTable()

	ht.Put("key1", "value1")
	ht.Put("key2", "value2")
	ht.Put("key3", "value3")

	ht.Clear()

	if !ht.IsEmpty() {
		t.Error("Expected hash table to be empty after clear")
	}
	if ht.Size() != 0 {
		t.Errorf("Expected size 0, got %d", ht.Size())
	}
	if ht.Contains("key1") {
		t.Error("Expected 'key1' to be removed after clear")
	}
}

// TestResize 测试自动扩容
func TestResize(t *testing.T) {
	ht := NewHashTable(4, 0.75) // 小容量，容易触发扩容

	initialCapacity := ht.Capacity()

	// 插入足够多的元素触发扩容
	for i := 0; i < 10; i++ {
		ht.Put(fmt.Sprintf("key%d", i), fmt.Sprintf("value%d", i))
	}

	if ht.Capacity() <= initialCapacity {
		t.Errorf("Expected capacity to increase from %d, got %d", initialCapacity, ht.Capacity())
	}

	// 验证所有元素仍然存在
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("key%d", i)
		expectedValue := fmt.Sprintf("value%d", i)
		value, err := ht.Get(key)
		if err != nil {
			t.Errorf("Key %s not found after resize", key)
		}
		if value != expectedValue {
			t.Errorf("Expected %s, got %v for key %s", expectedValue, value, key)
		}
	}
}

// TestBucketDistribution 测试桶分布
func TestBucketDistribution(t *testing.T) {
	ht := NewDefaultHashTable()

	// 插入一些元素
	for i := 0; i < 20; i++ {
		ht.Put(fmt.Sprintf("key%d", i), i)
	}

	distribution := ht.GetBucketDistribution()
	if len(distribution) != ht.Capacity() {
		t.Errorf("Expected distribution length %d, got %d", ht.Capacity(), len(distribution))
	}

	// 验证分布总数等于元素总数
	totalElements := 0
	for _, count := range distribution {
		totalElements += count
	}
	if totalElements != ht.Size() {
		t.Errorf("Expected total elements %d, got %d", ht.Size(), totalElements)
	}

	// 测试最大链长度
	maxChainLength := ht.GetMaxChainLength()
	if maxChainLength < 0 {
		t.Errorf("Expected non-negative max chain length, got %d", maxChainLength)
	}
}

// TestString 测试字符串表示
func TestString(t *testing.T) {
	ht := NewDefaultHashTable()

	// 测试空哈希表
	str := ht.String()
	if str != "HashTable{}" {
		t.Errorf("Expected 'HashTable{}', got '%s'", str)
	}

	// 测试非空哈希表
	ht.Put("key1", "value1")
	str = ht.String()
	if !strings.Contains(str, "key1: value1") {
		t.Errorf("Expected string to contain 'key1: value1', got '%s'", str)
	}
	if !strings.Contains(str, "size: 1") {
		t.Errorf("Expected string to contain 'size: 1', got '%s'", str)
	}
}

// TestHashTableWithOpenAddressing 测试开放地址法哈希表
func TestHashTableWithOpenAddressing(t *testing.T) {
	ht := NewHashTableWithOpenAddressing(8, 0.5)

	// 测试基本操作
	ht.Put("key1", "value1")
	ht.Put("key2", "value2")
	ht.Put("key3", "value3")

	if ht.Size() != 3 {
		t.Errorf("Expected size 3, got %d", ht.Size())
	}

	// 测试获取
	value, err := ht.Get("key1")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if value != "value1" {
		t.Errorf("Expected 'value1', got %v", value)
	}

	// 测试删除
	removedValue, err := ht.Remove("key2")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if removedValue != "value2" {
		t.Errorf("Expected 'value2', got %v", removedValue)
	}
	if ht.Size() != 2 {
		t.Errorf("Expected size 2, got %d", ht.Size())
	}

	// 验证删除后无法获取
	_, err = ht.Get("key2")
	if err == nil {
		t.Error("Expected error for deleted key")
	}

	// 测试更新
	ht.Put("key1", "updated_value1")
	if ht.Size() != 2 {
		t.Errorf("Expected size to remain 2, got %d", ht.Size())
	}

	updatedValue, err := ht.Get("key1")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if updatedValue != "updated_value1" {
		t.Errorf("Expected 'updated_value1', got %v", updatedValue)
	}
}

// TestOpenAddressingResize 测试开放地址法扩容
func TestOpenAddressingResize(t *testing.T) {
	ht := NewHashTableWithOpenAddressing(4, 0.5)

	initialCapacity := ht.capacity

	// 插入足够多的元素触发扩容
	for i := 0; i < 6; i++ {
		ht.Put(fmt.Sprintf("key%d", i), i)
	}

	if ht.capacity <= initialCapacity {
		t.Errorf("Expected capacity to increase from %d, got %d", initialCapacity, ht.capacity)
	}

	// 验证所有元素仍然存在
	for i := 0; i < 6; i++ {
		key := fmt.Sprintf("key%d", i)
		value, err := ht.Get(key)
		if err != nil {
			t.Errorf("Key %s not found after resize", key)
		}
		if value != i {
			t.Errorf("Expected %d, got %v for key %s", i, value, key)
		}
	}
}

// TestHashCollision 测试哈希冲突处理
func TestHashCollision(t *testing.T) {
	ht := NewHashTable(2, 0.75) // 小容量起步，插入过程会自动扩容

	// 持续插入直到发生哈希冲突：FNV-32a 分布均匀，冲突虽不保证出现在
	// 前几个键中，但循环探测能保证测试确定性（一般几十个键内即出现冲突）
	const maxKeys = 10000
	i := 0
	for ; i < maxKeys; i++ {
		ht.Put(fmt.Sprintf("key%d", i), i)
		if ht.GetMaxChainLength() > 1 {
			break
		}
	}
	if i >= maxKeys {
		t.Fatalf("插入 %d 个键仍未发生冲突", maxKeys)
	}

	// 验证已插入的所有元素都能正确获取
	for j := 0; j <= i; j++ {
		key := fmt.Sprintf("key%d", j)
		value, err := ht.Get(key)
		if err != nil {
			t.Errorf("Key %s not found", key)
		}
		if value != j {
			t.Errorf("Expected %d, got %v for key %s", j, value, key)
		}
	}

	// 验证链表长度（已发生冲突，链长必然 > 1）
	maxChainLength := ht.GetMaxChainLength()
	if maxChainLength <= 1 {
		t.Errorf("Expected chain length > 1 due to collisions, got %d", maxChainLength)
	}
}

// BenchmarkHashTablePut 基准测试：插入操作
func BenchmarkHashTablePut(b *testing.B) {
	ht := NewDefaultHashTable()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("key%d", i)
		ht.Put(key, i)
	}
}

// BenchmarkHashTableGet 基准测试：获取操作
func BenchmarkHashTableGet(b *testing.B) {
	ht := NewDefaultHashTable()

	// 预先插入数据
	for i := 0; i < 10000; i++ {
		key := fmt.Sprintf("key%d", i)
		ht.Put(key, i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("key%d", i%10000)
		ht.Get(key)
	}
}

// BenchmarkHashTableRemove 基准测试：删除操作
func BenchmarkHashTableRemove(b *testing.B) {
	ht := NewDefaultHashTable()

	// 预先插入数据
	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("key%d", i)
		ht.Put(key, i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("key%d", i)
		ht.Remove(key)
	}
}

// BenchmarkOpenAddressingPut 基准测试：开放地址法插入
func BenchmarkOpenAddressingPut(b *testing.B) {
	ht := NewHashTableWithOpenAddressing(16, 0.5)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("key%d", i)
		ht.Put(key, i)
	}
}

// BenchmarkOpenAddressingGet 基准测试：开放地址法获取
func BenchmarkOpenAddressingGet(b *testing.B) {
	ht := NewHashTableWithOpenAddressing(16, 0.5)

	// 预先插入数据
	for i := 0; i < 10000; i++ {
		key := fmt.Sprintf("key%d", i)
		ht.Put(key, i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("key%d", i%10000)
		ht.Get(key)
	}
}
