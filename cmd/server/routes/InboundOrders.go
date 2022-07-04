package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	controllers "github.com/vinigracindo/mercado-fresco-stranger-strings/cmd/server/controllers/inbound_orders"
	EmployeeRepository "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/employees/repository"
	InboundOrdersRepository "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/inboud_orders/repository"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/inboud_orders/service"
)

func InboundOrdersRoutes(routes *gin.RouterGroup, db *sql.DB) {
	employeeRepository := EmployeeRepository.NewMariaDBEmployeeRepository(db)
	inboundOrdersRepository := InboundOrdersRepository.NewMariaDBInboundRepositoryRepository(db)

	inboundOrdersService := service.NewInboundOrderService(inboundOrdersRepository, employeeRepository)
	inboundOrdersController := controllers.NewInboundOrdersController(inboundOrdersService)

	routes.POST("/", inboundOrdersController.Create())
}
