package stdlib

import (
	"reflect"
	"testing"
)

func TestSortByAgeAsc(t *testing.T) {
	people := []Person{{Name: "C", Age: 30}, {Name: "A", Age: 10}, {Name: "B", Age: 20}}
	SortByAgeAsc(people)
	if people[0].Age != 10 || people[1].Age != 20 || people[2].Age != 30 {
		t.Errorf("Expected ages sorted 10/20/30, got %v", people)
	}
}

func TestSortByNameLen(t *testing.T) {
	people := []Person{{Name: "abc", Age: 1}, {Name: "a", Age: 2}, {Name: "ab", Age: 3}}
	SortByNameLen(people)
	if people[0].Name != "a" || people[1].Name != "ab" || people[2].Name != "abc" {
		t.Errorf("Expected names sorted by length a/ab/abc, got %v", people)
	}
}

func TestSortInts(t *testing.T) {
	nums := []int{3, 1, 2}
	SortInts(nums)
	if !reflect.DeepEqual(nums, []int{1, 2, 3}) {
		t.Errorf("Expected [1 2 3], got %v", nums)
	}
}

func TestStableSortByAge(t *testing.T) {
	people := []Person{
		{Name: "张三", Age: 25},
		{Name: "李四", Age: 20},
		{Name: "王五", Age: 25},
		{Name: "赵六", Age: 20},
	}
	StableSortByAge(people)

	// 相同年龄者保持原始相对顺序：李四先于赵六，张三先于王五
	want := []string{"李四", "赵六", "张三", "王五"}
	for i, name := range want {
		if people[i].Name != name {
			t.Errorf("At index %d, expected %s, got %s (stability broken)", i, name, people[i].Name)
		}
	}
}

func TestSearchInt(t *testing.T) {
	sorted := []int{1, 3, 3, 5, 7, 9}

	// 命中
	if got := SearchInt(sorted, 5); got != 3 {
		t.Errorf("Expected index 3, got %d", got)
	}
	// 重复元素返回第一个
	if got := SearchInt(sorted, 3); got != 1 {
		t.Errorf("Expected index 1, got %d", got)
	}
	// 未命中
	if got := SearchInt(sorted, 4); got != -1 {
		t.Errorf("Expected -1, got %d", got)
	}
	// 小于所有元素 / 大于所有元素
	if got := SearchInt(sorted, 0); got != -1 {
		t.Errorf("Expected -1, got %d", got)
	}
	if got := SearchInt(sorted, 10); got != -1 {
		t.Errorf("Expected -1, got %d", got)
	}
	// 空切片
	if got := SearchInt([]int{}, 1); got != -1 {
		t.Errorf("Expected -1 for empty slice, got %d", got)
	}
}

func TestSearchInsertPosition(t *testing.T) {
	sorted := []int{1, 3, 5, 7}

	if got := SearchInsertPosition(sorted, 4); got != 2 {
		t.Errorf("Expected 2, got %d", got)
	}
	if got := SearchInsertPosition(sorted, 1); got != 0 {
		t.Errorf("Expected 0, got %d", got)
	}
	if got := SearchInsertPosition(sorted, 8); got != 4 {
		t.Errorf("Expected 4, got %d", got)
	}
	if got := SearchInsertPosition(sorted, 0); got != 0 {
		t.Errorf("Expected 0, got %d", got)
	}
	if got := SearchInsertPosition([]int{}, 5); got != 0 {
		t.Errorf("Expected 0 for empty slice, got %d", got)
	}
}
