package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/carry/domain"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/pkg/httputil"
)

type RequestCarryPost struct {
	Cid         int64  `json:"cid" binding:"required"`
	CompanyName string `json:"company_name" binding:"required"`
	Address     string `json:"address" binding:"required"`
	Telephone   string `json:"telephone" binding:"required"`
	LocalityID  int64  `json:"locality_id" binding:"required"`
}

type Carry struct {
	service domain.CarryService
}

func NewCarryController(s domain.CarryService) *Carry {
	return &Carry{
		service: s,
	}
}

// Carry godoc
// @Summary      Create carry
// @Description  create carry
// @Tags         carry
// @Accept       json
// @Produce      json
// @Param carry body RequestCarryPost true "Create carry"
// @Success      201  {object}  domain.CarryModel
// @Failure      409  {object}  httputil.HTTPError
// @Failure      422  {object}  httputil.HTTPError
// @Router /carries [post]
func (c Carry) CreateCarry() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var body RequestCarryPost

		if err := ctx.ShouldBindJSON(&body); err != nil {
			httputil.NewError(ctx, http.StatusUnprocessableEntity, err)
			return
		}

		model := &domain.CarryModel{
			Cid:         body.Cid,
			Address:     body.Address,
			CompanyName: body.CompanyName,
			Telephone:   body.Telephone,
			LocalityID:  body.LocalityID,
		}

		newCarry, err := c.service.Create(ctx.Request.Context(), model)

		if err != nil {
			httputil.NewError(ctx, http.StatusConflict, err)
			return
		}

		httputil.NewResponse(ctx, http.StatusCreated, newCarry)
	}
}
