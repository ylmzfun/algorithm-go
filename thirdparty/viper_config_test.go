package thirdparty

import (
	"strings"
	"testing"

	"github.com/spf13/viper"
)

func TestLoadConfigFromYAML(t *testing.T) {
	yamlData := `
host: 127.0.0.1
port: 9090
debug: true
database:
  driver: sqlite
  dsn: ":memory:"
  max_open_conns: 10
`
	cfg, err := LoadConfigFromYAML(yamlData)
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	if cfg.Host != "127.0.0.1" {
		t.Errorf("Expected host 127.0.0.1, got %s", cfg.Host)
	}
	if cfg.Port != 9090 {
		t.Errorf("Expected port 9090, got %d", cfg.Port)
	}
	if !cfg.Debug {
		t.Error("Expected debug true")
	}
	if cfg.Database.Driver != "sqlite" {
		t.Errorf("Expected driver 'sqlite', got '%s'", cfg.Database.Driver)
	}
	if cfg.Database.DSN != ":memory:" {
		t.Errorf("Expected dsn ':memory:', got '%s'", cfg.Database.DSN)
	}
	if cfg.Database.MaxOpenConns != 10 {
		t.Errorf("Expected max_open_conns 10, got %d", cfg.Database.MaxOpenConns)
	}
}

func TestLoadConfigFromYAMLDefaults(t *testing.T) {
	// 只提供 host，port 与 debug 应走 SetDefault 默认值
	yamlData := "host: localhost\n"
	cfg, err := LoadConfigFromYAML(yamlData)
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	if cfg.Host != "localhost" {
		t.Errorf("Expected host 'localhost', got '%s'", cfg.Host)
	}
	if cfg.Port != 8080 {
		t.Errorf("Expected default port 8080, got %d", cfg.Port)
	}
	if cfg.Debug {
		t.Error("Expected default debug false")
	}
}

func TestLoadConfigFromYAMLInvalid(t *testing.T) {
	// tab 开头的缩进是非法的 YAML，ReadConfig 应返回错误
	_, err := LoadConfigFromYAML("\tport: 8080\n")
	if err == nil {
		t.Error("Expected error for invalid yaml")
	}
}

func TestGetServerPort(t *testing.T) {
	v := viper.New()
	v.SetConfigType("yaml")
	if err := v.ReadConfig(strings.NewReader("port: 3000\n")); err != nil {
		t.Fatalf("Failed to read config: %v", err)
	}

	if port := GetServerPort(v); port != 3000 {
		t.Errorf("Expected port 3000, got %d", port)
	}
}
