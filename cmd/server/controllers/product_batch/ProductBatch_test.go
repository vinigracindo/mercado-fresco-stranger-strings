package controllers_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/go-playground/assert/v2"
	controllers "github.com/vinigracindo/mercado-fresco-stranger-strings/cmd/server/controllers/product_batch"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/product_batch/domain"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/product_batch/domain/mocks"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/pkg/testutil"
)

var (
	EndpointProductBatch    = "/api/v1/productBatches"
	invalidProductBatchJSON = []byte(`{invalid json}`)
)

func makeProductBatchRequest() controllers.RequestProductBatchPost {
	return controllers.RequestProductBatchPost{
		BatchNumber:        1,
		CurrentQuantity:    1,
		CurrentTemperature: 1.0,
		DueDate:            "2022-01-01",
		InitialQuantity:    1.0,
		ManufacturingDate:  "2022-01-01",
		ManufacturingHour:  1,
		MinumumTemperature: 1.0,
		ProductId:          1,
		SectionId:          1,
	}
}

func makeProductBatch() domain.ProductBatch {
	date := time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC)
	return domain.ProductBatch{
		BatchNumber:        1,
		CurrentQuantity:    1,
		CurrentTemperature: 1.0,
		DueDate:            date,
		InitialQuantity:    1.0,
		ManufacturingDate:  date,
		ManufacturingHour:  1,
		MinumumTemperature: 1.0,
		ProductId:          1,
		SectionId:          1,
	}
}

func TestProductBatch_Create(t *testing.T) {
	expectedProductBatch := makeProductBatch()
	requestProductBatch := makeProductBatchRequest()

	mockService := mocks.NewProductBatchService(t)
	controller := controllers.NewProductBatchController(mockService)
	router := testutil.SetUpRouter()
	router.POST(EndpointProductBatch, controller.Create())

	t.Run("create_ok: when data entry is successful, should return code 201. The object must be returned.", func(t *testing.T) {
		mockService.
			On("Create", context.TODO(), &expectedProductBatch).
			Return(&expectedProductBatch, nil).
			Once()

		reqBody, _ := json.Marshal(requestProductBatch)
		response := testutil.ExecuteTestRequest(router, http.MethodPost, EndpointProductBatch, reqBody)

		assert.Equal(t, http.StatusCreated, response.Code)
	})

	t.Run("invalid_json: when the request body is invalid, should return code 422. The error must be returned.", func(t *testing.T) {
		reqBody, _ := json.Marshal(invalidProductBatchJSON)
		response := testutil.ExecuteTestRequest(router, http.MethodPost, EndpointProductBatch, reqBody)

		assert.Equal(t, http.StatusUnprocessableEntity, response.Code)
	})

	t.Run("invalid_time_parse_dueDate: when the request body is invalid, should return code 400. The error must be returned.", func(t *testing.T) {
		newRequestProductBatch := makeProductBatchRequest()
		newRequestProductBatch.DueDate = "invalid date"
		reqBody, _ := json.Marshal(newRequestProductBatch)
		response := testutil.ExecuteTestRequest(router, http.MethodPost, EndpointProductBatch, reqBody)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})

	t.Run("invalid_time_parse_manufacturingDate: when the request body is invalid, should return code 400. The error must be returned.", func(t *testing.T) {
		newRequestProductBatch := makeProductBatchRequest()
		newRequestProductBatch.ManufacturingDate = "invalid date"
		reqBody, _ := json.Marshal(newRequestProductBatch)
		response := testutil.ExecuteTestRequest(router, http.MethodPost, EndpointProductBatch, reqBody)

		assert.Equal(t, http.StatusBadRequest, response.Code)

	})

	t.Run("not_found_relations: when the section is not found, should return code 409. The error must be returned.", func(t *testing.T) {
		mockService.
			On("Create", context.TODO(), &expectedProductBatch).
			Return(nil, fmt.Errorf("not found")).
			Once()

		reqBody, _ := json.Marshal(requestProductBatch)
		response := testutil.ExecuteTestRequest(router, http.MethodPost, EndpointProductBatch, reqBody)

		assert.Equal(t, http.StatusConflict, response.Code)

	})

}
