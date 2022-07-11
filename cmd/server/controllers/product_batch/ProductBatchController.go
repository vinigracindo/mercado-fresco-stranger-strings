package controllers

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/product_batch/domain"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/pkg/httputil"
)

type RequestProductBatchPost struct {
	BatchNumber        int64   `json:"batch_number" binding:"required"`
	CurrentQuantity    int64   `json:"current_quantity" binding:"required"`
	CurrentTemperature float64 `json:"current_temperature" binding:"required"`
	DueDate            string  `json:"due_date" binding:"required"`
	InitialQuantity    int64   `json:"initial_quantity" binding:"required"`
	ManufacturingDate  string  `json:"manufacturing_date" binding:"required"`
	ManufacturingHour  int64   `json:"manufacturing_hour" binding:"required"`
	MinumumTemperature float64 `json:"minumum_temperature" binding:"required"`
	ProductId          int64   `json:"product_id" binding:"required"`
	SectionId          int64   `json:"section_id" binding:"required"`
}

type ProductBatchController struct {
	service domain.ProductBatchService
}

func NewProductBatchController(service domain.ProductBatchService) *ProductBatchController {
	return &(ProductBatchController{
		service: service})
}

// Create godoc
// @Summary      Create product batch
// @Description  create product batch
// @Tags         Product batches
// @Accept       json
// @Produce      json
// @Param ProductBatch body RequestProductBatchPost true "Create product batch"
// @Success      201  {object} domain.ProductBatch
// @Failure      400  {object}  httputil.HTTPError
// @Failure      422  {object}  httputil.HTTPError
// @Router /productBatches [post]
func (c *ProductBatchController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var request RequestProductBatchPost

		if err := ctx.ShouldBindJSON(&request); err != nil {
			httputil.NewError(ctx, http.StatusUnprocessableEntity, errors.New("invalid input. Check the data entered"))
			return
		}

		dueDate, err := time.Parse("2006-01-02", request.DueDate)
		if err != nil {
			httputil.NewError(ctx, http.StatusBadRequest, err)
			return
		}
		manufacturingDate, err := time.Parse("2006-01-02", request.ManufacturingDate)
		if err != nil {
			httputil.NewError(ctx, http.StatusBadRequest, err)
			return
		}

		model := domain.ProductBatch{
			BatchNumber:        request.BatchNumber,
			CurrentQuantity:    request.CurrentQuantity,
			CurrentTemperature: request.CurrentTemperature,
			DueDate:            dueDate,
			InitialQuantity:    request.InitialQuantity,
			ManufacturingDate:  manufacturingDate,
			ManufacturingHour:  request.ManufacturingHour,
			MinumumTemperature: request.MinumumTemperature,
			ProductId:          request.ProductId,
			SectionId:          request.SectionId,
		}

		newProductBatch, err := c.service.Create(ctx.Request.Context(), &model)

		if err != nil {
			httputil.NewError(ctx, http.StatusConflict, err)
			return
		}

		httputil.NewResponse(ctx, http.StatusCreated, newProductBatch)
	}
}
