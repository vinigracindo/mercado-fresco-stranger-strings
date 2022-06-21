package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/cmd/server/controllers"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/seller"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/seller/mocks"
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

const ENDPOINT = "/api/v1/sellers"

var body = seller.Seller{
	Cid:         123,
	CompanyName: "Mercado Livre",
	Address:     "Osasco, SP",
	Telephone:   "11 99999999",
}

var bodyFail = seller.Seller{
	Cid:         0,
	CompanyName: "",
	Address:     "",
	Telephone:   "",
}

var expectedSeller = seller.Seller{
	Id:          1,
	Cid:         123,
	CompanyName: "Mercado Livre",
	Address:     "Osasco, SP",
	Telephone:   "11 99999999",
}

func Test_Controller_Create(t *testing.T) {
	service := mocks.NewService(t)

	t.Run("create_ok: Quando a entrada de dados for bem-sucedida, um código 201 será retornado junto com o objeto inserido", func(t *testing.T) {
		service.On("Create", expectedSeller.Cid, expectedSeller.CompanyName, expectedSeller.Address, expectedSeller.Telephone).Return(expectedSeller, nil)

		controller := controllers.NewSeller(service)

		requestBody, _ := json.Marshal(body)

		r := SetUpRouter()

		r.POST(ENDPOINT, controller.Create())

		response := CreateRequestTest(r, "POST", ENDPOINT, requestBody)

		assert.Equal(t, http.StatusCreated, response.Code)
	})
}
