package routers

import (
	"todo/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// health check
	r.GET("/check", controllers.Check)

	// todo routers
	todoAPI := r.Group("/api/todos")
	{
		todoAPI.POST("/", controllers.CreateTodo)
		// todoAPI.GET("/:id", controllers.GetTodoInfo)
		// todoAPI.GET("/", controllers.GetTodoList)
		// todoAPI.PUT("/:id", controllers.UpdateTodo)
		// todoAPI.DELETE("/:id", controllers.RemoveTodo)
	}

	return r
}
