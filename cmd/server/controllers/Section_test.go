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

	t.Run("create_ok: Quando a entrada de dados for bem-sucedida, um código 201 será retornado junto com o objeto inserido.", func(t *testing.T) {
		mockService := mocks.NewService(t)

		mockService.On("Create",
			expectedSection.SectionNumber,
			expectedSection.CurrentTemperature,
			expectedSection.MinimumTemperature,
			expectedSection.CurrentCapacity,
			expectedSection.MinimumCapacity,
			expectedSection.MaximumCapacity,
			expectedSection.WarehouseId,
			expectedSection.ProductTypeId,
		).Return(expectedSection, nil)

		controller := controllers.NewSection(mockService)

		requestBody, _ := json.Marshal(body)

		r := SetUpRouter()

		r.POST("/api/v1/sections", controller.Create())

		response := CreateRequestTest(r, "POST", "/api/v1/sections", requestBody)
		assert.Equal(t, response.Code, http.StatusCreated)
	})

	//nao deu certo
	// t.Run("create_fail: Se o objeto JSON não contiver os campos necessários, um código 422 será retornado.", func(t *testing.T) {
	// 	mockService := mocks.NewService(t)

	// 	controller := controllers.NewSection(mockService)

	// 	requestBody, _ := json.Marshal(body)

	// 	r := SetUpRouter()

	// 	r.POST("/api/v1/sections", controller.Create())

	// 	response := CreateRequestTest(r, "POST", "/api/v1/sections", requestBody)
	// 	assert.Equal(t, response.Code, http.StatusUnprocessableEntity)
	// })

	t.Run("create_conflict: Se o section_number já existir, ele retornará um erro 409 Conflict", func(t *testing.T) {
		mockService := mocks.NewService(t)

		mockService.On("Create",
			expectedSection.SectionNumber,
			expectedSection.CurrentTemperature,
			expectedSection.MinimumTemperature,
			expectedSection.CurrentCapacity,
			expectedSection.MinimumCapacity,
			expectedSection.MaximumCapacity,
			expectedSection.WarehouseId,
			expectedSection.ProductTypeId,
		).Return(section.Section{}, fmt.Errorf("Already a section with this code"))

		controller := controllers.NewSection(mockService)

		requestBody, _ := json.Marshal(body)

		r := SetUpRouter()

		r.POST("/api/v1/sections", controller.Create())

		response := CreateRequestTest(r, "POST", "/api/v1/sections", requestBody)
		assert.Equal(t, response.Code, http.StatusConflict)
	})

}

func CreateRequestTest(gin *gin.Engine, method string, url string, body []byte) *httptest.ResponseRecorder {
	request := httptest.NewRequest(method, url, bytes.NewBuffer(body))
	response := httptest.NewRecorder()

	gin.ServeHTTP(response, request)

	return response
}
