package controllers_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	controllers "github.com/vinigracindo/mercado-fresco-stranger-strings/cmd/server/controllers/warehouse"
	warehouse "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/warehouse/domain"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/warehouse/domain/mocks"

	"github.com/vinigracindo/mercado-fresco-stranger-strings/pkg/testutil"
)

const EndpointWarehouse = "/api/v1/warehouses"

var listPossiblesWarehouses = []warehouse.WarehouseModel{
	{
		Id:                 0,
		Address:            "Avenida Teste",
		Telephone:          "31 999999999",
		WarehouseCode:      "AZADAS30",
		MinimunCapacity:    10,
		MinimunTemperature: 9,
		LocalityID:         1,
	},
	{
		Id:                 1,
		Address:            "Avenida Teste Segunda",
		Telephone:          "31 77777777",
		WarehouseCode:      "od78",
		MinimunCapacity:    5555555,
		MinimunTemperature: 444444,
		LocalityID:         2,
	},
}

func Test_Controller_Warehouse_CreateWarehouse(t *testing.T) {

	body := controllers.RequestWarehousePost{
		Address:            "Avenida Teste",
		Telephone:          "31 999999999",
		WarehouseCode:      "AZADAS30",
		MinimunCapacity:    10,
		MinimunTemperature: 9,
		LocalityID:         1,
	}

	t.Run("create_ok: if warehouses was successfully created", func(t *testing.T) {

		service := mocks.NewWarehouseService(t)

		service.On("Create",
			context.TODO(),
			body.Address,
			body.Telephone,
			body.WarehouseCode,
			body.MinimunTemperature,
			body.MinimunCapacity,
			body.LocalityID).Return(listPossiblesWarehouses[0], nil)

		controller := controllers.NewWarehouse(service)

		requestBody, _ := json.Marshal(body)

		r := testutil.SetUpRouter()

		r.POST(EndpointWarehouse, controller.CreateWarehouse())

		response := testutil.ExecuteTestRequest(r, http.MethodPost, EndpointWarehouse, requestBody)

		expect := map[string]interface{}{
			"data": listPossiblesWarehouses[0],
		}

		assert.Equal(t, http.StatusCreated, response.Code)
		assert.JSONEq(t, CreateStringJSON(expect), response.Body.String())
	})

	t.Run("create_fail: return 409, because the is already an warehouse with that code", func(t *testing.T) {

		service := mocks.NewWarehouseService(t)

		errMsg := fmt.Errorf("error: already a warehouse with the code: %s", body.WarehouseCode)

		service.On("Create",
			context.TODO(),
			body.Address,
			body.Telephone,
			body.WarehouseCode,
			body.MinimunTemperature,
			body.MinimunCapacity,
			body.LocalityID).Return(warehouse.WarehouseModel{}, errMsg)

		controller := controllers.NewWarehouse(service)

		requestBody, _ := json.Marshal(body)

		r := testutil.SetUpRouter()

		r.POST(EndpointWarehouse, controller.CreateWarehouse())

		response := testutil.ExecuteTestRequest(r, http.MethodPost, EndpointWarehouse, requestBody)

		assert.Equal(t, http.StatusConflict, response.Code)
	})

	t.Run("create_fail: when json object do not have all necessary fields, return 422 code", func(t *testing.T) {

		controller := controllers.NewWarehouse(nil)

		r := testutil.SetUpRouter()

		r.POST(EndpointWarehouse, controller.CreateWarehouse())

		response := testutil.ExecuteTestRequest(r, http.MethodPost, EndpointWarehouse, []byte{})

		assert.Equal(t, http.StatusUnprocessableEntity, response.Code)
	})
}

