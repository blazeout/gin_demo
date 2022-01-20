package main

import "fmt"

func CreateTest() {
	//dbRes := GlobalDb.Omit("Name").Create(&TestUser{
	//	Age: 20,
	//})
	dbRes := GlobalDb.Create(&[]TestUser{
		{Name: "王嘉豪", Age: 19},
		{Name: "halo", Age: 18},
		{Name: "sbb", Age: 20},
		{Name: "nitm", Age: 9},
		{Name: "zhiz", Age: 29},
	})
	fmt.Println(dbRes.Error, " ", dbRes.RowsAffected)
	if dbRes.Error != nil {
		fmt.Println("创建失败了.孙贼")
	} else {
		fmt.Println("创建成功")
	}
}
