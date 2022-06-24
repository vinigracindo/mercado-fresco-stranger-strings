package routes

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/vinigracindo/mercado-fresco-stranger-strings/cmd/server/controllers/buyer"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/buyer/repository"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/buyer/service"
)

func BuyerRoutes(routes *gin.RouterGroup) {
	buyerRepository := repository.NewBuyerRepository()
	buyerService := service.NewBuyerService(buyerRepository)
	buyerController := controllers.NewBuyer(buyerService)

	routes.GET("/", buyerController.GetAll())
	routes.GET("/:id", buyerController.GetId())
	routes.POST("/", buyerController.Create())
	routes.PATCH("/:id", buyerController.UpdateCardNumberLastName())
	routes.DELETE("/:id", buyerController.DeleteBuyer())

}
