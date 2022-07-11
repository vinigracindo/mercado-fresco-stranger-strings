package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	controllers "github.com/vinigracindo/mercado-fresco-stranger-strings/cmd/server/controllers/section"
	repository "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/section/repository/mariadb"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/section/service"
)

func SectionRoutes(routes *gin.RouterGroup, db *sql.DB) {
	sectionRepository := repository.NewMariadbSectionRepository(db)
	sectionService := service.NewServiceSection(sectionRepository)
	sectionController := controllers.NewSection(sectionService)

	//report product by section route
	routes.GET("/reportProducts", sectionController.GetReportProductsBySection())

	routes.DELETE("/:id", sectionController.Delete())
	routes.PATCH("/:id", sectionController.UpdateCurrentCapacity())
	routes.POST("/", sectionController.Create())
	routes.GET("/:id", sectionController.GetById())
	routes.GET("/", sectionController.GetAll())
}
