package advanced

import (
	"context"
	"time"
)

// context（上下文）是 Go 并发编程中传递"请求级信息"的标准方式：
// 1. Go 惯例：context 必须作为函数的第一个参数，参数名约定为 ctx
// 2. 作用：传递取消信号、超时/截止时间，以及少量请求级键值数据（如请求 ID、trace ID）
// 3. 原则：不要将 context 存进结构体；不要传 nil context（不确定时用 context.Background()）
// 4. 取消可传递：父 context 取消时，所有派生（子）context 一并取消
// 5. 使用 context 的函数应响应 ctx.Done()，实现协作式取消（取消是"请求"而非"强制终止"）

// requestIDKey 用于在 context 中存取请求 ID 的私有键类型
// （context 要求键使用自定义类型，避免不同包间键冲突）
type requestIDKey struct{}

// ProcessItems 逐个处理字符串列表，响应 context 取消
// 思路：每处理一项前检查 ctx.Err()，检测到取消立即返回已处理的部分结果
// 作用：演示协作式取消——任务在处理间隙"查看"取消信号
// 业务场景：批量任务处理（批量发邮件、批量导入）支持提前终止
// 时间复杂度：O(n)，n 为处理的元素个数
func ProcessItems(ctx context.Context, items []string) []string {
	result := make([]string, 0, len(items))
	for _, item := range items {
		if ctx.Err() != nil {
			// 检测到取消（Canceled 或 DeadlineExceeded），提前返回已处理部分
			return result
		}
		result = append(result, "processed:"+item)
	}
	return result
}

// WithRequestID 将请求 ID 写入 context 并返回新的 context
// 思路：context.WithValue 以键值对形式携带数据，键使用私有类型避免冲突
// 作用：演示 WithValue 传值（请求级元数据在调用链中的传递）
// 业务场景：日志关联（trace ID）、鉴权信息透传、网关透传请求头
func WithRequestID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, requestIDKey{}, id)
}

// GetRequestID 从 context 中读取请求 ID
// 思路：使用与写入时相同的键查询，类型断言失败时返回空串
// 作用：演示从 context 取值的安全写法
// 业务场景：日志中间件取出 trace ID、鉴权中间件取出用户 ID
func GetRequestID(ctx context.Context) string {
	if id, ok := ctx.Value(requestIDKey{}).(string); ok {
		return id
	}
	return ""
}

// CancelPropagation 演示取消传播：取消父 context，子 context 同步取消
// 思路：WithCancel 从父 context 派生子 context；父 cancel 后，子 ctx.Done() 被关闭
// 作用：演示取消信号的层级传播
// 业务场景：HTTP 服务优雅退出、分布式任务整体取消
func CancelPropagation() (childErr error) {
	parent, cancelParent := context.WithCancel(context.Background())
	defer cancelParent()

	child, cancelChild := context.WithCancel(parent)
	defer cancelChild()

	// 启动 goroutine 等待子 context 的取消信号
	done := make(chan struct{})
	go func() {
		<-child.Done() // 父取消后此处立即返回
		close(done)
	}()

	cancelParent() // 取消父 context

	<-done // 确定性同步：等待子 context 的 Done 被关闭（channel 同步，非 sleep）
	return child.Err()
}

// RunWithTimeout 演示 WithTimeout：任务超过时限后自动取消
// 思路：WithTimeout 内部启动计时器，到期自动取消 context；任务内检查 ctx.Err() 感知超时
// 作用：演示为任务设置执行时限（deadline）的标准方式
// 业务场景：HTTP 请求超时、数据库查询超时、外部服务调用限时
func RunWithTimeout(timeout time.Duration, process func(ctx context.Context) []string) (results []string, timedOut bool) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel() // 及时释放计时器资源

	results = process(ctx)
	return results, ctx.Err() != nil
}
