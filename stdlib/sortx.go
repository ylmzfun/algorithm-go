package stdlib

// 思路：sort 提供基于比较的排序与二分查找通用算法，slices（Go 1.21 泛型）提供
// 针对切片类型的排序与查找工具。
// 作用：演示自定义排序、泛型排序、稳定排序与二分查找。
// 业务场景：
// 1. 排行榜按分数排序
// 2. 订单按金额/时间排序
// 3. 有序集合中快速查找（二分查找）

import (
	"slices"
	"sort"
)

// Person 人员，用于演示自定义排序
type Person struct {
	Name string
	Age  int
}

// SortByAgeAsc 按年龄升序排序（sort.Slice 示例，不稳定）
// 时间复杂度：O(n log n)
func SortByAgeAsc(people []Person) {
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
}

// SortByNameLen 按姓名长度升序排序（slices.SortFunc 示例）
// 时间复杂度：O(n log n)
func SortByNameLen(people []Person) {
	slices.SortFunc(people, func(a, b Person) int {
		return len(a.Name) - len(b.Name)
	})
}

// SortInts 整数切片升序排序（slices.Sort 示例）
func SortInts(nums []int) {
	slices.Sort(nums)
}

// StableSortByAge 按年龄稳定排序（sort.SliceStable 示例）
// 相同年龄的元素保持原始相对顺序
// 时间复杂度：O(n log n)
func StableSortByAge(people []Person) {
	sort.SliceStable(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
}

// SearchInt 在升序整数切片中二分查找 target
// 返回第一个等于 target 的下标；未找到返回 -1
// 时间复杂度：O(log n)
func SearchInt(sorted []int, target int) int {
	i := sort.Search(len(sorted), func(i int) bool {
		return sorted[i] >= target
	})
	if i < len(sorted) && sorted[i] == target {
		return i
	}
	return -1
}

// SearchInsertPosition 返回 target 应插入的位置（lower bound）
// 即第一个 >= target 的下标，插入后切片仍有序
// 时间复杂度：O(log n)
func SearchInsertPosition(sorted []int, target int) int {
	return sort.Search(len(sorted), func(i int) bool {
		return sorted[i] >= target
	})
}
