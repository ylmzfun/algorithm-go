package corealgo

import (
	"reflect"
	"testing"
)

func TestActivitySelection(t *testing.T) {
	cases := []struct {
		start []int
		end   []int
		want  int
	}{
		// 经典用例：最多可选 4 个活动
		{[]int{1, 3, 0, 5, 3, 5, 6, 8, 8, 2, 12}, []int{4, 5, 6, 7, 9, 9, 10, 11, 12, 13, 14}, 4},
		{[]int{1, 2, 3}, []int{2, 3, 4}, 3}, // 全部互不冲突
		{[]int{1, 2, 2}, []int{2, 2, 3}, 3}, // 结束时间相同的并列情况
		{[]int{5, 3, 1}, []int{6, 4, 2}, 3}, // 输入无序时也可正确排序选择
		{[]int{}, []int{}, 0},               // 空输入
	}
	for _, tc := range cases {
		got := ActivitySelection(tc.start, tc.end)
		if len(got) != tc.want {
			t.Errorf("ActivitySelection(%v, %v): Expected %d activities, got %d (%v)",
				tc.start, tc.end, tc.want, len(got), got)
			continue
		}
		// 校验所选活动按选择顺序互不冲突
		lastEnd := -1 << 31
		for _, idx := range got {
			if tc.start[idx] < lastEnd {
				t.Errorf("ActivitySelection(%v, %v): selected activities conflict at index %d",
					tc.start, tc.end, idx)
			}
			lastEnd = tc.end[idx]
		}
	}
}

func TestCanJump(t *testing.T) {
	cases := []struct {
		nums []int
		want bool
	}{
		{[]int{2, 3, 1, 1, 4}, true},
		{[]int{3, 2, 1, 0, 4}, false}, // 卡在下标 3 无法前进
		{[]int{0}, true},              // 已在末尾
		{[]int{1}, true},
		{[]int{0, 2, 3}, false}, // 第一步就无法出发
		{[]int{}, false},
	}
	for _, tc := range cases {
		got := CanJump(tc.nums)
		if got != tc.want {
			t.Errorf("CanJump(%v): Expected %v, got %v", tc.nums, tc.want, got)
		}
	}
}

func TestMinJump(t *testing.T) {
	cases := []struct {
		nums []int
		want int
	}{
		{[]int{2, 3, 1, 1, 4}, 2},
		{[]int{1, 2, 3}, 2},
		{[]int{1, 1, 1, 1}, 3},
		{[]int{2, 3, 0, 1, 4}, 2},
		{[]int{0}, 0},
		{[]int{1}, 0},
		{[]int{1, 2}, 1},
		{[]int{3, 2, 1, 0, 4}, -1}, // 无法到达末尾
	}
	for _, tc := range cases {
		got := MinJump(tc.nums)
		if got != tc.want {
			t.Errorf("MinJump(%v): Expected %d, got %d", tc.nums, tc.want, got)
		}
	}
}

func TestCoinChangeGreedy(t *testing.T) {
	cases := []struct {
		coins  []int
		amount int
		count  int
		ok     bool
	}{
		{[]int{1, 5, 10, 25}, 30, 2, true},  // 25 + 5
		{[]int{1, 5, 10, 25}, 93, 8, true},  // 3*25 + 10 + 5 + 3*1
		{[]int{1, 5, 10, 25}, 0, 0, true},   // 金额为 0 无需找零
		{[]int{1, 5, 10, 25}, -5, 0, false}, // 非法金额
		{[]int{2}, 3, 0, false},             // 无法恰好凑出
		{[]int{5, 10, 25, 1}, 30, 2, true},  // 输入无序也可正确求解
		// 贪心反例：[1,3,4] 凑 6，贪心得 4+1+1=3 枚，最优为 3+3=2 枚
		{[]int{1, 3, 4}, 6, 3, true},
	}
	for _, tc := range cases {
		orig := make([]int, len(tc.coins))
		copy(orig, tc.coins)
		count, ok := CoinChangeGreedy(tc.coins, tc.amount)
		if count != tc.count || ok != tc.ok {
			t.Errorf("CoinChangeGreedy(%v, %d): Expected (%d, %v), got (%d, %v)",
				tc.coins, tc.amount, tc.count, tc.ok, count, ok)
		}
		if !reflect.DeepEqual(tc.coins, orig) {
			t.Errorf("CoinChangeGreedy(%v, %d): input coins were modified, got %v",
				orig, tc.amount, tc.coins)
		}
	}
}
