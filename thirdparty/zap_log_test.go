package thirdparty

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// newMemoryLogger 构造写入内存缓冲区的 JSON 日志器
// 作用：测试中捕获日志输出并断言内容，不依赖任何外部服务
func newMemoryLogger(t *testing.T) (*zap.Logger, *bytes.Buffer) {
	t.Helper()
	var buf bytes.Buffer
	encoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	core := zapcore.NewCore(encoder, zapcore.AddSync(&buf), zapcore.DebugLevel)
	logger := zap.New(core)
	return logger, &buf
}

func TestNewDevelopmentLogger(t *testing.T) {
	logger, err := NewDevelopmentLogger()
	if err != nil {
		t.Fatalf("Failed to create development logger: %v", err)
	}
	defer logger.Sync()

	if logger == nil {
		t.Error("Expected non-nil development logger")
	}
}

func TestNewProductionLogger(t *testing.T) {
	logger, err := NewProductionLogger()
	if err != nil {
		t.Fatalf("Failed to create production logger: %v", err)
	}
	defer logger.Sync()

	if logger == nil {
		t.Error("Expected non-nil production logger")
	}
}

func TestNewSugaredLogger(t *testing.T) {
	sugar, err := NewSugaredLogger()
	if err != nil {
		t.Fatalf("Failed to create sugared logger: %v", err)
	}
	defer sugar.Sync()

	if sugar == nil {
		t.Error("Expected non-nil sugared logger")
	}
}

func TestLogUserAction(t *testing.T) {
	logger, buf := newMemoryLogger(t)
	LogUserAction(logger, "login", 42, true)
	logger.Sync()

	out := buf.String()
	wants := []string{`"level":"info"`, `"action":"login"`, `"user_id":42`, `"success":true`}
	for _, want := range wants {
		if !strings.Contains(out, want) {
			t.Errorf("Expected log contains %s, got %s", want, out)
		}
	}
}

func TestLogError(t *testing.T) {
	logger, buf := newMemoryLogger(t)
	LogError(logger, errors.New("connection refused"), "db-init")
	logger.Sync()

	out := buf.String()
	wants := []string{`"level":"error"`, `"context":"db-init"`, `"error":"connection refused"`}
	for _, want := range wants {
		if !strings.Contains(out, want) {
			t.Errorf("Expected log contains %s, got %s", want, out)
		}
	}
}
