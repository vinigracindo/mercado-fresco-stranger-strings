package httputil

import (
	"github.com/gin-gonic/gin"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/libs/logger"
)

type response struct {
	Data interface{} `json:"data"`
}

func NewResponse(ctx *gin.Context, status int, data interface{}) {
	res := response{
		Data: data,
	}
	ctx.JSON(status, res)

	logger.Logger.Info(ctx, ctx.Request.Method, ctx.Request.RequestURI, "", ctx.Writer.Status())
}
