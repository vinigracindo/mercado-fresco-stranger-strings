package controllers_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	controllers "github.com/vinigracindo/mercado-fresco-stranger-strings/cmd/server/controllers/section"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/section/domain"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/section/domain/mocks"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/pkg/testutil"
)

var bodySection = domain.SectionModel{
	SectionNumber:      1,
	CurrentTemperature: 1,
	MinimumTemperature: 1,
	CurrentCapacity:    int64(1),
	MinimumCapacity:    int64(1),
	MaximumCapacity:    int64(1),
	WarehouseId:        int64(1),
	ProductTypeId:      int64(1),
}

var bodyFailSection = domain.SectionModel{
	SectionNumber:      0,
	CurrentTemperature: 0,
	MinimumTemperature: 0,
	CurrentCapacity:    0,
	MinimumCapacity:    0,
	MaximumCapacity:    0,
	WarehouseId:        0,
	ProductTypeId:      0,
}

var EndpointSection = "/api/v1/sections"

var expectedSection = domain.SectionModel{
	Id:                 1,
	SectionNumber:      1,
	CurrentTemperature: 1,
	MinimumTemperature: 1,
	CurrentCapacity:    1,
	MinimumCapacity:    1,
	MaximumCapacity:    1,
	WarehouseId:        1,
	ProductTypeId:      1,
}

var expectedRecordProductBySection = domain.ReportProductsModel{
	Id:            int64(1),
	SectionNumber: int64(1),
	ProductsCount: int64(200),
}

var ctx = context.Background()

func TestSectionController_Create(t *testing.T) {
	mockService := mocks.NewSectionService(t)

	t.Run("create_ok: when data entry is successful, should return code 201", func(t *testing.T) {
		mockService.
			On("Create",
				ctx,
				expectedSection.SectionNumber,
				expectedSection.CurrentTemperature,
				expectedSection.MinimumTemperature,
				expectedSection.CurrentCapacity,
				expectedSection.MinimumCapacity,
				expectedSection.MaximumCapacity,
				expectedSection.WarehouseId,
				expectedSection.ProductTypeId,
			).Return(expectedSection, nil).
			Once()

		controller := controllers.NewSection(mockService)

		requestBody, _ := json.Marshal(bodySection)

		r := testutil.SetUpRouter()

		r.POST(EndpointSection, controller.Create())

		response := testutil.ExecuteTestRequest(r, http.MethodPost, EndpointSection, requestBody)

		assert.Equal(t, http.StatusCreated, response.Code)
		assert.JSONEq(t, "{\"data\":{\"id\":1,\"section_number\":1,\"current_temperature\":1,\"minimum_temperature\":1,\"current_capacity\":1,\"minimum_capacity\":1,\"maximum_capacity\":1,\"warehouse_id\":1,\"product_type_id\":1}}", response.Body.String())
	})

	t.Run("create_fail: when the JSON does not contain the required fields, should return code 422", func(t *testing.T) {
		controller := controllers.NewSection(nil)

		r := testutil.SetUpRouter()
		r.POST(EndpointSection, controller.Create())
		response := testutil.ExecuteTestRequest(r, http.MethodPost, EndpointSection, []byte{})

		assert.Equal(t, http.StatusUnprocessableEntity, response.Code)
		assert.JSONEq(t, "{\"code\":422,\"message\":\"EOF\"}", response.Body.String())
	})

	t.Run("create_conflict: when the section_number already exists, should return code 409", func(t *testing.T) {
		mockService.
			On("Create",
				ctx,
				expectedSection.SectionNumber,
				expectedSection.CurrentTemperature,
				expectedSection.MinimumTemperature,
				expectedSection.CurrentCapacity,
				expectedSection.MinimumCapacity,
				expectedSection.MaximumCapacity,
				expectedSection.WarehouseId,
				expectedSection.ProductTypeId,
			).
			Return(domain.SectionModel{}, fmt.Errorf("already a section with this code")).
			Once()

		controller := controllers.NewSection(mockService)

		requestBody, _ := json.Marshal(bodySection)

		r := testutil.SetUpRouter()

		r.POST(EndpointSection, controller.Create())

		response := testutil.ExecuteTestRequest(r, http.MethodPost, EndpointSection, requestBody)

		assert.Equal(t, http.StatusConflict, response.Code)
		assert.JSONEq(t, "{\"code\":409,\"message\":\"already a section with this code\"}", response.Body.String())
	})
}

