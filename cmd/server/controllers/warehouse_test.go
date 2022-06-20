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

func Test_Controller_Warehouse_CreateWarehouse_Ok(t *testing.T) {

	ENDPOINT := "/api/v1/warehouses"

	expectedWh := warehouse.WarehouseModel{
		Id:                 0,
		Address:            "Avenida Teste",
		Telephone:          "31 999999999",
		WarehouseCode:      "AZADAS30",
		MinimunCapacity:    10,
		MinimunTemperature: 9,
	}

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
			expectedWh.Address,
			expectedWh.Telephone,
			expectedWh.WarehouseCode,
			expectedWh.MinimunTemperature,
			expectedWh.MinimunCapacity).Return(expectedWh, nil)

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
			expectedWh.Address,
			expectedWh.Telephone,
			expectedWh.WarehouseCode,
			expectedWh.MinimunTemperature,
			expectedWh.MinimunCapacity).Return(warehouse.WarehouseModel{}, errMsg)

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

func CreateRequestTest(gin *gin.Engine, method string, url string, body []byte) *httptest.ResponseRecorder {

	request := httptest.NewRequest(method, url, bytes.NewBuffer(body))

	response := httptest.NewRecorder()

	gin.ServeHTTP(response, request)

	return response
}
