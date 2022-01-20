package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var GlobalDb *gorm.DB

func main() {
	dsn := "root:admin@tcp(127.0.0.1:3306)/goday16?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("connect sql failed")
	}
	sqlDB, err := db.DB()
	if err != nil {
		return
	}
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(time.Hour)
	GlobalDb = db
	one2One()
}
