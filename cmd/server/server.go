package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/cmd/server/controllers"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/cmd/server/routes"
	docs "github.com/vinigracindo/mercado-fresco-stranger-strings/docs"

	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/buyer"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/employees"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/product"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/seller"
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
	employeeRepository := employees.NewRepository()
	employeeService := employees.NewService(employeeRepository)
	employeeController := controllers.NewEmployee(employeeService)

	employeeGroup := apiV1.Group("/employees")
	employeeGroup.GET("/", employeeController.GetAll())
	employeeGroup.GET("/:id", employeeController.GetById())
	employeeGroup.POST("/", employeeController.Create())
	employeeGroup.PATCH("/:id", employeeController.UpdateFullname())
	employeeGroup.DELETE("/:id", employeeController.Delete())

	// Product routes
	productRepository := product.CreateRepository()
	productService := product.CreateService(productRepository)
	productController := controllers.CreateProductController(productService)

	productGroup := apiV1.Group("/products")
	productGroup.GET("/", productController.GetAll())
	productGroup.GET("/:id", productController.GetById())
	productGroup.POST("/", productController.Create())
	productGroup.PATCH("/:id", productController.UpdateDescription())
	productGroup.DELETE("/:id", productController.Delete())

	//Warehouse routes
	routes.WarehouseRoutes(apiV1.Group("/warehouses"))

	//Seller routes
	sellerRepository := seller.NewRepository()
	sellerService := seller.NewService(sellerRepository)
	sellerController := controllers.NewSeller(sellerService)

	sellerGroup := apiV1.Group("/sellers")
	sellerGroup.GET("/", sellerController.GetAll())
	sellerGroup.GET("/:id", sellerController.GetById())
	sellerGroup.POST("/", sellerController.Create())
	sellerGroup.PATCH("/:id", sellerController.Update())
	sellerGroup.DELETE("/:id", sellerController.Delete())

	//Buyer routes
	buyerRepository := buyer.NewRepository()
	buyerService := buyer.NewService(buyerRepository)
	buyerController := controllers.NewBuyer(buyerService)

	buyerGroup := apiV1.Group("buyers")
	buyerGroup.GET("/", buyerController.GetAll())
	buyerGroup.GET("/:id", buyerController.GetId())
	buyerGroup.POST("/", buyerController.Create())
	buyerGroup.PATCH("/:id", buyerController.UpdateCardNumberLastName())
	buyerGroup.DELETE("/:id", buyerController.DeleteBuyer())

	router.Run(fmt.Sprintf(":%d", port))
}
