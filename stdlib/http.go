package stdlib

// 思路：net/http 基于 Handler 接口与 ServeMux 路由分发请求，handler 向 ResponseWriter
// 写入响应，客户端通过 http.Client 发起请求。
// 作用：演示 HTTP 服务端（路由、JSON 接口）与客户端（带超时的 Get/Post）的完整用法。
// 业务场景：
// 1. RESTful API 服务
// 2. 微服务之间的 HTTP 调用
// 3. 爬虫、监控告警系统的抓取客户端

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Message 统一的 JSON 响应结构
type Message struct {
	Code int         `json:"code"`           // 业务码，0 表示成功
	Msg  string      `json:"msg"`            // 提示信息
	Data interface{} `json:"data,omitempty"` // 业务数据
}

// NewMux 构建演示用的 HTTP 路由
// 注册的接口：
//
//	GET  /api/hello  返回欢迎信息
//	GET  /api/greet  读取 ?name= 查询参数并返回问候语
//	GET  /api/user   返回用户 JSON
//	POST /api/echo   解析请求体 JSON，回显 message 字段
func NewMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/hello", helloHandler)
	mux.HandleFunc("/api/greet", greetHandler)
	mux.HandleFunc("/api/user", userHandler)
	mux.HandleFunc("/api/echo", echoHandler)
	return mux
}

// NewHTTPServer 创建带演示路由的 HTTP 服务器
// addr: 监听地址，如 ":8080"
func NewHTTPServer(addr string) *http.Server {
	return &http.Server{
		Addr:              addr,
		Handler:           NewMux(),
		ReadHeaderTimeout: 5 * time.Second,
	}
}

// JSONResponse 写入成功的 JSON 响应（HTTP 200，业务码 0）
func JSONResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(Message{Code: 0, Msg: "success", Data: data})
}

// JSONError 写入失败的 JSON 响应
// status: HTTP 状态码，如 400/404/500
func JSONError(w http.ResponseWriter, status int, msg string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(Message{Code: status, Msg: msg})
}

// NewClient 创建带超时的 HTTP 客户端
// timeout: 单次请求总超时（含连接建立、重定向与响应体读取）
func NewClient(timeout time.Duration) *http.Client {
	return &http.Client{Timeout: timeout}
}

// GetJSON 发送 GET 请求并解析 JSON 响应体到 out
// url: 请求地址；timeout: 超时时间；out: 解析目标（须为指针）
func GetJSON(url string, timeout time.Duration, out interface{}) error {
	resp, err := NewClient(timeout).Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status: %s", resp.Status)
	}
	return json.NewDecoder(resp.Body).Decode(out)
}

// PostJSON 发送 JSON 请求体，并解析 JSON 响应体到 out
// payload: 请求体对象，会被 json.Marshal 序列化
func PostJSON(url string, payload interface{}, timeout time.Duration, out interface{}) error {
	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	resp, err := NewClient(timeout).Post(url, "application/json", bytes.NewReader(body))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status: %s", resp.Status)
	}
	return json.NewDecoder(resp.Body).Decode(out)
}

// helloHandler 返回欢迎信息
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		JSONError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	JSONResponse(w, "hello，欢迎使用 Go 标准库 net/http")
}

// greetHandler 读取查询参数 name 并返回问候语
func greetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		JSONError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "匿名用户"
	}
	JSONResponse(w, map[string]string{"greeting": "你好，" + name})
}

// userHandler 返回用户 JSON
func userHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		JSONError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	JSONResponse(w, map[string]interface{}{
		"id":    1,
		"name":  "小明",
		"email": "xiaoming@example.com",
		"tags":  []string{"golang", "algorithm"},
	})
}

// echoRequest 回显接口的请求体结构
type echoRequest struct {
	Message string `json:"message"`
}

// echoHandler 解析请求体 JSON 并回显 message 字段
func echoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		JSONError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	var req echoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		JSONError(w, http.StatusBadRequest, "invalid json body")
		return
	}
	JSONResponse(w, map[string]string{"echo": req.Message})
}
