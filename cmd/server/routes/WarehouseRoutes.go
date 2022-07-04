package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	controllers "github.com/vinigracindo/mercado-fresco-stranger-strings/cmd/server/controllers/warehouse"
	repository "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/warehouse/repository/mariadb"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/warehouse/services"
)

func WarehouseRoutes(routes *gin.RouterGroup, db *sql.DB) {
	warehouseRepository := repository.NewMariadbWarehouseRepository(db)
	warehouseService := services.NewWarehouseService(warehouseRepository)
	warehouseController := controllers.NewWarehouse(warehouseService)

	routes.GET("/", warehouseController.GetAllWarehouse())
	routes.GET("/:id", warehouseController.GetWarehouseByID())
	routes.POST("/", warehouseController.CreateWarehouse())
	routes.DELETE("/:id", warehouseController.DeleteWarehouse())
	routes.PATCH("/:id", warehouseController.UpdateWarehouse())
}
