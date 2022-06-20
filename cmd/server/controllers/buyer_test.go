package controllers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/cmd/server/controllers"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/buyer"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/buyer/mocks"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func CreateRequestTest(gin *gin.Engine, method string, url string, body []byte) *httptest.ResponseRecorder {

	request := httptest.NewRequest(method, url, bytes.NewBuffer(body))
	response := httptest.NewRecorder()

	gin.ServeHTTP(response, request)

	return response
}

func Test_Controller_Buyer_CreateBuyer(t *testing.T) {
	expectBuyer := &buyer.Buyer{

		Id:           0,
		CardNumberId: 402323,
		FirstName:    "FirstNameTest",
		LastName:     "LastNameTest",
	}

	body := &buyer.Buyer{
		Id:           0,
		CardNumberId: 402323,
		FirstName:    "FirstNameTest",
		LastName:     "LastNameTest",
	}

	bodyFail := &buyer.Buyer{
		Id:           0,
		CardNumberId: 0,
		FirstName:    "",
		LastName:     "",
	}

	t.Run("create_ok", func(t *testing.T) {

		service := mocks.NewService(t)

		service.On("Create",
			expectBuyer.CardNumberId,
			expectBuyer.FirstName,
			expectBuyer.LastName).Return(expectBuyer, nil)

		controller := controllers.NewBuyer(service)
		requestBody, _ := json.Marshal(body)

		r := SetUpRouter()
		r.POST("/api/v1/buyers", controller.Create())
		response := CreateRequestTest(r, "POST", "/api/v1/buyers", requestBody)

		assert.Equal(t, response.Code, http.StatusCreated)
	})

	t.Run("create_fail", func(t *testing.T) {

		service := mocks.NewService(t)
		controller := controllers.NewBuyer(service)
		requestBody, _ := json.Marshal(bodyFail)

		r := SetUpRouter()
		r.POST("/api/v1/buyers", controller.Create())
		response := CreateRequestTest(r, "POST", "/api/v1/buyers", requestBody)

		assert.Equal(t, response.Code, http.StatusUnprocessableEntity)
	})

	t.Run("create_conflict", func(t *testing.T) {

		service := mocks.NewService(t)

		service.On("Create",
			expectBuyer.CardNumberId,
			expectBuyer.FirstName,
			expectBuyer.LastName).Return(nil, fmt.Errorf("buyer already registered: %d", expectBuyer.CardNumberId))

		controller := controllers.NewBuyer(service)
		requestBody, _ := json.Marshal(body)

		r := SetUpRouter()
		r.POST("/api/v1/buyers", controller.Create())
		response := CreateRequestTest(r, "POST", "/api/v1/buyers", requestBody)

		assert.Equal(t, response.Code, http.StatusConflict)
	})
}
