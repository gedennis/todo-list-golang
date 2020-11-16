package controllers

import (
	"net/http"
	"time"
	"todo/models"

	"github.com/gin-gonic/gin"
	"github.com/lithammer/shortuuid"
)

type CreateTodoRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}

type UpdateTodoRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func CreateTodo(c *gin.Context) {
	var todoItem CreateTodoRequest
	if err := c.ShouldBindJSON(&todoItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Create todo Failed",
		})
		return
	}
	var item = models.Todo{
		ID:          shortuuid.New(),
		CreateTime:  time.Now(),
		Title:       todoItem.Title,
		Description: todoItem.Description,
	}
	models.Insert(item)
	c.JSON(http.StatusOK, gin.H{
		"message": "Create todo OK",
	})

}
func GetTodoInfo(c *gin.Context) {
	id := c.Param("id")
	todo := models.Get(id)
	c.JSON(http.StatusOK, gin.H{
		"message": "Get todo list ok.",
		"data":    todo,
	})
}

func GetTodoList(c *gin.Context) {
	todos := models.GetAll()

	c.JSON(http.StatusOK, gin.H{
		"message": "Get todo list ok.",
		"data":    todos,
	})
}
func UpdateTodo(c *gin.Context) {
	var todoItem UpdateTodoRequest
	id := c.Param("id")
	if err := c.ShouldBindJSON(&todoItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Update todo Failed",
		})
		return
	}
	models.Update(id, todoItem.Title, todoItem.Description)
	c.JSON(http.StatusOK, gin.H{
		"message": "Update todo OK",
	})

}
func RemoveTodo(c *gin.Context) {
	id := c.Param("id")
	models.Delete(id)
	c.JSON(http.StatusOK, gin.H{
		"message": "Remove todo ok.",
		"data":    nil,
	})

}
