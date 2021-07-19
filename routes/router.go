package routes

import (
	Controller "Todo/app/v1/controller"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.GET("/ping", Controller.Ping)

	user := router.Group("/user")
	{
		user.POST("/register", Controller.Register)
		user.POST("/login", Controller.DoLogin)
		user.POST("/logout", Controller.DoLogout)
	}

	todo := router.Group("/todo")
	{
		todo.POST("/add", Controller.AddTodo)
		todo.GET("/list", Controller.GetUsersTodoList)
		todo.PUT("/:id", Controller.UpdateTodoItem)
		todo.DELETE("/:id", Controller.DeleteTodoItem)
	}

	TodoGroup := router.Group("/todo_group")
	{
		TodoGroup.POST("/add", Controller.AddTodoGroup)
		TodoGroup.GET("/list", Controller.GetAllTodoGroup)
		TodoGroup.PUT("/:id", Controller.UpdateTodoGroup)
		TodoGroup.DELETE("/:id", Controller.DeleteTodoGroup)
	}

}
