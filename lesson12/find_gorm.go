package main

import (
	"fmt"
)

type UserInfo struct {
	Name string
	Age  uint8
}

func FindTest() {
	// gorm 查询
	// 1. 使用map接收查询返回值
	//ret := map[string]interface{}{}
	// 2. 使用结构体接收
	//var User []TestUser
	var U []UserInfo
	// 查询方法
	//dbRes := GlobalDb.Where("name = ? and age = 22" ,"wjh").Last(&User)
	//dbRes := GlobalDb.Where(&TestUser{Name: "wjh", Age: 222}).Or("name = ?", "王嘉豪").First(&User)

	// 内联条件
	// 1. string 形式
	// GlobalDb.First(&User, "name = ?", "wjh")
	// 2.map形式
	// GlobalDb.First(&User,map[string]interface{}{"name" : "wjh"})
	// 3.结构体形式
	// GlobalDb.First(&User, TestUser{Name: "wjh"})
	//dbRes := GlobalDb.Where("name <> ?", "王嘉豪").Find(&User)
	dbRes := GlobalDb.Model(&TestUser{}).Where("name like ?", "%%豪").Find(&U)
	fmt.Println(U)
	fmt.Println(dbRes.Error)
	// 检索条件的方法
	//is := errors.Is(dbRes.Error, gorm.ErrRecordNotFound)
	//fmt.Println(is)
}
