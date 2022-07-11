package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct{}

func NewController() *Controller {
	return &Controller{}
}

// HandlePing returns a successful pong answer to all HTTP requests
// @Summary Healthcheck endpoint
// @Description Returns a successful pong answer to all HTTP requests
// @Tags misc
// @Produce json
// @Success 200 {string} string "pong"
// @Failure 500 {object} httputil.HTTPError
// @Router /ping [get]
func (controller *Controller) HandlePing(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}
