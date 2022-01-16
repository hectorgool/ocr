package controller

import (
	"ocr/schema"
	"ocr/utils"

	"github.com/gin-gonic/gin"
)

func GetRoot(c *gin.Context) {

	c.JSON(200, gin.H{
		"data": "ocr",
	})

}

func GetAccountNumber(c *gin.Context) {

	utils.PutInStruct()
	c.JSON(200, gin.H{
		"number": utils.PrintNumber(utils.GetNumber(schema.ACTNumber)),
	})

}
