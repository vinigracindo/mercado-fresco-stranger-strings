package controllers

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	EmployeesDomain "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/employees/domain"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/inboud_orders/domain"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/pkg/httputil"
)

type InboundOrdersController struct {
	service domain.InboundOrdersService
}

func NewInboundOrdersController(service domain.InboundOrdersService) InboundOrdersController {
	return InboundOrdersController{service: service}
}

type requestInboundOrdersPost struct {
	OrderDate      string `json:"order_date" binding:"required"`
	OrderNumber    string `json:"order_number" binding:"required"`
	EmployeeId     int64  `json:"employee_id" binding:"required"`
	ProductBatchId int64  `json:"product_batch_id" binding:"required"`
	WarehouseId    int64  `json:"warehouse_id" binding:"required"`
}

// Create godoc
// @Summary      Create InboundOrder
// @Description  create inbound order
// @Tags         InboundOrders
// @Accept       json
// @Produce      json
// @Param Employee body requestEmployeePost true "Create inbound orders"
// @Success      201  {object}  domain.InboundOrder
// @Failure      409  {object}  httputil.HTTPError
// @Failure      422  {object}  httputil.HTTPError
// @Router /inboundOrders [post]
func (controller InboundOrdersController) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request requestInboundOrdersPost

		if err := c.ShouldBindJSON(&request); err != nil {
			httputil.NewError(c, http.StatusBadRequest, err)
			return
		}

		date, err := time.Parse("2006-01-02", request.OrderDate)
		if err != nil {
			httputil.NewError(c, http.StatusBadRequest, err)
			return
		}

		inboundOrders, err := controller.service.Create(
			c.Request.Context(),
			date,
			request.OrderNumber,
			request.EmployeeId,
			request.ProductBatchId,
			request.WarehouseId,
		)

		if errors.Is(err, EmployeesDomain.ErrEmployeeNotFound) {
			httputil.NewError(c, http.StatusConflict, EmployeesDomain.ErrEmployeeNotFound)
			return
		}

		if err != nil {
			httputil.NewError(c, http.StatusInternalServerError, err)
			return
		}
		httputil.NewResponse(c, http.StatusCreated, inboundOrders)
	}
}