func Test_Controller_Warehouse_GetAllWarehouse(t *testing.T) {
	t.Run("find_all: return a list with all warehouses storages", func(t *testing.T) {

		service := mocks.NewWarehouseService(t)

		service.On("GetAll", context.TODO()).Return(listPossiblesWarehouses, nil)

		controller := controllers.NewWarehouse(service)

		r := testutil.SetUpRouter()

		r.GET(EndpointWarehouse, controller.GetAllWarehouse())

		response := testutil.ExecuteTestRequest(r, http.MethodGet, EndpointWarehouse, []byte{})

		expect := map[string]interface{}{
			"data": listPossiblesWarehouses,
		}

		assert.Equal(t, http.StatusOK, response.Code)
		assert.JSONEq(t, CreateStringJSON(expect), response.Body.String())
	})

	t.Run("find_all_error: when an error ocorrency in the server", func(t *testing.T) {
		service := mocks.NewWarehouseService(t)

		service.On("GetAll", context.TODO()).Return([]warehouse.WarehouseModel{}, fmt.Errorf("error: internal error"))

		controller := controllers.NewWarehouse(service)

		r := testutil.SetUpRouter()

		r.GET(EndpointWarehouse, controller.GetAllWarehouse())

		response := testutil.ExecuteTestRequest(r, http.MethodGet, EndpointWarehouse, []byte{})

		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})
}

