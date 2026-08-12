package advanced

import (
	"fmt"
)

// defer / panic / recover 是 Go 的错误与清理机制：
// 1. defer 延迟执行：函数返回前执行注册的函数，多个 defer 按"后进先出"（LIFO）顺序执行
// 2. defer 常见用途：释放资源（关闭文件/连接）、解锁、记录耗时、配合 recover
// 3. panic 使程序进入恐慌状态，沿调用栈向上传播，沿途执行各层注册的 defer
// 4. recover 只能在 defer 中调用，用于捕获 panic 并恢复执行
// 5. 选择原则：预期内的失败用 error 返回值处理；panic 只用于"不可能发生"或程序
//    无法继续的情况（如前置条件不满足、不可达分支），绝不用 panic 做常规错误处理

// DeferOrder 演示 defer 的后进先出（LIFO）执行顺序
// 思路：利用命名返回值，defer 在函数返回前依次执行并追加到返回值
// 作用：演示 defer 的执行顺序，以及"命名返回值会被 defer 修改"的机制
// 业务场景：资源清理的逆序释放（如先关连接再关文件）
func DeferOrder() (order []int) {
	defer func() { order = append(order, 1) }() // 最先注册，最后执行
	defer func() { order = append(order, 2) }() // 第二个注册，第二个执行
	defer func() { order = append(order, 3) }() // 最后注册，最先执行
	return
}

// SafeDivide 安全的除法：除数为 0 时 panic，由 recover 捕获并转为 error
// 思路：命名返回值 + defer recover，将 panic 转换为 error 返回给调用方
// 作用：演示 panic/recover 的错误恢复模式
// 注意：这是教学示例；生产代码应直接用 if b == 0 { return 0, err } 处理
// 业务场景：调用可能 panic 的第三方库时做兜底，防止程序崩溃
func SafeDivide(a, b int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic recovered: %v", r)
		}
	}()

	if b == 0 {
		panic("division by zero")
	}
	return a / b, nil
}

// Must 断言式辅助函数：条件不成立时 panic
// 思路：用于表达"此处条件必须成立"的编程假设（前置条件/不变量）
// 作用：演示 panic 的典型使用场景——不可能发生的情况
// 注意：只用于"绝不应该发生"的情况；可预期的失败请返回 error
// 业务场景：初始化校验（配置必须非空）、不可达分支断言
func Must(condition bool, msg string) {
	if !condition {
		panic(msg)
	}
}

// CallWithRecover 通用错误恢复包装器：执行 fn，将其 panic 转为 error 返回
// 思路：在 defer 中调用 recover，借助命名返回值承载 error
// 作用：演示"panic 边界"模式——对不可控代码统一做兜底
// 业务场景：调用插件/用户自定义回调/第三方库时防止 panic 击穿整个服务
func CallWithRecover(fn func()) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic recovered: %v", r)
		}
	}()
	fn()
	return nil
}
