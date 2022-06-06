package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/seller"
)

type requestSellerPost struct {
	Id          int64  `json:"id, omitempty"`
	Cid         int64  `json:"cid" binding:"required"`
	CompanyName string `json:"company_name" binding:"required"`
	Address     string `json:"address" binding:"required"`
	Telephone   string `json:"telephone" binding:"required"`
}

type requestSellerPatch struct {
	Id        int64  `json:"id"`
	Address   string `json:"address" binding:"required"`
	Telephone string `json:"telephone" binding:"required"`
}

type SellerController struct {
	service seller.Service
}

func NewSeller(s seller.Service) SellerController {
	return SellerController{
		service: s,
	}
}

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

func (c SellerController) Get() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		seller, err := c.service.Get(id)
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

func (c SellerController) CreateSeller() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req requestSellerPost
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
				"error":   err.Error(),
				"message": "invalid request",
			})
			return
		}
		seller, err := c.service.CreateSeller(req.Cid, req.CompanyName, req.Address, req.Telephone)
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

func (c SellerController) UpdateSellerAddresAndTel() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req requestSellerPatch
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
				"error":   err.Error(),
				"message": "invalid request",
			})
			return
		}
		seller, err := c.service.UpdateSellerAddresAndTel(req.Id, req.Address, req.Telephone)
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

func (c SellerController) DeleteSeller() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		err = c.service.DeleteSeller(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{})
		}
		ctx.JSON(http.StatusNoContent, gin.H{
			"data": "Seller deleted",
		})
	}
}
