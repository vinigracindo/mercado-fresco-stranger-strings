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
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/buyer"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/buyer/mocks"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func ExecuteTestRequest(router *gin.Engine, method string, path string, body []byte) *httptest.ResponseRecorder {

	request := httptest.NewRequest(method, path, bytes.NewBuffer(body))
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	return response
}

var ENDPOINT = "/api/v1/buyers"

var expectBuyer = &buyer.Buyer{
	Id:           0,
	CardNumberId: 402323,
	FirstName:    "FirstNameTest",
	LastName:     "LastNameTest",
}

var body = &buyer.Buyer{
	CardNumberId: 402323,
	FirstName:    "FirstNameTest",
	LastName:     "LastNameTest",
}

func Test_Controller_Create(t *testing.T) {
	service := mocks.NewService(t)

	t.Run("create_ok: when data entry is successful, should return code 201.", func(t *testing.T) {

		service.
			On("Create",
				expectBuyer.CardNumberId,
				expectBuyer.FirstName,
				expectBuyer.LastName).
			Return(expectBuyer, nil).
			Once()

		controller := controllers.NewBuyer(service)
		requestBody, _ := json.Marshal(body)

		r := SetUpRouter()
		r.POST(ENDPOINT, controller.Create())
		response := ExecuteTestRequest(r, http.MethodPost, ENDPOINT, requestBody)

		assert.Equal(t, http.StatusCreated, response.Code)

	})

	t.Run("create_fail: when the JSON does not contain the required fields, should return code 422", func(t *testing.T) {

		service := mocks.NewService(t)
		controller := controllers.NewBuyer(service)

		r := SetUpRouter()
		r.POST(ENDPOINT, controller.Create())
		response := ExecuteTestRequest(r, http.MethodPost, ENDPOINT, []byte{})

		assert.Equal(t, http.StatusUnprocessableEntity, response.Code)
	})

	t.Run("create_conflict: when the card_number already exists, should return code 409.", func(t *testing.T) {

		service.
			On("Create",
				expectBuyer.CardNumberId,
				expectBuyer.FirstName,
				expectBuyer.LastName).
			Return(nil, fmt.Errorf("buyer already registered %d", expectBuyer.CardNumberId)).
			Once()

		controller := controllers.NewBuyer(service)
		requestBody, _ := json.Marshal(body)

		r := SetUpRouter()
		r.POST(ENDPOINT, controller.Create())
		response := ExecuteTestRequest(r, http.MethodPost, ENDPOINT, requestBody)

		assert.Equal(t, http.StatusConflict, response.Code)

	})
}

func Test_Controller_GetAll(t *testing.T) {
	service := mocks.NewService(t)

	t.Run("find_all: when data entry is successful, should return code 200.", func(t *testing.T) {

		service.
			On("GetAll").
			Return([]buyer.Buyer{*expectBuyer}, nil).
			Once()

		controller := controllers.NewBuyer(service)
		requestBody, _ := json.Marshal(body)

		r := SetUpRouter()
		r.GET(ENDPOINT, controller.GetAll())
		response := ExecuteTestRequest(r, http.MethodGet, ENDPOINT, requestBody)

		assert.Equal(t, http.StatusOK, response.Code)

		assert.JSONEq(t, "{\"data\":[{\"id\":0,\"card_number_id\":402323,\"first_name\":\"FirstNameTest\",\"last_name\":\"LastNameTest\"}]}", response.Body.String())
	})

	t.Run("find_all_fail: when GetAll fail, should return code 400.", func(t *testing.T) {

		service.
			On("GetAll").
			Return([]buyer.Buyer{}, fmt.Errorf("error")).
			Once()

		controller := controllers.NewBuyer(service)
		requestBody, _ := json.Marshal(body)

		r := SetUpRouter()
		r.GET(ENDPOINT, controller.GetAll())
		response := ExecuteTestRequest(r, http.MethodGet, ENDPOINT, requestBody)

		assert.Equal(t, http.StatusBadRequest, response.Code)

	})
}

