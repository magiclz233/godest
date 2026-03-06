package config

import (
	"strings"

	"godest/pkg/log"

	"go.uber.org/zap"
	"github.com/spf13/viper"
)

type Config struct {
	App      AppConfig      `mapstructure:"app"`
	Database DatabaseConfig `mapstructure:"database"`
	Redis    RedisConfig    `mapstructure:"redis"`
	JWT      JWTConfig      `mapstructure:"jwt"`
}

type AppConfig struct {
	Name string `mapstructure:"name"`
	Port string `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}

type DatabaseConfig struct {
	Driver string `mapstructure:"driver"`
	Source string `mapstructure:"source"`
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
	SSLMode string `mapstructure:"sslmode"`
}

type RedisConfig struct {
	Addr     string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type JWTConfig struct {
	Secret string `mapstructure:"secret"`
	Expire int64  `mapstructure:"expire"`
}

var GlobalConfig *Config

func LoadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./internal/config")
	viper.AddConfigPath("./config")

	viper.SetDefault("app.name", "godest")
	viper.SetDefault("app.port", ":8080")
	viper.SetDefault("app.mode", "debug")
	viper.SetDefault("database.driver", "sqlite")
	viper.SetDefault("database.source", "godest.db")
	viper.SetDefault("redis.addr", "localhost:6379")
	viper.SetDefault("redis.db", 0)
	viper.SetDefault("jwt.expire", 24)

	viper.SetEnvPrefix("GODEST")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Warn("read config file failed, using defaults/env", zap.Error(err))
	}

	if err := viper.Unmarshal(&GlobalConfig); err != nil {
		log.Fatal("unmarshal config failed", zap.Error(err))
	}
}
