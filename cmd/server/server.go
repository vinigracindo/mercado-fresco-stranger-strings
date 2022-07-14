package server

import (
	"fmt"
	docs "github.com/vinigracindo/mercado-fresco-stranger-strings/docs/specs"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/libs/logger"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/cmd/server/http/ping"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/cmd/server/routes"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/config"
)

type APIServer struct{}

func NewAPIServer() APIServer {
	return APIServer{}
}

func (api *APIServer) Run(port int) {

	db := config.ConnectDb("mysql")
	defer db.Close()

	logger.InitializeLogger(db)

	router := gin.Default()

	// Swagger
	docs.SwaggerInfo.BasePath = "/api/v1"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Health check
	pingController := ping.NewController()
	router.GET("/ping", pingController.HandlePing)

	apiV1 := router.Group("api/v1")
	routes.SectionRoutes(apiV1.Group("/sections"), db)
	routes.EmployeeRoutes(apiV1.Group("/employees"), db)
	routes.InboundOrdersRoutes(apiV1.Group("/inboundOrders"), db)
	routes.ProductRoutes(apiV1.Group("/products"), db)
	routes.ProductRecordsRoutes(apiV1.Group("/productRecords"), db)
	routes.WarehouseRoutes(apiV1.Group("/warehouses"), db)
	routes.SellerRoutes(apiV1.Group("/sellers"), db)
	routes.BuyerRoutes(apiV1.Group("/buyers"), db)
	routes.CarryRoutes(apiV1.Group("/carries"), db)
	routes.LocalityRoutes(apiV1.Group("/localities"), db)
	routes.ProductBatchRoutes(apiV1.Group("/productBatches"), db)
	routes.PurchaseOrdersRoutes(apiV1.Group("/purchaseOrders"), db)

	router.Run(fmt.Sprintf(":%d", port))
}
