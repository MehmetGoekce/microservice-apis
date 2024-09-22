package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"web-service-gin/api"
)

// @title Orders API
// @version 1.0
// @description API that allows you to manage orders for Kaffeehaus
// @host localhost:8080
// @BasePath /api
func main() {
	router := gin.Default()
	// Mount Swagger UI
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// Define API routes
	api.MountRoutes(router)
	router.Run(":8080")
}
