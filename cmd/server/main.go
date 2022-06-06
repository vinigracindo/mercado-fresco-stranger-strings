package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/cmd/server/controllers"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/product"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/employees"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/section"
)

func main() {

	router := gin.Default()


	groupV1 := router.Group("api/v1")

	// Section routes
	sectionRepository := section.NewRepository()
	sectionService := section.NewService(sectionRepository)
	sectionController := controllers.NewSection(sectionService)

	sectionGroup := groupV1.Group("/sections")
	sectionGroup.DELETE("/:id", sectionController.Delete())
	sectionGroup.PATCH("/:id", sectionController.UpdateCurrentCapacity())
	sectionGroup.POST("/", sectionController.CreateSection())
	sectionGroup.GET("/:id", sectionController.GetById())
	sectionGroup.GET("/", sectionController.GetAll())

	// Employee routes
	employeeRepository := employees.NewRepository()
	employeeService := employees.NewService(employeeRepository)
	employeeController := controllers.NewEmployee(employeeService)

	employeeGroup := groupV1.Group("/employees")
	employeeGroup.GET("/", employeeController.GetAll())
	employeeGroup.GET("/:id", employeeController.GetById())
	employeeGroup.POST("/", employeeController.Store())
	employeeGroup.PATCH("/:id", employeeController.UpdateFullname())
	employeeGroup.DELETE("/:id", employeeController.Delete())
  
  // Product routes
	productRepository := product.CreateRepository()
	productService := product.CreateService(productRepository)
	productController := controllers.CreateProductController(productService)

	productGroup := groupV1.Group("/products")
	productGroup.GET("/", productController.GetAll())
	productGroup.GET("/:id", productController.GetById())
	productGroup.POST("/", productController.Create())
	productGroup.PATCH("/:id", productController.UpdateDescription())
	productGroup.DELETE("/:id", productController.Delete())


	router.Run()
}
