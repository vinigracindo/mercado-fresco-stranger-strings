package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/cmd/server/controllers"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/product"
)

func main() {
	router := gin.Default()

	// define o base path
	groupV1 := router.Group("/api/v1/")

	// Product routes
	productRepository := product.NewRepository()
	productService := product.NewService(productRepository)
	productController := controllers.NewProduct(productService)

	groupProduct := groupV1.Group("products")
	groupProduct.GET("/", productController.GetAll())
	groupProduct.GET("/:id", productController.GetById())
	router.Run()
}
