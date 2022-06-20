package controllers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
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

func ExecuteTestRequest(router *gin.Engine, method string, path string, body []byte) *httptest.ResponseRecorder {

	request := httptest.NewRequest(method, path, bytes.NewBuffer(body))

	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	return response
}

const ENDPOINT = "/api/v1/products"

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

		mockService.
			On("Create", expectedProduct.ProductCode, expectedProduct.Description, expectedProduct.Width, expectedProduct.Height,
				expectedProduct.Length, expectedProduct.NetWeight, expectedProduct.ExpirationRate, expectedProduct.RecommendedFreezingTemperature,
				expectedProduct.FreezingRate, expectedProduct.ProductTypeId, expectedProduct.SellerId).
			Return(&expectedProduct, nil).
			Once()

		controller := controllers.CreateProductController(mockService)

		requestBody, _ := json.Marshal(body)

		router := SetUpRouter()
		router.POST(ENDPOINT, controller.Create())

		response := ExecuteTestRequest(router, http.MethodPost, ENDPOINT, requestBody)

		assert.Equal(t, http.StatusCreated, response.Code)
		assert.JSONEq(t, "{\"data\":{\"id\":1,\"product_code\":\"PROD02\",\"description\":\"Yogurt\",\"width\":1.2,"+
			"\"height\":6.4,\"length\":4.5,\"net_weight\":3.4,\"expiration_rate\":1.5,\"recommended_freezing_temperature\":1.3,"+
			"\"freezing_rate\":2,\"product_type_id\":2,\"seller_id\":2}}", response.Body.String())

		fmt.Println(response.Body.String())
	})

	t.Run("create_fail: quando o objeto JSON não contiver os campos necessários, um código 422 será retornado.", func(t *testing.T) {

		controller := controllers.CreateProductController(nil)

		router := SetUpRouter()
		router.POST(ENDPOINT, controller.Create())

		response := ExecuteTestRequest(router, http.MethodPost, ENDPOINT, []byte{})

		assert.Equal(t, http.StatusUnprocessableEntity, response.Code)
		assert.Equal(t, "{\"code\":422,\"message\":\"invalid input. Check the data entered\"}", response.Body.String())
	})

	t.Run("create_conflict: quando o product_code já existir, ele retornará um erro 409 Conflict.", func(t *testing.T) {

		expectedError := errors.New("the product code has already been registered")
		mockService.
			On("Create", expectedProduct.ProductCode, expectedProduct.Description, expectedProduct.Width,
				expectedProduct.Height, expectedProduct.Length, expectedProduct.NetWeight, expectedProduct.ExpirationRate,
				expectedProduct.RecommendedFreezingTemperature, expectedProduct.FreezingRate, expectedProduct.ProductTypeId,
				expectedProduct.SellerId).
			Return(nil, expectedError).
			Once()

		controller := controllers.CreateProductController(mockService)

		requestBody, _ := json.Marshal(body)

		router := SetUpRouter()
		router.POST(ENDPOINT, controller.Create())

		response := ExecuteTestRequest(router, http.MethodPost, ENDPOINT, requestBody)

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

		expectedError := errors.New("the request sent to the server is invalid or corrupted")
		mockService.
			On("GetAll").
			Return(nil, expectedError).
			Once()

		controller := controllers.CreateProductController(mockService)

		requestBody, _ := json.Marshal(body)

		router := SetUpRouter()
		router.GET(ENDPOINT, controller.GetAll())

		response := ExecuteTestRequest(router, http.MethodGet, ENDPOINT, requestBody)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
		assert.Equal(t, "{\"code\":500,\"message\":\"the request sent to the server is invalid or corrupted\"}", response.Body.String())
	})

	t.Run("find_all: quando a solicitação for bem-sucedida, o back-end retornará uma lista de todos os produtos existentes.", func(t *testing.T) {

		mockService.
			On("GetAll").
			Return(expectedProduct, nil).
			Once()

		controller := controllers.CreateProductController(mockService)

		requestBody, _ := json.Marshal(body)

		router := SetUpRouter()
		router.GET(ENDPOINT, controller.GetAll())

		response := ExecuteTestRequest(router, http.MethodGet, ENDPOINT, requestBody)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.JSONEq(t, "{\"data\":[{\"id\":1,\"product_code\":\"PROD02\",\"description\":\"Yogurt\",\"width\":1.2,\"height\":6.4,"+
			"\"length\":4.5,\"net_weight\":3.4,\"expiration_rate\":1.5,\"recommended_freezing_temperature\":1.3,\"freezing_rate\":2,"+
			"\"product_type_id\":2,\"seller_id\":2},{\"id\":2,\"product_code\":\"PROD03\",\"description\":\"Yogurt light\",\"width\":1.5,"+
			"\"height\":5.4,\"length\":3.5,\"net_weight\":4.4,\"expiration_rate\":1.8,\"recommended_freezing_temperature\":1.2,"+
			"\"freezing_rate\":2,\"product_type_id\":3,\"seller_id\":3}]}", response.Body.String())
	})
}

