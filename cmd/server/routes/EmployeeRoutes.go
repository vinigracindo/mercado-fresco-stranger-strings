package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	controllers "github.com/vinigracindo/mercado-fresco-stranger-strings/cmd/server/controllers/employees"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/employees/repository"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/employees/service"
)

func EmployeeRoutes(routes *gin.RouterGroup, db *sql.DB) {
	employeeRepository := repository.NewMariaDBEmployeeRepository(db)
	employeeService := service.NewEmployeeService(employeeRepository)
	employeeController := controllers.NewEmployeeController(employeeService)

	// Inbound Orders Report
	routes.GET("/reportInboundOrders", employeeController.ReportInboundOrders())

	// Employee routes
	routes.GET("/", employeeController.GetAll())
	routes.GET("/:id", employeeController.GetById())
	routes.POST("/", employeeController.Create())
	routes.PATCH("/:id", employeeController.UpdateFullname())
	routes.DELETE("/:id", employeeController.Delete())
}
