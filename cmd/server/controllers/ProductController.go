package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/product"
)

type requestProductPost struct {
	ProductCode                    string  `json:"product_code" binding:"required"`
	Description                    string  `json:"description" binding:"required"`
	Width                          float64 `json:"width" binding:"required"`
	Height                         float64 `json:"height" binding:"required"`
	Length                         float64 `json:"length" binding:"required"`
	NetWeight                      float64 `json:"net_weight" binding:"required"`
	ExpirationRate                 float64 `json:"expiration_rate" binding:"required"`
	RecommendedFreezingTemperature float64 `json:"recommended_freezing_temperature" binding:"required"`
	FreezingRate                   int     `json:"freezing_rate" binding:"required"`
	ProductTypeId                  int     `json:"product_type_id" binding:"required"`
	SellerId                       int     `json:"seller_id" binding:"required"`
}
type requestProductPatch struct {
	Description string `json:"description" binding:"required"`
}

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
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"data": products})
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
		ctx.JSON(http.StatusOK, gin.H{"data": *product})
	}
}

func (c *ProductController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req requestProductPost
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity,
				gin.H{
					"message": "Invalid input. Check the data entered",
				})
			return
		}
		newProduct, err := c.service.Create(req.ProductCode, req.Description, req.Width, req.Height, req.Length, req.NetWeight,
			req.ExpirationRate, req.RecommendedFreezingTemperature, req.FreezingRate, req.ProductTypeId, req.SellerId)

		if err != nil {
			ctx.JSON(http.StatusConflict, gin.H{
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
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
			return
		}
		var req requestProductPatch
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if req.Description == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Description field is required"})
			return
		}
		product, err := c.service.UpdateDescription(id, req.Description)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, product)
	}
}

func (c *ProductController) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
			return
		}

		err = c.service.Delete(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusNoContent, gin.H{})
	}
}
