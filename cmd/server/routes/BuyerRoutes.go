package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	controllers "github.com/vinigracindo/mercado-fresco-stranger-strings/cmd/server/controllers/buyer"
	repository "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/buyer/repository/mariaDB"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/buyer/service"
	purchaseOrdersRecords "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/purchase_orders/repository/mariaDB"
)

func BuyerRoutes(routes *gin.RouterGroup, db *sql.DB) {
	buyerRepository := repository.NewmariadbBuyerRepository(db)
	purchaseOrdersRepository := purchaseOrdersRecords.NewMariadbPurchaseOrdersRepository(db)

	buyerService := service.NewBuyerService(buyerRepository, purchaseOrdersRepository)
	buyerController := controllers.NewBuyerController(buyerService)

	routes.GET("/reportPurchaseOrders", buyerController.GetPurchaseOrdersReports())

	routes.GET("/", buyerController.GetAll())
	routes.GET("/:id", buyerController.GetId())
	routes.POST("/", buyerController.Create())
	routes.PATCH("/:id", buyerController.UpdateCardNumberLastName())
	routes.DELETE("/:id", buyerController.DeleteBuyer())

}
