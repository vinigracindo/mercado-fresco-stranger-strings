package controllers_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	controllers "github.com/vinigracindo/mercado-fresco-stranger-strings/cmd/server/controllers/inbound_orders"
	EmployeeDomain "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/employees/domain"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/inboud_orders/domain"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/inboud_orders/domain/mocks"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/pkg/testutil"
)

var (
	EndpointInboundOrders    = "/api/v1/inbound_orders"
	invalidInboundOrdersJSON = []byte(`{invalid json}`)
)

func makeInboundOrderRequest() controllers.RequestInboundOrdersPost {
	return controllers.RequestInboundOrdersPost{
		OrderDate:      "2020-01-01",
		OrderNumber:    "order#1",
		EmployeeId:     1,
		ProductBatchId: 1,
		WarehouseId:    1,
	}
}

func makeInboundOrder() domain.InboundOrders {
	date := time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)
	return domain.InboundOrders{
		OrderDate:      date,
		OrderNumber:    "order#1",
		EmployeeId:     1,
		ProductBatchId: 1,
		WarehouseId:    1,
	}
}

func TestInboundOrders_Create(t *testing.T) {
	anyId := int64(1)
	expectedInboundOrders := makeInboundOrder()
	requestInbound := makeInboundOrderRequest()
	date, _ := time.Parse("2006-01-02", requestInbound.OrderDate)

	mockService := mocks.NewInboundOrdersService(t)
	controller := controllers.NewInboundOrdersController(mockService)
	router := testutil.SetUpRouter()
	router.POST(EndpointInboundOrders, controller.Create())

	t.Run("create_ok: when data entry is successful, should return code 201. The object must be returned.", func(t *testing.T) {
		mockService.
			On("Create", context.TODO(), date, "order#1", anyId, anyId, anyId).
			Return(expectedInboundOrders, nil).
			Once()

		reqBody, _ := json.Marshal(requestInbound)
		response := testutil.ExecuteTestRequest(router, http.MethodPost, EndpointInboundOrders, reqBody)

		assert.Equal(t, http.StatusCreated, response.Code)
	})

	t.Run("invalid_json: when the request body is invalid, should return code 422. The error must be returned.", func(t *testing.T) {
		reqBody, _ := json.Marshal(invalidInboundOrdersJSON)
		response := testutil.ExecuteTestRequest(router, http.MethodPost, EndpointInboundOrders, reqBody)

		assert.Equal(t, http.StatusUnprocessableEntity, response.Code)
	})

	t.Run("invalid_time_parse: when the request body is invalid, should return code 400. The error must be returned.", func(t *testing.T) {
		newRequestInbound := makeInboundOrderRequest()
		newRequestInbound.OrderDate = "invalid date"
		reqBody, _ := json.Marshal(newRequestInbound)
		response := testutil.ExecuteTestRequest(router, http.MethodPost, EndpointInboundOrders, reqBody)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})

	t.Run("employee_not_found: when the employee is not found, should return code 409. The error must be returned.", func(t *testing.T) {
		mockService.
			On("Create", context.TODO(), date, "order#1", anyId, anyId, anyId).
			Return(domain.InboundOrders{}, EmployeeDomain.ErrEmployeeNotFound).
			Once()

		reqBody, _ := json.Marshal(requestInbound)
		response := testutil.ExecuteTestRequest(router, http.MethodPost, EndpointInboundOrders, reqBody)

		assert.Equal(t, http.StatusConflict, response.Code)
	})

	t.Run("create_failed: when the data entry is failed, should return code 500. The error must be returned.", func(t *testing.T) {
		mockService.
			On("Create", context.TODO(), date, "order#1", anyId, anyId, anyId).
			Return(domain.InboundOrders{}, fmt.Errorf("An error")).
			Once()

		reqBody, _ := json.Marshal(requestInbound)
		response := testutil.ExecuteTestRequest(router, http.MethodPost, EndpointInboundOrders, reqBody)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})
}
