package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/cmd/server/controllers"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/section"
)

func main() {

	router := gin.Default()

	groupV1 := router.Group("api/v1/")

	// Section routes
	repositorySection := section.NewRepository()
	serviceSection := section.NewService(repositorySection)
	controllerSection := controllers.NewSection(serviceSection)

	groupSection := groupV1.Group("sections")
	groupSection.POST("/", controllerSection.CreateSection())

	router.Run()
}
