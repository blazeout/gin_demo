package main

//type Info struct {
//	gorm.Model
//	Money uint
//	DogID uint
//}
//
//type Dog struct {
//	gorm.Model
//	Name string
//	Infos Info
//	GirlGods []GirlGod `gorm:",many2many:dog_girl_god"`
//}
//
//type GirlGod struct {
//	gorm.Model
//	Name string
//	Dogs []Dog `gorm:"many2many:girl_god_dog"`
//}

//func OneToMany(){
//	//d1 := Dog{
//	//	Model : gorm.Model{ID: 1},
//	//	Name: "汪汪1",
//	//}
//	//d2 := Dog{
//	//	Model : gorm.Model{ID: 2},
//	//	Name: "汪汪2",
//	//}
//	//g := GirlGod{
//	//	Model : gorm.Model{
//	//		ID: 1,
//	//	},
//	//	Name: "女神1",
//	//	Dogs: []Dog{d1, d2},
//	//}
//	var girl GirlGod
//	// 此时没有带出来2个狗狗 需要预加载
//	// 汪汪1只带了狗没有带钱, 因为 在预加载的时候money>100
//	GlobalDb.Preload("Dogs", func(db *gorm.DB) *gorm.DB {
//		return db.Joins("Infos").Where("money > 200")
//	}).First(&girl)
//	fmt.Println(girl)
//
//}
