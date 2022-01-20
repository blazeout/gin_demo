package main

// 更新

func UpdateTest() {
	// update 只更新你选择的字段
	// updates 更新所有字段 此时有两种形式 一种为map 一种为结构体 零值不参与更新
	// save 不论如何都更新 所有的内容 包括0值

	// 1. update : GlobalDb.Model(&TestUser{}).Where("age = ? ", "20").Update("name", "王嘉豪")

	// 2. save : var User []TestUser
	// dbRes := GlobalDb.Where("name = ?", "王嘉豪").Find(&User)
	// for k  := range User {
	//	User[k].Age = 10
	// }
	//dbRes.Save(&User)

	// 3.updates
	// Map
	// struct
	var user []TestUser
	GlobalDb.Find(&user).Updates(map[string]interface{}{"name": " ", "age": 0})

}
