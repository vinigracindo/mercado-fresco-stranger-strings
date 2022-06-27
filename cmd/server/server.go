package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/cmd/server/controllers"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/cmd/server/routes"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/docs"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/buyer"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/employees"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/product"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/section"
)

type APIServer struct{}

func NewAPIServer() APIServer {
	return APIServer{}
}

func (api *APIServer) Run(port int) {
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
	routes.SellerRoutes(apiV1.Group("/sellers"))

	//Buyer routes
	routes.BuyerRoutes(apiV1.Group("/buyers"))

	router.Run(fmt.Sprintf(":%d", port))
}
