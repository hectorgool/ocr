package controller

import (
	"net/http"
	"ocr/schema"
	"ocr/utils"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// PostNumbers godoc
// @Summary send data to the server to proces the numbers of the OCR
// @Schemes LogDB
// @Description Send data to the server to proces the numbers of the OCR, the list of numbers must be code in base64 for proces
// @Tags bunsan
// @Accept json
// @Produce json
// @Success 200 {object} c.JSON(http.StatusOK, jsonOutput)
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

// GetLogs godoc
// @Summary Get a json with the fields in the table log_db of database
// @Schemes LogDB
// @Description Get a json with, id, Endpoint, Method, Input, Output, CreatedOn, stored in log_db of database
// @Tags bunsan
// @Produce json
// @Success 200 {object} GetLogs
// @Router /logs [get]
func GetLogs(c *gin.Context) {

	c.JSON(200, gin.H{
		"logs": utils.GetLogRows(),
	})

}
