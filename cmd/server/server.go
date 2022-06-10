package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/cmd/server/controllers"
	docs "github.com/vinigracindo/mercado-fresco-stranger-strings/docs"

	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/buyer"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/employees"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/product"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/section"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/seller"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/warehouse"
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

	groupV1 := router.Group("api/v1")

	// Section routes
	sectionRepository := section.NewRepository()
	sectionService := section.NewService(sectionRepository)
	sectionController := controllers.NewSection(sectionService)

	sectionGroup := groupV1.Group("/sections")
	sectionGroup.DELETE("/:id", sectionController.Delete())
	sectionGroup.PATCH("/:id", sectionController.UpdateCurrentCapacity())
	sectionGroup.POST("/", sectionController.CreateSection())
	sectionGroup.GET("/:id", sectionController.GetById())
	sectionGroup.GET("/", sectionController.GetAll())

	// Employee routes
	employeeRepository := employees.NewRepository()
	employeeService := employees.NewService(employeeRepository)
	employeeController := controllers.NewEmployee(employeeService)

	employeeGroup := groupV1.Group("/employees")
	employeeGroup.GET("/", employeeController.GetAll())
	employeeGroup.GET("/:id", employeeController.GetById())
	employeeGroup.POST("/", employeeController.Store())
	employeeGroup.PATCH("/:id", employeeController.UpdateFullname())
	employeeGroup.DELETE("/:id", employeeController.Delete())

	// Product routes
	productRepository := product.CreateRepository()
	productService := product.CreateService(productRepository)
	productController := controllers.CreateProductController(productService)

	productGroup := groupV1.Group("/products")
	productGroup.GET("/", productController.GetAll())
	productGroup.GET("/:id", productController.GetById())
	productGroup.POST("/", productController.Create())
	productGroup.PATCH("/:id", productController.UpdateDescription())
	productGroup.DELETE("/:id", productController.Delete())

	//Warehouse routes
	warehouseRepository := warehouse.NewRepository()
	warehouseService := warehouse.NewService(warehouseRepository)
	warehouseController := controllers.NewWarehouse(warehouseService)

	warehouseGroup := groupV1.Group("/warehouses")
	warehouseGroup.GET("/", warehouseController.GetAllWarehouse())
	warehouseGroup.GET("/:id", warehouseController.GetWarehouseByID())
	warehouseGroup.POST("/", warehouseController.CreateWarehouse())
	warehouseGroup.DELETE("/:id", warehouseController.DeleteWarehouse())
	warehouseGroup.PATCH("/:id", warehouseController.UpdateWarehouse())

	//Seller routes
	sellerRepository := seller.NewRepository()
	sellerService := seller.NewService(sellerRepository)
	sellerController := controllers.NewSeller(sellerService)

	sellerGroup := groupV1.Group("/seller")
	sellerGroup.GET("/", sellerController.GetAll())
	sellerGroup.GET("/:id", sellerController.GetById())
	sellerGroup.POST("/", sellerController.CreateSeller())
	sellerGroup.PATCH("/:id", sellerController.UpdateSellerAddresAndTel())
	sellerGroup.DELETE("/:id", sellerController.DeleteSeller())

	//Buyer routes
	buyerRepository := buyer.NewRepository()
	buyerService := buyer.NewService(buyerRepository)
	buyerController := controllers.NewBuyer(buyerService)

	buyerGroup := groupV1.Group("buyers")
	buyerGroup.GET("/", buyerController.GetAll())
	buyerGroup.GET("/:id", buyerController.GetId())
	buyerGroup.POST("/", buyerController.Store())
	buyerGroup.PATCH("/:id", buyerController.CardNumberLastName())
	buyerGroup.DELETE("/:id", buyerController.Delete())

	router.Run(fmt.Sprintf(":%d", port))
}
