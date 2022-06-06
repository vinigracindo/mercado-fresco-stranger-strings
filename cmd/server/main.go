package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/cmd/server/controllers"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/employees"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/section"
)

func main() {
	router := gin.Default()
  
  groupV1 := router.Group("api/v1")

  // Section routes
	repositorySection := section.NewRepository()
	serviceSection := section.NewService(repositorySection)
	controllerSection := controllers.NewSection(serviceSection)
  
  groupSection := groupV1.Group("/sections")
	groupSection.DELETE("/:id", controllerSection.Delete())
	groupSection.PATCH("/:id", controllerSection.UpdateCurrentCapacity())
	groupSection.POST("/", controllerSection.CreateSection())
	groupSection.GET("/:id", controllerSection.GetById())
	groupSection.GET("/", controllerSection.GetAll())
  
  // Employee routes
	employeeRepository := employees.NewRepository()
	employeeService := employees.NewService(employeeRepository)
	employeeController := controllers.NewEmployee(employeeService)

	employeeRouter := groupV1.Group("/employees")
	employeeRouter.GET("/", employeeController.GetAll())
	employeeRouter.GET("/:id", employeeController.GetById())
	employeeRouter.POST("/", employeeController.Store())
	employeeRouter.PATCH("/:id", employeeController.UpdateFullname())
	employeeRouter.DELETE("/:id", employeeController.Delete())

	router.Run()
}