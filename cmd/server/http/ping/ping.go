package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct{}

func NewController() *Controller {
	return &Controller{}
}

func (controller *Controller) HandlePing(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}
