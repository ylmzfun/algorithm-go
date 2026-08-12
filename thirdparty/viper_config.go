package thirdparty

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

// ServerConfig 服务配置结构
// 思路：用结构体承载 viper 解析后的配置，字段通过 mapstructure tag 与配置键对应
// 作用：将 YAML 配置映射为强类型结构体，避免业务代码中散落魔法字符串
// 业务场景：Web 服务启动参数、数据库连接信息等应用配置
type ServerConfig struct {
	Host     string         `mapstructure:"host"` // 监听地址
	Port     int            `mapstructure:"port"` // 监听端口
	Debug    bool           `mapstructure:"debug"`
	Database DatabaseConfig `mapstructure:"database"`
}

// DatabaseConfig 数据库配置
// 作用：嵌套结构体，演示 viper.UnmarshalKey 对嵌套配置的解析
type DatabaseConfig struct {
	Driver       string `mapstructure:"driver"`         // 数据库驱动
	DSN          string `mapstructure:"dsn"`            // 数据源名称
	MaxOpenConns int    `mapstructure:"max_open_conns"` // 最大连接数
}

// LoadConfigFromYAML 从 YAML 字符串加载配置
// 思路：使用 viper.New 创建独立实例，用 SetDefault 设置默认值，
// 再通过 ReadConfig 从内存读取 YAML，最后以 Get/UnmarshalKey 组合取值
// 作用：不依赖真实配置文件即可完成配置解析，便于单元测试
// 业务场景：测试环境、配置中心下发配置字符串等场景
func LoadConfigFromYAML(data string) (*ServerConfig, error) {
	v := viper.New()
	v.SetConfigType("yaml")

	// 设置默认值：配置缺省时兜底
	v.SetDefault("port", 8080)
	v.SetDefault("debug", false)

	if err := v.ReadConfig(strings.NewReader(data)); err != nil {
		return nil, fmt.Errorf("读取配置失败: %w", err)
	}

	// 顶层字段演示 viper.GetString / GetInt / GetBool 直接取值
	cfg := &ServerConfig{
		Host:  v.GetString("host"),
		Port:  v.GetInt("port"),
		Debug: v.GetBool("debug"),
	}
	// 嵌套结构体使用 UnmarshalKey 自动映射
	if err := v.UnmarshalKey("database", &cfg.Database); err != nil {
		return nil, fmt.Errorf("解析 database 配置失败: %w", err)
	}
	return cfg, nil
}

// GetServerPort 从配置中读取端口
// 作用：演示对外暴露 GetInt 便捷方法，供业务代码直接使用
// 复杂度：O(1)
func GetServerPort(v *viper.Viper) int {
	return v.GetInt("port")
}
