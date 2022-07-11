package controllers_test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	controllers "github.com/vinigracindo/mercado-fresco-stranger-strings/cmd/server/controllers/employees"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/employees/domain"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/employees/domain/mocks"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/pkg/testutil"
)

var (
	EndpointEmployee = "/api/v1/employees"
	invalidJSON      = []byte(`{invalid json}`)
)

func makeEmployee() domain.Employee {
	return domain.Employee{
		Id:           1,
		CardNumberId: "123456",
		FirstName:    "John",
		LastName:     "Doe",
		WarehouseId:  1,
	}
}

func makeEmployeeInboundOrdersReport() domain.EmployeeInboundOrdersReport {
	return domain.EmployeeInboundOrdersReport{
		Employee: makeEmployee(),
		Count:    10,
	}
}

func TestEmployeeController_Create(t *testing.T) {
	expectedEmployee := makeEmployee()
	mockService := mocks.NewEmployeeService(t)
	controller := controllers.NewEmployeeController(mockService)
	router := testutil.SetUpRouter()
	router.POST(EndpointEmployee, controller.Create())

	t.Run("create_ok: when data entry is successful, should return code 201. The object must be returned.", func(t *testing.T) {
		mockService.
			On("Create", context.TODO(), "123456", "John", "Doe", int64(1)).
			Return(expectedEmployee, nil).
			Once()

		reqBody, _ := json.Marshal(expectedEmployee)
		response := testutil.ExecuteTestRequest(router, http.MethodPost, EndpointEmployee, reqBody)

		assert.Equal(t, http.StatusCreated, response.Code)
	})

	t.Run("create_fail: when the JSON does not contain the required fields, should return code 422.", func(t *testing.T) {
		reqBody, _ := json.Marshal(invalidJSON)
		response := testutil.ExecuteTestRequest(router, http.MethodPost, EndpointEmployee, reqBody)

		assert.Equal(t, http.StatusUnprocessableEntity, response.Code)
	})

	t.Run("create_conflict: when the card_number already exists, should return code 409.", func(t *testing.T) {
		mockService.
			On("Create", context.TODO(), "123456", "John", "Doe", int64(1)).
			Return(domain.Employee{}, domain.ErrCardNumberMustBeUnique).
			Once()

		reqBody, _ := json.Marshal(expectedEmployee)
		response := testutil.ExecuteTestRequest(router, http.MethodPost, EndpointEmployee, reqBody)

		assert.Equal(t, http.StatusConflict, response.Code)
	})

}

func TestEmployeeController_GetAll(t *testing.T) {
	mockService := mocks.NewEmployeeService(t)
	controller := controllers.NewEmployeeController(mockService)
	router := testutil.SetUpRouter()
	router.GET(EndpointEmployee, controller.GetAll())

	expectedEmployees := []domain.Employee{
		makeEmployee(),
		makeEmployee(),
	}

	t.Run("find_all: when data entry is successful, should return code 200.", func(t *testing.T) {
		mockService.
			On("GetAll", context.TODO()).
			Return(expectedEmployees, nil).
			Once()

		response := testutil.ExecuteTestRequest(router, http.MethodGet, EndpointEmployee, nil)

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
		mockService.On("GetAll", context.TODO()).Return([]domain.Employee{}, errors.New("internal server error.")).Once()
		response := testutil.ExecuteTestRequest(router, http.MethodGet, EndpointEmployee, nil)
		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})
}

