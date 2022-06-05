package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/cmd/server/controllers"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/employees"
)

func main() {
	// Employee
	employeeRepository := employees.NewRepository()
	employeeService := employees.NewService(employeeRepository)
	employeeController := controllers.NewEmployee(employeeService)

	router := gin.Default()

	apiv1 := router.Group("/api/v1")
	employeeRouter := apiv1.Group("/employees")
	employeeRouter.GET("/", employeeController.GetAll())
	employeeRouter.GET("/:id", employeeController.Get())
	employeeRouter.POST("/", employeeController.Store())
	employeeRouter.PATCH("/:id", employeeController.Update())

	router.Run()
}
