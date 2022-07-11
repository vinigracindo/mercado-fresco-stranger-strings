package controllers_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	controllers "github.com/vinigracindo/mercado-fresco-stranger-strings/cmd/server/controllers/buyer"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/buyer/domain"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/buyer/domain/mocks"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/pkg/testutil"
)

var EndpointBuyer = "/api/v1/buyers"
var EndPointPurchaserOrders = "/api/buyers/purchaseOrders"

var expectBuyer = &domain.Buyer{
	Id:           0,
	CardNumberId: "402323",
	FirstName:    "FirstNameTest",
	LastName:     "LastNameTest",
}

var bodyBuyer = &domain.Buyer{
	CardNumberId: "402323",
	FirstName:    "FirstNameTest",
	LastName:     "LastNameTest",
}

var ctx = context.Background()

func TestBuyerController_Create(t *testing.T) {
	service := mocks.NewBuyerService(t)

	t.Run("create_ok: when data entry is successful, should return code 201.", func(t *testing.T) {

		service.
			On("Create",
				ctx,
				expectBuyer.CardNumberId,
				expectBuyer.FirstName,
				expectBuyer.LastName).
			Return(expectBuyer, nil).
			Once()

		controller := controllers.NewBuyerController(service)
		requestBody, _ := json.Marshal(bodyBuyer)

		r := testutil.SetUpRouter()
		r.POST(EndpointBuyer, controller.Create())
		response := testutil.ExecuteTestRequest(r, http.MethodPost, EndpointBuyer, requestBody)

		assert.Equal(t, http.StatusCreated, response.Code)

	})

	t.Run("create_fail: when the JSON does not contain the required fields, should return code 422", func(t *testing.T) {

		service := mocks.NewBuyerService(t)
		controller := controllers.NewBuyerController(service)

		r := testutil.SetUpRouter()
		r.POST(EndpointBuyer, controller.Create())
		response := testutil.ExecuteTestRequest(r, http.MethodPost, EndpointBuyer, []byte{})

		assert.Equal(t, http.StatusUnprocessableEntity, response.Code)
	})

	t.Run("create_conflict: when the card_number already exists, should return code 409.", func(t *testing.T) {

		service.
			On("Create",
				ctx,
				expectBuyer.CardNumberId,
				expectBuyer.FirstName,
				expectBuyer.LastName).
			Return(nil, fmt.Errorf("buyer already registered %s", expectBuyer.CardNumberId)).
			Once()

		controller := controllers.NewBuyerController(service)
		requestBody, _ := json.Marshal(bodyBuyer)

		r := testutil.SetUpRouter()
		r.POST(EndpointBuyer, controller.Create())
		response := testutil.ExecuteTestRequest(r, http.MethodPost, EndpointBuyer, requestBody)

		assert.Equal(t, http.StatusConflict, response.Code)

	})
}

func TestBuyerController_GetAll(t *testing.T) {
	service := mocks.NewBuyerService(t)

	t.Run("find_all: when data entry is successful, should return code 200.", func(t *testing.T) {

		service.
			On("GetAll", ctx).
			Return(&[]domain.Buyer{*expectBuyer}, nil).
			Once()

		controller := controllers.NewBuyerController(service)
		requestBody, _ := json.Marshal(bodyBuyer)

		r := testutil.SetUpRouter()
		r.GET(EndpointBuyer, controller.GetAll())
		response := testutil.ExecuteTestRequest(r, http.MethodGet, EndpointBuyer, requestBody)

		assert.Equal(t, http.StatusOK, response.Code)

		assert.JSONEq(t, "{\"data\":[{\"id\":0,\"card_number_id\":\"402323\",\"first_name\":\"FirstNameTest\",\"last_name\":\"LastNameTest\"}]}", response.Body.String())
	})

	t.Run("find_all_fail: when GetAll fail, should return code 400.", func(t *testing.T) {

		service.
			On("GetAll", ctx).
			Return(&[]domain.Buyer{}, fmt.Errorf("error")).
			Once()

		controller := controllers.NewBuyerController(service)
		requestBody, _ := json.Marshal(bodyBuyer)

		r := testutil.SetUpRouter()
		r.GET(EndpointBuyer, controller.GetAll())
		response := testutil.ExecuteTestRequest(r, http.MethodGet, EndpointBuyer, requestBody)

		assert.Equal(t, http.StatusBadRequest, response.Code)

	})
}

