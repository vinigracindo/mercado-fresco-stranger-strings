package controllers

import (
	"log"
	"net/http"
	"strconv"

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

func (w Warehouse) GetWarehouseByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if paramId, check := ctx.Params.Get("id"); check {
			id, err := strconv.Atoi(paramId)

			if err != nil {
				ctx.JSON(http.StatusInternalServerError, "erro: internal error")
				log.Println(err)
				return
			}

			wh, err := w.service.GetById(int64(id))

			if err != nil {
				ctx.JSON(http.StatusBadRequest, err.Error())
				return
			}

			ctx.JSON(http.StatusOK, wh)
		}
	}

}

func (w Warehouse) DeleteWarehouse() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if paramId, check := ctx.Params.Get("id"); check {
			id, err := strconv.Atoi(paramId)

			if err != nil {
				ctx.JSON(http.StatusInternalServerError, "erro: internal error")
				log.Println(err)
				return
			}

			err = w.service.Delete(int64(id))

			if err != nil {
				ctx.JSON(http.StatusBadRequest, err.Error())
				return
			}

			ctx.JSON(http.StatusNoContent, gin.H{})
		}
	}

}
