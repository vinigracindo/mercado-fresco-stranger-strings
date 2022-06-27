package routes

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/vinigracindo/mercado-fresco-stranger-strings/cmd/server/controllers/warehouse"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/warehouse/repository"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/warehouse/services"
)

func WarehouseRoutes(routes *gin.RouterGroup) {
	warehouseRepository := repository.NewWarehouseRepository()
	warehouseService := services.NewWarehouseService(warehouseRepository)
	warehouseController := controllers.NewWarehouse(warehouseService)

	routes.GET("/", warehouseController.GetAllWarehouse())
	routes.GET("/:id", warehouseController.GetWarehouseByID())
	routes.POST("/", warehouseController.CreateWarehouse())
	routes.DELETE("/:id", warehouseController.DeleteWarehouse())
	routes.PATCH("/:id", warehouseController.UpdateWarehouse())
}
