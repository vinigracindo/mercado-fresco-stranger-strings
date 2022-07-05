package server

import (
	"fmt"

	"github.com/vinigracindo/mercado-fresco-stranger-strings/config"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/cmd/server/routes"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/docs"
)

type APIServer struct{}

func NewAPIServer() APIServer {
	return APIServer{}
}

func (api *APIServer) Run(port int) {

	db := config.ConnectDb()
	defer db.Close()

	router := gin.Default()

	// Swagger
	docs.SwaggerInfo.BasePath = "/api/v1"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	apiV1 := router.Group("api/v1")

	//Section routes
	routes.SectionRoutes(apiV1.Group("/sections"))

	// Employee routes
	routes.WarehouseRoutes(apiV1.Group("/employees"))

	// Product routes
	routes.ProductRoutes(apiV1.Group("/products"))

	//Warehouse routes
	routes.WarehouseRoutes(apiV1.Group("/warehouses"))

	//Seller routes
	routes.SellerRoutes(apiV1.Group("/sellers"), db)

	//Buyer routes
	routes.BuyerRoutes(apiV1.Group("/buyers"))

	router.Run(fmt.Sprintf(":%d", port))
}
