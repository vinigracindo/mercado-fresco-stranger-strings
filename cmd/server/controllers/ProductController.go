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

func CreateProductController(prodService product.Service) *ProductController {
	return &(ProductController{service: prodService})
}

func (c *ProductController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		products, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, products)
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

func (c *ProductController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity,
				gin.H{
					"message": "Input inválido. Verifique os dados digitados",
				})
			return
		}
		newProduct, err := c.service.Create(req.ProductCode, req.Description, req.Width, req.Height, req.Length, req.NetWeight,
			req.ExpirationRate, req.RecommendedFreezingTemperature, req.FreezingRate, req.ProductTypeId, req.SellerId)

		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusCreated, newProduct)
	}
}

func (c *ProductController) UpdateDescription() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Id inválido"})
			return
		}
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if req.Description == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "o campo descrição é obrigatório"})
			return
		}
		product, err := c.service.UpdateDescription(int64(int(id)), req.Description)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, product)
	}
}

type request struct {
	ProductCode                    string  `json:"productCode" binding:"required"`
	Description                    string  `json:"description,required" binding:"required"`
	Width                          float64 `json:"width,required" binding:"required"`
	Height                         float64 `json:"height,required" binding:"required"`
	Length                         float64 `json:"length,required" binding:"required"`
	NetWeight                      float64 `json:"netWeight,required" binding:"required"`
	ExpirationRate                 float64 `json:"expirationRate,required" binding:"required"`
	RecommendedFreezingTemperature float64 `json:"recommendedFreezingTemperature,required" binding:"required"`
	FreezingRate                   int     `json:"freezingRate,required" binding:"required"`
	ProductTypeId                  int     `json:"productTypeId,required" binding:"required"`
	SellerId                       int     `json:"sellerId,required" binding:"required"`
}
