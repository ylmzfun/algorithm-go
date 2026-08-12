package corealgo

import (
	"reflect"
	"testing"
)

func TestMaxSubArrayDivide(t *testing.T) {
	cases := []struct {
		nums []int
		want int
	}{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6}, // 经典用例：4,-1,2,1
		{[]int{-1, -2, -3}, -1},                   // 全负数，取最大单个元素
		{[]int{5}, 5},                             // 单元素
		{[]int{1, 2, 3}, 6},                       // 全正数，取整个数组
		{[]int{-2, 1}, 1},
		{[]int{2, -1, 2}, 3},
		{[]int{}, 0}, // 空数组
	}
	for _, tc := range cases {
		got := MaxSubArrayDivide(tc.nums)
		if got != tc.want {
			t.Errorf("MaxSubArrayDivide(%v): Expected %d, got %d", tc.nums, tc.want, got)
		}
	}
}

func TestCountInversions(t *testing.T) {
	cases := []struct {
		nums []int
		want int
	}{
		{[]int{3, 1, 2}, 2},
		{[]int{5, 4, 3, 2, 1}, 10}, // 完全逆序
		{[]int{1, 2, 3, 4, 5}, 0},  // 完全有序
		{[]int{1, 3, 2}, 1},
		{[]int{2, 1, 2}, 1}, // 相等元素不算逆序对
		{[]int{}, 0},
		{[]int{1}, 0},
	}
	for _, tc := range cases {
		orig := make([]int, len(tc.nums))
		copy(orig, tc.nums)
		got := CountInversions(tc.nums)
		if got != tc.want {
			t.Errorf("CountInversions(%v): Expected %d, got %d", tc.nums, tc.want, got)
		}
		if !reflect.DeepEqual(tc.nums, orig) {
			t.Errorf("CountInversions(%v): input slice was modified, got %v", orig, tc.nums)
		}
	}
}
