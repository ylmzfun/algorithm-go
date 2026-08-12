package stdlib

import (
	"sync/atomic"
	"testing"
	"time"
)

func TestFormatRFC3339(t *testing.T) {
	tm := time.Date(2024, 1, 15, 10, 30, 0, 0, time.FixedZone("CST", 8*3600))
	if got := FormatRFC3339(tm); got != "2024-01-15T10:30:00+08:00" {
		t.Errorf("Expected \"2024-01-15T10:30:00+08:00\", got %q", got)
	}
}

func TestParseRFC3339(t *testing.T) {
	tm, err := ParseRFC3339("2024-01-15T10:30:00Z")
	if err != nil {
		t.Fatalf("ParseRFC3339 failed: %v", err)
	}
	want := time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC)
	if !tm.Equal(want) {
		t.Errorf("Expected %v, got %v", want, tm)
	}

	if _, err := ParseRFC3339("2024-01-15 10:30:00"); err == nil {
		t.Error("Expected error for non-RFC3339 layout")
	}
}

func TestFormatCustomAndParseCustom(t *testing.T) {
	tm := time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC)
	if got := FormatCustom(tm); got != "2024-01-15 10:30:00" {
		t.Errorf("Expected \"2024-01-15 10:30:00\", got %q", got)
	}
	back, err := ParseCustom("2024-01-15 10:30:00")
	if err != nil {
		t.Fatalf("ParseCustom failed: %v", err)
	}
	if !back.Equal(tm) {
		t.Errorf("Expected %v, got %v", tm, back)
	}
	if _, err := ParseCustom("2024/01/15 10:30:00"); err == nil {
		t.Error("Expected error for wrong layout")
	}
}

func TestRunTicker(t *testing.T) {
	var count atomic.Int32
	stop := RunTicker(10*time.Millisecond, func() {
		count.Add(1)
	})

	// 等待定时器至少触发 3 次
	deadline := time.Now().Add(2 * time.Second)
	for count.Load() < 3 && time.Now().Before(deadline) {
		time.Sleep(5 * time.Millisecond)
	}
	if count.Load() < 3 {
		t.Fatalf("Expected at least 3 ticks, got %d", count.Load())
	}

	// 停止后定时器不再触发
	stop()
	time.Sleep(30 * time.Millisecond) // 等待可能已排队的最后一次回调
	after := count.Load()
	time.Sleep(40 * time.Millisecond)
	if got := count.Load(); got != after {
		t.Errorf("Expected ticker stopped (%d), got %d", after, got)
	}
}

func TestWaitTimeout(t *testing.T) {
	// 通道已关闭 -> 立即返回 true
	done := make(chan struct{})
	close(done)
	if !WaitTimeout(done, time.Second) {
		t.Error("Expected true for closed channel")
	}

	// 通道未关闭 -> 超时返回 false
	pending := make(chan struct{})
	if WaitTimeout(pending, 20*time.Millisecond) {
		t.Error("Expected false after timeout")
	}
}

func TestDayRange(t *testing.T) {
	tm := time.Date(2024, 1, 15, 14, 30, 0, 0, time.UTC)
	start, end := DayRange(tm)
	if got := start.Format(CustomLayout); got != "2024-01-15 00:00:00" {
		t.Errorf("Expected start 2024-01-15 00:00:00, got %s", got)
	}
	if got := end.Format(CustomLayout); got != "2024-01-16 00:00:00" {
		t.Errorf("Expected end 2024-01-16 00:00:00, got %s", got)
	}
	if !end.After(start) {
		t.Error("Expected end after start")
	}
}

func TestCalculateAge(t *testing.T) {
	birth := time.Date(2000, 2, 29, 8, 0, 0, 0, time.UTC)
	now := time.Date(2024, 3, 1, 12, 0, 0, 0, time.UTC)
	if age := CalculateAge(birth, now); age != 24 {
		t.Errorf("Expected 24, got %d", age)
	}
	// 今年生日未到
	now = time.Date(2024, 2, 28, 12, 0, 0, 0, time.UTC)
	if age := CalculateAge(birth, now); age != 23 {
		t.Errorf("Expected 23, got %d", age)
	}
	// 出生日期晚于当前时间
	if age := CalculateAge(time.Date(2030, 1, 1, 0, 0, 0, 0, time.UTC), now); age != 0 {
		t.Errorf("Expected 0, got %d", age)
	}
}
