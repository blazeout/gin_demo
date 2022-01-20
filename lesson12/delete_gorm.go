package main

func DeleteTest() {
	var user []TestUser
	GlobalDb.Unscoped().Where("age = ?", "0").Delete(&user)
}
