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

		// PREPARAÇÃO
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

		// EXECUÇÃO
		response := ExecuteTestRequest(router, "POST", "/api/v1/products", requestBody)

		// VALIDAÇÃO
		assert.Equal(t, http.StatusCreated, response.Code)
	})

	t.Run("create_fail: quando o objeto JSON não contiver os campos necessários, um código 422 será retornado.", func(t *testing.T) {
		controller := controllers.CreateProductController(nil)

		router := SetUpRouter()
		router.POST("/api/v1/products", controller.Create())

		// EXECUÇÃO
		response := ExecuteTestRequest(router, "POST", "/api/v1/products", []byte{})

		// VALIDAÇÃO
		assert.Equal(t, http.StatusUnprocessableEntity, response.Code)
		assert.Equal(t, "{\"message\":\"invalid input. Check the data entered\"}", response.Body.String())
	})

	t.Run("create_conflict: quando o product_code já existir, ele retornará um erro 409 Conflict.", func(t *testing.T) {

		// PREPARAÇÃO
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

		// EXECUÇÃO
		response := ExecuteTestRequest(router, "POST", "/api/v1/products", requestBody)

		// VALIDAÇÃO
		assert.Equal(t, http.StatusConflict, response.Code)
		assert.Equal(t, "{\"code\":409,\"message\":\"the product code has already been registered\"}", response.Body.String())
	})
}

func TestProductController_GetAll(t *testing.T) {

	mockService := mocks.NewProductService(t)

	expectedProduct := []product.Product{
		{
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
		},
		{
			Id:                             2,
			ProductCode:                    "PROD03",
			Description:                    "Yogurt light",
			Width:                          1.5,
			Height:                         5.4,
			Length:                         3.5,
			NetWeight:                      4.4,
			ExpirationRate:                 1.8,
			RecommendedFreezingTemperature: 1.2,
			FreezingRate:                   2,
			ProductTypeId:                  3,
			SellerId:                       3,
		},
	}

	body := []product.Product{
		{
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
		},
		{
			ProductCode:                    "PROD03",
			Description:                    "Yogurt light",
			Width:                          1.5,
			Height:                         5.4,
			Length:                         3.5,
			NetWeight:                      4.4,
			ExpirationRate:                 1.8,
			RecommendedFreezingTemperature: 1.2,
			FreezingRate:                   2,
			ProductTypeId:                  3,
			SellerId:                       3,
		},
	}

	t.Run("find_all_bad_request: quando a solicitação não for bem-sucedida, o back-end retornará um erro 400 BadRequest.", func(t *testing.T) {

		// PREPARAÇÃO
		expectedError := errors.New("the request sent to the server is invalid or corrupted")
		mockService.
			On("GetAll").
			Return(nil, expectedError).
			Once()

		controller := controllers.CreateProductController(mockService)

		requestBody, _ := json.Marshal(body)

		// configura engine de rotas
		router := SetUpRouter()
		router.GET("/api/v1/products", controller.GetAll())

		// EXECUÇÃO
		response := ExecuteTestRequest(router, "GET", "/api/v1/products", requestBody)

		// VALIDAÇÃO
		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.Equal(t, "{\"code\":400,\"message\":\"the request sent to the server is invalid or corrupted\"}", response.Body.String())
	})

	t.Run("find_all: quando a solicitação for bem-sucedida, o back-end retornará uma lista de todos os produtos existentes.", func(t *testing.T) {

		// PREPARACÃO
		mockService.
			On("GetAll").
			Return(expectedProduct, nil).
			Once()

		controller := controllers.CreateProductController(mockService)

		requestBody, _ := json.Marshal(body)

		router := SetUpRouter()
		router.GET("/api/v1/products", controller.GetAll())

		//  EXECUCAO
		response := ExecuteTestRequest(router, "GET", "/api/v1/products", requestBody)

		// VALIDACAO
		assert.Equal(t, http.StatusOK, response.Code)
	})
}

