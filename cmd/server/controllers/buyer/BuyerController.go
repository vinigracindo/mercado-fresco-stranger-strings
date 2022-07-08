package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/buyer/domain"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/pkg/httputil"
)

type requestBuyerPost struct {
	CardNumberId string `json:"card_number_id" binding:"required"`
	FirstName    string `json:"first_name"    binding:"required"`
	LastName     string `json:"last_name"     binding:"required"`
}

type requestBuyerPatch struct {
	CardNumberId string `json:"card_number_id" binding:"required"`
	LastName     string `json:"last_name"     binding:"required"`
}

type BuyerController struct {
	service domain.BuyerService
}

func NewBuyerController(buyerService domain.BuyerService) BuyerController {
	return BuyerController{service: buyerService}
}

// Buyers godoc
// @Summary      Create buyer
// @Description  create buyer
// @Tags         Buyers
// @Accept       json
// @Produce      json
// @Param Buyer body requestBuyerPost true "Create buyer"
// @Success      201  {object} domain.Buyer
// @Failure      409  {object}  httputil.HTTPError
// @Failure      422  {object}  httputil.HTTPError
// @Router /buyers [post]
func (c *BuyerController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req requestBuyerPost
		if err := ctx.ShouldBindJSON(&req); err != nil {
			httputil.NewError(ctx, http.StatusUnprocessableEntity, err)
			return
		}

		buyer, err := c.service.Create(ctx.Request.Context(), req.CardNumberId, req.FirstName, req.LastName)
		if err != nil {
			httputil.NewError(ctx, http.StatusConflict, err)
			return
		}
		httputil.NewResponse(ctx, http.StatusCreated, buyer)
	}
}

// Buyers godoc
// @Summary      List all buyers
// @Description  get buyers
// @Tags         Buyers
// @Accept       json
// @Produce      json
// @Success      200  {object} []domain.Buyer
// @Failure      404  {object}  httputil.HTTPError
// @Router /buyers [get]
func (c *BuyerController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		buyers, err := c.service.GetAll(ctx.Request.Context())

		if err != nil {
			httputil.NewError(ctx, http.StatusBadRequest, err)
			return
		}
		httputil.NewResponse(ctx, http.StatusOK, buyers)
	}
}

// Buyers godoc
// @Summary      List buyer by id
// @Description  get buyer by id
// @Tags         Buyers
// @Accept       json
// @Produce      json
// @Param id path int true "Buyer ID"
// @Success      200  {object} domain.Buyer
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Router /buyers/{id} [get]
func (c *BuyerController) GetId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			httputil.NewError(ctx, http.StatusBadRequest, err)
			return
		}
		buyer, err := c.service.GetId(ctx.Request.Context(), id)
		if err != nil {
			httputil.NewError(ctx, http.StatusNotFound, err)
			return
		}
		httputil.NewResponse(ctx, http.StatusOK, buyer)
	}
}

// Buyers godoc
// @Summary      Update UpdateCardNumberLastName
// @Description  Update UpdateCardNumberLastName field by id
// @Tags         Buyers
// @Accept       json
// @Produce      json
// @Param id path int true "Buyers ID"
// @Param Buyer body requestBuyerPatch true "Update field"
// @Success      200  {object} domain.Buyer
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Router /buyers/{id} [patch]
func (c *BuyerController) UpdateCardNumberLastName() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			httputil.NewError(ctx, http.StatusBadRequest, err)
			return
		}
		var req requestBuyerPatch
		if err := ctx.ShouldBindJSON(&req); err != nil {
			httputil.NewError(ctx, http.StatusBadRequest, err)
			return
		}

		buyer, err := c.service.Update(ctx.Request.Context(), id, req.CardNumberId, req.LastName)
		if err != nil {
			httputil.NewError(ctx, http.StatusNotFound, err)
			return
		}
		httputil.NewResponse(ctx, http.StatusOK, buyer)
	}
}

// Buyers godoc
// @Summary      DeleteBuyer buyer
// @Description  DeleteBuyer buyer by id
// @Tags         Buyers
// @Accept       json
// @Produce      json
// @Param id path int true "Buyer ID"
// @Success      204
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Router /buyers/{id} [delete]
func (c *BuyerController) DeleteBuyer() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			httputil.NewError(ctx, http.StatusBadRequest, err)
			return
		}

		err = c.service.Delete(ctx.Request.Context(), id)
		if err != nil {
			httputil.NewError(ctx, http.StatusNotFound, err)
			return
		}
		httputil.NewResponse(ctx, http.StatusNoContent, "")
	}
}

func (c *BuyerController) GetPurchaseOrdersReports() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		idParam, isPresent := ctx.GetQuery("id")
		if isPresent {
			c.GetPurchaseOrdersReportsBuyerId(ctx, idParam)
			return
		}
		c.GetAllPurchaseOrdersReports(ctx)

	}
}

func (c *BuyerController) GetPurchaseOrdersReportsBuyerId(ctx *gin.Context, idParam string) {
	productId, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	result, err := c.service.GetPurchaseOrdersReports(ctx.Request.Context(), productId)
	if err != nil {
		httputil.NewError(ctx, http.StatusNotFound, err)
		return
	}

	httputil.NewResponse(ctx, http.StatusOK, result)
}

func (c *BuyerController) GetAllPurchaseOrdersReports(ctx *gin.Context) {
	result, err := c.service.GetAllPurchaseOrdersReports(ctx.Request.Context())

	if err != nil {
		httputil.NewError(ctx, http.StatusInternalServerError, err)
		return
	}
	httputil.NewResponse(ctx, http.StatusOK, result)
}
