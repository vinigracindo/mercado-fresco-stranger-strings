package routes

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/vinigracindo/mercado-fresco-stranger-strings/cmd/server/controllers/employees"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/employees/repository"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/employees/service"
)

func EmployeeRoutes(routes *gin.RouterGroup) {
	employeeRepository := repository.NewEmployeeRepository()
	employeeService := service.NewEmployeeService(employeeRepository)
	employeeController := controllers.NewEmployeeController(employeeService)

	routes.GET("/", employeeController.GetAll())
	routes.GET("/:id", employeeController.GetById())
	routes.POST("/", employeeController.Create())
	routes.PATCH("/:id", employeeController.UpdateFullname())
	routes.DELETE("/:id", employeeController.Delete())
}
