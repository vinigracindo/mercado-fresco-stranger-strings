package controllers_test

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	controllers "github.com/vinigracindo/mercado-fresco-stranger-strings/cmd/server/controllers/product_records"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/product_records/domain"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/product_records/domain/mocks"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/pkg/testutil"
	"net/http"
	"testing"
	"time"
)

const EndpointProductRecords = "/api/v1/productRecords"

func TestProductRecordsController_Create(t *testing.T) {

	var expectedProductRecords = controllers.RequestProductRecordsPost{
		LastUpdateDate: "2022-04-04",
		PurchasePrice:  10.5,
		SalePrice:      15.2,
		ProductId:      1,
	}

	dateParse, _ := time.Parse("2006-01-02", expectedProductRecords.LastUpdateDate)

	var bodyProductRecords = domain.ProductRecords{
		LastUpdateDate: dateParse,
		PurchasePrice:  10.5,
		SalePrice:      15.2,
		ProductId:      1,
	}

	mockService := mocks.NewProductRecordsService(t)
	controller := controllers.CreateProductRecordsController(mockService)

	router := testutil.SetUpRouter()
	router.POST(EndpointProductRecords, controller.Create())

	t.Run("create_ok: when data entry is successful, should return code 201", func(t *testing.T) {

		mockService.
			On("Create", context.TODO(), &bodyProductRecords).
			Return(&bodyProductRecords, nil).
			Once()

		requestBody, _ := json.Marshal(expectedProductRecords)
		response := testutil.ExecuteTestRequest(router, http.MethodPost, EndpointProductRecords, requestBody)

		assert.Equal(t, http.StatusCreated, response.Code)
		assert.JSONEq(t, "{\"data\":{\"id\":0,\"last_update_date\":\"2022-04-04T00:00:00Z\",\"purchase_price\":10.5,\"sale_price\":15.2,\"product_id\":1}}",
			response.Body.String())

		fmt.Printf(response.Body.String())
	})

	t.Run("create_fail_invalid_json: when the JSON does not contain the required fields, should return code 422", func(t *testing.T) {

		requestBody, _ := json.Marshal([]byte{})
		response := testutil.ExecuteTestRequest(router, http.MethodPost, EndpointProductRecords, requestBody)

		assert.Equal(t, http.StatusUnprocessableEntity, response.Code)
		assert.Equal(t, "{\"code\":422,\"message\":\"invalid input. Check the data entered\"}", response.Body.String())
	})

	t.Run("create_fail_invalid_time_parse: when the request body is invalid, should return code 400", func(t *testing.T) {

		invalidDateProductRecord := controllers.RequestProductRecordsPost{
			LastUpdateDate: "invalid date",
			PurchasePrice:  expectedProductRecords.PurchasePrice,
			SalePrice:      expectedProductRecords.SalePrice,
			ProductId:      expectedProductRecords.ProductId,
		}

		requestBody, _ := json.Marshal(invalidDateProductRecord)
		response := testutil.ExecuteTestRequest(router, http.MethodPost, EndpointProductRecords, requestBody)

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.Equal(t, "{\"code\":400,\"message\":\"parsing time \\\"invalid date\\\" as \\\"2006-01-02\\\": cannot parse \\\"invalid date\\\" as \\\"2006\\\"\"}",
			response.Body.String())
	})

	t.Run("product_not_found: when the product id is not found, should return code 409", func(t *testing.T) {

		mockService.
			On("Create", context.TODO(), &bodyProductRecords).
			Return(nil, domain.ErrProductIdNotFound).
			Once()

		requestBody, _ := json.Marshal(expectedProductRecords)
		response := testutil.ExecuteTestRequest(router, http.MethodPost, EndpointProductRecords, requestBody)

		assert.Equal(t, http.StatusConflict, response.Code)
		assert.Equal(t, "{\"code\":409,\"message\":\"product id not found\"}", response.Body.String())
	})
}