func TestProductController_GetById(t *testing.T) {

	mockService := mocks.NewProductService(t)

	t.Run("find_by_id_parse_error: quando o id do produto não for parseado", func(t *testing.T) {

		controller := controllers.CreateProductController(mockService)

		router := SetUpRouter()
		router.GET(ENDPOINT+"/:id", controller.GetById())

		response := ExecuteTestRequest(router, http.MethodGet, ENDPOINT+"/abc", []byte{})

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.Equal(t, "{\"code\":400,\"message\":\"strconv.ParseInt: parsing \\\"abc\\\": invalid syntax\"}", response.Body.String())
	})

	t.Run("find_by_id_non_existent: quando o produto não existe, um código 404 será devolvido", func(t *testing.T) {

		expectedError := errors.New("the product id was not found")
		mockService.
			On("GetById", int64(5)).
			Return(nil, expectedError).
			Once()

		controller := controllers.CreateProductController(mockService)

		router := SetUpRouter()
		router.GET(ENDPOINT+"/:id", controller.GetById())

		response := ExecuteTestRequest(router, http.MethodGet, ENDPOINT+"/5", []byte{})

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

		mockService.
			On("GetById", int64(1)).
			Return(&expectedProduct, nil).
			Once()

		controller := controllers.CreateProductController(mockService)

		requestBody, _ := json.Marshal(body)

		router := SetUpRouter()
		router.GET(ENDPOINT+"/:id", controller.GetById())

		response := ExecuteTestRequest(router, http.MethodGet, ENDPOINT+"/1", requestBody)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.JSONEq(t, "{\"data\":{\"id\":1,\"product_code\":\"PROD02\",\"description\":\"Yogurt\",\"width\":1.2,\"height\":6.4,"+
			"\"length\":4.5,\"net_weight\":3.4,\"expiration_rate\":1.5,\"recommended_freezing_temperature\":1.3,\"freezing_rate\":2,"+
			"\"product_type_id\":2,\"seller_id\":2}}", response.Body.String())
	})
}

