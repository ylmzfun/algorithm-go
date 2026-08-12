package corealgo

import "sort"

// --- 贪心算法 ---

// ActivitySelection 活动选择问题：返回最多可选活动的原始下标列表（按选择顺序）
// 思路：按结束时间从小到大排序（结束越早，剩余可用时间越多），贪心地选择
// 第一个结束的活动，之后依次选择开始时间不早于上一个选中活动结束时间的活动。
// 该贪心策略可证明得到全局最优解
// 时间复杂度：O(n log n)（排序主导），n 为活动个数
// 空间复杂度：O(n)（排序用的下标数组）
// 适用场景：
// 1. 会议室/教室资源排期
// 2. 任务调度中最大化可完成的任务数量
// 3. 档期、时间段冲突最少的排班
func ActivitySelection(start, end []int) []int {
	n := len(start)
	if n == 0 || len(end) != n {
		return nil
	}
	order := make([]int, n)
	for i := range order {
		order[i] = i
	}
	// 按结束时间排序；结束时间相同时按开始时间排序，保证结果确定
	sort.Slice(order, func(i, j int) bool {
		a, b := order[i], order[j]
		if end[a] != end[b] {
			return end[a] < end[b]
		}
		return start[a] < start[b]
	})
	result := []int{order[0]}
	lastEnd := end[order[0]]
	for i := 1; i < n; i++ {
		idx := order[i]
		if start[idx] >= lastEnd {
			result = append(result, idx)
			lastEnd = end[idx]
		}
	}
	return result
}

// CanJump 跳跃游戏：判断能否从下标 0 跳到最后一个下标
// 思路：维护当前能到达的最远下标 reach，从左到右遍历，若当前下标超出
// reach 则说明中间断档无法继续；否则用 nums[i]+i 不断更新 reach
// 时间复杂度：O(n)，n 为数组长度
// 空间复杂度：O(1)
// 适用场景：
// 1. 资源/体力限制下判断路径是否可达
// 2. 跳跃类游戏的可达性判定
// 3. 网络转发中可达性检查
func CanJump(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	reach := 0
	for i := 0; i < len(nums); i++ {
		if i > reach {
			return false
		}
		if i+nums[i] > reach {
			reach = i + nums[i]
		}
	}
	return true
}

// MinJump 跳跃游戏 II：返回跳到最后一个下标所需的最少跳跃次数
// 思路：贪心。在当前可跳范围内遍历并记录能到达的最远位置，当遍历到当前
// 范围的边界时执行一次跳跃，并把范围扩展到记录的最远位置，直到覆盖终点
// 时间复杂度：O(n)
// 空间复杂度：O(1)
// 适用场景：
// 1. 最少中转次数问题（网络路由跳数最少、换乘最少）
// 2. 游戏中的最短跳跃路径
// 注意：无法到达末尾时返回 -1
func MinJump(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}
	jumps, curEnd, farthest := 0, 0, 0
	for i := 0; i < n-1; i++ {
		if i+nums[i] > farthest {
			farthest = i + nums[i]
		}
		if i == curEnd {
			jumps++
			curEnd = farthest
			if curEnd >= n-1 {
				return jumps
			}
		}
	}
	return -1
}

// CoinChangeGreedy 零钱找零（贪心）：返回凑出 amount 所需的最少硬币数及是否可行
// 思路：将硬币面额降序排列，每次优先使用面额最大的硬币，尽可能多用
// 贪心适用条件：硬币面额系统需满足"贪心最优"性质（canonical coin system），
// 经典例子：[1, 5, 10, 25]（美分）、人民币 [1, 2, 5, 10, 20, 50, 100]。
// 反例：[1, 3, 4] 凑 6 时贪心给出 4+1+1=3 枚，而最优解是 3+3=2 枚；
// 此类面额系统应改用动态规划（见 dp.go 的背包思想）求解
// 时间复杂度：O(n log n)（排序）+ O(n)（遍历），n 为硬币面额种类
// 空间复杂度：O(n)
// 适用场景：
// 1. 收银台、自动售货机、售票机找零
// 2. 零钱兑换类问题的快速近似求解
// 返回：count 为最少硬币数，ok 表示是否恰好凑出 amount（凑不出时返回 0, false）
func CoinChangeGreedy(coins []int, amount int) (int, bool) {
	if amount < 0 {
		return 0, false
	}
	if amount == 0 {
		return 0, true
	}
	sorted := make([]int, len(coins))
	copy(sorted, coins)
	sort.Sort(sort.Reverse(sort.IntSlice(sorted)))
	count := 0
	rest := amount
	for _, c := range sorted {
		if c <= 0 {
			continue
		}
		count += rest / c
		rest %= c
	}
	if rest != 0 {
		return 0, false
	}
	return count, true
}
