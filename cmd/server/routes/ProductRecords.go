package routes

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	controllers "github.com/vinigracindo/mercado-fresco-stranger-strings/cmd/server/controllers/product_records"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/product_records/repository/mariadb"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/product_records/service"

	mariadbProduct "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/product/repository/mariadb"
)

func ProductRecordsRoutes(routes *gin.RouterGroup, db *sql.DB) {

	productRecordsRepository := mariadb.CreateProductRecordsRepository(db)
	productRepository := mariadbProduct.CreateProductRepository(db)

	productRecordsService := service.CreateProductRecordsService(productRecordsRepository, productRepository)
	productRecordsController := controllers.CreateProductRecordsController(productRecordsService)

	routes.POST("/", productRecordsController.Create())

}
