package httputil

import (
	"github.com/gin-gonic/gin"
)

type response struct {
	Data any `json:"data"`
}

func NewResponse(ctx *gin.Context, status int, data interface{}) {
	res := response{
		Data: data,
	}
	ctx.JSON(status, res)
}