func TestProductController_GetById(t *testing.T) {

	mockService := mocks.NewProductService(t)

	t.Run("find_by_id_parse_error: quando o id do produto não for parseado", func(t *testing.T) {

		controller := controllers.CreateProductController(mockService)

		// configura engine de rotas
		router := SetUpRouter()
		router.GET("/api/v1/products/:id", controller.GetById())

		// EXECUÇÃO
		response := ExecuteTestRequest(router, "GET", "/api/v1/products/abc", []byte{})

		// VALIDAÇÃO
		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.Equal(t, "{\"code\":400,\"message\":\"strconv.ParseInt: parsing \\\"abc\\\": invalid syntax\"}", response.Body.String())
	})

	t.Run("find_by_id_non_existent: quando o produto não existe, um código 404 será devolvido", func(t *testing.T) {

		// PREPARAÇÃO
		expectedError := errors.New("the product id was not found")
		mockService.
			On("GetById", int64(5)).
			Return(nil, expectedError).
			Once()

		controller := controllers.CreateProductController(mockService)

		// configura engine de rotas
		router := SetUpRouter()
		router.GET("/api/v1/products/:id", controller.GetById())

		// EXECUÇÃO
		response := ExecuteTestRequest(router, "GET", "/api/v1/products/5", []byte{})

		// VALIDAÇÃO
		assert.Equal(t, http.StatusNotFound, response.Code)
		assert.Equal(t, "{\"code\":404,\"message\":\"the product id was not found\"}", response.Body.String())
	})

	t.Run("find_by_id_existent: quando a solicitação for bem-sucedida, o backend retornará as informações do produto solicitado.", func(t *testing.T) {

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

		// PREPARACÃO
		mockService.
			On("GetById", int64(1)).
			Return(&expectedProduct, nil).
			Once()

		controller := controllers.CreateProductController(mockService)

		requestBody, _ := json.Marshal(body)

		router := SetUpRouter()
		router.GET("/api/v1/products/:id", controller.GetById())

		//  EXECUCAO
		response := ExecuteTestRequest(router, "GET", "/api/v1/products/1", requestBody)

		// VALIDACAO
		assert.Equal(t, http.StatusOK, response.Code)
	})
}

