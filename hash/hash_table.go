package hash

import (
	"errors"
	"fmt"
	"hash/fnv"
	"strings"
)

// HashTable 哈希表实现（使用链地址法解决冲突）
// 思路：通过哈希函数将键映射到数组索引，使用链表处理哈希冲突
// 作用：提供平均O(1)时间复杂度的插入、删除、查找操作
// 业务场景：
// 1. 数据库索引：快速查找记录
// 2. 缓存系统：Redis、Memcached的核心数据结构
// 3. 编程语言：变量符号表，对象属性存储
// 4. 网络路由：IP地址到MAC地址的映射
// 5. 去重系统：快速检查元素是否存在
// 6. 分布式系统：一致性哈希，负载均衡
// 7. 搜索引擎：倒排索引的实现
type HashTable struct {
	buckets  []*HashNode // 桶数组，每个桶是一个链表的头节点
	size     int         // 当前存储的键值对数量
	capacity int         // 桶的数量
	loadFactor float64  // 负载因子阈值
}

// HashNode 哈希表节点（链表节点）
type HashNode struct {
	Key   string      // 键
	Value interface{} // 值
	Next  *HashNode   // 指向下一个节点
}

// NewHashTable 创建新的哈希表
// initialCapacity: 初始桶数量
// loadFactor: 负载因子阈值，超过此值时进行扩容
func NewHashTable(initialCapacity int, loadFactor float64) *HashTable {
	if initialCapacity <= 0 {
		initialCapacity = 16
	}
	if loadFactor <= 0 || loadFactor > 1 {
		loadFactor = 0.75
	}
	
	return &HashTable{
		buckets:    make([]*HashNode, initialCapacity),
		size:       0,
		capacity:   initialCapacity,
		loadFactor: loadFactor,
	}
}

// NewDefaultHashTable 创建默认配置的哈希表
func NewDefaultHashTable() *HashTable {
	return NewHashTable(16, 0.75)
}

// hash 哈希函数，将字符串键映射到桶索引
func (ht *HashTable) hash(key string) int {
	h := fnv.New32a()
	h.Write([]byte(key))
	return int(h.Sum32()) % ht.capacity
}

// Put 插入或更新键值对
// 时间复杂度：平均O(1)，最坏O(n)
func (ht *HashTable) Put(key string, value interface{}) {
	// 检查是否需要扩容
	if float64(ht.size) >= float64(ht.capacity)*ht.loadFactor {
		ht.resize()
	}
	
	index := ht.hash(key)
	head := ht.buckets[index]
	
	// 如果桶为空，直接插入
	if head == nil {
		ht.buckets[index] = &HashNode{
			Key:   key,
			Value: value,
			Next:  nil,
		}
		ht.size++
		return
	}
	
	// 遍历链表，查找是否已存在该键
	current := head
	for current != nil {
		if current.Key == key {
			// 键已存在，更新值
			current.Value = value
			return
		}
		current = current.Next
	}
	
	// 键不存在，在链表头部插入新节点
	newNode := &HashNode{
		Key:   key,
		Value: value,
		Next:  head,
	}
	ht.buckets[index] = newNode
	ht.size++
}

// Get 获取指定键的值
// 时间复杂度：平均O(1)，最坏O(n)
func (ht *HashTable) Get(key string) (interface{}, error) {
	index := ht.hash(key)
	current := ht.buckets[index]
	
	// 遍历链表查找键
	for current != nil {
		if current.Key == key {
			return current.Value, nil
		}
		current = current.Next
	}
	
	return nil, errors.New("key not found")
}

// Contains 检查是否包含指定键
// 时间复杂度：平均O(1)，最坏O(n)
func (ht *HashTable) Contains(key string) bool {
	_, err := ht.Get(key)
	return err == nil
}

// Remove 删除指定键的键值对
// 时间复杂度：平均O(1)，最坏O(n)
func (ht *HashTable) Remove(key string) (interface{}, error) {
	index := ht.hash(key)
	head := ht.buckets[index]
	
	if head == nil {
		return nil, errors.New("key not found")
	}
	
	// 如果要删除的是头节点
	if head.Key == key {
		ht.buckets[index] = head.Next
		ht.size--
		return head.Value, nil
	}
	
	// 遍历链表查找要删除的节点
	current := head
	for current.Next != nil {
		if current.Next.Key == key {
			nodeToRemove := current.Next
			current.Next = nodeToRemove.Next
			ht.size--
			return nodeToRemove.Value, nil
		}
		current = current.Next
	}
	
	return nil, errors.New("key not found")
}

