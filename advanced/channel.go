package advanced

// channel（通道）是 goroutine 之间通信的主要方式，遵循"通过通信共享内存"的思想：
// 1. 无缓冲 channel：发送与接收必须同时就绪，否则发送方/接收方会阻塞，形成同步握手
// 2. 有缓冲 channel：缓冲区未满时发送不阻塞，缓冲区非空时接收不阻塞
// 3. close 关闭 channel 后，接收方仍可读走缓冲区中剩余的值，读完后再接收得到零值
// 4. 单向 channel（chan<- 只写 / <-chan 只读）用于在 API 边界表达意图，防止误用

// Produce 生产者函数：将 nums 依次发送到 channel，发送完毕后 close
// 思路：在函数内部创建 channel 并启动 goroutine 发送，返回时收缩为只读单向 channel
// 作用：演示单向 channel（<-chan）的用途——调用方只能接收，不能发送或关闭
// 业务场景：生产者-消费者模型中的生产者端（日志采集、任务分发）
func Produce(nums []int) <-chan int {
	ch := make(chan int)
	go func() {
		for _, n := range nums {
			ch <- n // 无缓冲 channel：每次发送都要等接收方就绪
		}
		close(ch) // 发送完毕必须 close，接收方才能结束 range
	}()
	return ch
}

// ConsumeAll 消费者函数：通过 range 从 channel 读取所有元素，直到 channel 被 close
// 思路：range ch 自动感知 channel 的关闭，读取完所有值后退出循环
// 作用：演示 close + range 遍历 channel 的标准模式
// 业务场景：消费生产者产出的数据流（日志处理、消息消费）
// 时间复杂度：O(n)，n 为接收的元素个数
func ConsumeAll(ch <-chan int) []int {
	result := make([]int, 0)
	for v := range ch {
		result = append(result, v)
	}
	return result
}

// BufferedProduce 有缓冲 channel 的生产者：将 nums 全部写入容量为 capacity 的缓冲 channel，然后 close
// 思路：容量足够的缓冲 channel 允许生产者在消费者尚未就绪时先行写入，写入不阻塞
// 作用：演示有缓冲 channel 的"异步缓冲"特性，解耦生产与消费的节奏
// 注意：为便于演示，当 capacity 小于元素个数时自动扩容为 len(nums)；
// 真实场景中应谨慎设计缓冲容量，缓冲写满时发送方会被阻塞
// 业务场景：任务队列、事件缓冲、流量削峰
func BufferedProduce(nums []int, capacity int) <-chan int {
	if capacity < len(nums) {
		capacity = len(nums)
	}
	ch := make(chan int, capacity)
	for _, n := range nums {
		ch <- n // 缓冲区未满时不会阻塞，因此无需 goroutine
	}
	close(ch)
	return ch
}

// UnbufferedHandshake 演示无缓冲 channel 的同步握手特性
// 思路：两个 goroutine 通过同一个无缓冲 channel 互相交换数据，
// 每次发送必须等到对方接收，每次接收必须等到对方发送
// 作用：演示无缓冲 channel 的同步语义——发送与接收互为对方"就绪"的条件
// 业务场景：两个协程之间的精确同步、回合制协作
func UnbufferedHandshake() (int, int) {
	ch := make(chan int)
	var a, b int

	done := make(chan struct{})
	go func() {
		ch <- 1     // 发送 1，阻塞直到主 goroutine 接收
		a = <-ch    // 接收 2
		close(done) // 通知主 goroutine 本协程已完成
	}()

	b = <-ch // 主 goroutine 先接收 1
	ch <- 2  // 再发送 2
	<-done   // 等协程完成，保证 a 已被赋值（channel 同步，非 sleep）
	return a, b
}
