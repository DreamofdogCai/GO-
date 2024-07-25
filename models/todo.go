package models

import "bubble_practice2/dao"

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func CreateATable() (err error) {
	var todo Todo
	err = dao.DB.AutoMigrate(&todo)
	return
}

func CreateATodo(todo *Todo) (err error) {
	err = dao.DB.Create(todo).Error
	return
}

func SelectTodos() (todos *[]Todo, err error) {
	err = dao.DB.Find(&todos).Error
	if err != nil {
		return nil, err
	}
	return
}

func SelectById(id string) (todo *Todo, err error) {
	err = dao.DB.Where("id = ?", id).First(&todo).Error
	if err != nil {
		return nil, err
	}
	return
}

func SaveATodo(todo *Todo) (err error) {
	err = dao.DB.Save(todo).Error
	return
}

func DeleteById(id string) (err error) {
	err = dao.DB.Where("id = ?", id).Delete(&Todo{}).Error
	return
}