// Size 返回哈希表的大小
func (ht *HashTable) Size() int {
	return ht.size
}

// IsEmpty 判断哈希表是否为空
func (ht *HashTable) IsEmpty() bool {
	return ht.size == 0
}

// Capacity 返回当前容量（桶数量）
func (ht *HashTable) Capacity() int {
	return ht.capacity
}

// LoadFactor 返回当前负载因子
func (ht *HashTable) LoadFactor() float64 {
	return float64(ht.size) / float64(ht.capacity)
}

// Keys 返回所有键的切片
// 时间复杂度：O(n)
func (ht *HashTable) Keys() []string {
	keys := make([]string, 0, ht.size)
	
	for i := 0; i < ht.capacity; i++ {
		current := ht.buckets[i]
		for current != nil {
			keys = append(keys, current.Key)
			current = current.Next
		}
	}
	
	return keys
}

// Values 返回所有值的切片
// 时间复杂度：O(n)
func (ht *HashTable) Values() []interface{} {
	values := make([]interface{}, 0, ht.size)
	
	for i := 0; i < ht.capacity; i++ {
		current := ht.buckets[i]
		for current != nil {
			values = append(values, current.Value)
			current = current.Next
		}
	}
	
	return values
}

// Entries 返回所有键值对
// 时间复杂度：O(n)
func (ht *HashTable) Entries() []Entry {
	entries := make([]Entry, 0, ht.size)
	
	for i := 0; i < ht.capacity; i++ {
		current := ht.buckets[i]
		for current != nil {
			entries = append(entries, Entry{
				Key:   current.Key,
				Value: current.Value,
			})
			current = current.Next
		}
	}
	
	return entries
}

// Entry 键值对结构
type Entry struct {
	Key   string
	Value interface{}
}

// Clear 清空哈希表
func (ht *HashTable) Clear() {
	for i := 0; i < ht.capacity; i++ {
		ht.buckets[i] = nil
	}
	ht.size = 0
}

// resize 扩容，容量翻倍并重新哈希所有元素
func (ht *HashTable) resize() {
	oldBuckets := ht.buckets
	oldCapacity := ht.capacity
	
	// 创建新的桶数组
	ht.capacity *= 2
	ht.buckets = make([]*HashNode, ht.capacity)
	ht.size = 0
	
	// 重新插入所有元素
	for i := 0; i < oldCapacity; i++ {
		current := oldBuckets[i]
		for current != nil {
			ht.Put(current.Key, current.Value)
			current = current.Next
		}
	}
}

// GetBucketDistribution 获取桶的分布情况（用于分析哈希函数性能）
func (ht *HashTable) GetBucketDistribution() []int {
	distribution := make([]int, ht.capacity)
	
	for i := 0; i < ht.capacity; i++ {
		count := 0
		current := ht.buckets[i]
		for current != nil {
			count++
			current = current.Next
		}
		distribution[i] = count
	}
	
	return distribution
}

// GetMaxChainLength 获取最长链表的长度
func (ht *HashTable) GetMaxChainLength() int {
	maxLength := 0
	
	for i := 0; i < ht.capacity; i++ {
		length := 0
		current := ht.buckets[i]
		for current != nil {
			length++
			current = current.Next
		}
		if length > maxLength {
			maxLength = length
		}
	}
	
	return maxLength
}

// String 字符串表示
func (ht *HashTable) String() string {
	if ht.IsEmpty() {
		return "HashTable{}"
	}
	
	var elements []string
	for i := 0; i < ht.capacity; i++ {
		current := ht.buckets[i]
		for current != nil {
			elements = append(elements, fmt.Sprintf("%s: %v", current.Key, current.Value))
			current = current.Next
		}
	}
	
	return fmt.Sprintf("HashTable{size: %d, capacity: %d, loadFactor: %.2f, elements: {%s}}", 
		ht.size, ht.capacity, ht.LoadFactor(), strings.Join(elements, ", "))
}

// HashTableWithOpenAddressing 使用开放地址法的哈希表实现
// 优点：内存使用更紧凑，缓存友好
// 缺点：删除操作复杂，负载因子不能太高
type HashTableWithOpenAddressing struct {
	keys     []string      // 键数组
	values   []interface{} // 值数组
	deleted  []bool        // 删除标记数组
	size     int           // 当前大小
	capacity int           // 容量
	loadFactor float64    // 负载因子阈值
}

