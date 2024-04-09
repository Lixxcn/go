package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// 定义数据表结构
type Message struct {
	gorm.Model
	Name    string `gorm:"type:varchar(255)"`
	Content string `gorm:"type:varchar(255)"`
}

// 数据库初始化设置
func Get_db() *gorm.DB {
	dsn := `root:1234@tcp(127.0.0.1:3306)/messagedb?charset=utf8mb4&parseTime=True&loc=Local`
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	sqlDB, _ := db.DB()
	// SetMaxIdleConns设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime设置连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Hour)
	db.AutoMigrate(&Message{})
	return db
}

// 定义全局变量Db
var Db = Get_db()
