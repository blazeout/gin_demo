package models

import "gin_demo/lesson15/database_models"

// Todo Model
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func CreateData(todo *Todo) (err error) {
	if err = database_models.GlobalDB.Create(&todo).Error; err != nil {
		return nil
	}
	return
}

func ShowData(todoList *[]Todo) (err error) {
	if err = database_models.GlobalDB.Find(&todoList).Error; err != nil {
		return nil
	}
	return
}

func UpdateTodo(todo *Todo, id string) (err error) {
	if err = database_models.GlobalDB.Where("id = ?", id).First(&todo).Error; err != nil {
		return nil
	}
	return
}

func Save(todo Todo) {
	database_models.GlobalDB.Save(todo)
}

func DeleteTodo(todo *Todo, id string) (err error) {
	if err = database_models.GlobalDB.Where("id = ?", id).Delete(&todo).Error; err != nil {
		return nil
	}
	return

}
