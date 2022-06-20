package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/warehouse"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/pkg/httputil"
)

type RequestWarehousePost struct {
	Address            string  `json:"address" binding:"required"`
	Telephone          string  `json:"telephone" binding:"required"`
	WarehouseCode      string  `json:"warehouse_code" binding:"required"`
	MinimunCapacity    int64   `json:"minimun_capacity" binding:"required"`
	MinimunTemperature float64 `json:"minimun_temperature" binding:"required"`
}

type RequestWarehousePatch struct {
	MinimunCapacity    int64   `json:"minimun_capacity"`
	MinimunTemperature float64 `json:"minimun_temperature"`
}

type Warehouse struct {
	service warehouse.Service
}

func NewWarehouse(s warehouse.Service) *Warehouse {
	return &Warehouse{
		service: s,
	}
}

// Warehouse godoc
// @Summary      Create warehouse
// @Description  create warehouse
// @Tags         Warehouse
// @Accept       json
// @Produce      json
// @Param Warehouse body requestWarehousePost true "Create warehouse"
// @Success      201  {object}  warehouse.WarehouseModel
// @Failure      409  {object}  httputil.HTTPError
// @Failure      422  {object}  httputil.HTTPError
// @Router /warehouses [post]
func (w Warehouse) CreateWarehouse() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var wh RequestWarehousePost

		if err := ctx.ShouldBindJSON(&wh); err != nil {
			httputil.NewError(ctx, http.StatusUnprocessableEntity, err)
			return
		}

		newWh, err := w.service.Create(wh.Address, wh.Telephone, wh.WarehouseCode, wh.MinimunTemperature, wh.MinimunCapacity)

		if err != nil {
			httputil.NewError(ctx, http.StatusConflict, err)
			return
		}

		httputil.NewResponse(ctx, http.StatusCreated, newWh)
	}
}

// Warehouse godoc
// @Summary      List all warehouse
// @Description  get Warehouse
// @Tags         Warehouse
// @Accept       json
// @Produce      json
// @Success      200  {object} []warehouse.WarehouseModel
// @Failure      404  {object}  httputil.HTTPError
// @Router /warehouses [get]
func (w Warehouse) GetAllWarehouse() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		shw, err := w.service.GetAll()

		if err != nil {
			httputil.NewError(ctx, http.StatusInternalServerError, err)
			return
		}

		httputil.NewResponse(ctx, http.StatusOK, shw)
	}
}

// Warehouse godoc
// @Summary      List Warehouse by id
// @Description  get Warehouse by id
// @Tags         Warehouse
// @Accept       json
// @Produce      json
// @Param id path int true "Warehouse ID"
// @Success      200  {object} warehouse.WarehouseModel
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Router /warehouses/{id} [get]
func (w Warehouse) GetWarehouseByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if paramId, check := ctx.Params.Get("id"); check {
			id, err := strconv.Atoi(paramId)

			if err != nil {
				httputil.NewError(ctx, http.StatusBadRequest, err)
				return
			}

			wh, err := w.service.GetById(int64(id))

			if err != nil {
				httputil.NewError(ctx, http.StatusNotFound, err)
				return
			}

			httputil.NewResponse(ctx, http.StatusOK, wh)

		}
	}

}

// Warehouse godoc
// @Summary      Delete Warehouse
// @Description  Delete Warehouse by id
// @Tags         Warehouse
// @Accept       json
// @Produce      json
// @Param id path int true "Warehouse ID"
// @Success      204
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Router /warehouses/{id} [delete]
func (w Warehouse) DeleteWarehouse() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if paramId, check := ctx.Params.Get("id"); check {
			id, err := strconv.Atoi(paramId)

			if err != nil {
				httputil.NewError(ctx, http.StatusInternalServerError, err)
				return
			}

			err = w.service.Delete(int64(id))

			if err != nil {
				httputil.NewError(ctx, http.StatusBadRequest, err)
				return
			}

			httputil.NewResponse(ctx, http.StatusNoContent, gin.H{})

		}
	}

}

// Warehouse godoc
// @Summary      Update warehouse
// @Description  Update warehouse
// @Tags         Warehouse
// @Accept       json
// @Produce      json
// @Param id path int true "Warehouse ID"
// @Param Warehouse body requestWarehousePatch true "Update warehouse"
// @Success      201  {object} warehouse.WarehouseModel
// @Failure      409  {object}  httputil.HTTPError
// @Failure      422  {object}  httputil.HTTPError
// @Router /warehouses/{id} [patch]
func (w Warehouse) UpdateWarehouse() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if paramId, check := ctx.Params.Get("id"); check {
			id, err := strconv.Atoi(paramId)

			if err != nil {
				httputil.NewError(ctx, http.StatusBadRequest, err)
			}

			var body RequestWarehousePatch
			var patchWh warehouse.WarehouseModel

			if err := ctx.ShouldBindJSON(&body); err != nil {
				httputil.NewError(ctx, http.StatusBadRequest, err)
				return
			}

			patchWh, err = w.service.UpdateTempAndCap(int64(id), body.MinimunTemperature, body.MinimunCapacity)

			if err != nil {
				httputil.NewError(ctx, http.StatusNotFound, err)
				return
			}

			httputil.NewResponse(ctx, http.StatusOK, patchWh)
			return
		}
	}
}
