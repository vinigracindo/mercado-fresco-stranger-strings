package controllers

import (
	"errors"
	httputil "github.com/vinigracindo/mercado-fresco-stranger-strings/pkg/httputil"
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
	service product.ProductService
}

func CreateProductController(prodService product.ProductService) *ProductController {
	return &(ProductController{service: prodService})
}

// GetAll godoc
// @Summary      List all products
// @Description  get all products
// @Tags         Products
// @Accept       json
// @Produce      json
// @Success      200  {array} product.Product
// @Failure      404  {object}  httputil.HTTPError
// @Router /products [get]
func (c *ProductController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		products, err := c.service.GetAll()

		if err != nil {
			httputil.NewError(ctx, http.StatusBadRequest, err)
			return
		}

		httputil.NewResponse(ctx, http.StatusOK, products)
	}
}

// GetById godoc
// @Summary Get product by ID
// @Tags Products
// @Description Get product by ID
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} product.Product
// @Failure 400  {object}  httputil.HTTPError
// @Failure 404  {object}  httputil.HTTPError
// @Router /products/{id} [get]
func (c *ProductController) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

		if err != nil {
			httputil.NewError(ctx, http.StatusBadRequest, err)
			return
		}

		productId, err := c.service.GetById(id)

		if err != nil {
			httputil.NewError(ctx, http.StatusNotFound, err)
			return
		}

		httputil.NewResponse(ctx, http.StatusOK, productId)
	}
}

// Create godoc
// @Summary      Create product
// @Description  create product
// @Tags         Products
// @Accept       json
// @Produce      json
// @Param Product body requestProductPost true "Create product"
// @Success      201  {object} product.Product
// @Failure      400  {object}  httputil.HTTPError
// @Failure      422  {object}  httputil.HTTPError
// @Router /products [post]
func (c *ProductController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req requestProductPost

		if err := ctx.ShouldBindJSON(&req); err != nil {
			httputil.NewError(ctx, http.StatusUnprocessableEntity, errors.New("invalid input. Check the data entered"))
			return
		}

		newProduct, err := c.service.
			Create(req.ProductCode, req.Description, req.Width, req.Height, req.Length, req.NetWeight, req.ExpirationRate,
				req.RecommendedFreezingTemperature, req.FreezingRate, req.ProductTypeId, req.SellerId)

		if err != nil {
			httputil.NewError(ctx, http.StatusConflict, err)
			return
		}

		httputil.NewResponse(ctx, http.StatusCreated, newProduct)
	}
}

// UpdateDescription godoc
// @Summary      Update product fullname
// @Description  Update product description
// @Tags         Products
// @Accept       json
// @Produce      json
// @Param id path int true "Product ID"
// @Param Product body requestProductPatch true "Update field"
// @Success      200  {object} product.Product
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Router /products/{id} [patch]
func (c *ProductController) UpdateDescription() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

		if err != nil {
			httputil.NewError(ctx, http.StatusBadRequest, errors.New("invalid id"))
			return
		}

		var req requestProductPatch

		if err := ctx.ShouldBindJSON(&req); err != nil {
			httputil.NewError(ctx, http.StatusBadRequest, err)
			return
		}

		if req.Description == "" {
			httputil.NewError(ctx, http.StatusBadRequest, errors.New("description field is required"))
			return
		}

		productUpdate, err := c.service.UpdateDescription(id, req.Description)

		if err != nil {
			httputil.NewError(ctx, http.StatusNotFound, err)
			return
		}

		httputil.NewResponse(ctx, http.StatusOK, productUpdate)

	}
}

// Delete godoc
// @Summary      Delete product
// @Description  Delete product by id
// @Tags         Products
// @Accept       json
// @Produce      json
// @Param id path int true "Product ID"
// @Success      204
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Router /products/{id} [delete]
func (c *ProductController) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}

		err = c.service.Delete(id)

		if err != nil {
			httputil.NewError(ctx, http.StatusNotFound, err)
			return
		}

		httputil.NewResponse(ctx, http.StatusNoContent, err)
	}
}
