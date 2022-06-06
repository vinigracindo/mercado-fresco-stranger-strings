package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/cmd/server/controllers"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/seller"
)

func main() {

	router := gin.Default()

	//Seller
	sellerRepository := seller.NewRepository()
	sellerService := seller.NewService(sellerRepository)
	sellerController := controllers.NewSeller(sellerService)
	v1 := router.Group("/api/v1")
	sellerRouter := v1.Group("/seller")

	sellerRouter.GET("/", sellerController.GetAll())
	sellerRouter.GET("/:id", sellerController.Get())
	sellerRouter.POST("/", sellerController.CreateSeller())
	sellerRouter.PATCH("/:id", sellerController.UpdateSeller())

	router.Run()
}
