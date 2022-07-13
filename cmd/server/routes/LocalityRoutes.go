package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	controllers "github.com/vinigracindo/mercado-fresco-stranger-strings/cmd/server/controllers/locality"
	repositoryLocality "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/locality/repository/mariadb"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/locality/services"
	repository "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/seller/repository/mariadb"
)

func LocalityRoutes(routes *gin.RouterGroup, db *sql.DB) {
	localityRepository := repositoryLocality.NewMariadbLocalityRepository(db)
	sellerRepository := repository.NewMariaDBSellerRepository(db)
	localityService := services.NewLocalityService(localityRepository, sellerRepository)
	localityController := controllers.NewLocalityController(localityService)

	routes.POST("/", localityController.CreateLocality())
	routes.GET("/reportCarries", localityController.ReportCarrie())
	routes.GET("/reportSellers", localityController.GetReportLocalities())
}
