package locality

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/locality/domain"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/pkg/httputil"
)

type RequestLocalityPost struct {
	LocalityName string `json:"locality_name" binding:"required"`
	ProvinceName string `json:"province_name" binding:"required"`
	CountryName  string `json:"country_name" binding:"required"`
	ProvinceId   int64  `json:"province_id" binding:"required"`
}

type Locality struct {
	service domain.LocalityService
}

func NewLocalityController(s domain.LocalityService) *Locality {
	return &Locality{
		service: s,
	}
}

func (l Locality) ReportCarrie() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// if query_id, check := ctx.GetQuery("id"); check {

		query_id, exist := ctx.GetQuery("id")

		if !exist {
			query_id = "0"
		}

		id, err := strconv.Atoi(query_id)

		if err != nil {
			httputil.NewError(ctx, http.StatusBadRequest, err)
			return
		}

		report_list, err := l.service.ReportCarrie(ctx.Request.Context(), int64(id))

		if err != nil {
			httputil.NewError(ctx, http.StatusNotFound, err)
			return
		}

		httputil.NewResponse(ctx, http.StatusOK, report_list)

	}
	// }
}

func (c Locality) CreateLocality() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req RequestLocalityPost

		if err := ctx.ShouldBindJSON(&req); err != nil {
			httputil.NewError(ctx, http.StatusUnprocessableEntity, err)
			return
		}

		locality := &domain.LocalityModel{
			LocalityName: req.LocalityName,
			ProvinceName: req.ProvinceName,
			CountryName:  req.CountryName,
			ProvinceId:   req.ProvinceId,
		}

		newLocality, err := c.service.CreateLocality(ctx.Request.Context(), locality)

		if err != nil {
			httputil.NewError(ctx, http.StatusConflict, err)
			return
		}
		httputil.NewResponse(ctx, http.StatusCreated, newLocality)
	}
}
