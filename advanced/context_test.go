package advanced

import (
	"context"
	"testing"
	"time"
)

func TestProcessItems(t *testing.T) {
	items := []string{"a", "b", "c"}
	got := ProcessItems(context.Background(), items)
	if len(got) != 3 {
		t.Fatalf("Expected 3 processed items, got %d", len(got))
	}
	for i, item := range items {
		expected := "processed:" + item
		if got[i] != expected {
			t.Errorf("At index %d, expected %q, got %q", i, expected, got[i])
		}
	}
}

func TestProcessItemsCanceled(t *testing.T) {
	// 处理开始前 context 已被取消：立即返回空结果（确定性，无需等待）
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	got := ProcessItems(ctx, []string{"a", "b", "c"})
	if len(got) != 0 {
		t.Errorf("Expected empty result for canceled context, got %v", got)
	}
}

func TestProcessItemsEmpty(t *testing.T) {
	got := ProcessItems(context.Background(), nil)
	if len(got) != 0 {
		t.Errorf("Expected empty result, got %v", got)
	}
}

func TestWithRequestID(t *testing.T) {
	ctx := WithRequestID(context.Background(), "req-12345")
	if id := GetRequestID(ctx); id != "req-12345" {
		t.Errorf("Expected request id 'req-12345', got %q", id)
	}

	// 未设置时返回空串
	if id := GetRequestID(context.Background()); id != "" {
		t.Errorf("Expected empty request id, got %q", id)
	}

	// 派生 context 也能读到（值随调用链传递）
	child, cancel := context.WithCancel(ctx)
	defer cancel()
	if id := GetRequestID(child); id != "req-12345" {
		t.Errorf("Expected request id inherited by child, got %q", id)
	}
}

func TestCancelPropagation(t *testing.T) {
	// 取消父 context 后，子 context 同步取消（channel 同步验证，非 sleep）
	err := CancelPropagation()
	if err != context.Canceled {
		t.Errorf("Expected context.Canceled, got %v", err)
	}
}

func TestRunWithTimeout(t *testing.T) {
	// 超时：时限为 0，WithTimeout 对已过去的 deadline 立即取消，任务必然超时
	results, timedOut := RunWithTimeout(0, func(ctx context.Context) []string {
		return ProcessItems(ctx, []string{"a", "b", "c"})
	})
	if !timedOut {
		t.Error("Expected timeout with zero duration")
	}
	if len(results) != 0 {
		t.Errorf("Expected no results on timeout, got %v", results)
	}

	// 正常完成：超时时间远大于任务耗时，不会触发超时
	results, timedOut = RunWithTimeout(time.Hour, func(ctx context.Context) []string {
		return ProcessItems(ctx, []string{"a", "b", "c"})
	})
	if timedOut {
		t.Error("Expected no timeout with 1 hour duration")
	}
	if len(results) != 3 {
		t.Errorf("Expected 3 results, got %d: %v", len(results), results)
	}
}