// NewHashTableWithOpenAddressing 创建开放地址法哈希表
func NewHashTableWithOpenAddressing(initialCapacity int, loadFactor float64) *HashTableWithOpenAddressing {
	if initialCapacity <= 0 {
		initialCapacity = 16
	}
	if loadFactor <= 0 || loadFactor > 0.5 { // 开放地址法负载因子不宜过高
		loadFactor = 0.5
	}
	
	return &HashTableWithOpenAddressing{
		keys:       make([]string, initialCapacity),
		values:     make([]interface{}, initialCapacity),
		deleted:    make([]bool, initialCapacity),
		size:       0,
		capacity:   initialCapacity,
		loadFactor: loadFactor,
	}
}

// hash 哈希函数
func (ht *HashTableWithOpenAddressing) hash(key string) int {
	h := fnv.New32a()
	h.Write([]byte(key))
	return int(h.Sum32()) % ht.capacity
}

// Put 插入或更新键值对（线性探测）
func (ht *HashTableWithOpenAddressing) Put(key string, value interface{}) {
	// 检查是否需要扩容
	if float64(ht.size) >= float64(ht.capacity)*ht.loadFactor {
		ht.resize()
	}
	
	index := ht.hash(key)
	
	// 线性探测找到空位或相同键
	for {
		if ht.keys[index] == "" || ht.deleted[index] {
			// 找到空位，插入新键值对
			if ht.keys[index] == "" || ht.deleted[index] {
				if ht.deleted[index] {
					ht.deleted[index] = false
				} else {
					ht.size++
				}
			}
			ht.keys[index] = key
			ht.values[index] = value
			return
		} else if ht.keys[index] == key {
			// 键已存在，更新值
			ht.values[index] = value
			return
		}
		
		// 继续探测下一个位置
		index = (index + 1) % ht.capacity
	}
}

// Get 获取指定键的值
func (ht *HashTableWithOpenAddressing) Get(key string) (interface{}, error) {
	index := ht.hash(key)
	
	// 线性探测查找键
	for i := 0; i < ht.capacity; i++ {
		currentIndex := (index + i) % ht.capacity
		
		if ht.keys[currentIndex] == "" && !ht.deleted[currentIndex] {
			// 遇到真正的空位，键不存在
			break
		}
		
		if ht.keys[currentIndex] == key && !ht.deleted[currentIndex] {
			return ht.values[currentIndex], nil
		}
	}
	
	return nil, errors.New("key not found")
}

// Remove 删除指定键（懒删除）
func (ht *HashTableWithOpenAddressing) Remove(key string) (interface{}, error) {
	index := ht.hash(key)
	
	// 线性探测查找键
	for i := 0; i < ht.capacity; i++ {
		currentIndex := (index + i) % ht.capacity
		
		if ht.keys[currentIndex] == "" && !ht.deleted[currentIndex] {
			// 遇到真正的空位，键不存在
			break
		}
		
		if ht.keys[currentIndex] == key && !ht.deleted[currentIndex] {
			// 找到键，标记为删除
			value := ht.values[currentIndex]
			ht.deleted[currentIndex] = true
			ht.size--
			return value, nil
		}
	}
	
	return nil, errors.New("key not found")
}

// Size 返回大小
func (ht *HashTableWithOpenAddressing) Size() int {
	return ht.size
}

// IsEmpty 判断是否为空
func (ht *HashTableWithOpenAddressing) IsEmpty() bool {
	return ht.size == 0
}

// resize 扩容
func (ht *HashTableWithOpenAddressing) resize() {
	oldKeys := ht.keys
	oldValues := ht.values
	oldDeleted := ht.deleted
	oldCapacity := ht.capacity
	
	// 创建新数组
	ht.capacity *= 2
	ht.keys = make([]string, ht.capacity)
	ht.values = make([]interface{}, ht.capacity)
	ht.deleted = make([]bool, ht.capacity)
	ht.size = 0
	
	// 重新插入所有有效元素
	for i := 0; i < oldCapacity; i++ {
		if oldKeys[i] != "" && !oldDeleted[i] {
			ht.Put(oldKeys[i], oldValues[i])
		}
	}
}

// 业务应用示例：
// 1. 数据库：索引结构，快速定位记录
// 2. 缓存系统：内存缓存，快速存取数据
// 3. 编程语言：对象属性存储，变量查找表
// 4. 网络系统：路由表，DNS解析
// 5. 安全系统：密码存储，会话管理
// 6. 搜索引擎：倒排索引，快速检索
// 7. 分布式系统：一致性哈希，数据分片
// 8. 去重系统：快速判断数据是否重复