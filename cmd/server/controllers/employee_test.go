package controllers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/cmd/server/controllers"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/employees"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/employees/mocks"
)

var (
	ENDPOINT    = "/api/v1/employees"
	invalidJSON = []byte(`{invalid json}`)
)

func makeEmployee() employees.Employee {
	return employees.Employee{
		Id:           1,
		CardNumberId: "123456",
		FirstName:    "John",
		LastName:     "Doe",
		WarehouseId:  1,
	}
}

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

func TestEmployeeController_Create(t *testing.T) {
	expectedEmployee := makeEmployee()
	mockService := mocks.NewService(t)
	controller := controllers.NewEmployee(mockService)
	router := SetUpRouter()
	router.POST(ENDPOINT, controller.Create())

	t.Run("create_ok: when data entry is successful, should return code 201. The object must be returned.", func(t *testing.T) {
		mockService.
			On("Create", "123456", "John", "Doe", int64(1)).
			Return(expectedEmployee, nil).
			Once()

		reqBody, _ := json.Marshal(expectedEmployee)
		response := ExecuteTestRequest(router, http.MethodPost, ENDPOINT, reqBody)

		assert.Equal(t, http.StatusCreated, response.Code)
	})

	t.Run("create_fail: when the JSON does not contain the required fields, should return code 422.", func(t *testing.T) {
		reqBody, _ := json.Marshal(invalidJSON)
		response := ExecuteTestRequest(router, http.MethodPost, ENDPOINT, reqBody)

		assert.Equal(t, http.StatusUnprocessableEntity, response.Code)
	})

	t.Run("create_conflict: when the card_number already exists, should return code 409.", func(t *testing.T) {
		mockService.
			On("Create", "123456", "John", "Doe", int64(1)).
			Return(employees.Employee{}, employees.ErrCardNumberMustBeUnique).
			Once()

		reqBody, _ := json.Marshal(expectedEmployee)
		response := ExecuteTestRequest(router, http.MethodPost, ENDPOINT, reqBody)

		assert.Equal(t, http.StatusConflict, response.Code)
	})

}

func TestEmployeeController_GetAll(t *testing.T) {
	mockService := mocks.NewService(t)
	controller := controllers.NewEmployee(mockService)
	router := SetUpRouter()
	router.GET(ENDPOINT, controller.GetAll())

	expectedEmployees := []employees.Employee{
		makeEmployee(),
		makeEmployee(),
	}

	t.Run("find_all: when data entry is successful, should return code 200.", func(t *testing.T) {
		mockService.
			On("GetAll").
			Return(expectedEmployees, nil).
			Once()

		response := ExecuteTestRequest(router, http.MethodGet, ENDPOINT, nil)

		expectedBody := `
			{"data":
				[
					{"id":1,"card_number_id":"123456","first_name":"John","last_name":"Doe","warehouse_id":1},
					{"id":1,"card_number_id":"123456","first_name":"John","last_name":"Doe","warehouse_id":1}
				]
			}`

		assert.Equal(t, http.StatusOK, response.Code)
		assert.JSONEq(t, expectedBody, response.Body.String())
	})

	t.Run("find_all_err: when internal error occurs, should return code 500.", func(t *testing.T) {
		mockService.On("GetAll").Return([]employees.Employee{}, errors.New("internal server error.")).Once()
		response := ExecuteTestRequest(router, http.MethodGet, ENDPOINT, nil)
		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})
}

