package thirdparty

import (
	"bytes"
	"strings"
	"testing"
)

// runCmd 执行命令并返回输出与错误
// 作用：把命令输出重定向到内存缓冲区，测试全程不触碰真实标准输出
func runCmd(t *testing.T, args ...string) (string, error) {
	t.Helper()
	cmd := NewRootCmd()

	var buf bytes.Buffer
	cmd.SetOut(&buf)
	cmd.SetErr(&buf)
	cmd.SetArgs(args)

	err := cmd.Execute()
	return buf.String(), err
}

func TestRootCmdNoArgs(t *testing.T) {
	out, err := runCmd(t)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !strings.Contains(out, "cli-demo") {
		t.Errorf("Expected usage contains 'cli-demo', got %s", out)
	}
}

func TestGreetCmd(t *testing.T) {
	out, err := runCmd(t, "greet", "--name", "Go")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !strings.Contains(out, "Hello, Go!") {
		t.Errorf("Expected 'Hello, Go!' in output, got %s", out)
	}
}

func TestGreetCmdTimes(t *testing.T) {
	out, err := runCmd(t, "greet", "--name", "Go", "--times", "3")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if count := strings.Count(out, "Hello, Go!"); count != 3 {
		t.Errorf("Expected 3 greetings, got %d (output: %s)", count, out)
	}
}

func TestGreetCmdEmptyName(t *testing.T) {
	_, err := runCmd(t, "greet")
	if err == nil {
		t.Error("Expected error for empty name")
	}
}

func TestCalcCmdAdd(t *testing.T) {
	out, err := runCmd(t, "calc", "--a", "2", "--b", "3", "--op", "add")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !strings.Contains(out, "2 add 3 = 5") {
		t.Errorf("Expected '2 add 3 = 5' in output, got %s", out)
	}
}

func TestCalcCmdDiv(t *testing.T) {
	out, err := runCmd(t, "calc", "--a", "10", "--b", "2", "--op", "div")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !strings.Contains(out, "10 div 2 = 5") {
		t.Errorf("Expected '10 div 2 = 5' in output, got %s", out)
	}
}

func TestCalcCmdDivByZero(t *testing.T) {
	_, err := runCmd(t, "calc", "--a", "10", "--b", "0", "--op", "div")
	if err == nil {
		t.Error("Expected error for division by zero")
	}
}

func TestCalcCmdUnknownOp(t *testing.T) {
	_, err := runCmd(t, "calc", "--a", "1", "--b", "1", "--op", "mod")
	if err == nil {
		t.Error("Expected error for unknown operator")
	}
}