func Test_Controller_GetById(t *testing.T) {
	service := mocks.NewService(t)

	t.Run("find_by_id_existent: when the request is successful, should return code 200", func(t *testing.T) {

		service.
			On("GetId", int64(1)).
			Return(body, nil).
			Once()

		controller := controllers.NewBuyer(service)
		requestBody, _ := json.Marshal(body)

		r := SetUpRouter()
		r.GET(ENDPOINT+"/:id", controller.GetId())
		response := ExecuteTestRequest(r, http.MethodGet, ENDPOINT+"/1", requestBody)
		assert.Equal(t, http.StatusOK, response.Code)
		assert.JSONEq(t, "{\"data\":{\"id\":0,\"card_number_id\":402323,\"first_name\":\"FirstNameTest\",\"last_name\":\"LastNameTest\"}}", response.Body.String())

	})

	t.Run("find_by_id_inexistent: when the buyer does not exist, should return code 404", func(t *testing.T) {

		service.
			On("GetId", int64(1)).
			Return(nil, fmt.Errorf("buyer not found")).
			Once()
		controller := controllers.NewBuyer(service)

		r := SetUpRouter()
		r.GET(ENDPOINT+"/:id", controller.GetId())
		response := ExecuteTestRequest(r, http.MethodGet, ENDPOINT+"/1", []byte{})
		assert.Equal(t, http.StatusNotFound, response.Code)

	})

	t.Run("find_by_id_parse_error: when buyer id is not parsed, should return code 400.", func(t *testing.T) {
		controller := controllers.NewBuyer(service)

		r := SetUpRouter()
		r.GET(ENDPOINT+"/:id", controller.GetId())
		response := ExecuteTestRequest(r, http.MethodGet, ENDPOINT+"/idInvalido", []byte{})

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
}

func Test_Controller_Update(t *testing.T) {
	service := mocks.NewService(t)

	updateBody := &buyer.Buyer{
		Id:           1,
		CardNumberId: 402324,
		LastName:     "LastNameTest 2",
	}

	t.Run("update_ok: when the request is successful, should return code 200", func(t *testing.T) {

		service.
			On("Update", int64(1), updateBody.CardNumberId, updateBody.LastName).
			Return(updateBody, nil).
			Once()

		controller := controllers.NewBuyer(service)
		requestBody, _ := json.Marshal(updateBody)

		r := SetUpRouter()
		r.PATCH(ENDPOINT+"/:id", controller.UpdateCardNumberLastName())
		response := ExecuteTestRequest(r, http.MethodPatch, ENDPOINT+"/1", requestBody)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.JSONEq(t, "{\"data\":{\"id\":1,\"card_number_id\":402324,\"first_name\":\"\",\"last_name\":\"LastNameTest 2\"}}", response.Body.String())
	})

	t.Run("update_non_existent: when the buyer does not exist, should return code 404.", func(t *testing.T) {
		service.
			On("Update", int64(1), updateBody.CardNumberId, updateBody.LastName).
			Return(nil, fmt.Errorf("buyer with id %d not found", int64(1))).
			Once()

		controller := controllers.NewBuyer(service)
		requestBody, _ := json.Marshal(updateBody)

		r := SetUpRouter()
		r.PATCH("/api/v1/buyers/:id", controller.UpdateCardNumberLastName())
		response := ExecuteTestRequest(r, http.MethodPatch, "/api/v1/buyers/1", requestBody)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	t.Run("update_id_parse_error: when buyer id is not parsed, should return code 400.", func(t *testing.T) {
		controller := controllers.NewBuyer(service)
		r := SetUpRouter()
		r.PATCH(ENDPOINT+"/:id", controller.UpdateCardNumberLastName())
		response := ExecuteTestRequest(r, http.MethodPatch, ENDPOINT+"/idInvalido", []byte{})

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})

	t.Run("update_body_invalid: when the request body is not valid json, should return code 400.", func(t *testing.T) {
		controller := controllers.NewBuyer(service)

		r := SetUpRouter()
		r.PATCH(ENDPOINT+"/:id", controller.UpdateCardNumberLastName())
		response := ExecuteTestRequest(r, http.MethodPatch, ENDPOINT+"/1", nil)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
}
func Test_Controller_Delete(t *testing.T) {
	service := mocks.NewService(t)

	t.Run("delete_non_existent: when the buyer does not exist, should return code 404", func(t *testing.T) {

		service.
			On("Delete", int64(1)).
			Return(fmt.Errorf("buyer with id not found")).
			Once()

		controller := controllers.NewBuyer(service)

		r := SetUpRouter()

		r.DELETE(ENDPOINT+"/:id", controller.DeleteBuyer())
		response := ExecuteTestRequest(r, http.MethodDelete, ENDPOINT+"/1", []byte{})

		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	t.Run("delete_ok: when the request is successful, should return code 204.", func(t *testing.T) {

		service.
			On("Delete", int64(1)).
			Return(nil).
			Once()

		controller := controllers.NewBuyer(service)

		r := SetUpRouter()
		r.DELETE(ENDPOINT+"/:id", controller.DeleteBuyer())
		response := ExecuteTestRequest(r, http.MethodDelete, ENDPOINT+"/1", []byte{})

		assert.Equal(t, http.StatusNoContent, response.Code)
	})

	t.Run("delete_id_parse_error: when buyer id is not parsed, should return code 400", func(t *testing.T) {

		controller := controllers.NewBuyer(service)
		r := SetUpRouter()
		r.DELETE(ENDPOINT+"/:id", controller.DeleteBuyer())

		response := ExecuteTestRequest(r, http.MethodDelete, ENDPOINT+"/idInvalido", []byte{})

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
}
