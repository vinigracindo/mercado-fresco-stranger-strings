package controllers_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/cmd/server/controllers/product"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/product/domain"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/product/domain/mocks"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/pkg/testutil"
)

const EndpointProduct = "/api/v1/products"

var expectedProduct = domain.Product{
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

var bodyProduct = domain.Product{
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

func TestProductController_Create(t *testing.T) {

	mockService := mocks.NewProductService(t)

	t.Run("create_ok: when data entry is successful, should return code 201", func(t *testing.T) {

		mockService.
			On("Create", expectedProduct.ProductCode, expectedProduct.Description, expectedProduct.Width, expectedProduct.Height,
				expectedProduct.Length, expectedProduct.NetWeight, expectedProduct.ExpirationRate, expectedProduct.RecommendedFreezingTemperature,
				expectedProduct.FreezingRate, expectedProduct.ProductTypeId, expectedProduct.SellerId).
			Return(&expectedProduct, nil).
			Once()

		controller := controllers.CreateProductController(mockService)

		requestBody, _ := json.Marshal(bodyProduct)

		router := testutil.SetUpRouter()
		router.POST(EndpointProduct, controller.Create())

		response := testutil.ExecuteTestRequest(router, http.MethodPost, EndpointProduct, requestBody)

		assert.Equal(t, http.StatusCreated, response.Code)
		assert.JSONEq(t, "{\"data\":{\"id\":1,\"product_code\":\"PROD02\",\"description\":\"Yogurt\",\"width\":1.2,"+
			"\"height\":6.4,\"length\":4.5,\"net_weight\":3.4,\"expiration_rate\":1.5,\"recommended_freezing_temperature\":1.3,"+
			"\"freezing_rate\":2,\"product_type_id\":2,\"seller_id\":2}}", response.Body.String())

		fmt.Println(response.Body.String())
	})

	t.Run("create_fail: when the JSON does not contain the required fields, should return code 422", func(t *testing.T) {

		controller := controllers.CreateProductController(nil)

		router := testutil.SetUpRouter()
		router.POST(EndpointProduct, controller.Create())

		response := testutil.ExecuteTestRequest(router, http.MethodPost, EndpointProduct, []byte{})

		assert.Equal(t, http.StatusUnprocessableEntity, response.Code)
		assert.Equal(t, "{\"code\":422,\"message\":\"invalid input. Check the data entered\"}", response.Body.String())
	})

	t.Run("create_conflict: when the product_code already exists, should return code 409", func(t *testing.T) {

		expectedError := errors.New("the product code has already been registered")
		mockService.
			On("Create", expectedProduct.ProductCode, expectedProduct.Description, expectedProduct.Width,
				expectedProduct.Height, expectedProduct.Length, expectedProduct.NetWeight, expectedProduct.ExpirationRate,
				expectedProduct.RecommendedFreezingTemperature, expectedProduct.FreezingRate, expectedProduct.ProductTypeId,
				expectedProduct.SellerId).
			Return(nil, expectedError).
			Once()

		controller := controllers.CreateProductController(mockService)

		requestBody, _ := json.Marshal(bodyProduct)

		router := testutil.SetUpRouter()
		router.POST(EndpointProduct, controller.Create())

		response := testutil.ExecuteTestRequest(router, http.MethodPost, EndpointProduct, requestBody)

		assert.Equal(t, http.StatusConflict, response.Code)
		assert.Equal(t, "{\"code\":409,\"message\":\"the product code has already been registered\"}", response.Body.String())
	})
}

func TestProductController_GetAll(t *testing.T) {

	mockService := mocks.NewProductService(t)

	expectedProductList := []domain.Product{expectedProduct, expectedProduct}

	bodyList := []domain.Product{bodyProduct, bodyProduct}

	t.Run("find_all_internal_server_error: when the request is not successful, should return code 500 ", func(t *testing.T) {

		expectedError := errors.New("the request sent to the server is invalid or corrupted")
		mockService.
			On("GetAll").
			Return(nil, expectedError).
			Once()

		controller := controllers.CreateProductController(mockService)

		requestBody, _ := json.Marshal(bodyList)

		router := testutil.SetUpRouter()
		router.GET(EndpointProduct, controller.GetAll())

		response := testutil.ExecuteTestRequest(router, http.MethodGet, EndpointProduct, requestBody)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
		assert.Equal(t, "{\"code\":500,\"message\":\"the request sent to the server is invalid or corrupted\"}", response.Body.String())
	})

	t.Run("find_all: when data entry is successful, should return code 200", func(t *testing.T) {

		mockService.
			On("GetAll").
			Return(expectedProductList, nil).
			Once()

		controller := controllers.CreateProductController(mockService)

		requestBody, _ := json.Marshal(bodyList)

		router := testutil.SetUpRouter()
		router.GET(EndpointProduct, controller.GetAll())

		response := testutil.ExecuteTestRequest(router, http.MethodGet, EndpointProduct, requestBody)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.JSONEq(t, "{\"data\":[{\"id\":1,\"product_code\":\"PROD02\",\"description\":\"Yogurt\",\"width\":1.2,\"height\":6.4,"+
			"\"length\":4.5,\"net_weight\":3.4,\"expiration_rate\":1.5,\"recommended_freezing_temperature\":1.3,\"freezing_rate\":2,"+
			"\"product_type_id\":2,\"seller_id\":2},{\"id\":1,\"product_code\":\"PROD02\",\"description\":\"Yogurt\",\"width\":1.2,"+
			"\"height\":6.4,\"length\":4.5,\"net_weight\":3.4,\"expiration_rate\":1.5,\"recommended_freezing_temperature\":1.3,"+
			"\"freezing_rate\":2,\"product_type_id\":2,\"seller_id\":2}]}", response.Body.String())
	})
}

func TestProductController_GetById(t *testing.T) {

	mockService := mocks.NewProductService(t)

	t.Run("find_by_id_parse_error: when product id is not parsed, should return code 400", func(t *testing.T) {

		controller := controllers.CreateProductController(mockService)

		router := testutil.SetUpRouter()
		router.GET(EndpointProduct+"/:id", controller.GetById())

		response := testutil.ExecuteTestRequest(router, http.MethodGet, EndpointProduct+"/abc", []byte{})

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.Equal(t, "{\"code\":400,\"message\":\"strconv.ParseInt: parsing \\\"abc\\\": invalid syntax\"}", response.Body.String())
	})

	t.Run("find_by_id_existent: when the request is successful, should return code 200", func(t *testing.T) {

		expectedError := errors.New("the product id was not found")
		mockService.
			On("GetById", int64(5)).
			Return(nil, expectedError).
			Once()

		controller := controllers.CreateProductController(mockService)

		router := testutil.SetUpRouter()
		router.GET(EndpointProduct+"/:id", controller.GetById())

		response := testutil.ExecuteTestRequest(router, http.MethodGet, EndpointProduct+"/5", []byte{})

		assert.Equal(t, http.StatusNotFound, response.Code)
		assert.Equal(t, "{\"code\":404,\"message\":\"the product id was not found\"}", response.Body.String())
	})

	t.Run("find_by_id_existent: when the request is successful, should return code 200", func(t *testing.T) {

		mockService.
			On("GetById", int64(1)).
			Return(&expectedProduct, nil).
			Once()

		controller := controllers.CreateProductController(mockService)

		requestBody, _ := json.Marshal(bodyProduct)

		router := testutil.SetUpRouter()
		router.GET(EndpointProduct+"/:id", controller.GetById())

		response := testutil.ExecuteTestRequest(router, http.MethodGet, EndpointProduct+"/1", requestBody)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.JSONEq(t, "{\"data\":{\"id\":1,\"product_code\":\"PROD02\",\"description\":\"Yogurt\",\"width\":1.2,\"height\":6.4,"+
			"\"length\":4.5,\"net_weight\":3.4,\"expiration_rate\":1.5,\"recommended_freezing_temperature\":1.3,\"freezing_rate\":2,"+
			"\"product_type_id\":2,\"seller_id\":2}}", response.Body.String())
	})
}

func TestProductController_UpdateDescription(t *testing.T) {

	mockService := mocks.NewProductService(t)

	t.Run("update_invalid_id_parse_error: when product id is not parsed, should return code 400", func(t *testing.T) {

		controller := controllers.CreateProductController(mockService)

		router := testutil.SetUpRouter()
		router.PATCH(EndpointProduct+"/:id", controller.UpdateDescription())

		response := testutil.ExecuteTestRequest(router, http.MethodPatch, EndpointProduct+"/abc", []byte{})

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.Equal(t, "{\"code\":400,\"message\":\"invalid id\"}", response.Body.String())
	})

	t.Run("update_invalid_body: when the body is invalid, should return code 400", func(t *testing.T) {

		controller := controllers.CreateProductController(nil)

		router := testutil.SetUpRouter()
		router.PATCH(EndpointProduct+"/:id", controller.UpdateDescription())

		response := testutil.ExecuteTestRequest(router, http.MethodPatch, EndpointProduct+"/1", []byte{})

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.Equal(t, "{\"code\":400,\"message\":\"EOF\"}", response.Body.String())
	})

	t.Run("update_invalid_field_value: when the field is empty,should return code 400", func(t *testing.T) {

		body := domain.Product{
			Description: "",
		}

		controller := controllers.CreateProductController(nil)

		requestBody, _ := json.Marshal(body)

		router := testutil.SetUpRouter()
		router.PATCH(EndpointProduct+"/:id", controller.UpdateDescription())

		response := testutil.ExecuteTestRequest(router, http.MethodPatch, EndpointProduct+"/1", requestBody)

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.Equal(t, "{\"code\":400,\"message\":\"Key: 'requestProductPatch.Description' Error:Field validation for 'Description' failed on the 'required' tag\"}", response.Body.String())
	})

	t.Run("update_non_existent: when the product does not exist, should return code 404", func(t *testing.T) {

		body := domain.Product{
			Description: "Yogurt",
		}
		expectedError := errors.New("the product id was not found")
		mockService.
			On("UpdateDescription", int64(8), body.Description).
			Return(nil, expectedError).
			Once()

		controller := controllers.CreateProductController(mockService)

		requestBody, _ := json.Marshal(body)

		router := testutil.SetUpRouter()
		router.PATCH(EndpointProduct+"/:id", controller.UpdateDescription())

		response := testutil.ExecuteTestRequest(router, http.MethodPatch, EndpointProduct+"/8", requestBody)

		assert.Equal(t, http.StatusNotFound, response.Code)
		assert.Equal(t, "{\"code\":404,\"message\":\"the product id was not found\"}", response.Body.String())
	})

	t.Run("update_ok: when the request is successful, should return code 200", func(t *testing.T) {

		expectedProduct := domain.Product{
			Id:                             1,
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

		body := domain.Product{
			Description: "Yogurt light",
		}

		mockService.
			On("UpdateDescription", expectedProduct.Id, expectedProduct.Description).
			Return(&expectedProduct, nil).
			Once()

		controller := controllers.CreateProductController(mockService)

		requestBody, _ := json.Marshal(body)

		router := testutil.SetUpRouter()
		router.PATCH(EndpointProduct+"/:id", controller.UpdateDescription())

		response := testutil.ExecuteTestRequest(router, http.MethodPatch, EndpointProduct+"/1", requestBody)

		assert.Equal(t, http.StatusOK, response.Code)

		assert.JSONEq(t, "{\"data\":{\"id\":1,\"product_code\":\"PROD02\",\"description\":\"Yogurt light\","+
			"\"width\":1.2,\"height\":6.4,\"length\":4.5,\"net_weight\":3.4,\"expiration_rate\":1.5,"+
			"\"recommended_freezing_temperature\":1.3,\"freezing_rate\":2,\"product_type_id\":2,\"seller_id\":2}}", response.Body.String())

		print(response.Body.String())
	})
}

func TestProductController_Delete(t *testing.T) {

	mockService := mocks.NewProductService(t)

	t.Run("delete_id_parse_error: when product id is not parsed, should return code 400", func(t *testing.T) {

		controller := controllers.CreateProductController(mockService)

		router := testutil.SetUpRouter()
		router.DELETE(EndpointProduct+"/:id", controller.Delete())

		response := testutil.ExecuteTestRequest(router, http.MethodDelete, EndpointProduct+"/abc", []byte{})

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.Equal(t, "{\"error\":\"invalid id\"}", response.Body.String())
	})

	t.Run("delete_non_existent: when the product does not exist, should return code 404", func(t *testing.T) {

		expectedError := errors.New("the product id was not found")
		mockService.
			On("Delete", int64(1)).
			Return(expectedError).
			Once()

		controller := controllers.CreateProductController(mockService)

		router := testutil.SetUpRouter()
		router.DELETE(EndpointProduct+"/:id", controller.Delete())

		response := testutil.ExecuteTestRequest(router, http.MethodDelete, EndpointProduct+"/1", []byte{})

		assert.Equal(t, http.StatusNotFound, response.Code)
		assert.Equal(t, "{\"code\":404,\"message\":\"the product id was not found\"}", response.Body.String())

	})

	t.Run("delete_ok: when the request is successful, should return code 204", func(t *testing.T) {

		mockService.
			On("Delete", int64(1)).
			Return(nil).
			Once()

		controller := controllers.CreateProductController(mockService)

		router := testutil.SetUpRouter()
		router.DELETE(EndpointProduct+"/:id", controller.Delete())

		response := testutil.ExecuteTestRequest(router, http.MethodDelete, EndpointProduct+"/1", []byte{})

		assert.Equal(t, http.StatusNoContent, response.Code)
	})
}
