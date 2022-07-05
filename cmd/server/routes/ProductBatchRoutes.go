package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	controllers "github.com/vinigracindo/mercado-fresco-stranger-strings/cmd/server/controllers/product_batch"
	productRepository "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/product/repository/mariadb"
	repository "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/product_batch/repository/mariadb"
	sectionRepository "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/section/repository/mariadb"

	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/product_batch/service"
)

func ProductBatchRoutes(routes *gin.RouterGroup, db *sql.DB) {
	productBatchRepository := repository.NewMariadbProductBatchRepository(db)
	productRepository := productRepository.CreateProductRepository(db)
	sectionRepository := sectionRepository.NewMariadbSectionRepository(db)

	productBatchService := service.NewProductBatchService(productBatchRepository, productRepository, sectionRepository)
	productBatchController := controllers.NewProductBatchController(productBatchService)

	routes.POST("/", productBatchController.Create())
}
