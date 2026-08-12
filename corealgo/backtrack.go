package corealgo

import (
	"sort"
	"strings"
)

// --- 回溯算法 ---

// SolveNQueens N 皇后问题：返回 n 皇后问题的所有合法解
// 思路：逐行放置皇后，用 cols、diag1、diag2 三个标记数组判断当前列与两条
// 对角线是否冲突；无冲突则放置皇后并递归下一行，回溯时撤销标记
// 时间复杂度：O(n!)，实际因剪枝远小于此
// 空间复杂度：O(n)（标记数组 + 递归栈，不含答案存储）
// 适用场景：
// 1. 八皇后类棋盘约束问题
// 2. 排课/排班中的冲突规避（行列、对角线约束建模）
// 3. 数独等约束满足问题（CSP）的求解模板
// 返回：棋盘字符串列表，'Q' 表示皇后，'.' 表示空位
func SolveNQueens(n int) [][]string {
	result := make([][]string, 0)
	if n <= 0 {
		return result
	}
	board := make([][]byte, n)
	for i := range board {
		board[i] = []byte(strings.Repeat(".", n))
	}
	cols := make([]bool, n)
	diag1 := make([]bool, 2*n-1) // 主对角线方向：r+c
	diag2 := make([]bool, 2*n-1) // 副对角线方向：r-c+n-1
	var backtrack func(r int)
	backtrack = func(r int) {
		if r == n {
			sol := make([]string, n)
			for i := range board {
				sol[i] = string(board[i])
			}
			result = append(result, sol)
			return
		}
		for c := 0; c < n; c++ {
			if cols[c] || diag1[r+c] || diag2[r-c+n-1] {
				continue
			}
			cols[c], diag1[r+c], diag2[r-c+n-1] = true, true, true
			board[r][c] = 'Q'
			backtrack(r + 1)
			board[r][c] = '.'
			cols[c], diag1[r+c], diag2[r-c+n-1] = false, false, false
		}
	}
	backtrack(0)
	return result
}

// Permute 全排列（自动去重）
// 思路：选择法 + 回溯。先将数组排序使相同元素相邻；回溯时若当前元素与
// 前一个元素相同且前一个未被使用，则跳过该分支，避免产生重复排列
// 时间复杂度：O(n*n!)，n 为数组长度（排列总数约 n!）
// 空间复杂度：O(n)（递归栈 + 标记数组，不含答案存储）
// 适用场景：
// 1. 密码、数字组合的穷举
// 2. 排班、比赛对阵的全排列枚举
// 3. 旅行商等排列类搜索的枚举基础
// 注意：nums 含重复元素时结果不会出现重复排列；内部对副本排序，不修改入参
func Permute(nums []int) [][]int {
	result := make([][]int, 0)
	if len(nums) == 0 {
		return result
	}
	vals := make([]int, len(nums))
	copy(vals, nums)
	sort.Ints(vals)
	used := make([]bool, len(vals))
	cur := make([]int, 0, len(vals))
	var backtrack func()
	backtrack = func() {
		if len(cur) == len(vals) {
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			result = append(result, tmp)
			return
		}
		for i := 0; i < len(vals); i++ {
			if used[i] {
				continue
			}
			if i > 0 && vals[i] == vals[i-1] && !used[i-1] {
				continue
			}
			used[i] = true
			cur = append(cur, vals[i])
			backtrack()
			cur = cur[:len(cur)-1]
			used[i] = false
		}
	}
	backtrack()
	return result
}

// Subsets 子集枚举（自动去重）
// 思路：回溯 + 组合式枚举。每层从当前下标向后选择元素组成子集；
// 排序后跳过与前一个位置相同的元素，保证含重复元素的输入不产生重复子集
// 时间复杂度：O(n*2^n)，n 为数组长度
// 空间复杂度：O(n)（递归深度，不含答案存储）
// 适用场景：
// 1. 商品组合、套餐搭配枚举
// 2. 特征选择中的子集搜索
// 3. 数据挖掘中的项集枚举（频繁项集挖掘）
// 注意：内部对副本排序，不修改入参；返回的每个子集内部元素升序
func Subsets(nums []int) [][]int {
	result := make([][]int, 0)
	vals := make([]int, len(nums))
	copy(vals, nums)
	sort.Ints(vals)
	cur := make([]int, 0)
	var backtrack func(idx int)
	backtrack = func(idx int) {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		result = append(result, tmp)
		for i := idx; i < len(vals); i++ {
			if i > idx && vals[i] == vals[i-1] {
				continue
			}
			cur = append(cur, vals[i])
			backtrack(i + 1)
			cur = cur[:len(cur)-1]
		}
	}
	backtrack(0)
	return result
}
