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
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/warehouse"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/warehouse/mocks"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

const ENDPOINT = "/api/v1/warehouses"

var listPossiblesWarehouses = []warehouse.WarehouseModel{
	{
		Id:                 0,
		Address:            "Avenida Teste",
		Telephone:          "31 999999999",
		WarehouseCode:      "AZADAS30",
		MinimunCapacity:    10,
		MinimunTemperature: 9,
	},
	{
		Id:                 1,
		Address:            "Avenida Teste Segunda",
		Telephone:          "31 77777777",
		WarehouseCode:      "od78",
		MinimunCapacity:    5555555,
		MinimunTemperature: 444444,
	},
}

func Test_Controller_Warehouse_CreateWarehouse(t *testing.T) {

	body := controllers.RequestWarehousePost{
		Address:            "Avenida Teste",
		Telephone:          "31 999999999",
		WarehouseCode:      "AZADAS30",
		MinimunCapacity:    10,
		MinimunTemperature: 9,
	}

	t.Run("create_ok: testar se a criação foi com sucessida", func(t *testing.T) {

		service := mocks.NewService(t)

		service.On("Create",
			listPossiblesWarehouses[0].Address,
			listPossiblesWarehouses[0].Telephone,
			listPossiblesWarehouses[0].WarehouseCode,
			listPossiblesWarehouses[0].MinimunTemperature,
			listPossiblesWarehouses[0].MinimunCapacity).Return(listPossiblesWarehouses[0], nil)

		controller := controllers.NewWarehouse(service)

		requestBody, _ := json.Marshal(body)

		r := SetUpRouter()

		r.POST(ENDPOINT, controller.CreateWarehouse())

		response := CreateRequestTest(r, http.MethodPost, ENDPOINT, requestBody)

		assert.Equal(t, http.StatusCreated, response.Code)
	})

	t.Run("create_fail: retornar um erro 409, porque ja existe uma warehouse com o codigo", func(t *testing.T) {

		service := mocks.NewService(t)

		errMsg := fmt.Errorf("error: already a warehouse with the code: %s", body.WarehouseCode)

		service.On("Create",
			listPossiblesWarehouses[0].Address,
			listPossiblesWarehouses[0].Telephone,
			listPossiblesWarehouses[0].WarehouseCode,
			listPossiblesWarehouses[0].MinimunTemperature,
			listPossiblesWarehouses[0].MinimunCapacity).Return(warehouse.WarehouseModel{}, errMsg)

		controller := controllers.NewWarehouse(service)

		requestBody, _ := json.Marshal(body)

		r := SetUpRouter()

		r.POST(ENDPOINT, controller.CreateWarehouse())

		response := CreateRequestTest(r, http.MethodPost, ENDPOINT, requestBody)

		assert.Equal(t, http.StatusConflict, response.Code)
	})

	t.Run("create_fail: quando o objeto JSON não contiver os campos necessários, um código 422 será retornado", func(t *testing.T) {

		controller := controllers.NewWarehouse(nil)

		r := SetUpRouter()

		r.POST(ENDPOINT, controller.CreateWarehouse())

		response := CreateRequestTest(r, http.MethodPost, ENDPOINT, []byte{})

		assert.Equal(t, http.StatusUnprocessableEntity, response.Code)
	})
}

func Test_Controller_Warehouse_GetAllWarehouse(t *testing.T) {
	t.Run("create_ok: testar se a criação foi com sucessida", func(t *testing.T) {

		service := mocks.NewService(t)

		service.On("GetAll").Return(listPossiblesWarehouses, nil)

		controller := controllers.NewWarehouse(service)

		r := SetUpRouter()

		r.GET(ENDPOINT, controller.GetAllWarehouse())

		response := CreateRequestTest(r, http.MethodGet, ENDPOINT, []byte{})

		assert.Equal(t, http.StatusOK, response.Code)
	})
}

