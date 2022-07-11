package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	controllers "github.com/vinigracindo/mercado-fresco-stranger-strings/cmd/server/controllers/carry"
	repository "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/carry/repository/mariadb"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/carry/services"
)

func CarryRoutes(routes *gin.RouterGroup, db *sql.DB) {
	repo := repository.NewMariadbCarryRepository(db)
	service := services.NewCarryService(repo)
	controller := controllers.NewCarryController(service)

	routes.POST("/", controller.CreateCarry())
}
