package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	controllers "github.com/vinigracindo/mercado-fresco-stranger-strings/cmd/server/controllers/buyer"
	repository "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/buyer/repository/mariaDB"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/buyer/service"
)

func BuyerRoutes(routes *gin.RouterGroup, db *sql.DB) {
	buyerRepository := repository.NewmariadbBuyerRepository(db)
	buyerService := service.NewBuyerService(buyerRepository)
	buyerController := controllers.NewBuyerController(buyerService)

	routes.GET("/", buyerController.GetAll())
	routes.GET("/:id", buyerController.GetId())
	routes.POST("/", buyerController.Create())
	routes.PATCH("/:id", buyerController.UpdateCardNumberLastName())
	routes.DELETE("/:id", buyerController.DeleteBuyer())

}
