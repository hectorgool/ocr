package main

import (
	"ocr/schema"
	"ocr/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": "ocr",
		})
	})

	r.GET("/number", func(c *gin.Context) {
		utils.PutInStruct()
		c.JSON(200, gin.H{
			"number": utils.PrintNumber(utils.GetNumber(schema.ACTNumber)),
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
