package sort

// Sort 排序算法集合
// 包含常见的排序算法及其实现，每种算法包含详细的时间/空间复杂度分析和适用场景说明

import (
	"fmt"
)

// --- 冒泡排序 ---

// BubbleSort 冒泡排序
// 思路：重复遍历数组，比较相邻元素，如果顺序错误则交换，每次遍历将最大元素"冒泡"到末尾
// 时间复杂度：O(n²)，最好 O(n)（已排序且有提前终止优化）
// 空间复杂度：O(1)
// 稳定性：稳定
// 适用场景：小规模数据、近乎有序的数据
func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

// --- 选择排序 ---

// SelectionSort 选择排序
// 思路：每次从未排序部分找到最小元素，放到已排序部分的末尾
// 时间复杂度：O(n²)
// 空间复杂度：O(1)
// 稳定性：不稳定
// 适用场景：小规模数据，对交换次数敏感的场景
func SelectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

// --- 插入排序 ---

// InsertionSort 插入排序
// 思路：将未排序元素依次插入到已排序部分的正确位置
// 时间复杂度：O(n²)，最好 O(n)（已排序）
// 空间复杂度：O(1)
// 稳定性：稳定
// 适用场景：小规模数据、近乎有序的数据、在线排序（数据逐步到达）
func InsertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

// --- 希尔排序 ---

// ShellSort 希尔排序
// 思路：插入排序的改进版，通过比较相距一定间隔的元素，使数组逐渐有序
// 使用 Knuth 序列：gap = 3*gap + 1
// 时间复杂度：O(n log n) ~ O(n^(3/2))，取决于 gap 序列
// 空间复杂度：O(1)
// 稳定性：不稳定
// 适用场景：中等规模数据，需要原地排序且不想用快速排序的场景
func ShellSort(arr []int) {
	n := len(arr)
	// Knuth 序列
	gap := 1
	for gap < n/3 {
		gap = 3*gap + 1
	}
	for gap > 0 {
		for i := gap; i < n; i++ {
			temp := arr[i]
			j := i
			for j >= gap && arr[j-gap] > temp {
				arr[j] = arr[j-gap]
				j -= gap
			}
			arr[j] = temp
		}
		gap /= 3
	}
}

// --- 归并排序 ---

// MergeSort 归并排序（递归实现）
// 思路：分治法，将数组递归分成两半，分别排序后合并
// 时间复杂度：O(n log n)
// 空间复杂度：O(n)（需要辅助数组）
// 稳定性：稳定
// 适用场景：大规模数据排序，需要稳定排序时，外部排序
func MergeSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	mergeSortHelper(arr, 0, len(arr)-1)
}

func mergeSortHelper(arr []int, left, right int) {
	if left >= right {
		return
	}
	mid := left + (right-left)/2
	mergeSortHelper(arr, left, mid)
	mergeSortHelper(arr, mid+1, right)
	merge(arr, left, mid, right)
}

func merge(arr []int, left, mid, right int) {
	temp := make([]int, right-left+1)
	i, j, k := left, mid+1, 0
	for i <= mid && j <= right {
		if arr[i] <= arr[j] {
			temp[k] = arr[i]
			i++
		} else {
			temp[k] = arr[j]
			j++
		}
		k++
	}
	for i <= mid {
		temp[k] = arr[i]
		i++
		k++
	}
	for j <= right {
		temp[k] = arr[j]
		j++
		k++
	}
	copy(arr[left:right+1], temp)
}

// --- 快速排序 ---

// QuickSort 快速排序
// 思路：选取 pivot，将数组分为小于和大于 pivot 的两部分，递归排序
// 三数取中法选择 pivot，避免最坏情况
// 时间复杂度：平均 O(n log n)，最坏 O(n²)
// 空间复杂度：O(log n)（递归栈）
// 稳定性：不稳定
// 适用场景：通用排序的首选，大多数场景下表现最优
func QuickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	quickSortHelper(arr, 0, len(arr)-1)
}

func quickSortHelper(arr []int, low, high int) {
	for low < high {
		// 小数组使用插入排序
		if high-low < 10 {
			insertionSortRange(arr, low, high)
			return
		}
		pi := partition(arr, low, high)
		// 尾递归优化：先排序较小的部分
		if pi-low < high-pi {
			quickSortHelper(arr, low, pi-1)
			low = pi + 1
		} else {
			quickSortHelper(arr, pi+1, high)
			high = pi - 1
		}
	}
}

