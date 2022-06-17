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
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/warehouse"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/warehouse/mocks"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func Test_Controller_Warehouse_CreateWarehouse(t *testing.T) {

	expectedWh := warehouse.WarehouseModel{
		Id:                 0,
		Address:            "Avenida Teste",
		Telephone:          "31 999999999",
		WarehouseCode:      "30",
		MinimunCapacity:    10,
		MinimunTemperature: 9,
	}

	body := warehouse.WarehouseModel{
		Address:            "Avenida Teste",
		Telephone:          "31 999999999",
		WarehouseCode:      "30",
		MinimunCapacity:    10,
		MinimunTemperature: 9,
	}

	t.Run("create_ok", func(t *testing.T) {

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

		r.POST("/api/v1/warehouses", controller.CreateWarehouse())

		response := CreateRequestTest(r, "POST", "/api/v1/warehouses", requestBody)

		assert.Equal(t, response.Code, http.StatusCreated)
	})
}

func CreateRequestTest(gin *gin.Engine, method string, url string, body []byte) *httptest.ResponseRecorder {

	request := httptest.NewRequest(method, url, bytes.NewBuffer(body))

	response := httptest.NewRecorder()

	gin.ServeHTTP(response, request)

	return response
}
