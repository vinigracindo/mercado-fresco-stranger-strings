package routes

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	controllers "github.com/vinigracindo/mercado-fresco-stranger-strings/cmd/server/controllers/product"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/product/repository/mariadb"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/product/service"

	mariadbProductRecords "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/product_records/repository/mariadb"
)

func ProductRoutes(routes *gin.RouterGroup, db *sql.DB) {

	productRepository := mariadb.CreateProductRepository(db)
	productRecordsRepository := mariadbProductRecords.CreateProductRecordsRepository(db)

	productService := service.CreateProductService(productRepository, productRecordsRepository)
	productController := controllers.CreateProductController(productService)

	routes.GET("/", productController.GetAll())
	routes.GET("/:id", productController.GetById())
	routes.POST("/", productController.Create())
	routes.PATCH("/:id", productController.UpdateDescription())
	routes.DELETE("/:id", productController.Delete())

	routes.GET("/reportRecords", productController.GetReportProductRecords())

}
