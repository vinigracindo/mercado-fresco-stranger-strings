package routes

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/vinigracindo/mercado-fresco-stranger-strings/cmd/server/controllers/section"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/section/repository"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/section/service"
)

func SectionRoutes(routes *gin.RouterGroup) {
	sectionRepository := repository.NewRepositorySection()
	sectionService := service.NewServiceSection(sectionRepository)
	sectionController := controllers.NewSection(sectionService)

	routes.DELETE("/:id", sectionController.Delete())
	routes.PATCH("/:id", sectionController.UpdateCurrentCapacity())
	routes.POST("/", sectionController.Create())
	routes.GET("/:id", sectionController.GetById())
	routes.GET("/", sectionController.GetAll())
}
