package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/section"
)

type Controller struct {
	service section.Service
}

func NewSection(s section.Service) *Controller {
	return &Controller{
		service: s,
	}
}

func (c *Controller) GetAll() gin.HandlerFunc {
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
