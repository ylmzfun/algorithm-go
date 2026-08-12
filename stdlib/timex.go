package stdlib

// 思路：time.Time 表示带时区的时间点；Go 的时间布局使用参考时间 2006-01-02 15:04:05
// 作为模板，格式化与解析共用同一套布局。定时器与超时控制依赖通道。
// 作用：演示时间格式化、解析、定时器与超时模式。
// 业务场景：
// 1. 日志时间戳与接口响应时间
// 2. 定时任务与心跳检测（Ticker）
// 3. 请求超时控制（time.After）

import (
	"sync"
	"time"
)

// CustomLayout 自定义时间布局（基于参考时间 2006-01-02 15:04:05）
const CustomLayout = "2006-01-02 15:04:05"

// FormatRFC3339 将时间格式化为 RFC3339（如 "2024-01-15T10:30:00+08:00"）
func FormatRFC3339(t time.Time) string {
	return t.Format(time.RFC3339)
}

// ParseRFC3339 解析 RFC3339 格式的时间字符串
// 格式非法时返回错误
func ParseRFC3339(s string) (time.Time, error) {
	return time.Parse(time.RFC3339, s)
}

// FormatCustom 使用自定义布局（2006-01-02 15:04:05）格式化时间
func FormatCustom(t time.Time) string {
	return t.Format(CustomLayout)
}

// ParseCustom 使用自定义布局解析时间字符串
func ParseCustom(s string) (time.Time, error) {
	return time.Parse(CustomLayout, s)
}

// RunTicker 定时器模式：每隔 interval 调用一次 fn
// 返回 stop 函数用于停止定时器（可安全重复调用）
// 注意：fn 执行耗时超过 interval 时会跳过中间周期
func RunTicker(interval time.Duration, fn func()) func() {
	ticker := time.NewTicker(interval)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				fn()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
	var once sync.Once
	return func() {
		once.Do(func() { close(quit) })
	}
}

// WaitTimeout 等待 done 通道关闭；在 timeout 内未关闭则返回 false
// 经典超时模式：select + time.After
func WaitTimeout(done <-chan struct{}, timeout time.Duration) bool {
	select {
	case <-done:
		return true
	case <-time.After(timeout):
		return false
	}
}

// DayRange 返回 t 所在自然日 [当天零点, 次日零点) 的起止时间
// 常用于按天统计的起止时间计算
func DayRange(t time.Time) (start, end time.Time) {
	start = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	end = start.AddDate(0, 0, 1)
	return start, end
}

// CalculateAge 计算出生日期 birth 在 now 时刻的周岁年龄
// now 由调用方传入（不依赖真实时钟，便于测试）
func CalculateAge(birth, now time.Time) int {
	if now.Before(birth) {
		return 0
	}
	age := now.Year() - birth.Year()
	// 今年生日未到则减一
	if now.Month() < birth.Month() ||
		(now.Month() == birth.Month() && now.Day() < birth.Day()) {
		age--
	}
	return age
}
