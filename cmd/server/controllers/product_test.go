package controllers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/cmd/server/controllers"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/product"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/product/mocks"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

// simula uma requisição HTTP para a engine de rotas (gin) com o método HTTP, path e body  passados. Retorna a resposta da API.
func ExecuteTestRequest(router *gin.Engine, method string, path string, body []byte) *httptest.ResponseRecorder {

	request := httptest.NewRequest(method, path, bytes.NewBuffer(body))

	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	return response
}

func TestProductController_Create(t *testing.T) {

	mockService := mocks.NewProductService(t)

	expectedProduct := product.Product{
		Id:                             1,
		ProductCode:                    "PROD02",
		Description:                    "Yogurt",
		Width:                          1.2,
		Height:                         6.4,
		Length:                         4.5,
		NetWeight:                      3.4,
		ExpirationRate:                 1.5,
		RecommendedFreezingTemperature: 1.3,
		FreezingRate:                   2,
		ProductTypeId:                  2,
		SellerId:                       2,
	}

	body := product.Product{
		ProductCode:                    "PROD02",
		Description:                    "Yogurt",
		Width:                          1.2,
		Height:                         6.4,
		Length:                         4.5,
		NetWeight:                      3.4,
		ExpirationRate:                 1.5,
		RecommendedFreezingTemperature: 1.3,
		FreezingRate:                   2,
		ProductTypeId:                  2,
		SellerId:                       2,
	}

	t.Run("create_ok: quando a entrada de dados for bem-sucedida, um código 201 será retornado junto com o objeto inserido", func(t *testing.T) {

		// PREPARACAO
		mockService.
			On("Create", expectedProduct.ProductCode, expectedProduct.Description, expectedProduct.Width, expectedProduct.Height,
				expectedProduct.Length, expectedProduct.NetWeight, expectedProduct.ExpirationRate, expectedProduct.RecommendedFreezingTemperature,
				expectedProduct.FreezingRate, expectedProduct.ProductTypeId, expectedProduct.SellerId).
			Return(&expectedProduct, nil).
			Once()

		controller := controllers.CreateProductController(mockService)

		requestBody, _ := json.Marshal(body)

		router := SetUpRouter()
		router.POST("/api/v1/products", controller.Create())

		//  EXECUCAO
		response := ExecuteTestRequest(router, "POST", "/api/v1/products", requestBody)

		// VALIDACAO
		assert.Equal(t, http.StatusCreated, response.Code)
	})

	t.Run("create_fail: quando o objeto JSON não contiver os campos necessários, um código 422 será retornado.", func(t *testing.T) {
		controller := controllers.CreateProductController(nil)

		router := SetUpRouter()
		router.POST("/api/v1/products", controller.Create())

		//  EXECUCAO
		response := ExecuteTestRequest(router, "POST", "/api/v1/products", []byte{})

		// VALIDACAO
		assert.Equal(t, http.StatusUnprocessableEntity, response.Code)
		assert.Equal(t, "{\"message\":\"invalid input. Check the data entered\"}", response.Body.String())
	})

	t.Run("create_conflict: quando o product_code já existir, ele retornará um erro 409 Conflict.", func(t *testing.T) {

		// PREPARACAO

		expectedError := errors.New("the product code has already been registered")
		mockService.
			On("Create", expectedProduct.ProductCode, expectedProduct.Description, expectedProduct.Width, expectedProduct.Height,
				expectedProduct.Length, expectedProduct.NetWeight, expectedProduct.ExpirationRate, expectedProduct.RecommendedFreezingTemperature,
				expectedProduct.FreezingRate, expectedProduct.ProductTypeId, expectedProduct.SellerId).
			Return(nil, expectedError).
			Once()

		controller := controllers.CreateProductController(mockService)

		requestBody, _ := json.Marshal(body)

		// configura engine de rotas
		router := SetUpRouter()
		router.POST("/api/v1/products", controller.Create())

		//  EXECUCAO -
		response := ExecuteTestRequest(router, "POST", "/api/v1/products", requestBody)

		// VALIDACAO
		assert.Equal(t, http.StatusConflict, response.Code)
		assert.Equal(t, "{\"code\":409,\"message\":\"the product code has already been registered\"}", response.Body.String())
	})
}
