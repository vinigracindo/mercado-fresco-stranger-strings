package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/product"
	"net/http"
	"strconv"
)

type ProductController struct {
	service product.Service
}

func NewProduct(p product.Service) *ProductController {
	return &ProductController{
		service: p,
	}
}

func (c ProductController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		product, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, product)
	}
}

func (c *ProductController) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		product, err := c.service.GetById(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, *product)
	}
}
