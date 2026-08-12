package advanced

import (
	"strings"
	"testing"
)

func TestDeferOrder(t *testing.T) {
	// defer 按后进先出（LIFO）执行：注册顺序 3、2、1，执行顺序 1、2、3
	got := DeferOrder()
	expected := []int{3, 2, 1}
	if len(got) != 3 {
		t.Fatalf("Expected 3 items, got %d", len(got))
	}
	for i := range expected {
		if got[i] != expected[i] {
			t.Errorf("At index %d, expected %d, got %d", i, expected[i], got[i])
		}
	}
}

func TestSafeDivide(t *testing.T) {
	// 正常除法
	result, err := SafeDivide(10, 2)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != 5 {
		t.Errorf("Expected 5, got %d", result)
	}

	// 除数为 0：panic 被 recover，转为 error 返回
	result, err = SafeDivide(10, 0)
	if err == nil {
		t.Error("Expected error for division by zero")
	}
	if result != 0 {
		t.Errorf("Expected 0 on error, got %d", result)
	}
	if !strings.Contains(err.Error(), "division by zero") {
		t.Errorf("Expected error mentioning 'division by zero', got %v", err)
	}
}

func TestMust(t *testing.T) {
	// 条件成立：不 panic
	Must(true, "should not panic")

	// 条件不成立：触发 panic
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Error("Expected panic from Must(false)")
			}
		}()
		Must(false, "condition failed")
	}()
}

func TestCallWithRecover(t *testing.T) {
	// 正常执行：无 panic
	err := CallWithRecover(func() {})
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// panic 被转换为 error
	err = CallWithRecover(func() { panic("boom") })
	if err == nil {
		t.Error("Expected error from panicking function")
	}
	if !strings.Contains(err.Error(), "boom") {
		t.Errorf("Expected error mentioning 'boom', got %v", err)
	}
}