func TestSectionController_GetAll(t *testing.T) {
	mockService := mocks.NewSectionService(t)

	t.Run("find_all: when data entry is successful, should return code 200", func(t *testing.T) {
		mockService.
			On("GetAll", ctx).
			Return([]domain.SectionModel{expectedSection}, nil).
			Once()
		controller := controllers.NewSection(mockService)

		requestBody, _ := json.Marshal(bodySection)

		r := testutil.SetUpRouter()
		r.GET(EndpointSection, controller.GetAll())
		response := testutil.ExecuteTestRequest(r, http.MethodGet, EndpointSection, requestBody)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.JSONEq(t, "{\"data\":[{\"id\":1,\"section_number\":1,\"current_temperature\":1,\"minimum_temperature\":1,\"current_capacity\":1,\"minimum_capacity\":1,\"maximum_capacity\":1,\"warehouse_id\":1,\"product_type_id\":1}]}", response.Body.String())
	})

	t.Run("get_all_fail: when GetAll fail, should return code 400", func(t *testing.T) {
		mockService.
			On("GetAll", ctx).
			Return([]domain.SectionModel{}, fmt.Errorf("any error"))
		controller := controllers.NewSection(mockService)

		requestBody, _ := json.Marshal(bodySection)

		r := testutil.SetUpRouter()
		r.GET(EndpointSection, controller.GetAll())
		response := testutil.ExecuteTestRequest(r, http.MethodGet, EndpointSection, requestBody)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
		assert.JSONEq(t, "{\"code\":500,\"message\":\"any error\"}", response.Body.String())
	})
}

func TestSectionController_GetById(t *testing.T) {
	mockService := mocks.NewSectionService(t)

	t.Run("find_by_id_existent: when the request is successful, should return code 200", func(t *testing.T) {
		mockService.
			On("GetById", ctx, int64(1)).
			Return(expectedSection, nil)
		controller := controllers.NewSection(mockService)

		requestBody, _ := json.Marshal(bodySection)

		r := testutil.SetUpRouter()
		r.GET(EndpointSection+"/:id", controller.GetById())
		response := testutil.ExecuteTestRequest(r, http.MethodGet, EndpointSection+"/1", requestBody)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.JSONEq(t, "{\"data\":{\"id\":1,\"section_number\":1,\"current_temperature\":1,\"minimum_temperature\":1,\"current_capacity\":1,\"minimum_capacity\":1,\"maximum_capacity\":1,\"warehouse_id\":1,\"product_type_id\":1}}", response.Body.String())
	})

	t.Run("find_by_id_non_existent: when the section does not exist, should return code 404", func(t *testing.T) {
		mockService := mocks.NewSectionService(t)
		mockService.On("GetById", ctx, int64(1)).Return(domain.SectionModel{}, fmt.Errorf("section not found"))
		controller := controllers.NewSection(mockService)

		requestBody, _ := json.Marshal(bodySection)

		r := testutil.SetUpRouter()
		r.GET(EndpointSection+"/:id", controller.GetById())
		response := testutil.ExecuteTestRequest(r, http.MethodGet, EndpointSection+"/1", requestBody)

		assert.Equal(t, http.StatusNotFound, response.Code)
		assert.JSONEq(t, "{\"code\":404,\"message\":\"section not found\"}", response.Body.String())
	})

	t.Run("find_by_id_parse_error: when section id is not parsed, should return code 400", func(t *testing.T) {
		controller := controllers.NewSection(nil)

		r := testutil.SetUpRouter()
		r.GET(EndpointSection+"/:id", controller.GetById())
		response := testutil.ExecuteTestRequest(r, http.MethodGet, EndpointSection+"/idInvalid", []byte{})

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.JSONEq(t, "{\"code\":400,\"message\":\"strconv.ParseInt: parsing \\\"idInvalid\\\": invalid syntax\"}", response.Body.String())
	})
}

