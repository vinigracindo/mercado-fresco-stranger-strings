package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/buyer"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/pkg/httputil"
)

type requestBuyerPost struct {
	CardNumberId int64  `json:"card_number_id" binding:"required"`
	FirstName    string `json:"first_name"    binding:"required"`
	LastName     string `json:"last_name"     binding:"required"`
}

type requestBuyerPatch struct {
	CardNumberId int64  `json:"card_number_id" binding:"required"`
	LastName     string `json:"last_name"     binding:"required"`
}

type BuyerController struct {
	service buyer.Service
}

func NewBuyer(service buyer.Service) BuyerController {
	return BuyerController{service: service}
}

// Buyers godoc
// @Summary      Create buyer
// @Description  create buyer
// @Tags         Buyers
// @Accept       json
// @Produce      json
// @Param Buyer body requestBuyerPost true "Create buyer"
// @Success      201  {object} buyer.Buyer
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

		buyer, err := c.service.Create(req.CardNumberId, req.FirstName, req.LastName)
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
// @Success      200  {object} []buyer.Buyer
// @Failure      404  {object}  httputil.HTTPError
// @Router /buyers [get]
func (c *BuyerController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		buyers, err := c.service.GetAll()

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
// @Success      200  {object} buyer.Buyer
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
		buyer, err := c.service.GetId(id)
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
// @Success      200  {object} buyer.Buyer
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

		buyer, err := c.service.Update(id, req.CardNumberId, req.LastName)
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

		err = c.service.Delete(int64(id))
		if err != nil {
			httputil.NewError(ctx, http.StatusNotFound, err)
			return
		}
		httputil.NewResponse(ctx, http.StatusNoContent, err)
	}
}
