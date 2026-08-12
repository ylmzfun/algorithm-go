package corealgo

// --- 分治算法 ---

// MaxSubArrayDivide 最大子数组和（分治解法）
// 思路：将数组从中间一分为二，最大子数组要么完全在左半、要么完全在右半、
// 要么跨越中点。递归求解左右两半，再从中点向两侧扩展线性求出跨中点的最大
// 和，三者取最大值即为答案
// 时间复杂度：O(n log n)
// 空间复杂度：O(log n)（递归栈）
// 另：本问题存在 Kadane 算法 O(n) 线性解（动态规划思想，单次遍历），
// 需要 O(n) 时优先使用 Kadane
// 适用场景：
// 1. 股票价格最大连续涨幅
// 2. 信号、序列中的最大连续能量段
// 3. 数据流中的最大连续收益
func MaxSubArrayDivide(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	return maxSubArrayDC(nums, 0, len(nums)-1)
}

// maxSubArrayDC 递归求解 nums[lo:hi+1] 的最大子数组和
func maxSubArrayDC(nums []int, lo, hi int) int {
	if lo == hi {
		return nums[lo]
	}
	mid := lo + (hi-lo)/2
	left := maxSubArrayDC(nums, lo, mid)
	right := maxSubArrayDC(nums, mid+1, hi)
	cross := maxCrossSum(nums, lo, mid, hi)
	if left >= right && left >= cross {
		return left
	}
	if right >= cross {
		return right
	}
	return cross
}

// maxCrossSum 求跨越中点的最大子数组和：从中点向两侧扩展取最大值并相加
func maxCrossSum(nums []int, lo, mid, hi int) int {
	leftSum := nums[mid]
	sum := 0
	for i := mid; i >= lo; i-- {
		sum += nums[i]
		if sum > leftSum {
			leftSum = sum
		}
	}
	rightSum := nums[mid+1]
	sum = 0
	for i := mid + 1; i <= hi; i++ {
		sum += nums[i]
		if sum > rightSum {
			rightSum = sum
		}
	}
	return leftSum + rightSum
}

// CountInversions 归并排序求逆序对数量
// 思路：在归并排序合并两个有序子数组时，若右半部分的元素先于左半部分元素
// 被取出，说明它比左半剩余的所有元素都小，这些元素与其构成逆序对，
// 累加 (mid-i+1) 即可。排序与计数在一次归并过程中同时完成
// 时间复杂度：O(n log n)，n 为数组长度
// 空间复杂度：O(n)（合并辅助数组）
// 适用场景：
// 1. 衡量数组接近有序的程度（逆序度）
// 2. 冒泡排序交换次数的事前估算
// 3. 推荐、榜单中元素错位程度的统计
// 注意：函数内部对入参副本排序，不修改原始数组
func CountInversions(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	arr := make([]int, len(nums))
	copy(arr, nums)
	tmp := make([]int, len(arr))
	count := 0
	var mergeCount func(lo, hi int)
	mergeCount = func(lo, hi int) {
		if lo >= hi {
			return
		}
		mid := lo + (hi-lo)/2
		mergeCount(lo, mid)
		mergeCount(mid+1, hi)
		i, j, k := lo, mid+1, lo
		for i <= mid && j <= hi {
			if arr[i] <= arr[j] {
				tmp[k] = arr[i]
				i++
			} else {
				tmp[k] = arr[j]
				count += mid - i + 1
				j++
			}
			k++
		}
		for i <= mid {
			tmp[k] = arr[i]
			i++
			k++
		}
		for j <= hi {
			tmp[k] = arr[j]
			j++
			k++
		}
		copy(arr[lo:hi+1], tmp[lo:hi+1])
	}
	mergeCount(0, len(arr)-1)
	return count
}
