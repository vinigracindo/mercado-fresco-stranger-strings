package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/cmd/server/controllers"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/product"
)

func main() {

	router := gin.Default()

	groupV1 := router.Group("/api/v1/")

	// Product routes
	productRepository := product.CreateRepository()
	productService := product.CreateService(productRepository)
	productController := controllers.CreateProductController(productService)

	productGroup := groupV1.Group("products")
	productGroup.GET("/", productController.GetAll())
	productGroup.GET("/:id", productController.GetById())
	productGroup.POST("/", productController.Create())
	router.Run()
}