func TestProductController_UpdateDescription(t *testing.T) {

	mockService := mocks.NewProductService(t)

	t.Run("update_description_parse_id_error: quando o id do produto não for parseado, um código 400 será retornado", func(t *testing.T) {

		controller := controllers.CreateProductController(mockService)

		router := SetUpRouter()
		router.PATCH(ENDPOINT+"/:id", controller.UpdateDescription())

		response := ExecuteTestRequest(router, http.MethodPatch, ENDPOINT+"/abc", []byte{})

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.Equal(t, "{\"code\":400,\"message\":\"invalid id\"}", response.Body.String())
	})

	t.Run("update_err: quando falhar o parse do body da requisição, um código 400 será retornado.", func(t *testing.T) {

		controller := controllers.CreateProductController(nil)

		router := SetUpRouter()
		router.PATCH(ENDPOINT+"/:id", controller.UpdateDescription())

		response := ExecuteTestRequest(router, http.MethodPatch, ENDPOINT+"/1", []byte{})

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
		router.PATCH(ENDPOINT+"/:id", controller.UpdateDescription())

		response := ExecuteTestRequest(router, http.MethodPatch, ENDPOINT+"/1", requestBody)

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.Equal(t, "{\"code\":400,\"message\":\"Key: 'requestProductPatch.Description' Error:Field validation for 'Description' failed on the 'required' tag\"}", response.Body.String())
	})

	t.Run("update_non_existent: quando o produto a ser atualizado não existir, um código 404 será devolvido.", func(t *testing.T) {

		body := product.Product{
			Description: "Yogurt",
		}
		expectedError := errors.New("the product id was not found")
		mockService.
			On("UpdateDescription", int64(8), body.Description).
			Return(nil, expectedError).
			Once()

		controller := controllers.CreateProductController(mockService)

		requestBody, _ := json.Marshal(body)

		router := SetUpRouter()
		router.PATCH(ENDPOINT+"/:id", controller.UpdateDescription())

		response := ExecuteTestRequest(router, http.MethodPatch, ENDPOINT+"/8", requestBody)

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

		mockService.
			On("UpdateDescription", expectedProduct.Id, expectedProduct.Description).
			Return(&expectedProduct, nil).
			Once()

		controller := controllers.CreateProductController(mockService)

		requestBody, _ := json.Marshal(body)

		router := SetUpRouter()
		router.PATCH(ENDPOINT+"/:id", controller.UpdateDescription())

		response := ExecuteTestRequest(router, http.MethodPatch, ENDPOINT+"/1", requestBody)

		assert.Equal(t, http.StatusOK, response.Code)

		assert.JSONEq(t, "{\"data\":{\"id\":1,\"product_code\":\"\",\"description\":\"Yogurt light\","+
			"\"width\":0,\"height\":0,\"length\":0,\"net_weight\":0,\"expiration_rate\":0,\"recommended_freezing_temperature\":0,"+
			"\"freezing_rate\":0,\"product_type_id\":0,\"seller_id\":0}}", response.Body.String())
	})
}

func TestProductController_Delete(t *testing.T) {

	mockService := mocks.NewProductService(t)

	t.Run("delete_parse_id_error: quando o id do produto não for parseado, um código 400 será retornado", func(t *testing.T) {

		controller := controllers.CreateProductController(mockService)

		router := SetUpRouter()
		router.DELETE(ENDPOINT+"/:id", controller.Delete())

		response := ExecuteTestRequest(router, http.MethodDelete, ENDPOINT+"/abc", []byte{})

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.Equal(t, "{\"error\":\"invalid id\"}", response.Body.String())
	})

	t.Run("delete_non_existent: quando o produto não existe, um código 404 será devolvido", func(t *testing.T) {

		expectedError := errors.New("the product id was not found")
		mockService.
			On("Delete", int64(1)).
			Return(expectedError).
			Once()

		controller := controllers.CreateProductController(mockService)

		router := SetUpRouter()
		router.DELETE(ENDPOINT+"/:id", controller.Delete())

		response := ExecuteTestRequest(router, http.MethodDelete, ENDPOINT+"/1", []byte{})

		assert.Equal(t, http.StatusNotFound, response.Code)
		assert.Equal(t, "{\"code\":404,\"message\":\"the product id was not found\"}", response.Body.String())

	})

	t.Run("delete_ok: quando a exclusão for bem-sucedida, um código 204 será retornado.", func(t *testing.T) {

		mockService.
			On("Delete", int64(1)).
			Return(nil).
			Once()

		controller := controllers.CreateProductController(mockService)

		router := SetUpRouter()
		router.DELETE(ENDPOINT+"/:id", controller.Delete())

		response := ExecuteTestRequest(router, http.MethodDelete, ENDPOINT+"/1", []byte{})

		assert.Equal(t, http.StatusNoContent, response.Code)
	})
}
