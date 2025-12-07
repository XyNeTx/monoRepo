package controllers

import (
	"gin-api/services"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	user, err := services.GetAllUsers()
	if err != nil {
		panic("Can't Get All Users")
	}
	c.JSON(200, gin.H{"data": user})
}
func CreateUsers(c *gin.Context) {
	err := services.CreateUsers()
	if err != nil {
		panic("Can't Create Users")
	}
	c.JSON(200, gin.H{"message": "Success"})
}