func Test_Controller_Warehouse_GetByID(t *testing.T) {

	t.Run("find_by_id_non_existent: if warehouse do not exist return 404 code", func(t *testing.T) {

		var id int64 = 99999
		url := fmt.Sprintf("%s/%d", EndpointWarehouse, id)
		errMsg := fmt.Errorf("erros: no warehouse was found with id %d", id)

		service := mocks.NewWarehouseService(t)
		service.On("GetById", context.TODO(), int64(id)).Return(warehouse.WarehouseModel{}, errMsg)
		controller := controllers.NewWarehouse(service)

		r := testutil.SetUpRouter()
		r.GET(EndpointWarehouse+"/:id", controller.GetWarehouseByID())
		response := testutil.ExecuteTestRequest(r, http.MethodGet, url, []byte{})

		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	t.Run("find_by_id_existent: when request was sucessufuly return an warehouse", func(t *testing.T) {
		service := mocks.NewWarehouseService(t)
		service.On("GetById", context.TODO(), int64(1)).Return(listPossiblesWarehouses[1], nil)
		controller := controllers.NewWarehouse(service)

		r := testutil.SetUpRouter()
		r.GET(EndpointWarehouse+"/:id", controller.GetWarehouseByID())
		response := testutil.ExecuteTestRequest(r, http.MethodGet, EndpointWarehouse+"/1", []byte{})

		expect := map[string]interface{}{
			"data": listPossiblesWarehouses[1],
		}

		assert.Equal(t, http.StatusOK, response.Code)
		assert.JSONEq(t, CreateStringJSON(expect), response.Body.String())
	})

	t.Run("find_by_id_non_id: if id does not exist return 422 code", func(t *testing.T) {
		url := fmt.Sprintf("%s/abc", EndpointWarehouse)
		controller := controllers.NewWarehouse(nil)

		r := testutil.SetUpRouter()

		r.GET(EndpointWarehouse+"/:id", controller.GetWarehouseByID())

		response := testutil.ExecuteTestRequest(r, http.MethodGet, url, []byte{})

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
}

func Test_Controller_Warehouse_Update(t *testing.T) {
	body := controllers.RequestWarehousePatch{
		MinimunCapacity:    66,
		MinimunTemperature: 999.0,
	}

	t.Run("update_ok: if warehouses was successfully updated return 200 code", func(t *testing.T) {

		var id int64 = 1
		url := fmt.Sprintf("%s/%d", EndpointWarehouse, id)

		service := mocks.NewWarehouseService(t)
		service.On("UpdateTempAndCap",
			context.TODO(),
			int64(id),
			999.0,
			int64(66)).Return(listPossiblesWarehouses[0], nil)

		controller := controllers.NewWarehouse(service)

		requestBody, _ := json.Marshal(body)

		r := testutil.SetUpRouter()

		r.PATCH(EndpointWarehouse+"/:id", controller.UpdateWarehouse())

		response := testutil.ExecuteTestRequest(r, http.MethodPatch, url, requestBody)

		expect := map[string]interface{}{
			"data": listPossiblesWarehouses[0],
		}

		assert.Equal(t, http.StatusOK, response.Code)
		assert.JSONEq(t, CreateStringJSON(expect), response.Body.String())
	})

	t.Run("update_non_existent: if does not find warehouses with the id, return 404 code", func(t *testing.T) {
		var id int64 = 9999
		url := fmt.Sprintf("%s/%d", EndpointWarehouse, id)
		errMsg := fmt.Errorf("erros: no warehouse was found with id %d", id)

		service := mocks.NewWarehouseService(t)
		service.On("UpdateTempAndCap",
			context.TODO(),
			int64(id),
			999.0,
			int64(66)).Return(listPossiblesWarehouses[0], errMsg)

		controller := controllers.NewWarehouse(service)

		requestBody, _ := json.Marshal(body)

		r := testutil.SetUpRouter()

		r.PATCH(EndpointWarehouse+"/:id", controller.UpdateWarehouse())

		response := testutil.ExecuteTestRequest(r, http.MethodPatch, url, requestBody)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	t.Run("update_non_id: return 422 code when id is of a invalid type", func(t *testing.T) {
		url := fmt.Sprintf("%s/abc", EndpointWarehouse)
		controller := controllers.NewWarehouse(nil)

		r := testutil.SetUpRouter()

		r.PATCH(EndpointWarehouse+"/:id", controller.UpdateWarehouse())

		response := testutil.ExecuteTestRequest(r, http.MethodPatch, url, []byte{})

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})

	t.Run("update_wrong_data_type: return 400 code when data are of the wrong type", func(t *testing.T) {
		var id int64 = 1
		url := fmt.Sprintf("%s/%d", EndpointWarehouse, id)

		controller := controllers.NewWarehouse(nil)

		requestBody, _ := json.Marshal("{\"minimun_capacity\":\"abc\",\"minimun_temperature\":\"abc2\"}")

		r := testutil.SetUpRouter()

		r.PATCH(EndpointWarehouse+"/:id", controller.UpdateWarehouse())

		response := testutil.ExecuteTestRequest(r, http.MethodPatch, url, requestBody)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
}

func Test_Controller_Warehouse_Delete(t *testing.T) {
	t.Run("delete_non_existent: return 404 code when no warehouses was found with the id", func(t *testing.T) {
		var id int64 = 1
		url := fmt.Sprintf("%s/%d", EndpointWarehouse, id)
		errMsg := fmt.Errorf("erros: no warehouse was found with id %d", id)

		service := mocks.NewWarehouseService(t)
		service.On("Delete", context.TODO(), id).Return(errMsg)

		controller := controllers.NewWarehouse(service)

		r := testutil.SetUpRouter()

		r.DELETE(EndpointWarehouse+"/:id", controller.DeleteWarehouse())

		response := testutil.ExecuteTestRequest(r, http.MethodDelete, url, []byte{})

		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	t.Run("delete_non_id: return 422 code when id is of a invalid type", func(t *testing.T) {
		url := fmt.Sprintf("%s/abc", EndpointWarehouse)
		controller := controllers.NewWarehouse(nil)

		r := testutil.SetUpRouter()

		r.DELETE(EndpointWarehouse+"/:id", controller.DeleteWarehouse())

		response := testutil.ExecuteTestRequest(r, http.MethodDelete, url, []byte{})

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})

	t.Run("delete_ok: return 204 code when an warehouses is successfully deleted", func(t *testing.T) {
		var id int64 = 1
		url := fmt.Sprintf("%s/%d", EndpointWarehouse, id)

		service := mocks.NewWarehouseService(t)
		service.On("Delete", context.TODO(), id).Return(nil)

		controller := controllers.NewWarehouse(service)

		r := testutil.SetUpRouter()

		r.DELETE(EndpointWarehouse+"/:id", controller.DeleteWarehouse())

		response := testutil.ExecuteTestRequest(r, http.MethodDelete, url, []byte{})

		assert.Equal(t, http.StatusNoContent, response.Code)
	})
}

func CreateStringJSON(obj interface{}) string {
	jsonObj, _ := json.Marshal(obj)
	return string(jsonObj)
}
