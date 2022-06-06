package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/seller"
)

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
