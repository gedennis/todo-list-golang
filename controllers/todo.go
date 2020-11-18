package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"todo/models"

	"github.com/gin-gonic/gin"
)

type CreateTodoRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}

// create todo
func CreateTodo(c *gin.Context) {
	// check request
	var todoItem CreateTodoRequest
	if err := c.ShouldBindJSON(&todoItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Invalid Request",
		})
		return
	}

	// create
	var todo = models.Todo{
		Title:       todoItem.Title,
		Description: todoItem.Description,
		Status:      0,
	}

	// response
	if err := models.CreateTodo(&todo); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Failed to create todo",
			"data":    todo,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Create todo OK",
		"data":    todo,
	})
}

func GetTodoInfo(c *gin.Context) {
	id := c.Param("id")

	var todo = models.Todo{}
	if err := models.GetTodo(&todo, id); err != nil {
		fmt.Println(todo)
		c.JSON(http.StatusOK, gin.H{
			"message": "Failed to get todo",
			"data":    nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Get todo info ok.",
		"data":    todo,
	})
}

func GetTodoList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "5"))

	todoList := []models.Todo{}
	if err := models.GetTodoList(&todoList, page, size); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Failed to get todo list",
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Get todo list ok.",
		"data":    todoList,
	})
}
func UpdateTodo(c *gin.Context) {
	// var todoItem UpdateTodoRequest
	// // id := c.Param("id")
	// if err := c.ShouldBindJSON(&todoItem); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"error":   err.Error(),
	// 		"message": "Update todo Failed",
	// 	})
	// 	return
	// }
	// models.Update(id, todoItem.Title, todoItem.Description)
	c.JSON(http.StatusOK, gin.H{
		"message": "Update todo OK",
	})

}
func RemoveTodo(c *gin.Context) {
	// id := c.Param("id")
	// models.Delete(id)
	c.JSON(http.StatusOK, gin.H{
		"message": "Remove todo ok.",
		"data":    nil,
	})

}
