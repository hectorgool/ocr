package controller

import (
	"net/http"
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

	c.JSON(200, gin.H{
		"accounts": utils.ArrayStringsValidate(),
	})

}

func PostNumbers(c *gin.Context) {

	var json schema.JsonRequest

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	output := utils.ArrayStringsValidateWithInput(utils.SplitStringByCharToArray(utils.DecodeBase64String(json.Numbers), "\n"))
	jsonResult := gin.H{"numbers": output}

	utils.CreateLog("/numbers", "POST", json, jsonResult)
	c.JSON(http.StatusOK, jsonResult)

}
