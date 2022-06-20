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
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/section"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/section/mocks"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

var body = section.Section{
	SectionNumber:      int64(1),
	CurrentTemperature: int64(1),
	MinimumTemperature: int64(1),
	CurrentCapacity:    int64(1),
	MinimumCapacity:    int64(1),
	MaximumCapacity:    int64(1),
	WarehouseId:        int64(1),
	ProductTypeId:      int64(1),
}

var bodyFail = section.Section{
	SectionNumber:      0,
	CurrentTemperature: 0,
	MinimumTemperature: 0,
	CurrentCapacity:    0,
	MinimumCapacity:    0,
	MaximumCapacity:    0,
	WarehouseId:        0,
	ProductTypeId:      0,
}

var expectedSection = section.Section{
	Id:                 int64(1),
	SectionNumber:      int64(1),
	CurrentTemperature: int64(1),
	MinimumTemperature: int64(1),
	CurrentCapacity:    int64(1),
	MinimumCapacity:    int64(1),
	MaximumCapacity:    int64(1),
	WarehouseId:        int64(1),
	ProductTypeId:      int64(1),
}

func TestSectionController_Create(t *testing.T) {
	mockService := mocks.NewService(t)

	t.Run("create_ok: when data entry is successful, should return code 201", func(t *testing.T) {
		mockService.
			On("Create",
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

		requestBody, _ := json.Marshal(body)

		r := SetUpRouter()

		r.POST("/api/v1/sections", controller.Create())

		response := CreateRequestTest(r, "POST", "/api/v1/sections", requestBody)
		assert.Equal(t, http.StatusCreated, response.Code)
	})

	t.Run("create_fail: when the JSON does not contain the required fields, should return code 422", func(t *testing.T) {
		controller := controllers.NewSection(nil)

		r := SetUpRouter()
		r.POST("/api/v1/sections", controller.Create())
		response := CreateRequestTest(r, "POST", "/api/v1/sections", []byte{})

		assert.Equal(t, http.StatusUnprocessableEntity, response.Code)
	})

	t.Run("create_conflict: when the section_number already exists, should return code 409", func(t *testing.T) {
		mockService.
			On("Create",
				expectedSection.SectionNumber,
				expectedSection.CurrentTemperature,
				expectedSection.MinimumTemperature,
				expectedSection.CurrentCapacity,
				expectedSection.MinimumCapacity,
				expectedSection.MaximumCapacity,
				expectedSection.WarehouseId,
				expectedSection.ProductTypeId,
			).
			Return(section.Section{}, fmt.Errorf("Already a section with this code")).
			Once()

		controller := controllers.NewSection(mockService)

		requestBody, _ := json.Marshal(body)

		r := SetUpRouter()

		r.POST("/api/v1/sections", controller.Create())

		response := CreateRequestTest(r, "POST", "/api/v1/sections", requestBody)
		assert.Equal(t, http.StatusConflict, response.Code)
	})
}

func TestSectionController_GetAll(t *testing.T) {
	mockService := mocks.NewService(t)

	t.Run("find_all: when data entry is successful, should return code 200", func(t *testing.T) {
		mockService.
			On("GetAll").
			Return([]section.Section{expectedSection}, nil).
			Once()
		controller := controllers.NewSection(mockService)

		requestBody, _ := json.Marshal(body)

		r := SetUpRouter()
		r.GET("/api/v1/sections", controller.GetAll())
		response := CreateRequestTest(r, "GET", "/api/v1/sections", requestBody)

		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("get_all_fail: when GetAll fail, should return code 400", func(t *testing.T) {
		mockService.
			On("GetAll").
			Return([]section.Section{}, fmt.Errorf("any error"))
		controller := controllers.NewSection(mockService)

		requestBody, _ := json.Marshal(body)

		r := SetUpRouter()
		r.GET("/api/v1/sections", controller.GetAll())
		response := CreateRequestTest(r, "GET", "/api/v1/sections", requestBody)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
}

func TestSectionController_GetById(t *testing.T) {
	mockService := mocks.NewService(t)

	t.Run("find_by_id_existent: when the request is successful, should return code 200", func(t *testing.T) {
		mockService.
			On("GetById", int64(1)).
			Return(expectedSection, nil)
		controller := controllers.NewSection(mockService)

		requestBody, _ := json.Marshal(body)

		r := SetUpRouter()
		r.GET("/api/v1/sections/:id", controller.GetById())
		response := CreateRequestTest(r, "GET", "/api/v1/sections/1", requestBody)

		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("find_by_id_non_existent: when the section does not exist, should return code 404", func(t *testing.T) {
		mockService := mocks.NewService(t)
		mockService.On("GetById", int64(1)).Return(section.Section{}, fmt.Errorf("Section not found"))
		controller := controllers.NewSection(mockService)

		requestBody, _ := json.Marshal(body)

		r := SetUpRouter()
		r.GET("/api/v1/sections/:id", controller.GetById())
		response := CreateRequestTest(r, "GET", "/api/v1/sections/1", requestBody)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	t.Run("find_by_id_parse_error: when section id is not parsed, should return code 400", func(t *testing.T) {
		controller := controllers.NewSection(nil)

		r := SetUpRouter()
		r.GET("/api/v1/sections/:id", controller.GetById())
		response := CreateRequestTest(r, "GET", "/api/v1/sections/idinvalid", []byte{})

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
}

func TestSectionController_Update(t *testing.T) {
	mockService := mocks.NewService(t)

	var bodyUpdate = section.Section{
		CurrentCapacity: int64(1),
	}

	var bodyFailUpdate = section.Section{
		CurrentCapacity: -1,
	}

	t.Run("update_ok: when the request is successful, should return code 200", func(t *testing.T) {
		mockService.
			On("UpdateCurrentCapacity", int64(1), int64(1)).
			Return(expectedSection, nil).
			Once()
		controller := controllers.NewSection(mockService)

		requestBody, _ := json.Marshal(bodyUpdate)

		r := SetUpRouter()
		r.PATCH("/api/v1/sections/:id", controller.UpdateCurrentCapacity())
		response := CreateRequestTest(r, "PATCH", "/api/v1/sections/1", requestBody)

		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("update_non_existent: when the section does not exist, should return code 404", func(t *testing.T) {
		mockService.
			On("UpdateCurrentCapacity", int64(1), int64(1)).
			Return(section.Section{}, fmt.Errorf("Section not found")).
			Once()
		controller := controllers.NewSection(mockService)

		requestBody, _ := json.Marshal(bodyUpdate)

		r := SetUpRouter()
		r.PATCH("/api/v1/sections/:id", controller.UpdateCurrentCapacity())
		response := CreateRequestTest(r, "PATCH", "/api/v1/sections/1", requestBody)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	t.Run("update_invalid_id_parse_error: when section id is not parsed, should return code 400", func(t *testing.T) {
		controller := controllers.NewSection(nil)

		r := SetUpRouter()
		r.PATCH("/api/v1/sections/:id", controller.UpdateCurrentCapacity())
		response := CreateRequestTest(r, "PATCH", "/api/v1/sections/idinvalid", []byte{})

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})

	t.Run("update_invalid_field_value: when the field is negative,should return code 400", func(t *testing.T) {
		controller := controllers.NewSection(nil)

		requestBody, _ := json.Marshal(bodyFailUpdate)

		r := SetUpRouter()
		r.PATCH("/api/v1/sections/:id", controller.UpdateCurrentCapacity())
		response := CreateRequestTest(r, "PATCH", "/api/v1/sections/1", requestBody)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})

	t.Run("update_invalid_body: when the body is invalid, should return code 400", func(t *testing.T) {
		controller := controllers.NewSection(mockService)

		requestBody, _ := json.Marshal(bodyFail)

		r := SetUpRouter()
		r.PATCH("/api/v1/sections/:id", controller.UpdateCurrentCapacity())
		response := CreateRequestTest(r, "PATCH", "/api/v1/sections/1", requestBody)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
}

func TestSectionController_Delete(t *testing.T) {
	mockService := mocks.NewService(t)

	t.Run("delete_ok: when the request is successful, should return code 204", func(t *testing.T) {
		mockService.
			On("Delete", int64(1)).
			Return(nil).
			Once()
		controller := controllers.NewSection(mockService)

		r := SetUpRouter()
		r.DELETE("/api/v1/sections/:id", controller.Delete())
		response := CreateRequestTest(r, "DELETE", "/api/v1/sections/1", []byte{})

		assert.Equal(t, http.StatusNoContent, response.Code)
	})

	t.Run("delete_non_existent: when the section does not exist, should return code 404", func(t *testing.T) {
		mockService.
			On("Delete", int64(1)).
			Return(fmt.Errorf("Section not found")).
			Once()

		controller := controllers.NewSection(mockService)

		r := SetUpRouter()
		r.DELETE("/api/v1/sections/:id", controller.Delete())
		response := CreateRequestTest(r, "DELETE", "/api/v1/sections/1", []byte{})

		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	t.Run("delete_id_parse_error: when section id is not parsed, should return code 400", func(t *testing.T) {
		controller := controllers.NewSection(nil)

		r := SetUpRouter()
		r.DELETE("/api/v1/sections/:id", controller.Delete())
		response := CreateRequestTest(r, "DELETE", "/api/v1/sections/idinvalid", []byte{})

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})

}

func CreateRequestTest(gin *gin.Engine, method string, url string, body []byte) *httptest.ResponseRecorder {
	request := httptest.NewRequest(method, url, bytes.NewBuffer(body))
	response := httptest.NewRecorder()

	gin.ServeHTTP(response, request)

	return response
}