func TestProductController_UpdateDescription(t *testing.T) {

	mockService := mocks.NewProductService(t)

	t.Run("update_description_parse_id_error: quando o id do produto não for parseado, um código 400 será retornado", func(t *testing.T) {

		controller := controllers.CreateProductController(mockService)

		// configura engine de rotas
		router := SetUpRouter()
		router.PATCH("/api/v1/products/:id", controller.UpdateDescription())

		// EXECUÇÃO
		response := ExecuteTestRequest(router, "PATCH", "/api/v1/products/abc", []byte{})

		// VALIDAÇÃO
		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.Equal(t, "{\"code\":400,\"message\":\"invalid id\"}", response.Body.String())
	})

	t.Run("update_err: quando falhar o parse do body da requisição, um código 400 será retornado.", func(t *testing.T) {

		controller := controllers.CreateProductController(nil)

		router := SetUpRouter()
		router.PATCH("/api/v1/products/:id", controller.UpdateDescription())

		// EXECUÇÃO
		response := ExecuteTestRequest(router, "PATCH", "/api/v1/products/1", []byte{})

		// VALIDAÇÃO
		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.Equal(t, "{\"code\":400,\"message\":\"EOF\"}", response.Body.String())
	})

	t.Run("update_err: quando o campo a ser atualizado estiver vazio, um código 400 será retornado.", func(t *testing.T) {

		body := product.Product{
			Description: "",
		}

		controller := controllers.CreateProductController(nil)

		requestBody, _ := json.Marshal(body)

		router := SetUpRouter()
		router.PATCH("/api/v1/products/:id", controller.UpdateDescription())

		// EXECUÇÃO
		response := ExecuteTestRequest(router, "PATCH", "/api/v1/products/1", requestBody)

		// VALIDAÇÃO
		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.Equal(t, "{\"code\":400,\"message\":\"Key: 'requestProductPatch.Description' Error:Field validation for 'Description' failed on the 'required' tag\"}", response.Body.String())
	})

	t.Run("update_non_existent: quando o produto a ser atualizado não existir, um código 404 será devolvido.", func(t *testing.T) {

		body := product.Product{
			Description: "Yogurt",
		}

		// PREPARAÇÃO
		expectedError := errors.New("the product id was not found")
		mockService.
			On("UpdateDescription", int64(8), body.Description).
			Return(nil, expectedError).
			Once()

		controller := controllers.CreateProductController(mockService)

		requestBody, _ := json.Marshal(body)

		// configura engine de rotas
		router := SetUpRouter()
		router.PATCH("/api/v1/products/:id", controller.UpdateDescription())

		// EXECUÇÃO
		response := ExecuteTestRequest(router, "PATCH", "/api/v1/products/8", requestBody)

		// VALIDAÇÃO
		assert.Equal(t, http.StatusNotFound, response.Code)
		assert.Equal(t, "{\"code\":404,\"message\":\"the product id was not found\"}", response.Body.String())
	})

	t.Run("update_ok: quando a atualização dos dados for bem-sucedida, o produto será devolvido com as informações atualizadas juntamente com um código 200.", func(t *testing.T) {

		expectedProduct := product.Product{
			Id:          1,
			Description: "Yogurt light",
		}

		body := product.Product{
			ProductCode:                    "PROD02",
			Description:                    "Yogurt light",
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

		// PREPARAÇÃO
		mockService.
			On("UpdateDescription", expectedProduct.Id, expectedProduct.Description).
			Return(&expectedProduct, nil).
			Once()

		controller := controllers.CreateProductController(mockService)

		requestBody, _ := json.Marshal(body)

		// configura engine de rotas
		router := SetUpRouter()
		router.PATCH("/api/v1/products/:id", controller.UpdateDescription())

		// EXECUÇÃO
		response := ExecuteTestRequest(router, "PATCH", "/api/v1/products/1", requestBody)

		// VALIDAÇÃO
		assert.Equal(t, http.StatusOK, response.Code)
	})
}

func TestProductController_Delete(t *testing.T) {

	mockService := mocks.NewProductService(t)

	t.Run("delete_parse_id_error: quando o id do produto não for parseado, um código 400 será retornado", func(t *testing.T) {

		controller := controllers.CreateProductController(mockService)

		// configura engine de rotas
		router := SetUpRouter()
		router.DELETE("/api/v1/products/:id", controller.Delete())

		// EXECUÇÃO
		response := ExecuteTestRequest(router, "DELETE", "/api/v1/products/abc", []byte{})

		// VALIDAÇÃO
		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.Equal(t, "{\"error\":\"invalid id\"}", response.Body.String())
	})

	t.Run("delete_non_existent: quando o produto não existe, um código 404 será devolvido", func(t *testing.T) {

		// PREPARAÇÃO
		expectedError := errors.New("the product id was not found")
		mockService.
			On("Delete", int64(1)).
			Return(expectedError).
			Once()

		controller := controllers.CreateProductController(mockService)

		// configura engine de rotas
		router := SetUpRouter()
		router.DELETE("/api/v1/products/:id", controller.Delete())

		// EXECUÇÃO
		response := ExecuteTestRequest(router, "DELETE", "/api/v1/products/1", []byte{})

		// VALIDAÇÃO
		assert.Equal(t, http.StatusNotFound, response.Code)
		assert.Equal(t, "{\"code\":404,\"message\":\"the product id was not found\"}", response.Body.String())

	})

	t.Run("delete_ok: quando a exclusão for bem-sucedida, um código 204 será retornado.", func(t *testing.T) {

		mockService.
			On("Delete", int64(1)).
			Return(nil).
			Once()

		controller := controllers.CreateProductController(mockService)

		// configura engine de rotas
		router := SetUpRouter()
		router.DELETE("/api/v1/products/:id", controller.Delete())

		// EXECUÇÃO
		response := ExecuteTestRequest(router, "DELETE", "/api/v1/products/1", []byte{})

		// VALIDAÇÃO
		assert.Equal(t, http.StatusNoContent, response.Code)
	})
}
