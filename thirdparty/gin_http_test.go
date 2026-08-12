package thirdparty

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

// init 将 gin 切换为测试模式，避免输出调试日志与默认中间件日志
func init() {
	gin.SetMode(gin.TestMode)
}

func TestSetupRouterPing(t *testing.T) {
	router := SetupRouter()

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/ping", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	var resp ApiResponse
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}
	if resp.Code != 0 {
		t.Errorf("Expected code 0, got %d", resp.Code)
	}
	if resp.Msg != "pong" {
		t.Errorf("Expected msg 'pong', got '%s'", resp.Msg)
	}

	// 验证中间件设置的响应头
	if w.Header().Get("X-Process-Time") == "" {
		t.Error("Expected X-Process-Time header set by middleware")
	}
}

func TestSetupRouterCreateUser(t *testing.T) {
	router := SetupRouter()

	body := bytes.NewBufferString(`{"name":"Alice","age":25}`)
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/users", body)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	var resp struct {
		Code int `json:"code"`
		Data struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		} `json:"data"`
	}
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}
	if resp.Code != 0 {
		t.Errorf("Expected code 0, got %d", resp.Code)
	}
	if resp.Data.Name != "Alice" {
		t.Errorf("Expected name 'Alice', got '%s'", resp.Data.Name)
	}
	if resp.Data.Age != 25 {
		t.Errorf("Expected age 25, got %d", resp.Data.Age)
	}
}

func TestSetupRouterCreateUserInvalid(t *testing.T) {
	router := SetupRouter()

	// 缺少必填字段 name，ShouldBindJSON 应返回校验错误
	body := bytes.NewBufferString(`{"age":25}`)
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/users", body)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %d", w.Code)
	}
}

func TestSetupRouterGetUser(t *testing.T) {
	router := SetupRouter()

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/users/42", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
	if !bytes.Contains(w.Body.Bytes(), []byte(`"id":"42"`)) {
		t.Errorf("Expected response contains id 42, got %s", w.Body.String())
	}
}
