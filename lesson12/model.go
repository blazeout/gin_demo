package main

import (
	"gorm.io/gorm"
)

type TestUser struct {
	gorm.Model
	Name string `gorm:"default:wjh"`
	Age  uint8  `gorm:"comment:年龄"`
}

func UserTest() {
	GlobalDb.AutoMigrate(&TestUser{})
}
