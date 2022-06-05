package main

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/cmd/server/controllers"
	docs "github.com/vinigracindo/mercado-fresco-stranger-strings/cmd/server/docs"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/warehouse"
)

func main() {
	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"

	warehouseRepo := warehouse.NewRepository()
	warehouseService := warehouse.NewService(warehouseRepo)
	warehouseController := controllers.NewWarehouse(warehouseService)

	routesWarehouse := router.Group("/api/v1/warehouses")
	{
		routesWarehouse.GET("/", warehouseController.GetAllWarehouse())
		routesWarehouse.GET("/:id", warehouseController.GetWarehouseByID())
		routesWarehouse.POST("/", warehouseController.CreateWarehouse())
		routesWarehouse.DELETE("/:id", warehouseController.DeleteWarehouse())
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run()
}
