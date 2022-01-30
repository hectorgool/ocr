package main

import (
	"ocr/controller"

	docs "ocr/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io:8080
// @BasePath /v2
func main() {

	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	bunsanAPI := r.Group("/api/v1/bunsan")
	{
		bunsanAPI.GET("/", controller.GetRoot)
		bunsanAPI.GET("/number", controller.GetAccountNumber)
		bunsanAPI.POST("/numbers", controller.PostNumbers)
		bunsanAPI.GET("/logs", controller.GetLogs)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
