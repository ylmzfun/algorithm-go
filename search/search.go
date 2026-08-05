package search

// Search 搜索算法集合
// 包含常见的搜索算法及其实现

// --- 线性搜索 ---

// LinearSearch 线性搜索
// 思路：从数组头到尾逐个比较，找到目标元素返回索引
// 时间复杂度：O(n)
// 空间复杂度：O(1)
// 适用场景：无序数组、小规模数据
func LinearSearch(arr []int, target int) int {
	for i, v := range arr {
		if v == target {
			return i
		}
	}
	return -1
}

// LinearSearchAll 找出所有匹配的索引
func LinearSearchAll(arr []int, target int) []int {
	result := make([]int, 0)
	for i, v := range arr {
		if v == target {
			result = append(result, i)
		}
	}
	return result
}

// --- 二分搜索 ---

// BinarySearch 二分搜索（迭代实现）
// 思路：每次将搜索范围减半，与中间元素比较后决定在左半部分还是右半部分继续搜索
// 前提：数组必须有序
// 时间复杂度：O(log n)
// 空间复杂度：O(1)
// 适用场景：有序数组的快速查找
func BinarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

// BinarySearchRecursive 二分搜索（递归实现）
func BinarySearchRecursive(arr []int, target int, left, right int) int {
	if left > right {
		return -1
	}
	mid := left + (right-left)/2
	if arr[mid] == target {
		return mid
	} else if arr[mid] < target {
		return BinarySearchRecursive(arr, target, mid+1, right)
	}
	return BinarySearchRecursive(arr, target, left, mid-1)
}

// BinarySearchFirst 查找第一次出现的位置（下界）
// 适用场景：有序数组中存在重复元素时，找到第一个等于 target 的索引
func BinarySearchFirst(arr []int, target int) int {
	left, right := 0, len(arr)-1
	result := -1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			result = mid
			right = mid - 1 // 继续在左侧搜索
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return result
}

// BinarySearchLast 查找最后一次出现的位置（上界）
func BinarySearchLast(arr []int, target int) int {
	left, right := 0, len(arr)-1
	result := -1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			result = mid
			left = mid + 1 // 继续在右侧搜索
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return result
}

// LowerBound 查找第一个 >= target 的位置
func LowerBound(arr []int, target int) int {
	left, right := 0, len(arr)
	for left < right {
		mid := left + (right-left)/2
		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

// UpperBound 查找第一个 > target 的位置
func UpperBound(arr []int, target int) int {
	left, right := 0, len(arr)
	for left < right {
		mid := left + (right-left)/2
		if arr[mid] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

// --- 跳跃搜索 ---

// JumpSearch 跳跃搜索
// 思路：将数组分成大小为 sqrt(n) 的块，先通过跳跃找到目标所在的块，再在块内线性搜索
// 前提：数组必须有序
// 时间复杂度：O(√n)
// 空间复杂度：O(1)
// 适用场景：有序数组，且二分搜索的常数因子较大时（如数据存储在慢速介质上）
func JumpSearch(arr []int, target int) int {
	n := len(arr)
	if n == 0 {
		return -1
	}
	step := 1
	for step*step < n {
		step++
	}
	prev := 0
	for arr[min(step, n)-1] < target {
		prev = step
		step += step
		if prev >= n {
			return -1
		}
	}
	for i := prev; i < min(step, n); i++ {
		if arr[i] == target {
			return i
		}
		if arr[i] > target {
			break
		}
	}
	return -1
}

// --- 插值搜索 ---

// InterpolationSearch 插值搜索
// 思路：根据目标值的大小估算其在数组中的大致位置（类似在字典中查单词）
// 前提：数组必须有序且均匀分布
// 时间复杂度：平均 O(log log n)，最坏 O(n)
// 空间复杂度：O(1)
// 适用场景：大规模有序且均匀分布的数据（如字典、电话簿）
func InterpolationSearch(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left <= right && target >= arr[left] && target <= arr[right] {
		if left == right {
			if arr[left] == target {
				return left
			}
			return -1
		}
		// 插值公式
		pos := left + ((right-left)/(arr[right]-arr[left]))*(target-arr[left])
		if pos < left || pos > right {
			return -1
		}
		if arr[pos] == target {
			return pos
		} else if arr[pos] < target {
			left = pos + 1
		} else {
			right = pos - 1
		}
	}
	return -1
}

// --- 指数搜索 ---

// ExponentialSearch 指数搜索
// 思路：扩展搜索范围（指数增长），找到超过目标值的范围后使用二分搜索
// 前提：数组必须有序
// 时间复杂度：O(log n)
// 空间复杂度：O(1)
// 适用场景：目标值在数组前面的概率较大时，或数组极大且大小未知时（无界搜索）
func ExponentialSearch(arr []int, target int) int {
	n := len(arr)
	if n == 0 {
		return -1
	}
	if arr[0] == target {
		return 0
	}
	// 以指数方式扩展范围
	bound := 1
	for bound < n && arr[bound] < target {
		bound *= 2
	}
	// 在确定的范围内二分搜索
	left := bound / 2
	right := bound
	if right >= n {
		right = n - 1
	}
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

// --- 斐波那契搜索 ---

// FibonacciSearch 斐波那契搜索
// 思路：使用斐波那契数列确定分割点，每次将搜索范围按斐波那契比例划分
// 前提：数组必须有序
// 时间复杂度：O(log n)
// 空间复杂度：O(1)
// 适用场景：有序数组，且二分搜索的除法操作开销较大时（在 CPU 不支持高效除法的场景）
func FibonacciSearch(arr []int, target int) int {
	n := len(arr)
	if n == 0 {
		return -1
	}
	// 生成斐波那契数列
	fib2 := 0   // fib(k-2)
	fib1 := 1   // fib(k-1)
	fib := fib1 + fib2 // fib(k)
	for fib < n {
		fib2 = fib1
		fib1 = fib
		fib = fib1 + fib2
	}
	offset := -1
	for fib > 1 {
		i := offset + fib2
		if i >= n {
			i = n - 1
		}
		if i < n && i >= 0 && arr[i] < target {
			fib = fib1
			fib1 = fib2
			fib2 = fib - fib1
			offset = i
		} else if i < n && i >= 0 && arr[i] > target {
			fib = fib2
			fib1 = fib1 - fib2
			fib2 = fib - fib1
		} else if i < n && i >= 0 && arr[i] == target {
			return i
		} else {
			return -1
		}
	}
	// 检查最后一个元素
	if fib1 == 1 && offset+1 < n && arr[offset+1] == target {
		return offset + 1
	}
	return -1
}

// --- 三路二分搜索 ---

// TernarySearch 三分搜索
// 思路：每次将数组分成三部分，通过两个中间点缩小范围
// 前提：数组必须有序
// 时间复杂度：O(log₃ n)
// 空间复杂度：O(1)
// 适用场景：在特定函数（凸函数）上找极值点
func TernarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid1 := left + (right-left)/3
		mid2 := right - (right-left)/3
		if arr[mid1] == target {
			return mid1
		}
		if arr[mid2] == target {
			return mid2
		}
		if target < arr[mid1] {
			right = mid1 - 1
		} else if target > arr[mid2] {
			left = mid2 + 1
		} else {
			left = mid1 + 1
			right = mid2 - 1
		}
	}
	return -1
}
