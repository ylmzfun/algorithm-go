package corealgo

import (
	"reflect"
	"sort"
	"testing"
)

// validBoard 校验 N 皇后解是否合法：每行、每列、每条对角线至多一个皇后
func validBoard(board []string) bool {
	n := len(board)
	cols := make([]bool, n)
	diag1 := make([]bool, 2*n-1)
	diag2 := make([]bool, 2*n-1)
	for r := 0; r < n; r++ {
		if len(board[r]) != n {
			return false
		}
		rowCount := 0
		for c := 0; c < n; c++ {
			switch board[r][c] {
			case 'Q':
				rowCount++
				if cols[c] || diag1[r+c] || diag2[r-c+n-1] {
					return false
				}
				cols[c], diag1[r+c], diag2[r-c+n-1] = true, true, true
			case '.':
			default:
				return false
			}
		}
		if rowCount != 1 { // 每行应恰有一个皇后
			return false
		}
	}
	return true
}

// isPermutation 判断 perm 是否为 nums 的一个排列（多重集相等）
func isPermutation(perm, nums []int) bool {
	if len(perm) != len(nums) {
		return false
	}
	count := make(map[int]int)
	for _, x := range nums {
		count[x]++
	}
	for _, x := range perm {
		count[x]--
		if count[x] < 0 {
			return false
		}
	}
	return true
}

// hasDuplicateRows 判断矩阵中是否存在完全相同的两行
func hasDuplicateRows(m [][]int) bool {
	for i := 0; i < len(m); i++ {
		for j := i + 1; j < len(m); j++ {
			if reflect.DeepEqual(m[i], m[j]) {
				return true
			}
		}
	}
	return false
}

// normalizeIntMatrix 将矩阵标准化：每行升序，整体按长度与字典序排序
func normalizeIntMatrix(m [][]int) [][]int {
	for _, row := range m {
		sort.Ints(row)
	}
	sort.Slice(m, func(i, j int) bool {
		a, b := m[i], m[j]
		if len(a) != len(b) {
			return len(a) < len(b)
		}
		for k := range a {
			if a[k] != b[k] {
				return a[k] < b[k]
			}
		}
		return false
	})
	return m
}

func TestSolveNQueens(t *testing.T) {
	cases := []struct {
		n    int
		want int
	}{
		{0, 0},
		{1, 1},
		{2, 0},
		{3, 0},
		{4, 2},
		{8, 92},
	}
	for _, tc := range cases {
		got := SolveNQueens(tc.n)
		if len(got) != tc.want {
			t.Errorf("SolveNQueens(%d): Expected %d solutions, got %d", tc.n, tc.want, len(got))
			continue
		}
		for _, board := range got {
			if !validBoard(board) {
				t.Errorf("SolveNQueens(%d): invalid solution %v", tc.n, board)
			}
		}
	}
}

func TestPermute(t *testing.T) {
	cases := []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 3}, 6},
		{[]int{1, 1, 2}, 3}, // 含重复元素，去重后 3 种
		{[]int{1, 1, 1}, 1}, // 全相同，仅 1 种
		{[]int{2}, 1},
		{[]int{}, 0},
	}
	for _, tc := range cases {
		orig := make([]int, len(tc.nums))
		copy(orig, tc.nums)
		got := Permute(tc.nums)
		if len(got) != tc.want {
			t.Errorf("Permute(%v): Expected %d permutations, got %d", tc.nums, tc.want, len(got))
			continue
		}
		if !reflect.DeepEqual(tc.nums, orig) {
			t.Errorf("Permute(%v): input slice was modified, got %v", orig, tc.nums)
		}
		if hasDuplicateRows(got) {
			t.Errorf("Permute(%v): result contains duplicate permutations", orig)
		}
		for _, p := range got {
			if !isPermutation(p, orig) {
				t.Errorf("Permute(%v): %v is not a valid permutation", orig, p)
			}
		}
	}
}

func TestSubsets(t *testing.T) {
	cases := []struct {
		nums []int
		want [][]int
	}{
		{[]int{1, 2, 3}, [][]int{{}, {1}, {2}, {3}, {1, 2}, {1, 3}, {2, 3}, {1, 2, 3}}},
		{[]int{1, 2, 2}, [][]int{{}, {1}, {2}, {1, 2}, {2, 2}, {1, 2, 2}}}, // 去重后 6 个
		{[]int{1, 1}, [][]int{{}, {1}, {1, 1}}},
		{[]int{}, [][]int{{}}}, // 空输入只有空子集
	}
	for _, tc := range cases {
		orig := make([]int, len(tc.nums))
		copy(orig, tc.nums)
		got := Subsets(tc.nums)
		if !reflect.DeepEqual(tc.nums, orig) {
			t.Errorf("Subsets(%v): input slice was modified, got %v", orig, tc.nums)
		}
		gotNorm := normalizeIntMatrix(got)
		wantNorm := normalizeIntMatrix(tc.want)
		if !reflect.DeepEqual(gotNorm, wantNorm) {
			t.Errorf("Subsets(%v): Expected %v, got %v", orig, wantNorm, gotNorm)
		}
	}
}
