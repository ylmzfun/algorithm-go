package thirdparty

import (
	"time"

	"go.uber.org/zap"
)

// NewDevelopmentLogger 创建开发环境日志器
// 思路：使用 zap.NewDevelopment 构建带彩色控制台输出的开发日志器
// 作用：日志格式便于人读，输出到标准错误，适合本地开发调试
// 业务场景：本地开发、联调阶段使用
func NewDevelopmentLogger() (*zap.Logger, error) {
	return zap.NewDevelopment()
}

// NewProductionLogger 创建生产环境日志器
// 思路：使用 zap.NewProduction 构建 JSON 格式日志器
// 作用：输出结构化 JSON 日志，便于日志采集系统（如 ELK、Loki）解析
// 业务场景：线上服务、容器化部署环境使用
func NewProductionLogger() (*zap.Logger, error) {
	return zap.NewProduction()
}

// NewSugaredLogger 创建 Sugar 风格日志器
// 思路：在开发日志器基础上调用 Sugar() 包装一层
// 作用：提供类似 fmt.Printf 的轻量级日志 API，代码更简洁
// 业务场景：对性能要求不高、追求编码效率的业务代码
func NewSugaredLogger() (*zap.SugaredLogger, error) {
	logger, err := NewDevelopmentLogger()
	if err != nil {
		return nil, err
	}
	return logger.Sugar(), nil
}

// LogUserAction 记录用户操作日志
// 作用：演示 zap.String/zap.Int/zap.Bool 等结构化字段，便于按字段过滤检索
// 复杂度：O(1)
func LogUserAction(logger *zap.Logger, action string, userID int, success bool) {
	logger.Info("用户操作",
		zap.String("action", action),
		zap.Int("user_id", userID),
		zap.Bool("success", success),
		zap.String("ts", time.Now().Format(time.RFC3339)),
	)
}

// LogError 记录错误日志
// 作用：演示 Error 级别日志与 zap.Error 字段（自动提取错误的 message 与堆栈）
// 复杂度：O(1)
func LogError(logger *zap.Logger, err error, context string) {
	logger.Error("操作失败",
		zap.String("context", context),
		zap.Error(err),
	)
}
