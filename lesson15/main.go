package main

import (
	"gin_demo/lesson15/database"
	"gin_demo/lesson15/database_models"
	"gin_demo/lesson15/router"
)

func main() {
	// 创建数据库
	// sql: create database bubble;
	// 程序启动时需要链接数据库
	database_models.ZhuanHuanG()
	database.InitMysql()
	r := router.SetUpRouter()
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
