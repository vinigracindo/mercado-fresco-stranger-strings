package testutil

import (
	"bytes"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	return router
}

func ExecuteTestRequest(router *gin.Engine, method string, path string, body []byte) *httptest.ResponseRecorder {
	request := httptest.NewRequest(method, path, bytes.NewBuffer(body))
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	return response
}
