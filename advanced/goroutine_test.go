package advanced

import (
	"testing"
)

func TestConcurrentSum(t *testing.T) {
	nums := make([]int, 0, 100)
	for i := 1; i <= 100; i++ {
		nums = append(nums, i)
	}

	// 多个 worker 并发求和（1+2+...+100 = 5050）
	got := ConcurrentSum(nums, 4)
	if got != 5050 {
		t.Errorf("Expected sum 5050, got %d", got)
	}

	// worker 数大于元素数
	got = ConcurrentSum(nums, 200)
	if got != 5050 {
		t.Errorf("Expected sum 5050, got %d", got)
	}

	// 非法 worker 数（<= 0 时按 1 处理）
	got = ConcurrentSum(nums, 0)
	if got != 5050 {
		t.Errorf("Expected sum 5050, got %d", got)
	}
}

func TestConcurrentSumEmpty(t *testing.T) {
	got := ConcurrentSum(nil, 4)
	if got != 0 {
		t.Errorf("Expected sum 0 for empty slice, got %d", got)
	}
}

func TestConcurrentSumSingle(t *testing.T) {
	got := ConcurrentSum([]int{42}, 3)
	if got != 42 {
		t.Errorf("Expected sum 42, got %d", got)
	}
}

func TestParallelMap(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8}
	double := func(v int) int { return v * 2 }

	got := ParallelMap(data, double, 3)
	expected := []int{2, 4, 6, 8, 10, 12, 14, 16}
	if len(got) != len(expected) {
		t.Fatalf("Expected %v, got %v", expected, got)
	}
	for i := range expected {
		if got[i] != expected[i] {
			t.Errorf("At index %d, expected %d, got %d", i, expected[i], got[i])
		}
	}
}

func TestParallelMapEdgeCases(t *testing.T) {
	// 空切片
	got := ParallelMap(nil, func(v int) int { return v }, 4)
	if len(got) != 0 {
		t.Errorf("Expected empty result, got %v", got)
	}

	// worker 数大于元素数
	got = ParallelMap([]int{1, 2}, func(v int) int { return v * 10 }, 5)
	if len(got) != 2 || got[0] != 10 || got[1] != 20 {
		t.Errorf("Expected [10 20], got %v", got)
	}

	// 非法 worker 数
	got = ParallelMap([]int{1, 2, 3}, func(v int) int { return v }, 0)
	if len(got) != 3 || got[0] != 1 || got[2] != 3 {
		t.Errorf("Expected [1 2 3], got %v", got)
	}
}
