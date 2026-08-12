package advanced

import (
	"errors"
	"time"
)

// select 用于在多个 channel 操作之间进行多路复用：
// 1. select 会阻塞，直到某个 case 的 channel 操作就绪
// 2. 若多个 case 同时就绪，select 会随机选择一个执行（Go 语言规范保证的均匀伪随机选择）
// 3. 常见用途：多 channel 竞争接收、超时控制（配合 time.After）、非阻塞尝试（配合 default）
// 4. 无 default 且所有 case 都不就绪时，select 会一直阻塞

// ErrTimeout select 超时错误
var ErrTimeout = errors.New("select timeout")

// SelectReceive 从两个 channel 中竞争接收一个值
// 思路：select 两个接收 case，谁先就绪就接收谁；若同时就绪则随机选择
// 作用：演示多 channel 竞争接收与 select 的随机选择特性
// 业务场景：多数据源（多路日志、多个上游服务）中取最先到达的结果
func SelectReceive(ch1, ch2 <-chan int) (value int, from int) {
	select {
	case v := <-ch1:
		return v, 1
	case v := <-ch2:
		return v, 2
	}
}

// SelectWithTimeout 带超时控制的接收
// 思路：select 同时监听数据 channel 与超时 channel（通常来自 time.After(d)），
// 谁先就绪就执行哪个 case
// 作用：演示 select 实现超时控制的标准模式
// 业务场景：网络请求超时、RPC 调用限时、任务执行 deadline
func SelectWithTimeout(ch <-chan int, timeout <-chan time.Time) (int, error) {
	select {
	case v := <-ch:
		return v, nil
	case <-timeout:
		return 0, ErrTimeout
	}
}

// SelectLoop 循环使用 select 从两个 channel 收集 count 个值
// 思路：在 for 循环中使用 select 多路接收；channel 关闭后将其置为 nil，
// 使该 case 不再参与 select（nil channel 永远不就绪）
// 作用：演示 select 在循环中的多路复用与 channel 关闭的处理
// 业务场景：合并多个数据流（多路传感器数据汇合、多队列消费）
func SelectLoop(ch1, ch2 <-chan int, count int) []int {
	result := make([]int, 0, count)
	for len(result) < count {
		select {
		case v, ok := <-ch1:
			if ok {
				result = append(result, v)
			} else {
				ch1 = nil // 已关闭：置 nil 禁用该 case
			}
		case v, ok := <-ch2:
			if ok {
				result = append(result, v)
			} else {
				ch2 = nil
			}
		}
		if ch1 == nil && ch2 == nil {
			break // 两个 channel 都已关闭，无法再收到数据
		}
	}
	return result
}

// TryReceive 非阻塞尝试接收：channel 无数据时立即返回而不阻塞
// 思路：select 的 default 分支在没有任何 case 就绪时立即执行
// 作用：演示 default 实现非阻塞操作
// 业务场景：轮询检查、优雅降级（数据未准备好时使用默认值）
func TryReceive(ch <-chan int) (value int, ok bool) {
	select {
	case v := <-ch:
		return v, true
	default:
		return 0, false
	}
}
