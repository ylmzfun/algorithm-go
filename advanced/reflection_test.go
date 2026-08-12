package advanced

import (
	"strings"
	"testing"
)

func TestTypeName(t *testing.T) {
	if got := TypeName(42); got != "int" {
		t.Errorf("TypeName(42): expected 'int', got %q", got)
	}
	if got := TypeName("hello"); got != "string" {
		t.Errorf("TypeName(hello): expected 'string', got %q", got)
	}

	user := User{Name: "tom", Age: 20}
	if got := TypeName(user); got != "User" {
		t.Errorf("TypeName(user): expected 'User', got %q", got)
	}
	if got := KindName(user); got != "struct" {
		t.Errorf("KindName(user): expected 'struct', got %q", got)
	}

	// 反射 nil 接口
	if got := TypeName(nil); got != "<nil>" {
		t.Errorf("TypeName(nil): expected '<nil>', got %q", got)
	}
}

func TestFieldValues(t *testing.T) {
	user := User{Name: "tom", Age: 20}
	infos, err := FieldValues(user)
	if err != nil {
		t.Fatalf("FieldValues failed: %v", err)
	}
	if len(infos) != 2 {
		t.Fatalf("Expected 2 fields, got %d", len(infos))
	}

	if infos[0].Name != "Name" || infos[0].Type != "string" {
		t.Errorf("Field 0: expected Name/string, got %s/%s", infos[0].Name, infos[0].Type)
	}
	if infos[0].Value != "tom" {
		t.Errorf("Field 0 value: expected 'tom', got %v", infos[0].Value)
	}
	if !strings.Contains(infos[0].Tag, `json:"name"`) {
		t.Errorf("Field 0 tag: expected json tag, got %q", infos[0].Tag)
	}

	if infos[1].Name != "Age" || infos[1].Type != "int" {
		t.Errorf("Field 1: expected Age/int, got %s/%s", infos[1].Name, infos[1].Type)
	}
	if infos[1].Value != 20 {
		t.Errorf("Field 1 value: expected 20, got %v", infos[1].Value)
	}

	// 非结构体输入报错
	if _, err := FieldValues(42); err == nil {
		t.Error("Expected error for non-struct input")
	}
}

func TestGetTagValue(t *testing.T) {
	user := User{Name: "tom", Age: 20}

	// 读取 json tag
	tag, err := GetTagValue(user, "Name", "json")
	if err != nil {
		t.Fatalf("GetTagValue failed: %v", err)
	}
	if tag != "name" {
		t.Errorf("Expected 'name', got %q", tag)
	}

	// 读取 validate tag
	tag, err = GetTagValue(user, "Age", "validate")
	if err != nil {
		t.Fatalf("GetTagValue failed: %v", err)
	}
	if !strings.Contains(tag, "gte=0") {
		t.Errorf("Expected validate tag containing 'gte=0', got %q", tag)
	}

	// 字段不存在：返回错误
	if _, err := GetTagValue(user, "Missing", "json"); err == nil {
		t.Error("Expected error for missing field")
	}

	// 字段存在但无该 tag：返回空字符串
	tag, err = GetTagValue(user, "Name", "unknown-tag")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if tag != "" {
		t.Errorf("Expected empty tag value, got %q", tag)
	}

	// 非结构体输入报错
	if _, err := GetTagValue(42, "Name", "json"); err == nil {
		t.Error("Expected error for non-struct input")
	}
}

func add(a, b int) int { return a + b }

func TestCallFunction(t *testing.T) {
	// 动态调用普通函数
	results, err := CallFunction(add, 3, 4)
	if err != nil {
		t.Fatalf("CallFunction failed: %v", err)
	}
	if len(results) != 1 || results[0] != 7 {
		t.Errorf("Expected [7], got %v", results)
	}

	// 非函数输入：返回错误
	if _, err := CallFunction(42); err == nil {
		t.Error("Expected error for non-function input")
	}

	// 参数个数不匹配：返回错误
	if _, err := CallFunction(add, 1, 2, 3); err == nil {
		t.Error("Expected error for wrong argument count")
	}

	// 参数类型不匹配：reflect 的 panic 被转为 error
	if _, err := CallFunction(add, "a", "b"); err == nil {
		t.Error("Expected error for wrong argument types")
	}
}

func TestCallMethod(t *testing.T) {
	calc := NewCalculator(10)

	// 动态调用结构体方法
	results, err := CallMethod(calc, "Add", 5)
	if err != nil {
		t.Fatalf("CallMethod failed: %v", err)
	}
	if len(results) != 1 || results[0] != 15 {
		t.Errorf("Expected [15], got %v", results)
	}

	// 方法不存在：返回错误
	if _, err := CallMethod(calc, "Subtract", 1); err == nil {
		t.Error("Expected error for missing method")
	}
}
