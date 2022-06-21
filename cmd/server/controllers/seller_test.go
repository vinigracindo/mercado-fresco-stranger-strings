package controllers_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/cmd/server/controllers"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/seller"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/seller/mocks"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/pkg/testutil"
)

const EndpointSeller = "/api/v1/sellers"

var bodySeller = seller.Seller{
	Cid:         123,
	CompanyName: "Mercado Livre",
	Address:     "Osasco, SP",
	Telephone:   "11 99999999",
}

var bodySellerFail = seller.Seller{
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

func TestSellerController_Create(t *testing.T) {
	service := mocks.NewService(t)

	t.Run("create_ok: when data entry is successful, should return code 201.", func(t *testing.T) {
		service.On("Create", expectedSeller.Cid, expectedSeller.CompanyName, expectedSeller.Address, expectedSeller.Telephone).
			Return(expectedSeller, nil).
			Once()

		controller := controllers.NewSeller(service)

		requestbodySeller, _ := json.Marshal(bodySeller)

		r := testutil.SetUpRouter()

		r.POST(EndpointSeller, controller.Create())

		response := testutil.ExecuteTestRequest(r, http.MethodPost, EndpointSeller, requestbodySeller)

		assert.Equal(t, http.StatusCreated, response.Code)
	})

	t.Run("create_fail: when the JSON does not contain the required fields, should return code 422", func(t *testing.T) {
		controller := controllers.NewSeller(nil)

		r := testutil.SetUpRouter()
		r.POST(EndpointSeller, controller.Create())
		response := testutil.ExecuteTestRequest(r, http.MethodPost, EndpointSeller, []byte{})

		assert.Equal(t, http.StatusUnprocessableEntity, response.Code)
	})

	t.Run("create_conflict: when the cid already exists, should return code 409.", func(t *testing.T) {
		service.On("Create", expectedSeller.Cid, expectedSeller.CompanyName, expectedSeller.Address, expectedSeller.Telephone).
			Return(seller.Seller{}, fmt.Errorf("Seller with this cid alredy exists")).
			Once()

		controller := controllers.NewSeller(service)
		requestbodySeller, _ := json.Marshal(bodySeller)
		r := testutil.SetUpRouter()
		r.POST(EndpointSeller, controller.Create())

		response := testutil.ExecuteTestRequest(r, http.MethodPost, EndpointSeller, requestbodySeller)

		assert.Equal(t, http.StatusConflict, response.Code)

	})
}

func TestSellerController_Get(t *testing.T) {
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
	t.Run("find_all: when data entry is successful, should return code 200", func(t *testing.T) {
		service.On("GetAll").Return(expectedListSeller, nil).Once()

		controller := controllers.NewSeller(service)
		r := testutil.SetUpRouter()
		r.GET(EndpointSeller, controller.GetAll())

		response := testutil.ExecuteTestRequest(r, http.MethodGet, EndpointSeller, []byte{})

		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("find_by_id_non_exitent: when the seller does not exist, should return code 404.", func(t *testing.T) {
		service.On("GetById", int64(9999)).Return(seller.Seller{}, fmt.Errorf("Seller not found")).Once()

		controller := controllers.NewSeller(service)
		r := testutil.SetUpRouter()
		r.GET(EndpointSeller+"/:id", controller.GetById())

		response := testutil.ExecuteTestRequest(r, http.MethodGet, EndpointSeller+"/9999", []byte{})

		assert.Equal(t, http.StatusNotFound, response.Code)

	})

	t.Run("find_by_id_exixtent: when the request is successful, should return code 200", func(t *testing.T) {
		service.On("GetById", int64(1)).Return(expectedListSeller[0], nil).Once()

		controller := controllers.NewSeller(service)
		requestbodySeller, _ := json.Marshal(bodySeller)
		r := testutil.SetUpRouter()
		r.GET(EndpointSeller+"/:id", controller.GetById())

		response := testutil.ExecuteTestRequest(r, http.MethodGet, EndpointSeller+"/1", requestbodySeller)

		assert.Equal(t, http.StatusOK, response.Code)
	})
}

func Test_Controller_Update(t *testing.T) {
	service := mocks.NewService(t)
	var bodySellerUpdate = seller.Seller{
		Address:   "Salvador, BA",
		Telephone: "71 88888888",
	}

	t.Run("update_ok: when the request is successful, should return code 200. The object must be returned.", func(t *testing.T) {
		service.On("Update", int64(1), "Salvador, BA", "71 88888888").
			Return(expectedSeller, nil).
			Once()

		controller := controllers.NewSeller(service)
		requestbodySeller, _ := json.Marshal(bodySellerUpdate)
		r := testutil.SetUpRouter()
		r.PATCH(EndpointSeller+"/:id", controller.Update())

		response := testutil.ExecuteTestRequest(r, http.MethodPatch, EndpointSeller+"/1", requestbodySeller)

		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("update_non_existent: when the employee does not exist, should return code 404.", func(t *testing.T) {
		service.On("Update", int64(9999), "Salvador, BA", "71 88888888").
			Return(seller.Seller{}, fmt.Errorf("Seller not found")).
			Once()

		controller := controllers.NewSeller(service)
		requestbodySeller, _ := json.Marshal(bodySellerUpdate)
		r := testutil.SetUpRouter()
		r.PATCH(EndpointSeller+"/:id", controller.Update())

		response := testutil.ExecuteTestRequest(r, http.MethodPatch, EndpointSeller+"/9999", requestbodySeller)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})
}

func TestSellerController_Delete(t *testing.T) {
	service := mocks.NewService(t)

	t.Run("delete_non_existent: when the seller does not exist, should return code 404.", func(t *testing.T) {
		service.On("Delete", int64(9999)).
			Return(fmt.Errorf("Seller not found")).
			Once()

		controller := controllers.NewSeller(service)
		r := testutil.SetUpRouter()

		r.DELETE(EndpointSeller+"/:id", controller.Delete())

		response := testutil.ExecuteTestRequest(r, http.MethodDelete, EndpointSeller+"/9999", []byte{})

		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	t.Run("delete_ok: when the request is successful, should return code 204.", func(t *testing.T) {
		service.On("Delete", int64(1)).
			Return(nil).
			Once()

		controller := controllers.NewSeller(service)
		r := testutil.SetUpRouter()
		r.DELETE(EndpointSeller+"/:id", controller.Delete())

		response := testutil.ExecuteTestRequest(r, http.MethodDelete, EndpointSeller+"/1", []byte{})

		assert.Equal(t, http.StatusNoContent, response.Code)

	})
}
