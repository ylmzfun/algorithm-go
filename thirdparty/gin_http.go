// Package thirdparty 演示 Go 语言知名第三方库的典型用法
// 本包包含 gin（HTTP 框架）、zap（日志）、cobra（CLI）、viper（配置）、
// gorm（ORM）五个库的封装示例，每个文件对应一个库及其配套测试。
package thirdparty

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// ApiResponse 统一 API 响应结构
// 思路：所有接口返回统一的 JSON 结构，便于前端统一处理
// 作用：约定响应格式 code/msg/data，code 为 0 表示成功，非 0 表示错误
// 业务场景：RESTful 服务中所有接口的响应体
type ApiResponse struct {
	Code int         `json:"code"`           // 业务状态码，0 表示成功
	Msg  string      `json:"msg"`            // 提示信息
	Data interface{} `json:"data,omitempty"` // 业务数据
}

// UserRequest 用户创建请求体
// 作用：演示 gin 的 ShouldBindJSON 参数绑定与 binding 标签校验
type UserRequest struct {
	Name string `json:"name" binding:"required"` // 用户名，必填
	Age  int    `json:"age"`                     // 年龄
}

// RequestLogger 请求日志中间件
// 作用：统计请求数量，并记录每个请求的处理耗时到响应头
// 复杂度：O(1)，对所有请求生效
func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next() // 继续执行后续处理器
		latency := time.Since(start)
		c.Header("X-Process-Time", latency.String())
	}
}

// createUserHandler 创建用户
// 作用：演示 ShouldBindJSON 参数绑定与参数校验失败时的错误响应
func createUserHandler(c *gin.Context) {
	var req UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ApiResponse{Code: 1, Msg: "参数错误: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, ApiResponse{
		Code: 0,
		Msg:  "创建成功",
		Data: gin.H{"name": req.Name, "age": req.Age},
	})
}

// getUserHandler 查询用户
// 作用：演示路径参数 c.Param 的取值方式
func getUserHandler(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, ApiResponse{
		Code: 0,
		Msg:  "查询成功",
		Data: gin.H{"id": id},
	})
}

// SetupRouter 构建 gin 路由引擎
// 思路：注册全局中间件、分组路由与各业务接口，把路由集中管理
// 作用：返回一个可复用的 *gin.Engine，既可通过 Run 监听真实端口，
// 也可在测试中配合 httptest.NewRecorder 直接调用，无需占用端口
// 业务场景：Web 服务入口，集中注册路由便于管理与测试
func SetupRouter() *gin.Engine {
	// gin.New 不含任何中间件，按需显式注册（gin.Default 则内置 Logger 与 Recovery）
	router := gin.New()
	router.Use(gin.Recovery()) // 兜底恢复，panic 时返回 500 而非崩溃
	router.Use(RequestLogger())

	// 探活接口
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, ApiResponse{Code: 0, Msg: "pong"})
	})

	// 分组路由：用户相关接口统一挂到 /api/v1/users 前缀下
	userGroup := router.Group("/api/v1/users")
	{
		userGroup.POST("", createUserHandler)
		userGroup.GET("/:id", getUserHandler)
	}

	return router
}
