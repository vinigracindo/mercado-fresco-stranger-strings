package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/warehouse"
)

type Warehouse struct {
	service warehouse.Service
}

func NewWarehouse(s warehouse.Service) *Warehouse {
	return &Warehouse{
		service: s,
	}
}

func (w Warehouse) CreateWarehouse() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var wh warehouse.WarehouseModel

		if err := ctx.BindJSON(&wh); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		newWh, err := w.service.Create(&wh)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		ctx.JSON(http.StatusCreated, &newWh)
	}
}

func (w Warehouse) GetAllWarehouse() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		shw, err := w.service.GetAll()

		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		ctx.JSON(http.StatusOK, &shw)
	}
}
