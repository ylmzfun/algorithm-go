// Package corealgo 核心算法集合
// 包含动态规划、贪心、回溯、分治、字符串算法、数论等经典算法的 Go 实现
// 每个算法均给出算法思路、时间复杂度、空间复杂度及典型业务场景说明
package corealgo

// --- 动态规划 ---

// Fibonacci 斐波那契数列第 n 项
// 思路：f(0)=0，f(1)=1，f(n)=f(n-1)+f(n-2)。使用滚动数组（两个变量）
// 滚动保存前两项结果，无需维护整个 dp 数组
// 时间复杂度：O(n)
// 空间复杂度：O(1)
// 适用场景：
// 1. 上楼梯问题（每次一阶或两阶）
// 2. 兔子繁殖、细胞分裂数量估算
// 3. 增长模型的递推计算
// 注意：n 较大时结果会溢出 int，请改用大数或取模运算
func Fibonacci(n int) int {
	if n < 0 {
		return 0
	}
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

// Knapsack01 0-1 背包问题：返回容量 capacity 内能装入的最大总价值
// 思路：dp[c] 表示容量为 c 时能获得的最大价值。对每个物品倒序更新容量，
// 倒序保证每个物品最多被选择一次（避免同一物品被重复使用）
// 状态转移：dp[c] = max(dp[c], dp[c-w[i]] + v[i])
// 时间复杂度：O(n*capacity)，n 为物品个数
// 空间复杂度：O(capacity)（一维滚动数组；二维写法为 O(n*capacity)）
// 适用场景：
// 1. 预算有限下的资源分配（采购、广告投放）
// 2. 行李箱/背包容量规划
// 3. 资源受限的任务收益最大化
// 注意：weights 与 values 长度不一致时返回 0
func Knapsack01(weights, values []int, capacity int) int {
	if len(weights) != len(values) {
		return 0
	}
	dp := make([]int, capacity+1)
	for i := 0; i < len(weights); i++ {
		w, v := weights[i], values[i]
		for c := capacity; c >= w; c-- {
			if dp[c-w]+v > dp[c] {
				dp[c] = dp[c-w] + v
			}
		}
	}
	return dp[capacity]
}

// LCSCount 最长公共子序列（LCS）长度
// 思路：dp[i][j] 表示 text1 前 i 个字符与 text2 前 j 个字符的 LCS 长度。
// 若 text1[i-1]==text2[j-1] 则 dp[i][j]=dp[i-1][j-1]+1，
// 否则取 dp[i-1][j] 与 dp[i][j-1] 的较大值
// 时间复杂度：O(m*n)，m、n 为两字符串长度
// 空间复杂度：O(m*n)，可进一步压缩为 O(min(m,n))
// 适用场景：
// 1. 文本相似度比对（diff、查重、代码抄袭检测）
// 2. 基因/DNA 序列比对
// 3. 版本差异分析
func LCSCount(text1, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else if dp[i-1][j] >= dp[i][j-1] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	return dp[m][n]
}

// LISLength 最长递增子序列（严格递增）长度，O(n log n) 版本
// 思路：维护 tails 数组，tails[k] 表示长度为 k+1 的递增子序列的最小末尾值。
// 对每个元素用二分查找（lower_bound）确定其在 tails 中的插入位置：
// 若能追加到末尾则序列变长，否则替换对应位置，从而保证后续可接出更长的序列
// 时间复杂度：O(n log n)，n 为数组长度
// 空间复杂度：O(n)
// 适用场景：
// 1. 最长上升股票收益序列
// 2. 排队、任务依赖链的最长长度
// 3. 数据流中单调递增趋势段的统计
func LISLength(nums []int) int {
	tails := make([]int, 0, len(nums))
	for _, x := range nums {
		lo, hi := 0, len(tails)
		for lo < hi {
			mid := lo + (hi-lo)/2
			if tails[mid] < x {
				lo = mid + 1
			} else {
				hi = mid
			}
		}
		if lo == len(tails) {
			tails = append(tails, x)
		} else {
			tails[lo] = x
		}
	}
	return len(tails)
}
