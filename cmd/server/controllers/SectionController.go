package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/section"
)

type ControllerSection struct {
	service section.Service
}

func NewSection(s section.Service) *ControllerSection {
	return &ControllerSection{
		service: s,
	}
}

func (c *ControllerSection) CreateSection() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var body section.Section

		if err := ctx.BindJSON(&body); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid inputs. Please check your inputs",
			})
			return
		}

		response, err := c.service.CreateSection(&body)
		if err != nil {
			ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, &response)
	}
}