func insertionSortRange(arr []int, low, high int) {
	for i := low + 1; i <= high; i++ {
		key := arr[i]
		j := i - 1
		for j >= low && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func partition(arr []int, low, high int) int {
	// 三数取中法选择 pivot
	mid := low + (high-low)/2
	if arr[mid] < arr[low] {
		arr[low], arr[mid] = arr[mid], arr[low]
	}
	if arr[high] < arr[low] {
		arr[low], arr[high] = arr[high], arr[low]
	}
	if arr[high] < arr[mid] {
		arr[mid], arr[high] = arr[high], arr[mid]
	}
	// 将 pivot 放到 right-1 位置，arr[right] >= pivot 保证不越界
	arr[mid], arr[high-1] = arr[high-1], arr[mid]
	pivot := arr[high-1]

	i := low
	j := high - 1
	for i < j {
		for i++; arr[i] < pivot; i++ {
		}
		for j--; arr[j] > pivot; j-- {
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i], arr[high-1] = arr[high-1], arr[i]
	return i
}

// --- 堆排序 ---

// HeapSort 堆排序
// 思路：利用最大堆的特性，每次将堆顶（最大值）与末尾交换，然后调整堆
// 时间复杂度：O(n log n)
// 空间复杂度：O(1)
// 稳定性：不稳定
// 适用场景：需要原地 O(n log n) 排序且不要求稳定性时，Top K 问题
func HeapSort(arr []int) {
	n := len(arr)
	// 构建最大堆
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}
	// 逐个提取元素
	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, i, 0)
	}
}

func heapify(arr []int, n int, root int) {
	largest := root
	left := 2*root + 1
	right := 2*root + 2

	if left < n && arr[left] > arr[largest] {
		largest = left
	}
	if right < n && arr[right] > arr[largest] {
		largest = right
	}
	if largest != root {
		arr[root], arr[largest] = arr[largest], arr[root]
		heapify(arr, n, largest)
	}
}

// --- 计数排序 ---

// CountingSort 计数排序
// 思路：统计每个元素出现的次数，根据计数结果重建有序数组
// 时间复杂度：O(n + k)，k 为数据范围
// 空间复杂度：O(k)
// 稳定性：稳定
// 适用场景：数据范围较小的整数排序，如成绩、年龄等
func CountingSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	min, max := arr[0], arr[0]
	for _, v := range arr {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	rangeSize := max - min + 1
	count := make([]int, rangeSize)
	output := make([]int, len(arr))

	for _, v := range arr {
		count[v-min]++
	}
	for i := 1; i < rangeSize; i++ {
		count[i] += count[i-1]
	}
	for i := len(arr) - 1; i >= 0; i-- {
		idx := arr[i] - min
		count[idx]--
		output[count[idx]] = arr[i]
	}
	copy(arr, output)
}

// --- 基数排序 ---

// RadixSort 基数排序（LSD 最低位优先）
// 思路：从最低位开始，对每一位使用计数排序进行稳定排序
// 时间复杂度：O(d * (n + k))，d 为最大位数
// 空间复杂度：O(n + k)
// 稳定性：稳定
// 适用场景：整数或定长字符串排序，位数固定且范围小的场景
func RadixSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	output := make([]int, len(arr))
	for exp := 1; max/exp > 0; exp *= 10 {
		count := make([]int, 10)
		for _, v := range arr {
			count[(v/exp)%10]++
		}
		for i := 1; i < 10; i++ {
			count[i] += count[i-1]
		}
		for i := len(arr) - 1; i >= 0; i-- {
			digit := (arr[i] / exp) % 10
			count[digit]--
			output[count[digit]] = arr[i]
		}
		copy(arr, output)
	}
}

// --- 桶排序 ---

// BucketSort 桶排序
// 思路：将数据分到有限数量的桶中，对每个桶分别排序，最后合并
// 时间复杂度：平均 O(n + k)，最坏 O(n²)（所有元素落入同一桶）
// 空间复杂度：O(n + k)
// 稳定性：取决于桶内排序算法的稳定性
// 适用场景：数据均匀分布在某个范围内
func BucketSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	min, max := arr[0], arr[0]
	for _, v := range arr {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	bucketCount := len(arr)
	buckets := make([][]int, bucketCount)
	for i := range buckets {
		buckets[i] = make([]int, 0)
	}
	rangeSize := float64(max-min) + 1
	for _, v := range arr {
		idx := int(float64(v-min) / rangeSize * float64(bucketCount-1))
		buckets[idx] = append(buckets[idx], v)
	}
	idx := 0
	for _, bucket := range buckets {
		// 桶内使用插入排序
		InsertionSort(bucket)
		for _, v := range bucket {
			arr[idx] = v
			idx++
		}
	}
}

// --- 实用工具函数 ---

// IsSorted 检查数组是否已排序
func IsSorted(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			return false
		}
	}
	return true
}

// FindKthSmallest 使用快速选择算法找到第 K 小的元素
// 时间复杂度：平均 O(n)
func FindKthSmallest(arr []int, k int) int {
	if k < 1 || k > len(arr) {
		return -1
	}
	// 复制一份避免修改原数组
	cpy := make([]int, len(arr))
	copy(cpy, arr)
	return quickSelect(cpy, 0, len(cpy)-1, k-1)
}

func quickSelect(arr []int, low, high, k int) int {
	if low == high {
		return arr[low]
	}
	pi := partition(arr, low, high)
	if pi == k {
		return arr[pi]
	} else if pi > k {
		return quickSelect(arr, low, pi-1, k)
	}
	return quickSelect(arr, pi+1, high, k)
}

// Reverse 反转数组
func Reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

// FormatArray 格式化数组输出
func FormatArray(arr []int) string {
	if len(arr) == 0 {
		return "[]"
	}
	s := "["
	for i, v := range arr {
		if i > 0 {
			s += ", "
		}
		s += fmt.Sprintf("%d", v)
	}
	s += "]"
	return s
}
