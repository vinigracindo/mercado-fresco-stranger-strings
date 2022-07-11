package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/product_records/domain"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/pkg/httputil"
	"net/http"
	"time"
)

type RequestProductRecordsPost struct {
	LastUpdateDate string  `json:"last_update_date" binding:"required"`
	PurchasePrice  float64 `json:"purchase_price" binding:"required"`
	SalePrice      float64 `json:"sale_price" binding:"required"`
	ProductId      int64   `json:"product_id" binding:"required"`
}

type ProductRecordsController struct {
	service domain.ProductRecordsService
}

func CreateProductRecordsController(prodService domain.ProductRecordsService) *ProductRecordsController {
	return &(ProductRecordsController{
		service: prodService})
}

// Create godoc
// @Summary      Create ProductRecords
// @Description  create product records
// @Tags         ProductRecords
// @Accept       json
// @Produce      json
// @Param ProductRecords body productRecordsDTO true "Create product records"
// @Success      201  {object}  domain.ProductRecords
// @Failure      409  {object}  httputil.HTTPError
// @Failure      422  {object}  httputil.HTTPError
// @Router /productRecords [post]
func (c *ProductRecordsController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var productRecordsDTO RequestProductRecordsPost

		if err := ctx.ShouldBindJSON(&productRecordsDTO); err != nil {
			httputil.NewError(ctx, http.StatusUnprocessableEntity, errors.New("invalid input. Check the data entered"))
			return
		}

		date, err := time.Parse("2006-01-02", productRecordsDTO.LastUpdateDate)

		if err != nil {
			httputil.NewError(ctx, http.StatusBadRequest, err)
			return
		}

		model := domain.ProductRecords{
			LastUpdateDate: date,
			PurchasePrice:  productRecordsDTO.PurchasePrice,
			SalePrice:      productRecordsDTO.SalePrice,
			ProductId:      productRecordsDTO.ProductId,
		}

		newProductRecords, err := c.service.Create(ctx.Request.Context(), &model)

		if err != nil {
			httputil.NewError(ctx, http.StatusConflict, err)
			return
		}

		httputil.NewResponse(ctx, http.StatusCreated, newProductRecords)
	}
}
