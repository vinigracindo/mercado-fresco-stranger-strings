package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/purchase_orders/domain"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/pkg/httputil"
)

type PurchaseOrdersCreate struct {
	OrderNumber     string `json:"order_number" binding:"required"`
	OrderDate       string `json:"order_date" binding:"required"`
	TrackingCode    string `json:"tracking_code" binding:"required"`
	BuyerId         int64  `json:"buyer_id" binding:"required"`
	ProductRecordId int64  `json:"product_record_id" binding:"required"`
	OrderStatusId   int64  `json:"order_status_id" binding:"required"`
}
type PurchaseOrdersController struct {
	service domain.PurchaseOrdersService
}

func NewPurchaseOrdersController(service domain.PurchaseOrdersService) PurchaseOrdersController {
	return PurchaseOrdersController{service: service}
}

// PurchaseOrders godoc
// @Summary      Create purchaseOrders
// @Description  create purchaseOrders
// @Tags         PurchaseOrders
// @Accept       json
// @Produce      json
// @Param purchaseOrders body requestpurchaseOrdersPost true "Create purchaseOrders"
// @Success      201  {object} domain.purchaseOrders
// @Failure      409  {object}  httputil.HTTPError
// @Failure      422  {object}  httputil.HTTPError
// @Router /purchaseOrders [post]
func (c *PurchaseOrdersController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req PurchaseOrdersCreate
		if err := ctx.ShouldBindJSON(&req); err != nil {
			httputil.NewError(ctx, http.StatusUnprocessableEntity, err)
			return
		}

		date, err := time.Parse("2006-01-02", req.OrderDate)
		if err != nil {
			httputil.NewError(ctx, http.StatusBadRequest, err)
			return
		}

		buyer, err := c.service.Create(
			ctx.Request.Context(),
			req.OrderNumber,
			date,
			req.TrackingCode,
			req.BuyerId,
			req.ProductRecordId,
			req.OrderStatusId,
		)
		if err != nil {
			httputil.NewError(ctx, http.StatusConflict, err)
			return
		}
		httputil.NewResponse(ctx, http.StatusCreated, buyer)
	}
}