func TestEmployeeController_GetById(t *testing.T) {
	mockService := mocks.NewService(t)
	controller := controllers.NewEmployee(mockService)
	router := SetUpRouter()
	router.GET(ENDPOINT+"/:id", controller.GetById())
	url := fmt.Sprintf("%s/%d", ENDPOINT, 1)

	t.Run("find_by_id_non_existent: when the employee does not exist, should return code 404.", func(t *testing.T) {
		mockService.
			On("GetById", int64(1)).
			Return(nil, employees.ErrEmployeeNotFound).
			Once()

		response := ExecuteTestRequest(router, http.MethodGet, url, nil)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	t.Run("find_by_id_existent: when the request is successful, should return code 200", func(t *testing.T) {
		expectedEmployee := makeEmployee()

		mockService.
			On("GetById", int64(1)).
			Return(&expectedEmployee, nil).
			Once()

		response := ExecuteTestRequest(router, http.MethodGet, url, nil)

		expectedBody := `
			{"data":
				{"id":1,"card_number_id":"123456","first_name":"John","last_name":"Doe","warehouse_id":1}
			}`

		assert.Equal(t, http.StatusOK, response.Code)
		assert.JSONEq(t, expectedBody, response.Body.String())
	})

	t.Run("invalid_id: when section id is not parsed, should return code 400.", func(t *testing.T) {
		url := fmt.Sprintf("%s/%s", ENDPOINT, "invalid_id")
		response := ExecuteTestRequest(router, http.MethodGet, url, nil)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
}

func TestEmployeeController_Update(t *testing.T) {
	mockService := mocks.NewService(t)
	controller := controllers.NewEmployee(mockService)
	router := SetUpRouter()
	router.PATCH(ENDPOINT+"/:id", controller.UpdateFullname())
	url := fmt.Sprintf("%s/%d", ENDPOINT, 1)

	t.Run("update_ok: when the request is successful, should return code 200. The object must be returned.", func(t *testing.T) {
		expectedEmployee := makeEmployee()
		expectedEmployee.FirstName = "Jane"
		expectedEmployee.LastName = "Doe"

		mockService.
			On("UpdateFullname", int64(1), "Jane", "Doe").
			Return(&expectedEmployee, nil).
			Once()

		reqBody, _ := json.Marshal(expectedEmployee)
		response := ExecuteTestRequest(router, http.MethodPatch, url, reqBody)

		expectedOutput := map[string]any{
			"data": expectedEmployee,
		}

		expectedResponseBody, _ := json.Marshal(expectedOutput)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.JSONEq(t, string(expectedResponseBody), response.Body.String())
	})

	t.Run("update_non_existent: when the employee does not exist, should return code 404.", func(t *testing.T) {
		mockService.
			On("UpdateFullname", int64(1), "John", "Doe").
			Return(nil, employees.ErrEmployeeNotFound).
			Once()

		reqBody, _ := json.Marshal(makeEmployee())
		response := ExecuteTestRequest(router, http.MethodPatch, url, reqBody)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	t.Run("invalid_id: when employee id is not parsed, should return code 400.", func(t *testing.T) {
		url := fmt.Sprintf("%s/%s", ENDPOINT, "invalid_id")
		response := ExecuteTestRequest(router, http.MethodPatch, url, nil)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})

	t.Run("invalid json: when the request body is not valid json, should return code 400.", func(t *testing.T) {
		response := ExecuteTestRequest(router, http.MethodPatch, url, invalidJSON)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
}

func TestEmployeeController_Delete(t *testing.T) {
	mockService := mocks.NewService(t)
	controller := controllers.NewEmployee(mockService)
	router := SetUpRouter()
	router.DELETE(ENDPOINT+"/:id", controller.Delete())
	url := fmt.Sprintf("%s/%d", ENDPOINT, 1)

	t.Run("delete_non_existent: when the employee does not exist, should return code 404.", func(t *testing.T) {
		mockService.
			On("Delete", int64(1)).
			Return(employees.ErrEmployeeNotFound).
			Once()

		response := ExecuteTestRequest(router, http.MethodDelete, url, nil)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	t.Run("delete_ok: when the request is successful, should return code 204.", func(t *testing.T) {
		mockService.
			On("Delete", int64(1)).
			Return(nil).
			Once()

		response := ExecuteTestRequest(router, http.MethodDelete, url, nil)

		assert.Equal(t, http.StatusNoContent, response.Code)
	})

	t.Run("invalid_id: when section id is not parsed, should return code 400.", func(t *testing.T) {
		url := fmt.Sprintf("%s/%s", ENDPOINT, "invalid_id")
		response := ExecuteTestRequest(router, http.MethodDelete, url, nil)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
}
