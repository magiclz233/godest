package database

import (
	"log"

	"go_pro/config"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

// DB 全局数据库连接对象
// Global database connection instance
var DB *gorm.DB

// Init 初始化数据库连接
// Init initializes the database connection
func Init() {
	var err error
	// 从配置中获取数据库源
	dsn := config.GlobalConfig.Database.Source

	// 打开数据库连接 (这里使用 SQLite，如果是 MySQL 可以换成 mysql.Open(dsn))
	// Open database connection
	DB, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("无法连接数据库: %v", err)
	}

	// 可以在这里设置连接池配置
	// sqlDB, _ := DB.DB()
	// sqlDB.SetMaxIdleConns(10)
	// sqlDB.SetMaxOpenConns(100)
}
