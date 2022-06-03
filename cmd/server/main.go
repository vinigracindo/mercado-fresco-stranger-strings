package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/cmd/server/controllers"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/section"
)

func main() {

	router := gin.Default()

	// Section routes
	repositorySection := section.NewRepository()
	serviceSection := section.NewService(repositorySection)
	controllerSection := controllers.NewSection(serviceSection)

	groupSection := router.Group("/sections")
	groupSection.GET("/", controllerSection.GetAll())

	router.Run()
}
