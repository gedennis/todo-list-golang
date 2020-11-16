package models

import (
	"time"
)

type Todo struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreateTime  time.Time `json:"create_time"`
}

var todos []Todo = make([]Todo, 0)

func GetAll() []Todo {
	return todos
}

func Insert(item Todo) bool {
	todos = append(todos, item)
	return true
}

func Get(id string) Todo {
	var result Todo
	for _, item := range todos {
		if item.ID == id {
			result = item
		}
	}
	return result
}

func Delete(id string) bool {
	var index int = -1
	for idx, item := range todos {
		if item.ID == id {
			index = idx
		}
	}
	if index == -1 {
		return true
	}
	todos = append(todos[:index], todos[index+1:]...)
	return true
}

func Update(id string, title string, description string) bool {
	var index int = -1
	for idx, item := range todos {
		if item.ID == id {
			index = idx
		}
	}
	if index == -1 {
		return true
	}
	todos[index].Title = title
	todos[index].Description = description

	return true
}
