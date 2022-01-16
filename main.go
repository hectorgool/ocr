package main

import (
	"ocr/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": "ocr",
		})
	})
	r.GET("/number", controller.GetAccountNumber)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
