package controllers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
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
		service.On("Create", expectedSeller.Cid, expectedSeller.CompanyName, expectedSeller.Address, expectedSeller.Telephone).Return(expectedSeller, nil).Once()

		controller := controllers.NewSeller(service)

		requestBody, _ := json.Marshal(body)

		r := SetUpRouter()

		r.POST(ENDPOINT, controller.Create())

		response := CreateRequestTest(r, "POST", ENDPOINT, requestBody)

		assert.Equal(t, http.StatusCreated, response.Code)
	})

	/*t.Run("create_bad_request: Quando o JSON tiver um campo incorreto, um código 400 será retornado", func(t *testing.T) {
		service.On("Create", 0, "", "", "").Return(seller.Seller{}, fmt.Errorf("Invalid Request"))

		controller := controllers.NewSeller(service)

		requestBody, _ := json.Marshal(bodyFail)

		r := SetUpRouter()

		r.POST(ENDPOINT, controller.Create())

		response := CreateRequestTest(r, "POST", ENDPOINT, requestBody)

		assert.Equal(t,  response.Code)

	})*/

	t.Run("create_fail: Se o objeto JSON não contiver os campos necessários, um código 422 será retornado", func(t *testing.T) {
		controller := controllers.NewSeller(nil)

		r := SetUpRouter()
		r.POST(ENDPOINT, controller.Create())
		response := CreateRequestTest(r, "POST", ENDPOINT, []byte{})

		assert.Equal(t, http.StatusUnprocessableEntity, response.Code)
	})

	t.Run("create_conflict: se o cid já existir, ele retornará um erro 409 conflict", func(t *testing.T) {
		service.On("Create", expectedSeller.Cid, expectedSeller.CompanyName, expectedSeller.Address, expectedSeller.Telephone).Return(seller.Seller{}, fmt.Errorf("Seller with this cid alredy exists")).Once()

		controller := controllers.NewSeller(service)
		requestBody, _ := json.Marshal(body)
		r := SetUpRouter()
		r.POST(ENDPOINT, controller.Create())

		response := CreateRequestTest(r, "POST", ENDPOINT, requestBody)

		assert.Equal(t, http.StatusConflict, response.Code)

	})
}

func Test_Controller_Get(t *testing.T) {
	service := mocks.NewService(t)
	expectedListSeller := []seller.Seller{
		{
			Id:          1,
			Cid:         123,
			CompanyName: "Mercado Livre",
			Address:     "Osasco, SP",
			Telephone:   "11 99999999",
		},
		{
			Id:          2,
			Cid:         1234,
			CompanyName: "Mercado Pago",
			Address:     "Salvador, BA",
			Telephone:   "11 88888888",
		},
	}
	t.Run("find_all: Quando a solicitação for bem sucedida, o back-end retornará uma lista de todos os vendedores existentes", func(t *testing.T) {
		service.On("GetAll").Return(expectedListSeller, nil).Once()

		controller := controllers.NewSeller(service)
		r := SetUpRouter()
		r.GET(ENDPOINT, controller.GetAll())

		response := CreateRequestTest(r, "GET", ENDPOINT, []byte{})

		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("find_by_id_non_exitent: Quando o vendedor não existir, um código 404 será devolvido", func(t *testing.T) {
		service.On("GetById", int64(9999)).Return(seller.Seller{}, fmt.Errorf("Seller not found")).Once()

		controller := controllers.NewSeller(service)
		r := SetUpRouter()
		r.GET(ENDPOINT+"/:id", controller.GetById())

		response := CreateRequestTest(r, "GET", ENDPOINT+"/9999", []byte{})

		assert.Equal(t, http.StatusNotFound, response.Code)

	})

	t.Run("find_by_id_exixtent: Quando a solicitação for bem sucedida, o back-end retornará as informações solicitadas do vendedor", func(t *testing.T) {
		service.On("GetById", int64(1)).Return(expectedListSeller[0], nil).Once()

		controller := controllers.NewSeller(service)
		requestBody, _ := json.Marshal(body)
		r := SetUpRouter()
		r.GET(ENDPOINT+"/:id", controller.GetById())

		response := CreateRequestTest(r, "GET", ENDPOINT+"/1", requestBody)

		assert.Equal(t, http.StatusOK, response.Code)
	})
}

func Test_Controller_Update(t *testing.T) {
	service := mocks.NewService(t)
	var bodyUpdate = seller.Seller{
		Address:   "Salvador, BA",
		Telephone: "71 88888888",
	}

	t.Run("update_ok: Quando a atualização dos dados for bem sucedida o vendedor será devolvido com as infomações atualizadas juntamente com um código 200", func(t *testing.T) {
		service.On("Update", int64(1), "Salvador, BA", "71 88888888").Return(expectedSeller, nil).Once()

		controller := controllers.NewSeller(service)
		requestBody, _ := json.Marshal(bodyUpdate)
		r := SetUpRouter()
		r.PATCH(ENDPOINT+"/:id", controller.Update())

		response := CreateRequestTest(r, "PATCH", ENDPOINT+"/1", requestBody)

		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("update_non_existent: Se o vendedor a ser atualizado não existir, um código 404 será devolvido", func(t *testing.T) {
		service.On("Update", int64(9999), "Salvador, BA", "71 88888888").Return(seller.Seller{}, fmt.Errorf("Seller not found")).Once()

		controller := controllers.NewSeller(service)
		requestBody, _ := json.Marshal(bodyUpdate)
		r := SetUpRouter()
		r.PATCH(ENDPOINT+"/:id", controller.Update())

		response := CreateRequestTest(r, "PATCH", ENDPOINT+"/9999", requestBody)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})
}

func Test_Controller_Delete(t *testing.T) {
	service := mocks.NewService(t)

	t.Run("delete_non_existent: Quando o vendedor não existir um código 404 deverá ser devolvido", func(t *testing.T) {
		service.On("Delete", int64(9999)).Return(fmt.Errorf("Seller not found")).Once()

		controller := controllers.NewSeller(service)
		requestBody, _ := json.Marshal(body)
		r := SetUpRouter()

		r.PATCH(ENDPOINT+"/:id", controller.Delete())

		response := CreateRequestTest(r, "PATCH", ENDPOINT+"/9999", requestBody)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})
}
