package httputil

import (
	"github.com/gin-gonic/gin"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/libs/logger"
)

func NewError(ctx *gin.Context, status int, err error) {
	er := HTTPError{
		Code:    status,
		Message: err.Error(),
	}
	ctx.JSON(status, er)

	logger.Logger.Error(ctx, ctx.Request.Method, ctx.Request.RequestURI, err.Error(), ctx.Writer.Status())
}

type HTTPError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
