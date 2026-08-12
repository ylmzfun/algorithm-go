package corealgo

import (
	"testing"
)

func TestFibonacci(t *testing.T) {
	cases := []struct {
		n    int
		want int
	}{
		{-5, 0},
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{10, 55},
		{20, 6765},
		{40, 102334155},
	}
	for _, tc := range cases {
		got := Fibonacci(tc.n)
		if got != tc.want {
			t.Errorf("Fibonacci(%d): Expected %d, got %d", tc.n, tc.want, got)
		}
	}
}

func TestKnapsack01(t *testing.T) {
	cases := []struct {
		weights  []int
		values   []int
		capacity int
		want     int
	}{
		{[]int{1, 3, 4, 5}, []int{1, 4, 5, 7}, 7, 9},      // 选物品 1、2：4+5=9
		{[]int{2, 3, 4, 5}, []int{3, 4, 5, 6}, 5, 7},      // 选物品 0、1：3+4=7
		{[]int{10, 20, 30}, []int{60, 100, 120}, 50, 220}, // 经典用例：100+120
		{[]int{}, []int{}, 10, 0},                         // 无物品
		{[]int{5}, []int{10}, 0, 0},                       // 容量为 0
		{[]int{5}, []int{10}, 4, 0},                       // 物品放不下
		{[]int{1, 2}, []int{5}, 3, 0},                     // weights 与 values 长度不一致
	}
	for _, tc := range cases {
		got := Knapsack01(tc.weights, tc.values, tc.capacity)
		if got != tc.want {
			t.Errorf("Knapsack01(%v, %v, %d): Expected %d, got %d",
				tc.weights, tc.values, tc.capacity, tc.want, got)
		}
	}
}

func TestLCSCount(t *testing.T) {
	cases := []struct {
		text1 string
		text2 string
		want  int
	}{
		{"abcde", "ace", 3},
		{"abc", "abc", 3},
		{"abc", "def", 0},
		{"", "abc", 0},
		{"abc", "", 0},
		{"aaaa", "aa", 2},
		{"AGGTAB", "GXTXAYB", 4},
	}
	for _, tc := range cases {
		got := LCSCount(tc.text1, tc.text2)
		if got != tc.want {
			t.Errorf("LCSCount(%q, %q): Expected %d, got %d", tc.text1, tc.text2, tc.want, got)
		}
	}
}

func TestLISLength(t *testing.T) {
	cases := []struct {
		nums []int
		want int
	}{
		{[]int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
		{[]int{0, 1, 0, 3, 2, 3}, 4},
		{[]int{7, 7, 7, 7}, 1}, // 严格递增，重复元素只算一次
		{[]int{3, 10, 2, 1, 20}, 3},
		{[]int{3, 2}, 1},
		{[]int{}, 0},
		{[]int{5}, 1},
	}
	for _, tc := range cases {
		got := LISLength(tc.nums)
		if got != tc.want {
			t.Errorf("LISLength(%v): Expected %d, got %d", tc.nums, tc.want, got)
		}
	}
}