func TestSectionController_Update(t *testing.T) {
	mockService := mocks.NewSectionService(t)

	var bodyUpdate = domain.SectionModel{
		CurrentCapacity: int64(1),
	}

	var bodyFailSectionUpdate = domain.SectionModel{
		CurrentCapacity: -1,
	}

	t.Run("update_ok: when the request is successful, should return code 200", func(t *testing.T) {
		mockService.
			On("UpdateCurrentCapacity", ctx, int64(1), int64(1)).
			Return(&expectedSection, nil).
			Once()
		controller := controllers.NewSection(mockService)

		requestBody, _ := json.Marshal(bodyUpdate)

		r := testutil.SetUpRouter()
		r.PATCH(EndpointSection+"/:id", controller.UpdateCurrentCapacity())
		response := testutil.ExecuteTestRequest(r, http.MethodPatch, EndpointSection+"/1", requestBody)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.JSONEq(t, "{\"data\":{\"id\":1,\"section_number\":1,\"current_temperature\":1,\"minimum_temperature\":1,\"current_capacity\":1,\"minimum_capacity\":1,\"maximum_capacity\":1,\"warehouse_id\":1,\"product_type_id\":1}}", response.Body.String())
	})

	t.Run("update_non_existent: when the section does not exist, should return code 404", func(t *testing.T) {
		mockService.
			On("UpdateCurrentCapacity", ctx, int64(1), int64(1)).
			Return(nil, fmt.Errorf("section not found")).
			Once()
		controller := controllers.NewSection(mockService)

		requestBody, _ := json.Marshal(bodyUpdate)

		r := testutil.SetUpRouter()
		r.PATCH(EndpointSection+"/:id", controller.UpdateCurrentCapacity())
		response := testutil.ExecuteTestRequest(r, http.MethodPatch, EndpointSection+"/1", requestBody)

		assert.Equal(t, http.StatusNotFound, response.Code)
		assert.JSONEq(t, "{\"code\":404,\"message\":\"section not found\"}", response.Body.String())
	})

	t.Run("update_invalid_id_parse_error: when section id is not parsed, should return code 400", func(t *testing.T) {
		controller := controllers.NewSection(nil)

		r := testutil.SetUpRouter()
		r.PATCH(EndpointSection+"/:id", controller.UpdateCurrentCapacity())
		response := testutil.ExecuteTestRequest(r, http.MethodPatch, EndpointSection+"/idInvalid", []byte{})

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.JSONEq(t, "{\"code\":400,\"message\":\"strconv.ParseInt: parsing \\\"idInvalid\\\": invalid syntax\"}", response.Body.String())
	})

	t.Run("update_invalid_field_value: when the field is negative,should return code 400", func(t *testing.T) {
		controller := controllers.NewSection(nil)

		requestBody, _ := json.Marshal(bodyFailSectionUpdate)

		r := testutil.SetUpRouter()
		r.PATCH(EndpointSection+"/:id", controller.UpdateCurrentCapacity())
		response := testutil.ExecuteTestRequest(r, http.MethodPatch, EndpointSection+"/1", requestBody)

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.JSONEq(t, "{\"code\":400,\"message\":\"the field CurrentCapacity invalid\"}", response.Body.String())
	})

	t.Run("update_invalid_body: when the body is invalid, should return code 400", func(t *testing.T) {
		controller := controllers.NewSection(mockService)

		requestBody, _ := json.Marshal(bodyFailSection)

		r := testutil.SetUpRouter()
		r.PATCH(EndpointSection+"/:id", controller.UpdateCurrentCapacity())
		response := testutil.ExecuteTestRequest(r, http.MethodPatch, EndpointSection+"/1", requestBody)

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.JSONEq(t, "{\"code\":400,\"message\":\"Key: 'requestSectionPatch.CurrentCapacity' Error:Field validation for 'CurrentCapacity' failed on the 'required' tag\"}", response.Body.String())
	})
}

