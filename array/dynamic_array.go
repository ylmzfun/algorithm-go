package array

import (
	"errors"
	"fmt"
)

// DynamicArray 动态数组实现
// 思路：使用切片作为底层存储，当容量不足时自动扩容
// 作用：提供类似于其他语言中ArrayList的功能，支持动态增长
// 业务场景：
// 1. 缓存系统中存储热点数据
// 2. 日志收集系统中临时存储日志条目
// 3. 购物车商品列表管理
// 4. 用户行为轨迹记录
type DynamicArray struct {
	data     []interface{} // 存储数据的切片
	size     int           // 当前元素数量
	capacity int           // 当前容量
}

// NewDynamicArray 创建新的动态数组
// initialCapacity: 初始容量，如果为0则使用默认值10
func NewDynamicArray(initialCapacity int) *DynamicArray {
	if initialCapacity <= 0 {
		initialCapacity = 10
	}
	return &DynamicArray{
		data:     make([]interface{}, initialCapacity),
		size:     0,
		capacity: initialCapacity,
	}
}

// Add 在数组末尾添加元素
// 时间复杂度：平均O(1)，最坏O(n)（需要扩容时）
func (da *DynamicArray) Add(element interface{}) {
	if da.size >= da.capacity {
		da.resize()
	}
	da.data[da.size] = element
	da.size++
}

// Insert 在指定位置插入元素
// 时间复杂度：O(n)
func (da *DynamicArray) Insert(index int, element interface{}) error {
	if index < 0 || index > da.size {
		return errors.New("index out of bounds")
	}
	
	if da.size >= da.capacity {
		da.resize()
	}
	
	// 将index位置及之后的元素向后移动
	for i := da.size; i > index; i-- {
		da.data[i] = da.data[i-1]
	}
	
	da.data[index] = element
	da.size++
	return nil
}

// Get 获取指定位置的元素
// 时间复杂度：O(1)
func (da *DynamicArray) Get(index int) (interface{}, error) {
	if index < 0 || index >= da.size {
		return nil, errors.New("index out of bounds")
	}
	return da.data[index], nil
}

// Set 设置指定位置的元素
// 时间复杂度：O(1)
func (da *DynamicArray) Set(index int, element interface{}) error {
	if index < 0 || index >= da.size {
		return errors.New("index out of bounds")
	}
	da.data[index] = element
	return nil
}

// Remove 删除指定位置的元素
// 时间复杂度：O(n)
func (da *DynamicArray) Remove(index int) (interface{}, error) {
	if index < 0 || index >= da.size {
		return nil, errors.New("index out of bounds")
	}
	
	element := da.data[index]
	
	// 将index之后的元素向前移动
	for i := index; i < da.size-1; i++ {
		da.data[i] = da.data[i+1]
	}
	
	da.size--
	da.data[da.size] = nil // 避免内存泄漏
	
	return element, nil
}

// RemoveLast 删除最后一个元素
// 时间复杂度：O(1)
func (da *DynamicArray) RemoveLast() (interface{}, error) {
	if da.size == 0 {
		return nil, errors.New("array is empty")
	}
	return da.Remove(da.size - 1)
}

// Size 返回数组大小
func (da *DynamicArray) Size() int {
	return da.size
}

// IsEmpty 判断数组是否为空
func (da *DynamicArray) IsEmpty() bool {
	return da.size == 0
}

// Capacity 返回当前容量
func (da *DynamicArray) Capacity() int {
	return da.capacity
}

// Clear 清空数组
func (da *DynamicArray) Clear() {
	for i := 0; i < da.size; i++ {
		da.data[i] = nil
	}
	da.size = 0
}

// Contains 检查是否包含指定元素
// 时间复杂度：O(n)
func (da *DynamicArray) Contains(element interface{}) bool {
	for i := 0; i < da.size; i++ {
		if da.data[i] == element {
			return true
		}
	}
	return false
}

// IndexOf 查找元素的索引
// 时间复杂度：O(n)
func (da *DynamicArray) IndexOf(element interface{}) int {
	for i := 0; i < da.size; i++ {
		if da.data[i] == element {
			return i
		}
	}
	return -1
}

// ToSlice 转换为切片
func (da *DynamicArray) ToSlice() []interface{} {
	result := make([]interface{}, da.size)
	copy(result, da.data[:da.size])
	return result
}

// String 字符串表示
func (da *DynamicArray) String() string {
	return fmt.Sprintf("DynamicArray{size: %d, capacity: %d, data: %v}", 
		da.size, da.capacity, da.data[:da.size])
}

// resize 扩容，容量翻倍
// 扩容策略：当前容量 * 2
func (da *DynamicArray) resize() {
	newCapacity := da.capacity * 2
	newData := make([]interface{}, newCapacity)
	copy(newData, da.data)
	da.data = newData
	da.capacity = newCapacity
}

// 业务应用示例：
// 1. 电商购物车：动态添加/删除商品
// 2. 消息队列：临时存储待处理消息
// 3. 缓存系统：存储热点数据，支持快速访问
// 4. 日志系统：收集日志条目，批量处理
// 5. 用户行为分析：记录用户操作序列