func TestEmployeeController_GetById(t *testing.T) {
	mockService := mocks.NewEmployeeService(t)
	controller := controllers.NewEmployeeController(mockService)
	router := testutil.SetUpRouter()
	router.GET(EndpointEmployee+"/:id", controller.GetById())
	url := fmt.Sprintf("%s/%d", EndpointEmployee, 1)

	t.Run("find_by_id_non_existent: when the employee does not exist, should return code 404.", func(t *testing.T) {
		mockService.
			On("GetById", context.TODO(), int64(1)).
			Return(nil, domain.ErrEmployeeNotFound).
			Once()

		response := testutil.ExecuteTestRequest(router, http.MethodGet, url, nil)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	t.Run("find_by_id_existent: when the request is successful, should return code 200", func(t *testing.T) {
		expectedEmployee := makeEmployee()

		mockService.
			On("GetById", context.TODO(), int64(1)).
			Return(&expectedEmployee, nil).
			Once()

		response := testutil.ExecuteTestRequest(router, http.MethodGet, url, nil)

		expectedBody := `
			{"data":
				{"id":1,"card_number_id":"123456","first_name":"John","last_name":"Doe","warehouse_id":1}
			}`

		assert.Equal(t, http.StatusOK, response.Code)
		assert.JSONEq(t, expectedBody, response.Body.String())
	})

	t.Run("invalid_id: when section id is not parsed, should return code 400.", func(t *testing.T) {
		url := fmt.Sprintf("%s/%s", EndpointEmployee, "invalid_id")
		response := testutil.ExecuteTestRequest(router, http.MethodGet, url, nil)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
}

func TestEmployeeController_Update(t *testing.T) {
	mockService := mocks.NewEmployeeService(t)
	controller := controllers.NewEmployeeController(mockService)
	router := testutil.SetUpRouter()
	router.PATCH(EndpointEmployee+"/:id", controller.UpdateFullname())
	url := fmt.Sprintf("%s/%d", EndpointEmployee, 1)

	t.Run("update_ok: when the request is successful, should return code 200. The object must be returned.", func(t *testing.T) {
		expectedEmployee := makeEmployee()
		expectedEmployee.FirstName = "Jane"
		expectedEmployee.LastName = "Doe"

		mockService.
			On("UpdateFullname", context.TODO(), int64(1), "Jane", "Doe").
			Return(&expectedEmployee, nil).
			Once()

		reqBody, _ := json.Marshal(expectedEmployee)
		response := testutil.ExecuteTestRequest(router, http.MethodPatch, url, reqBody)

		expectedOutput := map[string]any{
			"data": expectedEmployee,
		}

		expectedResponseBody, _ := json.Marshal(expectedOutput)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.JSONEq(t, string(expectedResponseBody), response.Body.String())
	})

	t.Run("update_non_existent: when the employee does not exist, should return code 404.", func(t *testing.T) {
		mockService.
			On("UpdateFullname", context.TODO(), int64(1), "John", "Doe").
			Return(nil, domain.ErrEmployeeNotFound).
			Once()

		reqBody, _ := json.Marshal(makeEmployee())
		response := testutil.ExecuteTestRequest(router, http.MethodPatch, url, reqBody)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	t.Run("invalid_id: when employee id is not parsed, should return code 400.", func(t *testing.T) {
		url := fmt.Sprintf("%s/%s", EndpointEmployee, "invalid_id")
		response := testutil.ExecuteTestRequest(router, http.MethodPatch, url, nil)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})

	t.Run("invalid json: when the request body is not valid json, should return code 400.", func(t *testing.T) {
		response := testutil.ExecuteTestRequest(router, http.MethodPatch, url, invalidJSON)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
}

func TestEmployeeController_Delete(t *testing.T) {
	mockService := mocks.NewEmployeeService(t)
	controller := controllers.NewEmployeeController(mockService)
	router := testutil.SetUpRouter()
	router.DELETE(EndpointEmployee+"/:id", controller.Delete())
	url := fmt.Sprintf("%s/%d", EndpointEmployee, 1)

	t.Run("delete_non_existent: when the employee does not exist, should return code 404.", func(t *testing.T) {
		mockService.
			On("Delete", context.TODO(), int64(1)).
			Return(domain.ErrEmployeeNotFound).
			Once()

		response := testutil.ExecuteTestRequest(router, http.MethodDelete, url, nil)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	t.Run("delete_ok: when the request is successful, should return code 204.", func(t *testing.T) {
		mockService.
			On("Delete", context.TODO(), int64(1)).
			Return(nil).
			Once()

		response := testutil.ExecuteTestRequest(router, http.MethodDelete, url, nil)

		assert.Equal(t, http.StatusNoContent, response.Code)
	})

	t.Run("invalid_id: when section id is not parsed, should return code 400.", func(t *testing.T) {
		url := fmt.Sprintf("%s/%s", EndpointEmployee, "invalid_id")
		response := testutil.ExecuteTestRequest(router, http.MethodDelete, url, nil)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
}

func TestEmployeeController_ReportInboundOrders(t *testing.T) {
	url := fmt.Sprintf("%s/%s", EndpointEmployee, "reportInboundOrders")

	mockService := mocks.NewEmployeeService(t)
	controller := controllers.NewEmployeeController(mockService)
	router := testutil.SetUpRouter()
	router.GET(url, controller.GetReportInboundOrders())

	t.Run("report_inbound_orders_ok: when the request is successful, should return code 200. The object must be returned.", func(t *testing.T) {
		expectedResult := []domain.EmployeeInboundOrdersReport{
			makeEmployeeInboundOrdersReport(),
		}
		mockService.
			On("GetAllReportInboundOrders", context.TODO()).
			Return(expectedResult, nil).
			Once()

		response := testutil.ExecuteTestRequest(router, http.MethodGet, url, nil)

		expectedBody := `
		{"data":
			[
				{"id":1,"card_number_id":"123456","first_name":"John","last_name":"Doe","warehouse_id":1, "inbound_orders_count":10}
			]
		}`

		assert.Equal(t, http.StatusOK, response.Code)
		assert.JSONEq(t, expectedBody, response.Body.String())

	})

	t.Run("report_service_fail: when the service fails, should return code 500.", func(t *testing.T) {
		mockService.
			On("GetAllReportInboundOrders", context.TODO()).
			Return(nil, fmt.Errorf("error")).
			Once()

		response := testutil.ExecuteTestRequest(router, http.MethodGet, url, nil)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})

	t.Run("invalid_query_params: when the query params are not valid, should return code 400.", func(t *testing.T) {
		url := fmt.Sprintf("%s/%s?id=invalid_value", EndpointEmployee, "reportInboundOrders")
		response := testutil.ExecuteTestRequest(router, http.MethodGet, url, nil)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})

	t.Run("internal_server_error: when the service fails on get report by id, should return code 500.", func(t *testing.T) {
		employeeID := int64(1)
		url_with_id := fmt.Sprintf("%s/%s?id=%d", EndpointEmployee, "reportInboundOrders", employeeID)
		mockService.
			On("GetReportInboundOrdersById", context.TODO(), employeeID).
			Return(domain.EmployeeInboundOrdersReport{}, fmt.Errorf("error")).
			Once()

		response := testutil.ExecuteTestRequest(router, http.MethodGet, url_with_id, nil)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})

	t.Run("report_inbound_orders_ok_with_id: when the request is successful, should return code 200. The object must be returned.", func(t *testing.T) {
		url_with_id := fmt.Sprintf("%s/%s?id=%d", EndpointEmployee, "reportInboundOrders", 1)
		expectedResult := makeEmployeeInboundOrdersReport()

		employeeID := int64(1)

		mockService.
			On("GetReportInboundOrdersById", context.TODO(), employeeID).
			Return(expectedResult, nil).
			Once()

		response := testutil.ExecuteTestRequest(router, http.MethodGet, url_with_id, nil)

		expectedBody := `
			{"data":
				{"id":1,"card_number_id":"123456","first_name":"John","last_name":"Doe","warehouse_id":1, "inbound_orders_count":10}
			}`

		assert.Equal(t, http.StatusOK, response.Code)
		assert.JSONEq(t, expectedBody, response.Body.String())
	})

	t.Run("employee_not_found: when the employee does not exist, should return code 404.", func(t *testing.T) {
		employeeID := int64(1)
		url_with_id := fmt.Sprintf("%s/%s?id=%d", EndpointEmployee, "reportInboundOrders", employeeID)
		mockService.
			On("GetReportInboundOrdersById", context.TODO(), employeeID).
			Return(domain.EmployeeInboundOrdersReport{}, domain.ErrEmployeeNotFound).
			Once()

		response := testutil.ExecuteTestRequest(router, http.MethodGet, url_with_id, nil)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})
}
