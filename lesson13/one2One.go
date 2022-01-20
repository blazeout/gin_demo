package main

import (
	"fmt"
	"gorm.io/gorm"
)

// belongsTo
type Dog struct {
	gorm.Model
	Name     string
	BeautyID uint // 狗链子

}

type Beauty struct {
	gorm.Model
	Name string
	Dog  Dog
}

func one2One() {
	//d := Dog{
	//	Model : gorm.Model{ID: 1},
	//}
	//d2 := Dog{
	//	Model : gorm.Model{ID: 2},
	//}
	b := Beauty{
		Model: gorm.Model{
			ID: 1,
		},
	}
	//b2 := Beauty{
	//	Model : gorm.Model{
	//		ID: 2,
	//	},
	//}
	//GlobalDb.Create(&b)
	//b := Beauty {
	//	Name: "猫猫",
	//	Dog : d,
	//}
	//GlobalDb.Create(&b)
	//var grail Beauty
	//GlobalDb.Preload("Dog").Find(&grail, 2)
	GlobalDb.Model(&b).Association("Dog").Clear()
	fmt.Println()
}
