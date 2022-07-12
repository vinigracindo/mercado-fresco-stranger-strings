package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/seller/domain"
	httputil "github.com/vinigracindo/mercado-fresco-stranger-strings/pkg/httputil"
)

type requestSellerPost struct {
	Id          int64  `json:"id"`
	Cid         int64  `json:"cid" binding:"required"`
	CompanyName string `json:"company_name" binding:"required"`
	Address     string `json:"address" binding:"required"`
	Telephone   string `json:"telephone" binding:"required"`
	LocalityId  int64  `json:"locality_id" binding:"required"`
}

type requestSellerPatch struct {
	Address   string `json:"address" binding:"required"`
	Telephone string `json:"telephone" binding:"required"`
}

type SellerController struct {
	service domain.ServiceSeller
}

func NewSeller(s domain.ServiceSeller) *SellerController {
	return &SellerController{
		service: s,
	}
}

// Seller godoc
// @Summary      List all seller
// @Description  get Seller
// @Tags         Seller
// @Accept       json
// @Produce      json
// @Success      200  {object} []domain.Seller
// @Failure      400  {object}  httputil.HTTPError
// @Router /sellers [get]
func (c SellerController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		seller, err := c.service.GetAll(ctx.Request.Context())
		if err != nil {
			httputil.NewError(ctx, http.StatusInternalServerError, err)
			return
		}
		httputil.NewResponse(ctx, http.StatusOK, seller)
	}
}

// Seller godoc
// @Summary      List Seller by id
// @Description  get Seller by id
// @Tags         Seller
// @Accept       json
// @Produce      json
// @Param id path int true "Seller ID"
// @Success      200  {object} domain.Seller
// @Failure      500  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Router /sellers/{id} [get]
func (c SellerController) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			httputil.NewError(ctx, http.StatusBadRequest, err)
			return
		}
		seller, err := c.service.GetById(ctx.Request.Context(), id)
		if err != nil {
			httputil.NewError(ctx, http.StatusNotFound, err)
			return
		}
		httputil.NewResponse(ctx, http.StatusOK, seller)
	}
}

// Seller godoc
// @Summary      Create seller
// @Description  create seller
// @Tags         Seller
// @Accept       json
// @Produce      json
// @Param Seller body requestSellerPost true "Create seller"
// @Success      201  {object}  domain.Seller
// @Failure      409  {object}  httputil.HTTPError
// @Failure      422  {object}  httputil.HTTPError
// @Router /sellers [post]
func (c SellerController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req requestSellerPost
		if err := ctx.ShouldBindJSON(&req); err != nil {
			httputil.NewError(ctx, http.StatusUnprocessableEntity, err)
			return
		}

		seller := domain.Seller{
			Cid:         req.Cid,
			CompanyName: req.CompanyName,
			Address:     req.Address,
			Telephone:   req.Telephone,
			LocalityId:  req.LocalityId,
		}

		newSeller, err := c.service.Create(ctx.Request.Context(), &seller)
		if err != nil {
			httputil.NewError(ctx, http.StatusConflict, err)
			return
		}
		httputil.NewResponse(ctx, http.StatusCreated, newSeller)
	}
}

// Seller godoc
// @Summary      Update seller
// @Description  Update seller
// @Tags         Seller
// @Accept       json
// @Produce      json
// @Param id path int true "Seller ID"
// @Param Warehouse body requestSellerPatch true "Update seller"
// @Success      200  {object} domain.Seller
// @Failure      404  {object}  httputil.HTTPError
// @Failure      422  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router /sellers/{id} [patch]
func (c SellerController) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			httputil.NewError(ctx, http.StatusBadRequest, err)
			return
		}

		var req requestSellerPatch

		if err := ctx.ShouldBindJSON(&req); err != nil {
			httputil.NewError(ctx, http.StatusUnprocessableEntity, err)
			return
		}

		sellerUpdate, err := c.service.Update(ctx.Request.Context(), id, req.Address, req.Telephone)

		if err != nil {
			httputil.NewError(ctx, http.StatusNotFound, err)
			return
		}

		httputil.NewResponse(ctx, http.StatusOK, sellerUpdate)
	}
}

// Seller godoc
// @Summary      Delete Seller
// @Description  Delete Seller by id
// @Tags         Seller
// @Accept       json
// @Produce      json
// @Param id path int true "Seller ID"
// @Success      204
// @Failure      500  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Router /sellers/{id} [delete]
func (c SellerController) Delete() gin.HandlerFunc {
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

		httputil.NewResponse(ctx, http.StatusNoContent, err)
	}
}
