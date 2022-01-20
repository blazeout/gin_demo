package database_models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Todo Model
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

var GlobalDB *gorm.DB

func ZhuanHuanG() {
	dsn := "root:root@tcp(127.0.0.1:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	GlobalDB = db
}