func Test_Controller_Warehouse_GetByID(t *testing.T) {

	t.Run("find_by_id_non_existent: Quando o armazém não existe, um código 404 será devolvido", func(t *testing.T) {

		var id int64 = 99999
		url := fmt.Sprintf("%s/%d", ENDPOINT, id)
		errMsg := fmt.Errorf("erros: no warehouse was found with id %d", id)

		service := mocks.NewService(t)
		service.On("GetById", int64(id)).Return(warehouse.WarehouseModel{}, errMsg)
		controller := controllers.NewWarehouse(service)

		r := SetUpRouter()
		r.GET(ENDPOINT+"/:id", controller.GetWarehouseByID())
		response := CreateRequestTest(r, http.MethodGet, url, []byte{})

		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	t.Run("find_by_id_existent: Quando a solicitação for bem-sucedida, o back-end retornará as informações solicitadas do armazém", func(t *testing.T) {
		service := mocks.NewService(t)
		service.On("GetById", int64(1)).Return(listPossiblesWarehouses[1], nil)
		controller := controllers.NewWarehouse(service)

		r := SetUpRouter()
		r.GET(ENDPOINT+"/:id", controller.GetWarehouseByID())
		response := CreateRequestTest(r, http.MethodGet, ENDPOINT+"/1", []byte{})

		assert.Equal(t, http.StatusOK, response.Code)
	})

}

func Test_Controller_Warehouse_Update(t *testing.T) {
	body := controllers.RequestWarehousePatch{
		MinimunCapacity:    66,
		MinimunTemperature: 999.0,
	}
	expectWarehouse := warehouse.WarehouseModel{
		Id:                 0,
		Address:            "Avenida Teste",
		Telephone:          "31 999999999",
		WarehouseCode:      "AZADAS30",
		MinimunCapacity:    66,
		MinimunTemperature: 999.0,
	}

	t.Run("update_ok: retornar o warehouses atualizados com o junto do codigo 200", func(t *testing.T) {

		var id int64 = 1
		url := fmt.Sprintf("%s/%d", ENDPOINT, id)

		service := mocks.NewService(t)
		service.On("UpdateTempAndCap", int64(id), 999.0, int64(66)).Return(expectWarehouse, nil)

		controller := controllers.NewWarehouse(service)

		requestBody, _ := json.Marshal(body)

		r := SetUpRouter()

		r.PATCH(ENDPOINT+"/:id", controller.UpdateWarehouse())

		response := CreateRequestTest(r, http.MethodPatch, url, requestBody)

		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("update_non_existent: Se o armazém a ser atualizado não existir, será retornado um código 404", func(t *testing.T) {
		var id int64 = 9999
		url := fmt.Sprintf("%s/%d", ENDPOINT, id)
		errMsg := fmt.Errorf("erros: no warehouse was found with id %d", id)

		service := mocks.NewService(t)
		service.On("UpdateTempAndCap", int64(id), 999.0, int64(66)).Return(expectWarehouse, errMsg)

		controller := controllers.NewWarehouse(service)

		requestBody, _ := json.Marshal(body)

		r := SetUpRouter()

		r.PATCH(ENDPOINT+"/:id", controller.UpdateWarehouse())

		response := CreateRequestTest(r, http.MethodPatch, url, requestBody)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	t.Run("update_non_id: Se não foi passado um ID, retonar um código 422", func(t *testing.T) {
		url := fmt.Sprintf("%s/abc", ENDPOINT)
		controller := controllers.NewWarehouse(nil)

		r := SetUpRouter()

		r.PATCH(ENDPOINT+"/:id", controller.UpdateWarehouse())

		response := CreateRequestTest(r, http.MethodPatch, url, []byte{})

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})

	t.Run("update_non_existent: Se o armazém a ser atualizado não existir, será retornado um código 404", func(t *testing.T) {
		var id int64 = 1
		url := fmt.Sprintf("%s/%d", ENDPOINT, id)

		controller := controllers.NewWarehouse(nil)

		requestBody, _ := json.Marshal("{\"minimun_capacity\":\"abc\",\"minimun_temperature\":\"abc2\"}")

		r := SetUpRouter()

		r.PATCH(ENDPOINT+"/:id", controller.UpdateWarehouse())

		response := CreateRequestTest(r, http.MethodPatch, url, requestBody)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})
}
func CreateRequestTest(gin *gin.Engine, method string, url string, body []byte) *httptest.ResponseRecorder {

	request := httptest.NewRequest(method, url, bytes.NewBuffer(body))

	response := httptest.NewRecorder()

	gin.ServeHTTP(response, request)

	return response
}
