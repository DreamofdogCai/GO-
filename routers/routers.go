package routers

import (
	"bubble_practice2/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouters() *gin.Engine {
	r := gin.Default()

	r.Static("/static", "static")

	r.LoadHTMLGlob("templates/*")

	r.GET("/", controller.IndexHandler)

	//v1
	v1Group := r.Group("/v1")
	//CRUD
	{
		v1Group.POST("/todo", controller.CreateTodo)

		v1Group.GET("/todo", controller.GetALLTodos)

		v1Group.PUT("/todo/:id", controller.UpdateATodo)

		v1Group.DELETE("/todo/:id", controller.DeleteATodo)
	}

	return r

}