func TestSectionController_Delete(t *testing.T) {
	mockService := mocks.NewSectionService(t)

	t.Run("delete_ok: when the request is successful, should return code 204", func(t *testing.T) {
		mockService.
			On("Delete", ctx, int64(1)).
			Return(nil).
			Once()
		controller := controllers.NewSection(mockService)

		r := testutil.SetUpRouter()
		r.DELETE(EndpointSection+"/:id", controller.Delete())
		response := testutil.ExecuteTestRequest(r, http.MethodDelete, EndpointSection+"/1", []byte{})

		assert.Equal(t, http.StatusNoContent, response.Code)
	})

	t.Run("delete_non_existent: when the section does not exist, should return code 404", func(t *testing.T) {
		mockService.
			On("Delete", ctx, int64(1)).
			Return(fmt.Errorf("section not found")).
			Once()

		controller := controllers.NewSection(mockService)

		r := testutil.SetUpRouter()
		r.DELETE(EndpointSection+"/:id", controller.Delete())
		response := testutil.ExecuteTestRequest(r, http.MethodDelete, EndpointSection+"/1", []byte{})

		assert.Equal(t, http.StatusNotFound, response.Code)
		assert.JSONEq(t, "{\"code\":404,\"message\":\"section not found\"}", response.Body.String())
	})

	t.Run("delete_id_parse_error: when section id is not parsed, should return code 400", func(t *testing.T) {
		controller := controllers.NewSection(nil)

		r := testutil.SetUpRouter()
		r.DELETE(EndpointSection+"/:id", controller.Delete())
		response := testutil.ExecuteTestRequest(r, http.MethodDelete, EndpointSection+"/idInvalid", []byte{})

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.JSONEq(t, "{\"code\":400,\"message\":\"strconv.ParseInt: parsing \\\"idInvalid\\\": invalid syntax\"}", response.Body.String())
	})
}

func TestSectionController_GetReportProductsBySection(t *testing.T) {
	EndpointReportProducts := EndpointSection + "/reportProducts"
	mockService := mocks.NewSectionService(t)
	controller := controllers.NewSection(mockService)

	r := testutil.SetUpRouter()
	r.GET(EndpointReportProducts, controller.GetReportProductsBySection())

	t.Run("get_all_report_product_records: when data entry is successful, should return code 200", func(t *testing.T) {
		mockService.
			On("GetAllProductCountBySection", ctx).
			Return(&[]domain.ReportProductsModel{expectedRecordProductBySection}, nil).
			Once()

		requestBody, _ := json.Marshal(expectedRecordProductBySection)
		response := testutil.ExecuteTestRequest(r, http.MethodGet, EndpointReportProducts, requestBody)

		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("get_all_report_product_records: when GetAll fail, should return code 500", func(t *testing.T) {
		mockService.
			On("GetAllProductCountBySection", ctx).
			Return(nil, fmt.Errorf("any error")).
			Once()

		response := testutil.ExecuteTestRequest(r, http.MethodGet, EndpointReportProducts, nil)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
		assert.JSONEq(t, "{\"code\":500,\"message\":\"any error\"}", response.Body.String())
	})

	t.Run("get_by_product_count_by_section: when data entry is successful, should return code 200", func(t *testing.T) {
		mockService.
			On("GetByIdProductCountBySection", ctx, int64(1)).
			Return(&expectedRecordProductBySection, nil).
			Once()

		requestBody, _ := json.Marshal(expectedRecordProductBySection)

		response := testutil.ExecuteTestRequest(r, http.MethodGet, EndpointReportProducts+"?id=1", requestBody)

		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("get_by_product_count_by_section: when the section does not exist, should return code 404", func(t *testing.T) {
		mockService.
			On("GetByIdProductCountBySection", ctx, int64(1)).
			Return(nil, fmt.Errorf("section not found")).
			Once()

		requestBody, _ := json.Marshal(expectedRecordProductBySection)

		response := testutil.ExecuteTestRequest(r, http.MethodGet, EndpointReportProducts+"?id=1", requestBody)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	t.Run("get_by_product_count_by_section_invalid_query_params: when the query params are not valid, should return code 400.", func(t *testing.T) {
		response := testutil.ExecuteTestRequest(r, http.MethodGet, EndpointReportProducts+"?id=abc", []byte{})

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.Equal(t, "{\"error\":\"invalid id\"}", response.Body.String())
	})
}
