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
