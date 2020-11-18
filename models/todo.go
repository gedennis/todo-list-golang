package models

import (
	"fmt"
	"todo/db"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title       string
	Description string
	Status      uint
}

func CreateTodo(todo *Todo) (err error) {
	if err = db.DB.Create(todo).Error; err != nil {
		fmt.Println("Create todo failed with error: ", err.Error())
		return err
	}

	return nil
}
