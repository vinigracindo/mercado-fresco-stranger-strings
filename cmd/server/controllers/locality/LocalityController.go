package locality

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/locality/domain"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/pkg/httputil"
)

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
		if paramId, check := ctx.Params.Get("id"); check {
			id, err := strconv.Atoi(paramId)

			if err != nil {
				httputil.NewError(ctx, http.StatusBadRequest, err)
				return
			}

			wh, err := l.service.ReportCarrie(ctx.Request.Context(), int64(id))

			if err != nil {
				httputil.NewError(ctx, http.StatusNotFound, err)
				return
			}

			httputil.NewResponse(ctx, http.StatusOK, wh)

		}
	}
}
