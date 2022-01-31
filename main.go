package main

import (
	"ocr/controller"

	docs "ocr/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Lector OCR API
// @version 1.0
// @description This is a sample Lector OCR server.
// @termsOfService http://swagger.io/terms/

// @contact.name Héctor González
// @contact.url http://github.com/hectorgool/ocr
// @contact.email hector.gonzalez@bunsan.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
func main() {

	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	bunsanAPI := r.Group("/api/v1/bunsan")
	{
		bunsanAPI.POST("/numbers", controller.PostNumbers)
		bunsanAPI.GET("/logs", controller.GetLogs)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
