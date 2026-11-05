package routes

import (
	"golang-user-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(
	router *gin.Engine,
	userController *controllers.UserController,
) {
	api := router.Group("/api/v1")

	{
		api.GET("/users", userController.GetAll)
		api.GET("/users/:id", userController.GetByID)
		api.POST("/users", userController.Create)
		api.PUT("/users/:id", userController.Update)
		api.DELETE("/users/:id", userController.Delete)
	}
}
