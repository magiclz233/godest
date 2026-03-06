package database

import (
	"fmt"

	"godest/internal/config"
	"godest/pkg/log"

	"github.com/glebarez/sqlite"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	var dialector gorm.Dialector
	conf := config.GlobalConfig.Database

	switch conf.Driver {
	case "postgres":
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
			conf.Host, conf.Username, conf.Password, conf.Database, conf.Port, conf.SSLMode)
		dialector = postgres.Open(dsn)
	case "sqlite":
		dialector = sqlite.Open(conf.Source)
	default:
		log.Fatal("unsupported database driver", zap.String("driver", conf.Driver))
	}
	var err error
	DB, err = gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		log.Fatal("connect database failed", zap.Error(err))
	}
}
