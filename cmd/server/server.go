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

	db := config.ConnectDb("mysql")
	defer db.Close()

	router := gin.Default()

	// Swagger
	docs.SwaggerInfo.BasePath = "/api/v1"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	apiV1 := router.Group("api/v1")

	routes.SectionRoutes(apiV1.Group("/sections"), db)
	routes.EmployeeRoutes(apiV1.Group("/employees"), db)
	routes.ProductRoutes(apiV1.Group("/products"), db)
	routes.WarehouseRoutes(apiV1.Group("/warehouses"), db)

	//Seller routes
	routes.SellerRoutes(apiV1.Group("/sellers"), db)
	routes.BuyerRoutes(apiV1.Group("/buyers"), db)
	routes.CarryRoutes(apiV1.Group("/carries"), db)
	routes.LocalityRoutes(apiV1.Group("/localities"), db)

	router.Run(fmt.Sprintf(":%d", port))
}
