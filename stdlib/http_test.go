package stdlib

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestJSONResponse(t *testing.T) {
	rec := httptest.NewRecorder()
	JSONResponse(rec, map[string]string{"k": "v"})

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rec.Code)
	}
	if ct := rec.Header().Get("Content-Type"); !strings.HasPrefix(ct, "application/json") {
		t.Errorf("Expected application/json content-type, got %s", ct)
	}

	var msg Message
	if err := json.Unmarshal(rec.Body.Bytes(), &msg); err != nil {
		t.Fatalf("Decode response failed: %v", err)
	}
	if msg.Code != 0 {
		t.Errorf("Expected code 0, got %d", msg.Code)
	}
	data, ok := msg.Data.(map[string]interface{})
	if !ok {
		t.Fatalf("Expected data map, got %v", msg.Data)
	}
	if data["k"] != "v" {
		t.Errorf("Expected data[k]=v, got %v", data["k"])
	}
}

func TestJSONError(t *testing.T) {
	rec := httptest.NewRecorder()
	JSONError(rec, http.StatusBadRequest, "bad request")

	if rec.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %d", rec.Code)
	}
	var msg Message
	if err := json.Unmarshal(rec.Body.Bytes(), &msg); err != nil {
		t.Fatalf("Decode response failed: %v", err)
	}
	if msg.Msg != "bad request" {
		t.Errorf("Expected msg \"bad request\", got %q", msg.Msg)
	}
}

func TestHelloEndpoint(t *testing.T) {
	srv := httptest.NewServer(NewMux())
	defer srv.Close()

	var msg Message
	if err := GetJSON(srv.URL+"/api/hello", 2*time.Second, &msg); err != nil {
		t.Fatalf("GetJSON failed: %v", err)
	}
	if msg.Code != 0 {
		t.Errorf("Expected code 0, got %d", msg.Code)
	}
	if data, ok := msg.Data.(string); !ok || data == "" {
		t.Errorf("Expected non-empty string data, got %v", msg.Data)
	}
}

func TestGreetEndpoint(t *testing.T) {
	srv := httptest.NewServer(NewMux())
	defer srv.Close()

	var msg Message
	if err := GetJSON(srv.URL+"/api/greet?name=%E5%BC%A0%E4%B8%89", 2*time.Second, &msg); err != nil {
		t.Fatalf("GetJSON failed: %v", err)
	}
	data, ok := msg.Data.(map[string]interface{})
	if !ok {
		t.Fatalf("Expected data map, got %v", msg.Data)
	}
	if data["greeting"] != "你好，张三" {
		t.Errorf("Expected \"你好，张三\", got %v", data["greeting"])
	}
}

func TestUserEndpoint(t *testing.T) {
	srv := httptest.NewServer(NewMux())
	defer srv.Close()

	var msg Message
	if err := GetJSON(srv.URL+"/api/user", 2*time.Second, &msg); err != nil {
		t.Fatalf("GetJSON failed: %v", err)
	}
	data, ok := msg.Data.(map[string]interface{})
	if !ok {
		t.Fatalf("Expected data map, got %v", msg.Data)
	}
	if data["name"] != "小明" {
		t.Errorf("Expected name 小明, got %v", data["name"])
	}
}

func TestEchoEndpoint(t *testing.T) {
	srv := httptest.NewServer(NewMux())
	defer srv.Close()

	payload := map[string]string{"message": "hi"}
	var msg Message
	if err := PostJSON(srv.URL+"/api/echo", payload, 2*time.Second, &msg); err != nil {
		t.Fatalf("PostJSON failed: %v", err)
	}
	data, ok := msg.Data.(map[string]interface{})
	if !ok {
		t.Fatalf("Expected data map, got %v", msg.Data)
	}
	if data["echo"] != "hi" {
		t.Errorf("Expected echo \"hi\", got %v", data["echo"])
	}
}

func TestMethodNotAllowed(t *testing.T) {
	srv := httptest.NewServer(NewMux())
	defer srv.Close()

	var msg Message
	err := PostJSON(srv.URL+"/api/hello", map[string]string{}, 2*time.Second, &msg)
	if err == nil {
		t.Fatal("Expected error for wrong method")
	}
	if !strings.Contains(err.Error(), "405") {
		t.Errorf("Expected 405 in error, got %v", err)
	}
}

func TestNotFound(t *testing.T) {
	srv := httptest.NewServer(NewMux())
	defer srv.Close()

	var msg Message
	err := GetJSON(srv.URL+"/api/nonexistent", 2*time.Second, &msg)
	if err == nil {
		t.Fatal("Expected error for unknown path")
	}
	if !strings.Contains(err.Error(), "404") {
		t.Errorf("Expected 404 in error, got %v", err)
	}
}

func TestEchoHandlerInvalidBody(t *testing.T) {
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/echo", strings.NewReader("{not-json"))
	NewMux().ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %d", rec.Code)
	}
}

func TestGetJSONTimeout(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/slow", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(200 * time.Millisecond)
		w.Write([]byte("ok"))
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()

	var out interface{}
	// 客户端 10ms 超时，远小于服务端 200ms 处理时间
	if err := GetJSON(srv.URL+"/slow", 10*time.Millisecond, &out); err == nil {
		t.Error("Expected timeout error")
	}
}
