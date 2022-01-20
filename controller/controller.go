package controller

import (
	"ocr/utils"

	"github.com/gin-gonic/gin"
)

func GetRoot(c *gin.Context) {

	c.JSON(200, gin.H{
		"data": "ocr",
	})

}

func GetAccountNumber(c *gin.Context) {

	c.JSON(200, gin.H{
		"accounts": utils.ArrayStringsValidate(),
	})

}
