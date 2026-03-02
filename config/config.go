package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

// Config 结构体用于映射配置文件内容
// Config struct maps the configuration file content
type Config struct {
	App      AppConfig      `mapstructure:"app"`
	Database DatabaseConfig `mapstructure:"database"`
	Redis    RedisConfig    `mapstructure:"redis"`
	JWT      JWTConfig      `mapstructure:"jwt"`
}

// AppConfig 应用配置
// AppConfig holds application settings
type AppConfig struct {
	Name string `mapstructure:"name"`
	Port string `mapstructure:"port"`
	Mode string `mapstructure:"mode"` // debug, release, test
}

// DatabaseConfig 数据库配置
// DatabaseConfig holds database connection settings
type DatabaseConfig struct {
	Driver string `mapstructure:"driver"`
	Source string `mapstructure:"source"`
}

// RedisConfig Redis 配置
// RedisConfig holds redis connection settings
type RedisConfig struct {
	Addr     string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

// JWTConfig JWT 配置
// JWTConfig holds jwt settings
type JWTConfig struct {
	Secret string `mapstructure:"secret"`
	Expire int64  `mapstructure:"expire"` // 过期时间 (小时)
}

// GlobalConfig 全局配置变量
var GlobalConfig *Config

// LoadConfig 加载配置
// LoadConfig reads configuration from file
func LoadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")

	// 默认值用于本地快速启动
	viper.SetDefault("app.name", "godest")
	viper.SetDefault("app.port", ":8080")
	viper.SetDefault("app.mode", "debug")
	viper.SetDefault("database.driver", "sqlite")
	viper.SetDefault("database.source", "godest.db")
	viper.SetDefault("redis.addr", "localhost:6379")
	viper.SetDefault("redis.db", 0)
	viper.SetDefault("jwt.expire", 24)

	// 支持环境变量覆盖，例如 GODEST_APP_PORT=:9090
	viper.SetEnvPrefix("GODEST")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("配置文件读取失败，将仅使用默认值和环境变量: %v", err)
	}

	if err := viper.Unmarshal(&GlobalConfig); err != nil {
		log.Fatalf("配置解析失败: %v", err)
	}
}
