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
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/employees"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/employees/mocks"
)

var (
	invalidJSON = []byte(`{invalid json}`)
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

func TestEmployeeController_Create(t *testing.T) {
	expectedEmployee := employees.Employee{
		CardNumberId: "123456",
		FirstName:    "John",
		LastName:     "Doe",
		WarehouseId:  1,
	}
	mockService := mocks.NewService(t)
	controller := controllers.NewEmployee(mockService)
	router := SetUpRouter()
	router.POST("/api/v1/employees", controller.Create())

	t.Run("create_ok: Quando a entrada de dados for bem-sucedida, um código 201 será retornado junto com o objeto inserido.", func(t *testing.T) {
		mockService.On("Create", "123456", "John", "Doe", int64(1)).Return(expectedEmployee, nil).Once()

		reqBody, _ := json.Marshal(expectedEmployee)
		response := ExecuteTestRequest(router, http.MethodPost, "/api/v1/employees", reqBody)

		assert.Equal(t, http.StatusCreated, response.Code)
	})

	t.Run("create_fail: Se o objeto JSON não contiver os campos necessários, um código 422 será	retornado.", func(t *testing.T) {
		reqBody, _ := json.Marshal(invalidJSON)
		response := ExecuteTestRequest(router, http.MethodPost, "/api/v1/employees", reqBody)

		assert.Equal(t, http.StatusUnprocessableEntity, response.Code)
	})

	t.Run("create_conflict: Se o card_number_id já existir, ele	retornará um erro 409 Conflict.", func(t *testing.T) {
		mockService.On("Create", "123456", "John", "Doe", int64(1)).Return(employees.Employee{}, employees.ErrCardNumberMustBeUnique).Once()
		reqBody, _ := json.Marshal(expectedEmployee)
		response := ExecuteTestRequest(router, http.MethodPost, "/api/v1/employees", reqBody)

		assert.Equal(t, http.StatusConflict, response.Code)
	})

}
