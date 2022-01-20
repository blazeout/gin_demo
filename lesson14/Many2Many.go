package main

import (
	"fmt"
	"gorm.io/gorm"
)

type Info struct {
	gorm.Model
	Money uint
	DogID uint
}

type Dog struct {
	gorm.Model
	Name     string
	Infos    Info
	GirlGods []GirlGod `gorm:"many2many:dog_girl_god"`
}

type GirlGod struct {
	gorm.Model
	Name string
	Dogs []Dog `gorm:"many2many:dog_girl_god"`
}

func ManyToMany() {
	i := Info{
		Money: 200,
	}
	//g1 := GirlGod{
	//	Model: gorm.Model{
	//		ID: 1,
	//	},
	//	Name: "绿茶1",
	//}
	//g2 := GirlGod{
	//	Model: gorm.Model{
	//		ID: 2,
	//	},
	//	Name: "绿茶2",
	//}
	//d1 := Dog{
	//	Model: gorm.Model{ID: 2},
	//	Name:  "汪汪2",
	//	GirlGods: []GirlGod{g1, g2},
	//	Infos: i,
	//}
	d := Dog{
		Model: gorm.Model{
			ID: 3,
		},
		Name:  "汪汪3",
		Infos: i,
	}
	//girls := make([]GirlGod,0)
	//GlobalDb.Model(&d).Preload("Dogs", func(db *gorm.DB) *gorm.DB {
	//	return db.Joins("Infos").Where("money < ?", 1000)
	//}).Association("GirlGods").Find(&girls)
	//fmt.Println(girls)
	err := GlobalDb.Model(&d).Association("GirlGods").Clear()
	if err != nil {
		fmt.Println("================")
		fmt.Println(err)
		return
	}
}
