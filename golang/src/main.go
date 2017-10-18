package main

import (
	"os"
  "github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hallo Docker Berlin! 🐳",
		})
	})

	r.Run(":" + port)
}
