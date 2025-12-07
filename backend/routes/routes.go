package routes

import (
	"gin-api/controllers"
	"gin-api/database"
	"gin-api/models"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	database.ConnectDB()

	database.DB.AutoMigrate(&models.User{})

	api := r.Group("/api")
	{
		api.GET("/users", controllers.GetUsers)
		api.POST("/users", controllers.CreateUsers)
	}
}
