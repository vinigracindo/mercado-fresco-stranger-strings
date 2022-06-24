package routes

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/vinigracindo/mercado-fresco-stranger-strings/cmd/server/controllers/product"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/product/repository"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/product/service"
)

func ProductRoutes(routes *gin.RouterGroup) {

	productRepository := repository.CreateProductRepository()
	productService := service.CreateProductService(productRepository)
	productController := controllers.CreateProductController(productService)

	routes.GET("/", productController.GetAll())
	routes.GET("/:id", productController.GetById())
	routes.POST("/", productController.Create())
	routes.PATCH("/:id", productController.UpdateDescription())
	routes.DELETE("/:id", productController.Delete())

}
