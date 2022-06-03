package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/cmd/server/controllers"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/warehouse"
)

func main() {
	router := gin.Default()

	warehouseRepo := warehouse.NewRepository()
	warehouseService := warehouse.NewService(warehouseRepo)
	warehouseController := controllers.NewWarehouse(warehouseService)

	routesWarehouse := router.Group("/api/v1/warehouses")
	{
		routesWarehouse.POST("/", warehouseController.CreateWarehouse())
	}

	router.Run()
}
