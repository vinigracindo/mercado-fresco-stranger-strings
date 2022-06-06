package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/section"
)

type requestSectionPost struct {
	Id                 int64 `json:"id"`
	SectionNumber      int64 `json:"section_number" binding:"required"`
	CurrentTemperature int64 `json:"current_temperature" binding:"required"`
	MinimumTemperature int64 `json:"minimum_temperature" binding:"required"`
	CurrentCapacity    int64 `json:"current_capacity" binding:"required"`
	MinimumCapacity    int64 `json:"minimum_capacity" binding:"required"`
	MaximumCapacity    int64 `json:"maximum_capacity" binding:"required"`
	WarehouseId        int64 `json:"warehouse_id" binding:"required"`
	ProductTypeId      int64 `json:"product_type_id" binding:"required"`
}

type requestSectionPatch struct {
	Id              int64 `json:"id"`
	CurrentCapacity int64 `json:"current_capacity" binding:"required"`
}

type ControllerSection struct {
	service section.Service
}

func NewSection(s section.Service) *ControllerSection {
	return &ControllerSection{
		service: s,
	}
}

func (c *ControllerSection) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		err = c.service.Delete(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusNoContent, gin.H{})
	}
}

func (c *ControllerSection) UpdateCurrentCapacity() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		var req requestSectionPatch
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if req.CurrentCapacity == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "The field CurrentCapacity is required "})
			return
		}

		section, err := c.service.UpdateCurrentCapacity(id, req.CurrentCapacity)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(200, gin.H{"data": section})
	}
}

func (c ControllerSection) CreateSection() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req requestSectionPost

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
				"error":   err.Error(),
				"message": "Invalid request",
			})
			return
		}

		response, err := c.service.CreateSection(
			req.SectionNumber,
			req.CurrentTemperature,
			req.MinimumTemperature,
			req.CurrentCapacity,
			req.MinimumCapacity,
			req.MaximumCapacity,
			req.WarehouseId,
			req.ProductTypeId,
		)

		if err != nil {
			ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, &response)
	}
}

func (c *ControllerSection) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		section, err := c.service.GetById(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": section,
		})
	}
}

func (c *ControllerSection) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		section, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, section)
	}
}