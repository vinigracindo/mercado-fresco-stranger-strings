package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	controllers "github.com/vinigracindo/mercado-fresco-stranger-strings/cmd/server/controllers/seller"
	repository "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/seller/repository/mariadb"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/seller/services"
)

func SellerRoutes(routes *gin.RouterGroup, db *sql.DB) {
	sellerRepository := repository.NewMariaDBSellerRepository(db)
	sellerService := services.NewSellerService(sellerRepository)
	sellerController := controllers.NewSeller(sellerService)

	routes.GET("/", sellerController.GetAll())
	routes.GET("/:id", sellerController.GetById())
	routes.POST("/", sellerController.Create())
	routes.PATCH("/:id", sellerController.Update())
	routes.DELETE("/:id", sellerController.Delete())
}
