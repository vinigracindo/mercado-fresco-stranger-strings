package controllers_test

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	controllers "github.com/vinigracindo/mercado-fresco-stranger-strings/cmd/server/controllers/purchase_orders"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/purchase_orders/domain"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/purchase_orders/domain/mocks"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/pkg/testutil"
)

var EndpointPurchaseOrders = "/api/v1/purchaseOrders"
var ctx = context.Background()
var orderDateNow = time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)

var expectedPurchaseOrders = &domain.PurchaseOrders{
	Id:              1,
	OrderNumber:     "order#1",
	OrderDate:       orderDateNow,
	TrackingCode:    "abscf123",
	BuyerId:         1,
	ProductRecordId: 1,
	OrderStatusId:   1,
}

var bodyPurchaseOrders = &domain.PurchaseOrders{
	OrderNumber:     "order#1",
	OrderDate:       orderDateNow,
	TrackingCode:    "abscf123",
	BuyerId:         1,
	ProductRecordId: 1,
	OrderStatusId:   1,
}

func TestPurchaseOrdersController_Create(t *testing.T) {
	service := mocks.NewPurchaseOrdersService(t)

	t.Run("create_ok: when data entry is successful, should return code 201.", func(t *testing.T) {
		service.On("Create",
			ctx,
			expectedPurchaseOrders.OrderNumber,
			expectedPurchaseOrders.OrderDate,
			expectedPurchaseOrders.TrackingCode,
			expectedPurchaseOrders.BuyerId,
			expectedPurchaseOrders.ProductRecordId,
			expectedPurchaseOrders.OrderStatusId,
		).Return(expectedPurchaseOrders, nil).Once()

		controller := controllers.NewPurchaseOrdersController(service)
		requestBody, _ := json.Marshal(bodyPurchaseOrders)

		r := testutil.SetUpRouter()
		r.POST(EndpointPurchaseOrders, controller.Create())
		response := testutil.ExecuteTestRequest(r, http.MethodPost, EndpointPurchaseOrders, requestBody)

		assert.Equal(t, http.StatusCreated, response.Code)

	})

}
