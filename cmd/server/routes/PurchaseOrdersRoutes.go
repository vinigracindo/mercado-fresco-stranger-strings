package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	controllers "github.com/vinigracindo/mercado-fresco-stranger-strings/cmd/server/controllers/purchase_orders"
	BuyerRepository "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/buyer/repository/mariaDB"
	PurchaseOrdersRepository "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/purchase_orders/repository/mariaDB"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/purchase_orders/service"
)

func PurchaseOrdersRoutes(routes *gin.RouterGroup, db *sql.DB) {

	buyerRepository := BuyerRepository.NewmariadbBuyerRepository(db)
	purchaseOrdersRepository := PurchaseOrdersRepository.NewMariadbPurchaseOrdersRepository(db)

	purchaseOrdersService := service.NewPurchaseOrdersService(purchaseOrdersRepository, buyerRepository)
	purchaseOrdersController := controllers.NewPurchaseOrdersController(purchaseOrdersService)

	routes.POST("/", purchaseOrdersController.Create())
}
