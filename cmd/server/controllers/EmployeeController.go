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

func (controller EmployeeController) GetById() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		employee, err := controller.service.GetById(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": employee,
		})
	}
}

func (controller EmployeeController) Store() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req requestEmployeePost
		if err := c.ShouldBindJSON(&req); err != nil {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
				"error":   err.Error(),
				"message": "Invalid request",
			})
			return
		}

		employee, err := controller.service.Store(req.CardNumberId, req.FirstName, req.LastName, req.WarehouseId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusCreated, gin.H{
			"data": employee,
		})
	}
}

func (controller EmployeeController) UpdateFullname() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		var req requestEmployeePatch
		if err := c.ShouldBindJSON(&req); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   err.Error(),
				"message": "Invalid request",
			})
			return
		}
		employee, err := controller.service.UpdateFullname(id, req.FirstName, req.LastName)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": employee,
		})
	}
}

func (controller EmployeeController) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		err = controller.service.Delete(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusNoContent, gin.H{
			"data": "Employee deleted",
		})
	}
}

type requestEmployeePost struct {
	Id           int64  `json:"id,omitempty"`
	CardNumberId string `json:"card_number_id" binding:"required"`
	FirstName    string `json:"first_name" binding:"required"`
	LastName     string `json:"last_name" binding:"required"`
	WarehouseId  int64  `json:"warehouse_id" binding:"required"`
}

type requestEmployeePatch struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
}