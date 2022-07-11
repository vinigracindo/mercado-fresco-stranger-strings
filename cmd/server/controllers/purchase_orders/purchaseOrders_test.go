package controllers_test

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	controllers "github.com/vinigracindo/mercado-fresco-stranger-strings/cmd/server/controllers/purchase_orders"
	buyerDomain "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/buyer/domain"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/purchase_orders/domain"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/purchase_orders/domain/mocks"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/pkg/testutil"
)

var EndpointPurchaseOrders = "/api/v1/purchaseOrders"

func TestPurchaseOrdersController_Create(t *testing.T) {

	var expectedPurchaseOrders = controllers.PurchaseOrdersCreate{
		OrderNumber:     "order#1",
		OrderDate:       "2022-06-06",
		TrackingCode:    "abscf123",
		BuyerId:         1,
		ProductRecordId: 1,
		OrderStatusId:   1,
	}

	orderDateNow, _ := time.Parse("2006-01-02", expectedPurchaseOrders.OrderDate)

	var bodyPurchaseOrders = &domain.PurchaseOrders{
		OrderNumber:     "order#1",
		OrderDate:       orderDateNow,
		TrackingCode:    "abscf123",
		BuyerId:         1,
		ProductRecordId: 1,
		OrderStatusId:   1,
	}
	service := mocks.NewPurchaseOrdersService(t)
	controller := controllers.NewPurchaseOrdersController(service)
	r := testutil.SetUpRouter()
	r.POST(EndpointPurchaseOrders, controller.Create())

	t.Run("create_ok: when data entry is successful, should return code 201.", func(t *testing.T) {
		service.On("Create",
			context.TODO(),
			bodyPurchaseOrders.OrderNumber,
			bodyPurchaseOrders.OrderDate,
			bodyPurchaseOrders.TrackingCode,
			bodyPurchaseOrders.BuyerId,
			bodyPurchaseOrders.ProductRecordId,
			bodyPurchaseOrders.OrderStatusId,
		).Return(bodyPurchaseOrders, nil).Once()

		requestBody, _ := json.Marshal(expectedPurchaseOrders)
		response := testutil.ExecuteTestRequest(r, http.MethodPost, EndpointPurchaseOrders, requestBody)

		assert.Equal(t, http.StatusCreated, response.Code)
	})

	t.Run("create_fail_invalid_json: when the JSON does not contain the required fields, should return code 422", func(t *testing.T) {

		requestBody, _ := json.Marshal([]byte{})
		response := testutil.ExecuteTestRequest(r, http.MethodPost, EndpointPurchaseOrders, requestBody)

		assert.Equal(t, http.StatusUnprocessableEntity, response.Code)

	})

	t.Run("create_not_found: when the buyer id is not found, should return code 409", func(t *testing.T) {

		service.On("Create",
			context.TODO(),
			bodyPurchaseOrders.OrderNumber,
			bodyPurchaseOrders.OrderDate,
			bodyPurchaseOrders.TrackingCode,
			bodyPurchaseOrders.BuyerId,
			bodyPurchaseOrders.ProductRecordId,
			bodyPurchaseOrders.OrderStatusId,
		).Return(nil, buyerDomain.ErrIDNotFound).Once()

		requestBody, _ := json.Marshal(expectedPurchaseOrders)
		response := testutil.ExecuteTestRequest(r, http.MethodPost, EndpointPurchaseOrders, requestBody)

		assert.Equal(t, http.StatusConflict, response.Code)

	})

	t.Run("create_fail_invalid_time_parse: when the request body is invalid, should return code 400", func(t *testing.T) {

		invalidDatePurchaseOrders := controllers.PurchaseOrdersCreate{
			OrderNumber:     "order#1",
			OrderDate:       "invalid date",
			TrackingCode:    "abscf123",
			BuyerId:         1,
			ProductRecordId: 1,
			OrderStatusId:   1,
		}

		requestBody, _ := json.Marshal(invalidDatePurchaseOrders)
		response := testutil.ExecuteTestRequest(r, http.MethodPost, EndpointPurchaseOrders, requestBody)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
}
