package database

import (
	"fmt"
	"gin_demo/lesson15/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var GlobalDB *gorm.DB

func InitMysql() {
	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root:root@tcp(127.0.0.1:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	sqlDB, err := db.DB()
	if err != nil {
		return
	}
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(time.Hour)
	GlobalDB = db
	// 模型绑定
	err = GlobalDB.AutoMigrate(&models.Todo{})
	if err != nil {
		return
	}
	fmt.Println("连接成功")
}
