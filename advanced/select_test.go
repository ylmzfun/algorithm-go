package advanced

import (
	"sort"
	"testing"
	"time"
)

func TestSelectReceive(t *testing.T) {
	// 只有 ch1 就绪：必然从 ch1 接收
	ch1 := make(chan int, 1)
	ch1 <- 100
	ch2 := make(chan int, 1)

	value, from := SelectReceive(ch1, ch2)
	if value != 100 || from != 1 {
		t.Errorf("Expected (100, 1), got (%d, %d)", value, from)
	}

	// 只有 ch2 就绪：必然从 ch2 接收
	ch3 := make(chan int, 1)
	ch4 := make(chan int, 1)
	ch4 <- 200

	value, from = SelectReceive(ch3, ch4)
	if value != 200 || from != 2 {
		t.Errorf("Expected (200, 2), got (%d, %d)", value, from)
	}
}

func TestSelectReceiveRandom(t *testing.T) {
	// 两个 channel 同时就绪：select 随机选择，但收到的值必然来自其中一方。
	// 注意：接收会从 channel 取走值，因此每次迭代都重建两个各含一个值的 channel
	for i := 0; i < 20; i++ {
		ch1 := make(chan int, 1)
		ch1 <- 1
		ch2 := make(chan int, 1)
		ch2 <- 2

		value, from := SelectReceive(ch1, ch2)
		if value != 1 && value != 2 {
			t.Errorf("Received unexpected value %d", value)
		}
		if from != 1 && from != 2 {
			t.Errorf("Received unexpected source %d", from)
		}
	}
}

func TestSelectWithTimeout(t *testing.T) {
	// 数据先就绪：正常返回数据（超时 channel 永不就绪）
	ch := make(chan int, 1)
	ch <- 42
	timeout := make(chan time.Time) // 无发送者，永不就绪

	value, err := SelectWithTimeout(ch, timeout)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if value != 42 {
		t.Errorf("Expected 42, got %d", value)
	}

	// 超时先就绪：返回 ErrTimeout
	// 确定性：直接构造已就绪的超时 channel，不依赖真实计时
	ch2 := make(chan int) // 无发送者，永不就绪
	timeout2 := make(chan time.Time, 1)
	timeout2 <- time.Now()

	value, err = SelectWithTimeout(ch2, timeout2)
	if err != ErrTimeout {
		t.Errorf("Expected ErrTimeout, got %v", err)
	}
	if value != 0 {
		t.Errorf("Expected 0 on timeout, got %d", value)
	}
}

func TestSelectLoop(t *testing.T) {
	// 两个 channel 各发 3 个值后关闭，收集全部 6 个值
	ch1 := make(chan int, 3)
	for _, v := range []int{1, 2, 3} {
		ch1 <- v
	}
	close(ch1)

	ch2 := make(chan int, 3)
	for _, v := range []int{4, 5, 6} {
		ch2 <- v
	}
	close(ch2)

	got := SelectLoop(ch1, ch2, 6)
	if len(got) != 6 {
		t.Fatalf("Expected 6 values, got %d: %v", len(got), got)
	}
	// 接收顺序不固定（select 随机选择），但值集合确定
	sort.Ints(got)
	expected := []int{1, 2, 3, 4, 5, 6}
	for i := range expected {
		if got[i] != expected[i] {
			t.Errorf("Expected %v, got %v", expected, got)
			break
		}
	}
}

func TestSelectLoopClosedEarly(t *testing.T) {
	// count 大于实际可收数据量时，两个 channel 关闭后循环自动结束
	ch1 := make(chan int, 1)
	ch1 <- 7
	close(ch1)
	ch2 := make(chan int)
	close(ch2)

	got := SelectLoop(ch1, ch2, 10)
	if len(got) != 1 || got[0] != 7 {
		t.Errorf("Expected [7], got %v", got)
	}
}

func TestTryReceive(t *testing.T) {
	// channel 有数据：成功接收
	ch := make(chan int, 1)
	ch <- 5
	value, ok := TryReceive(ch)
	if !ok || value != 5 {
		t.Errorf("Expected (5, true), got (%d, %v)", value, ok)
	}

	// channel 无数据：立即返回 false，不阻塞
	ch2 := make(chan int, 1)
	value, ok = TryReceive(ch2)
	if ok {
		t.Errorf("Expected ok = false, got %v (value = %d)", ok, value)
	}
}
