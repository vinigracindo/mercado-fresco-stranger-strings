package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/section"
	httputil "github.com/vinigracindo/mercado-fresco-stranger-strings/pkg/httputil"
)

type requestSectionPost struct {
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

// Sections godoc
// @Summary      Delete section
// @Description  Delete section by id
// @Tags         Sections
// @Accept       json
// @Produce      json
// @Param id path int true "Section ID"
// @Success      204
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Router /sections/{id} [delete]
func (c *ControllerSection) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			httputil.NewError(ctx, http.StatusBadRequest, err)
			return
		}

		err = c.service.Delete(id)
		if err != nil {
			httputil.NewError(ctx, http.StatusNotFound, err)
			return
		}
		httputil.NewResponse(ctx, http.StatusNoContent, "")
	}
}

// Sections godoc
// @Summary      Update currentCapacity
// @Description  Update currentCapacity field by id
// @Tags         Sections
// @Accept       json
// @Produce      json
// @Param id path int true "Section ID"
// @Param Section body requestSectionPatch true "Update field"
// @Success      200  {object} section.Section
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Router /sections/{id} [patch]
func (c *ControllerSection) UpdateCurrentCapacity() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			httputil.NewError(ctx, http.StatusBadRequest, err)
			return
		}

		var req requestSectionPatch
		if err := ctx.ShouldBindJSON(&req); err != nil {
			httputil.NewError(ctx, http.StatusBadRequest, err)
			return
		}

		if req.CurrentCapacity < 0 {
			httputil.NewError(ctx, http.StatusBadRequest, errors.New("The field CurrentCapacity invalid"))
			return
		}

		section, err := c.service.UpdateCurrentCapacity(id, req.CurrentCapacity)
		if err != nil {
			httputil.NewError(ctx, http.StatusNotFound, err)
			return
		}
		httputil.NewResponse(ctx, http.StatusOK, section)
	}
}

// Sections godoc
// @Summary      Create section
// @Description  create section
// @Tags         Sections
// @Accept       json
// @Produce      json
// @Param Section body requestSectionPost true "Create section"
// @Success      201  {object} section.Section
// @Failure      409  {object}  httputil.HTTPError
// @Failure      422  {object}  httputil.HTTPError
// @Router /sections [post]
func (c ControllerSection) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req requestSectionPost

		if err := ctx.ShouldBindJSON(&req); err != nil {
			httputil.NewError(ctx, http.StatusUnprocessableEntity, err)
			return
		}

		response, err := c.service.Create(
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
			httputil.NewError(ctx, http.StatusConflict, err)
			return
		}
		httputil.NewResponse(ctx, http.StatusCreated, &response)
	}
}

// Sections godoc
// @Summary      List section by id
// @Description  get section by id
// @Tags         Sections
// @Accept       json
// @Produce      json
// @Param id path int true "Section ID"
// @Success      200  {object} section.Section
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Router /sections/{id} [get]
func (c *ControllerSection) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			httputil.NewError(ctx, http.StatusBadRequest, err)
			return
		}

		section, err := c.service.GetById(id)
		if err != nil {
			httputil.NewError(ctx, http.StatusNotFound, err)
			return
		}

		httputil.NewResponse(ctx, http.StatusOK, section)
	}
}

// Sections godoc
// @Summary      List all sections
// @Description  get sections
// @Tags         Sections
// @Accept       json
// @Produce      json
// @Success      200  {object} []section.Section
// @Failure      400  {object}  httputil.HTTPError
// @Router /sections [get]
func (c *ControllerSection) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		section, err := c.service.GetAll()
		if err != nil {
			httputil.NewError(ctx, http.StatusBadRequest, err)
			return
		}
		httputil.NewResponse(ctx, http.StatusOK, section)
	}
}
