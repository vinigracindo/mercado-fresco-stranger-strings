package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/seller/domain"
)

type requestSellerPost struct {
	Id          int64  `json:"id"`
	Cid         int64  `json:"cid" binding:"required"`
	CompanyName string `json:"company_name" binding:"required"`
	Address     string `json:"address" binding:"required"`
	Telephone   string `json:"telephone" binding:"required"`
}

type requestSellerPatch struct {
	Address   string `json:"address" binding:"required"`
	Telephone string `json:"telephone" binding:"required"`
}

type SellerController struct {
	service domain.ServiceSeller
}

func NewSeller(s domain.ServiceSeller) SellerController {
	return SellerController{
		service: s,
	}
}

// Seller godoc
// @Summary      List all seller
// @Description  get Seller
// @Tags         Seller
// @Accept       json
// @Produce      json
// @Success      200  {object} []seller.Seller
// @Failure      400  {object}  httputil.HTTPError
// @Router /sellers [get]
func (c SellerController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		seller, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": seller,
		})
	}
}

// Seller godoc
// @Summary      List Seller by id
// @Description  get Seller by id
// @Tags         Seller
// @Accept       json
// @Produce      json
// @Param id path int true "Seller ID"
// @Success      200  {object} seller.Seller
// @Failure      500  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Router /sellers/{id} [get]
func (c SellerController) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		seller, err := c.service.GetById(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"data": seller,
		})
	}
}

// Seller godoc
// @Summary      Create seller
// @Description  create seller
// @Tags         Seller
// @Accept       json
// @Produce      json
// @Param Seller body requestSellerPost true "Create seller"
// @Success      201  {object}  seller.Seller
// @Failure      409  {object}  httputil.HTTPError
// @Failure      422  {object}  httputil.HTTPError
// @Router /sellers [post]
func (c SellerController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req requestSellerPost
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
				"error":   err.Error(),
				"message": "invalid request",
			})
			return
		}
		seller, err := c.service.Create(req.Cid, req.CompanyName, req.Address, req.Telephone)
		if err != nil {
			ctx.JSON(http.StatusConflict, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusCreated, gin.H{
			"data": seller,
		})
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
// @Success      200  {object} seller.Seller
// @Failure      404  {object}  httputil.HTTPError
// @Failure      422  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router /sellers/{id} [patch]
func (c SellerController) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		var req requestSellerPatch
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
				"error":   err.Error(),
				"message": "invalid request",
			})
			return
		}
		seller, err := c.service.Update(id, req.Address, req.Telephone)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"data": seller,
		})
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
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		err = c.service.Delete(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
		}
		ctx.JSON(http.StatusNoContent, gin.H{})
	}
}
