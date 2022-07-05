package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/employees/domain"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/pkg/httputil"
)

type EmployeeController struct {
	service domain.EmployeeService
}

func NewEmployeeController(service domain.EmployeeService) EmployeeController {
	return EmployeeController{service: service}
}

// GetAll godoc
// @Summary      List all employees
// @Description  get all employees
// @Tags         Employees
// @Accept       json
// @Produce      json
// @Success      200  {array} domain.Employee
// @Failure      404  {object}  httputil.HTTPError
// @Router /employees [get]
func (controller EmployeeController) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		employees, err := controller.service.GetAll(c.Request.Context())
		if err != nil {
			httputil.NewError(c, http.StatusInternalServerError, err)
			return
		}
		httputil.NewResponse(c, http.StatusOK, employees)
	}
}

// GetById godoc
// @Summary Get employee by ID
// @Tags Employees
// @Description Get employee by ID
// @Produce json
// @Param id path int true "Employee ID"
// @Success 200 {object} domain.Employee
// @Failure 400  {object}  httputil.HTTPError
// @Failure 404  {object}  httputil.HTTPError
// @Router /employees/{id} [get]
func (controller EmployeeController) GetById() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			httputil.NewError(c, http.StatusBadRequest, err)
			return
		}
		employee, err := controller.service.GetById(c.Request.Context(), id)
		if err != nil {
			httputil.NewError(c, http.StatusNotFound, err)
			return
		}
		httputil.NewResponse(c, http.StatusOK, employee)
	}
}

// Store godoc
// @Summary      Create employee
// @Description  create employee
// @Tags         Employees
// @Accept       json
// @Produce      json
// @Param Employee body requestEmployeePost true "Create employee"
// @Success      201  {object} domain.Employee
// @Failure      409  {object}  httputil.HTTPError
// @Failure      422  {object}  httputil.HTTPError
// @Router /employees [post]
func (controller EmployeeController) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req requestEmployeePost
		if err := c.ShouldBindJSON(&req); err != nil {
			httputil.NewError(c, http.StatusUnprocessableEntity, err)
			return
		}

		employee, err := controller.service.Create(c.Request.Context(), req.CardNumberId, req.FirstName, req.LastName, req.WarehouseId)
		if err != nil {
			httputil.NewError(c, http.StatusConflict, err)
			return
		}
		httputil.NewResponse(c, http.StatusCreated, employee)
	}
}

// UpdateFullname godoc
// @Summary      Update employee fullname
// @Description  Update employee first and last name field by id
// @Tags         Employees
// @Accept       json
// @Produce      json
// @Param id path int true "Employee ID"
// @Param Employee body requestEmployeePatch true "Update field"
// @Success      200  {object} domain.Employee
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Router /employees/{id} [patch]
func (controller EmployeeController) UpdateFullname() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			httputil.NewError(c, http.StatusBadRequest, err)
			return
		}

		var req requestEmployeePatch
		if err := c.ShouldBindJSON(&req); err != nil {
			httputil.NewError(c, http.StatusBadRequest, err)
			return
		}
		employee, err := controller.service.UpdateFullname(c.Request.Context(), id, req.FirstName, req.LastName)
		if err != nil {
			httputil.NewError(c, http.StatusNotFound, err)
			return
		}
		httputil.NewResponse(c, http.StatusOK, employee)
	}
}

// Delete godoc
// @Summary      Delete employee
// @Description  Delete employee by id
// @Tags         Employees
// @Accept       json
// @Produce      json
// @Param id path int true "Employee ID"
// @Success      204
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Router /employees/{id} [delete]
func (controller EmployeeController) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			httputil.NewError(c, http.StatusBadRequest, err)
			return
		}
		err = controller.service.Delete(c.Request.Context(), id)
		if err != nil {
			httputil.NewError(c, http.StatusNotFound, err)
			return
		}
		httputil.NewResponse(c, http.StatusNoContent, "Employee deleted")
	}
}

func (controller EmployeeController) ReportInboundOrders() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Query("id")
		id := new(int64)

		if idParam == "" {
			id = nil
		} else {
			idConv, err := strconv.ParseInt(idParam, 10, 64)
			if err != nil {
				httputil.NewError(c, http.StatusBadRequest, err)
				return
			}
			id = &idConv
		}

		result, err := controller.service.ReportInboundOrders(c.Request.Context(), id)
		if err != nil {
			httputil.NewError(c, http.StatusInternalServerError, err)
			return
		}
		httputil.NewResponse(c, http.StatusOK, result)
	}
}

type requestEmployeePost struct {
	CardNumberId string `json:"card_number_id" binding:"required"`
	FirstName    string `json:"first_name" binding:"required"`
	LastName     string `json:"last_name" binding:"required"`
	WarehouseId  int64  `json:"warehouse_id" binding:"required"`
}

type requestEmployeePatch struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
}
