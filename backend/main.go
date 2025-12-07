package main

import (
	"net/http"

	"github.com/gin-gonic/gin" // นำเข้าแพ็คเกจ Gin
)

func main() {
	router := gin.Default() // สร้าง router ด้วย Gin

	router.Use(func(c *gin.Context){
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello from Gin!"})
	})

	router.Run(":8080") // รันเซิร์ฟเวอร์บน port 8080
}
