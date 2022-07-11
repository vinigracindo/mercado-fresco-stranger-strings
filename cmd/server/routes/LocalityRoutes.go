package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	controllers "github.com/vinigracindo/mercado-fresco-stranger-strings/cmd/server/controllers/locality"
	repositoryLocality "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/locality/repository/mariadb"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/locality/services"
)

func LocalityRoutes(routes *gin.RouterGroup, db *sql.DB) {
	localityRepository := repositoryLocality.NewMariadbLocalityRepository(db)
	localityService := services.NewLocalityService(localityRepository)
	localityController := controllers.NewLocalityController(localityService)

	routes.POST("/", localityController.CreateLocality())
}