func TestBuyerController_GetById(t *testing.T) {
	service := mocks.NewBuyerService(t)

	t.Run("find_by_id_existent: when the request is successful, should return code 200", func(t *testing.T) {

		service.
			On("GetId", ctx, int64(1)).
			Return(bodyBuyer, nil).
			Once()

		controller := controllers.NewBuyerController(service)
		requestBody, _ := json.Marshal(bodyBuyer)

		r := testutil.SetUpRouter()
		r.GET(EndpointBuyer+"/:id", controller.GetId())
		response := testutil.ExecuteTestRequest(r, http.MethodGet, EndpointBuyer+"/1", requestBody)
		assert.Equal(t, http.StatusOK, response.Code)
		assert.JSONEq(t, "{\"data\":{\"id\":0,\"card_number_id\":\"402323\",\"first_name\":\"FirstNameTest\",\"last_name\":\"LastNameTest\"}}", response.Body.String())

	})

	t.Run("find_by_id_inexistent: when the buyer does not exist, should return code 404", func(t *testing.T) {

		service.
			On("GetId", ctx, int64(1)).
			Return(nil, fmt.Errorf("buyer not found")).
			Once()
		controller := controllers.NewBuyerController(service)

		r := testutil.SetUpRouter()
		r.GET(EndpointBuyer+"/:id", controller.GetId())
		response := testutil.ExecuteTestRequest(r, http.MethodGet, EndpointBuyer+"/1", []byte{})
		assert.Equal(t, http.StatusNotFound, response.Code)

	})

	t.Run("find_by_id_parse_error: when buyer id is not parsed, should return code 400.", func(t *testing.T) {
		controller := controllers.NewBuyerController(service)

		r := testutil.SetUpRouter()
		r.GET(EndpointBuyer+"/:id", controller.GetId())
		response := testutil.ExecuteTestRequest(r, http.MethodGet, EndpointBuyer+"/idInvalido", []byte{})

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
}

func TestBuyerController_Update(t *testing.T) {
	service := mocks.NewBuyerService(t)

	updateBody := &domain.Buyer{
		Id:           1,
		CardNumberId: "402324",
		LastName:     "LastNameTest 2",
	}

	t.Run("update_ok: when the request is successful, should return code 200", func(t *testing.T) {

		service.
			On("Update", ctx, int64(1), updateBody.CardNumberId, updateBody.LastName).
			Return(updateBody, nil).
			Once()

		controller := controllers.NewBuyerController(service)
		requestBody, _ := json.Marshal(updateBody)

		r := testutil.SetUpRouter()
		r.PATCH(EndpointBuyer+"/:id", controller.UpdateCardNumberLastName())
		response := testutil.ExecuteTestRequest(r, http.MethodPatch, EndpointBuyer+"/1", requestBody)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.JSONEq(t, "{\"data\":{\"id\":1,\"card_number_id\":\"402324\",\"first_name\":\"\",\"last_name\":\"LastNameTest 2\"}}", response.Body.String())
	})

	t.Run("update_non_existent: when the buyer does not exist, should return code 404.", func(t *testing.T) {
		service.
			On("Update", ctx, int64(1), updateBody.CardNumberId, updateBody.LastName).
			Return(nil, fmt.Errorf("buyer with id %d not found", int64(1))).
			Once()

		controller := controllers.NewBuyerController(service)
		requestBody, _ := json.Marshal(updateBody)

		r := testutil.SetUpRouter()
		r.PATCH("/api/v1/buyers/:id", controller.UpdateCardNumberLastName())
		response := testutil.ExecuteTestRequest(r, http.MethodPatch, "/api/v1/buyers/1", requestBody)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	t.Run("update_id_parse_error: when buyer id is not parsed, should return code 400.", func(t *testing.T) {
		controller := controllers.NewBuyerController(service)
		r := testutil.SetUpRouter()
		r.PATCH(EndpointBuyer+"/:id", controller.UpdateCardNumberLastName())
		response := testutil.ExecuteTestRequest(r, http.MethodPatch, EndpointBuyer+"/idInvalido", []byte{})

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})

	t.Run("update_body_invalid: when the request body is not valid json, should return code 400.", func(t *testing.T) {
		controller := controllers.NewBuyerController(service)

		r := testutil.SetUpRouter()
		r.PATCH(EndpointBuyer+"/:id", controller.UpdateCardNumberLastName())
		response := testutil.ExecuteTestRequest(r, http.MethodPatch, EndpointBuyer+"/1", nil)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
}
func TestBuyerController_Delete(t *testing.T) {
	service := mocks.NewBuyerService(t)

	t.Run("delete_non_existent: when the buyer does not exist, should return code 404", func(t *testing.T) {

		service.
			On("Delete", ctx, int64(1)).
			Return(fmt.Errorf("buyer with id not found")).
			Once()

		controller := controllers.NewBuyerController(service)

		r := testutil.SetUpRouter()

		r.DELETE(EndpointBuyer+"/:id", controller.DeleteBuyer())
		response := testutil.ExecuteTestRequest(r, http.MethodDelete, EndpointBuyer+"/1", []byte{})

		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	t.Run("delete_ok: when the request is successful, should return code 204.", func(t *testing.T) {

		service.
			On("Delete", ctx, int64(1)).
			Return(nil).
			Once()

		controller := controllers.NewBuyerController(service)

		r := testutil.SetUpRouter()
		r.DELETE(EndpointBuyer+"/:id", controller.DeleteBuyer())
		response := testutil.ExecuteTestRequest(r, http.MethodDelete, EndpointBuyer+"/1", []byte{})

		assert.Equal(t, http.StatusNoContent, response.Code)
	})

	t.Run("delete_id_parse_error: when buyer id is not parsed, should return code 400", func(t *testing.T) {

		controller := controllers.NewBuyerController(service)
		r := testutil.SetUpRouter()
		r.DELETE(EndpointBuyer+"/:id", controller.DeleteBuyer())

		response := testutil.ExecuteTestRequest(r, http.MethodDelete, EndpointBuyer+"/idInvalido", []byte{})

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
}

func TestBuyerController_GetPurchaseOrdersReports(t *testing.T) {
	service := mocks.NewBuyerService(t)
	controller := controllers.NewBuyerController(service)
	r := testutil.SetUpRouter()
	r.GET(EndPointPurchaserOrders, controller.GetPurchaseOrdersReports())

	expectResult := domain.PurchaseOrdersReport{
		Id:                 1,
		CardNumberId:       "40232212",
		FirstName:          "FirstNameTest",
		LastName:           "LastNameTest",
		CountBuyersRecords: 2,
	}

	expectBodyRecord := domain.PurchaseOrdersReport{
		Id:                 1,
		CardNumberId:       "40232212",
		FirstName:          "FirstNameTest",
		LastName:           "LastNameTest",
		CountBuyersRecords: 2,
	}

	var buyerId = &domain.Buyer{
		Id:           1,
		CardNumberId: "402323",
		FirstName:    "FirstNameTest",
		LastName:     "LastNameTest",
	}

	t.Run("GetId: Get purchase order reports", func(t *testing.T) {

		expectResult := []domain.PurchaseOrdersReport{expectResult}
		bodyList := []domain.PurchaseOrdersReport{expectBodyRecord}
		requestBody, _ := json.Marshal(bodyList)

		service.On("GetPurchaseOrdersReports", context.TODO(), buyerId.Id).
			Return(&expectResult, nil).
			Once()

		response := testutil.ExecuteTestRequest(r, http.MethodGet, EndPointPurchaserOrders+"?id=1", requestBody)

		assert.Equal(t, http.StatusOK, response.Code)

	})

	t.Run("GetId_err: purchase order reports by id", func(t *testing.T) {

		service.On("GetPurchaseOrdersReports", context.TODO(), buyerId.Id).
			Return(nil, domain.ErrIDNotFound).
			Once()

		response := testutil.ExecuteTestRequest(r, http.MethodGet, EndPointPurchaserOrders+"?id=1", []byte{})

		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	t.Run("invalidQuery: when the query params are not valid, should return code 400.", func(t *testing.T) {
		response := testutil.ExecuteTestRequest(r, http.MethodGet, EndPointPurchaserOrders+"?id=asd", []byte{})

		assert.Equal(t, http.StatusBadRequest, response.Code)

	})

	t.Run("GetAll: when the request is successful, should return code 200", func(t *testing.T) {

		service.On("GetAllPurchaseOrdersReports", context.TODO()).
			Return(&[]domain.PurchaseOrdersReport{expectResult}, nil).
			Once()

		requestBody, _ := json.Marshal(expectBodyRecord)

		response := testutil.ExecuteTestRequest(r, http.MethodGet, EndPointPurchaserOrders, requestBody)
		assert.Equal(t, http.StatusOK, response.Code)

	})

	t.Run("GetAll_err: when GetAll fail, should return code 500.", func(t *testing.T) {

		service.
			On("GetAllPurchaseOrdersReports", ctx).
			Return(&[]domain.PurchaseOrdersReport{}, fmt.Errorf("error")).
			Once()

		requestBody, _ := json.Marshal(expectBodyRecord)

		response := testutil.ExecuteTestRequest(r, http.MethodGet, EndPointPurchaserOrders, requestBody)
		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})

}
