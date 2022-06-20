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

func makeExpectedResponse(data any) string {
	expectedOutput := map[string]any{
		"data": data,
	}
	expectedResponseBody, _ := json.Marshal(expectedOutput)
	return string(expectedResponseBody)
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
	router.POST("/api/v1/employees", controller.Create())

	t.Run("create_ok: Quando a entrada de dados for bem-sucedida, um código 201 será retornado junto com o objeto inserido.", func(t *testing.T) {
		mockService.On("Create", "123456", "John", "Doe", int64(1)).Return(expectedEmployee, nil).Once()

		reqBody, _ := json.Marshal(expectedEmployee)
		response := ExecuteTestRequest(router, http.MethodPost, ENDPOINT, reqBody)

		assert.Equal(t, http.StatusCreated, response.Code)
	})

	t.Run("create_fail: Se o objeto JSON não contiver os campos necessários, um código 422 será	retornado.", func(t *testing.T) {
		reqBody, _ := json.Marshal(invalidJSON)
		response := ExecuteTestRequest(router, http.MethodPost, ENDPOINT, reqBody)

		assert.Equal(t, http.StatusUnprocessableEntity, response.Code)
	})

	t.Run("create_conflict: Se o card_number_id já existir, ele	retornará um erro 409 Conflict.", func(t *testing.T) {
		mockService.On("Create", "123456", "John", "Doe", int64(1)).Return(employees.Employee{}, employees.ErrCardNumberMustBeUnique).Once()
		reqBody, _ := json.Marshal(expectedEmployee)
		response := ExecuteTestRequest(router, http.MethodPost, ENDPOINT, reqBody)

		assert.Equal(t, http.StatusConflict, response.Code)
	})

}

func TestEmployeeController_GetAll(t *testing.T) {
	mockService := mocks.NewService(t)
	controller := controllers.NewEmployee(mockService)
	router := SetUpRouter()
	router.GET("/api/v1/employees", controller.GetAll())

	expectedEmployees := []employees.Employee{
		makeEmployee(),
		makeEmployee(),
	}

	t.Run("find_all: Quando a solicitação for bem-sucedida, o back-end retornará uma lista de todos os funcionários existentes.", func(t *testing.T) {
		mockService.On("GetAll").Return(expectedEmployees, nil).Once()
		response := ExecuteTestRequest(router, http.MethodGet, ENDPOINT, nil)

		expectedResponseBody := makeExpectedResponse(expectedEmployees)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.JSONEq(t, expectedResponseBody, response.Body.String())
	})
}

func TestEmployeeController_GetById(t *testing.T) {
	mockService := mocks.NewService(t)
	controller := controllers.NewEmployee(mockService)
	router := SetUpRouter()
	router.GET("/api/v1/employees/:id", controller.GetById())
	url := fmt.Sprintf("%s/%d", ENDPOINT, 1)

	t.Run("find_by_id_non_existent: Quando o funcionário não existir, um código 404 será retornado.", func(t *testing.T) {
		mockService.On("GetById", int64(1)).Return(nil, employees.ErrEmployeeNotFound).Once()
		response := ExecuteTestRequest(router, http.MethodGet, url, nil)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	t.Run("find_by_id_existent: Quando a solicitação for bem-sucedida, o back-end retornará as informações solicitadas do funcionário", func(t *testing.T) {
		expectedEmployee := makeEmployee()
		mockService.On("GetById", int64(1)).Return(&expectedEmployee, nil).Once()
		response := ExecuteTestRequest(router, http.MethodGet, url, nil)

		expectedResponseBody := makeExpectedResponse(expectedEmployee)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.JSONEq(t, string(expectedResponseBody), response.Body.String())
	})
}

func TestEmployeeController_Update(t *testing.T) {
	mockService := mocks.NewService(t)
	controller := controllers.NewEmployee(mockService)
	router := SetUpRouter()
	router.PATCH("/api/v1/employees/:id", controller.UpdateFullname())
	url := fmt.Sprintf("%s/%d", ENDPOINT, 1)

	t.Run("update_ok: Quando a atualização dos dados for bem sucedida, o funcionário será devolvido	com as informações atualizadas juntamente com um código 200",
		func(t *testing.T) {
			expectedEmployee := makeEmployee()
			expectedEmployee.FirstName = "Jane"
			expectedEmployee.LastName = "Doe"

			mockService.On("UpdateFullname", int64(1), "Jane", "Doe").Return(&expectedEmployee, nil).Once()
			reqBody, _ := json.Marshal(expectedEmployee)
			response := ExecuteTestRequest(router, http.MethodPatch, url, reqBody)

			expectedOutput := map[string]any{
				"data": expectedEmployee,
			}

			expectedResponseBody, _ := json.Marshal(expectedOutput)

			assert.Equal(t, http.StatusOK, response.Code)
			assert.JSONEq(t, string(expectedResponseBody), response.Body.String())
		})

	t.Run("update_non_existent: Se o funcionário a ser atualizado não existir, um código 404 será retornado.", func(t *testing.T) {
		mockService.On("UpdateFullname", int64(1), "John", "Doe").Return(nil, employees.ErrEmployeeNotFound).Once()
		reqBody, _ := json.Marshal(makeEmployee())
		response := ExecuteTestRequest(router, http.MethodPatch, url, reqBody)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})
}

func TestEmployeeController_Delete(t *testing.T) {
	mockService := mocks.NewService(t)
	controller := controllers.NewEmployee(mockService)
	router := SetUpRouter()
	router.DELETE("/api/v1/employees/:id", controller.Delete())
	url := fmt.Sprintf("%s/%d", ENDPOINT, 1)

	t.Run("delete_non_existent: Quando o funcionário não existir, um código 404 será retornado.", func(t *testing.T) {
		mockService.On("Delete", int64(1)).Return(employees.ErrEmployeeNotFound).Once()
		response := ExecuteTestRequest(router, http.MethodDelete, url, nil)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	t.Run("delete_ok: Quando a exclusão for bem-sucedida, um código 204 será retornado.", func(t *testing.T) {
		mockService.On("Delete", int64(1)).Return(nil).Once()
		response := ExecuteTestRequest(router, http.MethodDelete, url, nil)

		assert.Equal(t, http.StatusNoContent, response.Code)
	})
}
