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
	ProvinceId   int64  `json:"province_id"`
}

type Locality struct {
	service domain.LocalityService
}

func NewLocalityController(s domain.LocalityService) *Locality {
	return &Locality{
		service: s,
	}
}

// ReportCarrie godoc
// @Summary      Report carries
// @Description  Report carries
// @Tags         Localities
// @Accept       json
// @Produce      json
// @Param	id 	 query int false "locality ID"
// @Success      200
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Router /localities/reportCarries [get]
func (l Locality) ReportCarrie() gin.HandlerFunc {
	return func(ctx *gin.Context) {

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
}

// Store godoc
// @Summary      Create locality
// @Description  create localite
// @Tags         Localities
// @Accept       json
// @Produce      json
// @Param Locality body RequestLocalityPost true "Create locality"
// @Success      201  {object} domain.LocalityModel
// @Failure      409  {object}  httputil.HTTPError
// @Failure      422  {object}  httputil.HTTPError
// @Router /localities [post]
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

// GetReportLocalities godoc
// @Summary      Report localities by seller
// @Description  Report localities by seller
// @Tags         Localities
// @Accept       json
// @Produce      json
// @Param	id 	 query int false "Seller ID"
// @Success      204
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Router /localities/reportSellers [get]
func (c Locality) GetReportLocalities() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam, isPresent := ctx.GetQuery("id")
		if isPresent {
			c.getReportLocalitiesById(ctx, idParam)
			return
		}
		c.getAllReportLocalities(ctx)
	}
}

func (c Locality) getReportLocalitiesById(ctx *gin.Context, idParam string) {
	localitytID, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	result, err := c.service.GetByIdReportSeller(ctx.Request.Context(), localitytID)
	if err != nil {
		httputil.NewError(ctx, http.StatusNotFound, err)
		return
	}

	httputil.NewResponse(ctx, http.StatusOK, result)
}

func (c Locality) getAllReportLocalities(ctx *gin.Context) {
	result, err := c.service.GetAllReportSeller(ctx.Request.Context())

	if err != nil {
		httputil.NewError(ctx, http.StatusInternalServerError, err)
		return
	}

	httputil.NewResponse(ctx, http.StatusOK, result)
}
