package main

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/cmd/server/controllers"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/cmd/server/docs"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/section"
)

// @title           Swagger Mercado Fresco
// @version         1.0
// @description     This is a sample server celler server.

// @host      localhost:8080
// @BasePath  /api/v1

func main() {

	router := gin.Default()
	groupV1 := router.Group("api/v1/")

	// Swagger routes
	docs.SwaggerInfo.Host = "localhost:8080"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Section routes
	repositorySection := section.NewRepository()
	serviceSection := section.NewService(repositorySection)
	controllerSection := controllers.NewSection(serviceSection)

	groupSection := groupV1.Group("sections")
	groupSection.GET("/", controllerSection.GetAll())

	router.Run()
}
