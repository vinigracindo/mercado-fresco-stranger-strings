package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/employees"
)

type EmployeeController struct {
	service employees.Service
}

func NewEmployee(service employees.Service) EmployeeController {
	return EmployeeController{service: service}
}

func (controller EmployeeController) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		employees, err := controller.service.GetAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": employees,
		})
	}
}

func (controller EmployeeController) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		employee, err := controller.service.Get(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": employee,
		})
	}
}
