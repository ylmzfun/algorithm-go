package advanced

import (
	"sync"
)

// goroutine 是 Go 语言并发编程的核心概念：
// 1. 使用 go 关键字启动，开销极小（初始栈仅几 KB 且按需增长），可轻松创建成千上万个
// 2. 启动是非阻塞的：go f() 立即返回，f 在新 goroutine 中由 Go 运行时调度执行
// 3. goroutine 由 Go 运行时调度（M:N 模型），并非直接对应操作系统线程
// 4. 多个 goroutine 共享同一地址空间，通信推荐使用 channel 而非共享内存
// 5. 主 goroutine（main 函数）退出时，其他 goroutine 会立即终止
//
// 本文件演示 WaitGroup 对并发任务的编排：Add 登记任务数、Done 标记完成、Wait 等待全部结束

// ConcurrentSum 使用多个 goroutine 并发求和
// 思路：将切片按 workerCount 分成若干块，每个 goroutine 负责一块的求和，
// 通过 WaitGroup 等待所有 worker 完成后，主 goroutine 汇总各块的局部和
// 作用：演示 WaitGroup 并发任务编排的标准模式
// 业务场景：大数据量求和、日志统计、指标聚合等"分而治之"型任务
// 时间复杂度：串行 O(n)，并行约 O(n/workerCount)
func ConcurrentSum(nums []int, workerCount int) int {
	if workerCount <= 0 {
		workerCount = 1
	}
	if len(nums) == 0 {
		return 0
	}
	if workerCount > len(nums) {
		workerCount = len(nums)
	}

	var wg sync.WaitGroup
	chunkSum := make([]int, workerCount)
	chunkSize := (len(nums) + workerCount - 1) / workerCount // 向上取整，保证每块尽量均匀

	for i := 0; i < workerCount; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if end > len(nums) {
			end = len(nums)
		}
		wg.Add(1)
		go func(id, start, end int) {
			defer wg.Done()
			sum := 0
			for _, v := range nums[start:end] {
				sum += v
			}
			chunkSum[id] = sum // 每个 worker 只写自己的槽位，无数据竞争
		}(i, start, end)
	}
	wg.Wait()

	total := 0
	for _, s := range chunkSum {
		total += s
	}
	return total
}

// ParallelMap 并行地对每个元素执行转换函数
// 思路：将数据分块，每个 goroutine 处理一块，结果写入预先分配好的切片对应槽位，
// 最后通过 WaitGroup 等待全部 worker 完成
// 作用：演示"并行处理 + 结果收集"模式
// 业务场景：批量图片压缩、批量文本清洗、批量数据脱敏等 CPU 密集型转换
// 时间复杂度：串行 O(n)，并行约 O(n/workerCount)
func ParallelMap(data []int, fn func(int) int, workerCount int) []int {
	if workerCount <= 0 {
		workerCount = 1
	}
	result := make([]int, len(data))
	if len(data) == 0 {
		return result
	}
	if workerCount > len(data) {
		workerCount = len(data)
	}

	var wg sync.WaitGroup
	chunkSize := (len(data) + workerCount - 1) / workerCount
	for i := 0; i < workerCount; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if end > len(data) {
			end = len(data)
		}
		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()
			for j := start; j < end; j++ {
				result[j] = fn(data[j]) // 不同 goroutine 写不同下标，无数据竞争
			}
		}(start, end)
	}
	wg.Wait()
	return result
}

// 业务应用示例：
// 1. 高并发请求的统计聚合（ConcurrentSum）
// 2. 大数据集的批量转换/清洗（ParallelMap）
// 3. 多协程任务编排（WaitGroup 等待所有子任务完成）
