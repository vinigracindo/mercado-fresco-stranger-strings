package server

import (
	"database/sql"
	"fmt"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/config"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/cmd/server/routes"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/docs"
)

type APIServer struct {
	dbconnection *sql.DB
}

func NewAPIServer(dbconnection *sql.DB) APIServer {
	return APIServer{
		dbconnection: dbconnection,
	}
}

func (api *APIServer) Run(port int) {

	defer api.dbconnection.Close()

	router := gin.Default()

	// Swagger
	docs.SwaggerInfo.BasePath = "/api/v1"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	apiV1 := router.Group("api/v1")

	//Section routes
	routes.SectionRoutes(apiV1.Group("/sections"), db)

	// Employee routes
	routes.EmployeeRoutes(apiV1.Group("/employees"), db)

	// Product routes
	routes.ProductRoutes(apiV1.Group("/products"), db)

	//Warehouse routes
	routes.WarehouseRoutes(api.dbconnection, apiV1.Group("/warehouses"))

	//Seller routes
	routes.SellerRoutes(apiV1.Group("/sellers"))

	//Buyer routes
	routes.BuyerRoutes(apiV1.Group("/buyers"))

	router.Run(fmt.Sprintf(":%d", port))
}
