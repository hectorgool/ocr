package controller

import (
	"net/http"
	"ocr/schema"
	"ocr/utils"

	"github.com/gin-gonic/gin"
	// gin-swagger middleware
	//"github.com/swaggo/gin-swagger" // gin-swagger middleware
	//"github.com/swaggo/files" // gin-swagger middleware
)

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router / [get]
func GetRoot(c *gin.Context) {

	c.JSON(200, gin.H{
		"data": "ocr",
	})

}

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /number [get]
func GetAccountNumber(c *gin.Context) {

	c.JSON(200, gin.H{
		"accounts": utils.ArrayStringsValidate(),
	})

}

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /numbers [post]
func PostNumbers(c *gin.Context) {

	var json schema.JsonRequest

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	output := utils.ArrayStringsValidateWithInput(utils.SplitStringByCharToArray(utils.DecodeBase64String(json.Numbers), "\n"))
	jsonInput := gin.H{"input": json}
	jsonOutput := gin.H{"output": output}

	utils.CreateLog("/numbers", "POST", jsonInput, jsonOutput)
	c.JSON(http.StatusOK, jsonOutput)

}

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /logs [get]
func GetLogs(c *gin.Context) {

	c.JSON(200, gin.H{
		"logs": utils.GetLogRows(),
	})

}
