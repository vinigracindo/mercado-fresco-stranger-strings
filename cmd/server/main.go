package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/cmd/server/controllers"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/buyer"
)

func main() {

	router := gin.Default()

	// define o base path
	groupV1 := router.Group("/api/v1/")

	// Buyer routes
	buyerRepository := buyer.NewRepository()
	buyerService := buyer.NewService(buyerRepository)
	buyerController := controllers.NewBuyer(buyerService)

	groupProduct := groupV1.Group("buyers")
	groupProduct.GET("/", buyerController.GetAll())
	groupProduct.GET("/:id", buyerController.GetId())
	groupProduct.POST("/", buyerController.Store())

	router.Run()
}
