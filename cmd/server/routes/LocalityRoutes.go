package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	controllers "github.com/vinigracindo/mercado-fresco-stranger-strings/cmd/server/controllers/locality"
	repositoryCarry "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/carry/repository/mariadb"
	repositoryLocality "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/locality/repository/mariadb"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/locality/services"
)

func LocalityRoutes(routes *gin.RouterGroup, db *sql.DB) {
	localityRepository := repositoryLocality.NewMariadbLocalityRepository(db)
	carryRepository := repositoryCarry.NewMariadbCarryRepository(db)
	localityService := services.NewLocalityService(localityRepository, carryRepository)
	localityController := controllers.NewLocalityController(localityService)

	routes.GET("/reportCarrie/:id", localityController.ReportCarrie())

}
