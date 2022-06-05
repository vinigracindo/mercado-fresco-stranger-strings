package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/buyer"
)

type BuyerController struct {
	service buyer.Service
}

func NewBuyer(service buyer.Service) BuyerController {
	return BuyerController{service: service}
}

func (c *BuyerController) Store() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{
					"message": "valor de entrada inválido. por favor revise os dados"})
			return
		}

		buyer, err := c.service.Store(req.Id, req.CardNumberId, req.FirstName, req.LastName)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(
			http.StatusOK,
			gin.H{
				"data": buyer,
			})
	}
}

type request struct {
	Id           int64  `json:"id"`
	CardNumberId int64  `json:"cardNumberId" binding:"required"`
	FirstName    string `json:"firstName"    binding:"required"`
	LastName     string `json:"lastName"     binding:"required"`
}

func (c *BuyerController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		buyers, err := c.service.GetAll()

		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		ctx.JSON(http.StatusOK, &buyers)
	}
}

func (c *BuyerController) GetId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		buyer, err := c.service.GetId(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"data": buyer,
		})
	}
}

func (c *BuyerController) Update() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{
					"message": "valor de entrada inválido. por favor revise os dados"})
			return
		}

		buyer, err := c.service.Update(req.Id, req.CardNumberId, req.FirstName, req.LastName)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(
			http.StatusOK,
			gin.H{
				"data": buyer,
			})
	}
}

func (c *BuyerController) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
			return
		}

		err = c.service.Delete(int64(id))
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusNoContent, gin.H{})
	}
}
