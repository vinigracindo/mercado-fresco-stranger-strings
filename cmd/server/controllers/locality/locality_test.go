package locality_test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	controllers "github.com/vinigracindo/mercado-fresco-stranger-strings/cmd/server/controllers/locality"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/locality/domain"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/locality/domain/mocks"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/pkg/testutil"
)

const EndpointLocality = "/api/v1/locality"

var expectedLocality = domain.LocalityModel{
	Id:           1,
	LocalityName: "Salvador",
	ProvinceName: "Bahia",
	CountryName:  "Brasil",
	ProvinceId:   1,
}

var bodyLocality = domain.LocalityModel{
	LocalityName: "Salvador",
	ProvinceName: "Bahia",
	CountryName:  "Brasil",
	ProvinceId:   1,
}

var expectedReportSeller = []domain.ReportSeller{
	{
		LocalityId:   1,
		LocalityName: "Salvador",
		SellerCount:  1,
	},
}

var expectedBodyReportSeller = []domain.ReportSeller{
	{
		LocalityId:   1,
		LocalityName: "Salvador",
		SellerCount:  1,
	},
}

func Test_CreateLocalityController(t *testing.T) {
	service := mocks.NewLocalityService(t)
	ctx := context.Background()

	t.Run("create_ok: when data entry is successful, should return code 201.", func(t *testing.T) {
		service.
			On("CreateLocality", ctx, &bodyLocality).
			Return(&expectedLocality, nil).
			Once()

		controller := controllers.NewLocalityController(service)

		requestbodyLocality, _ := json.Marshal(expectedLocality)

		r := testutil.SetUpRouter()

		r.POST(EndpointLocality, controller.CreateLocality())

		response := testutil.ExecuteTestRequest(r, http.MethodPost, EndpointLocality, requestbodyLocality)

		assert.Equal(t, http.StatusCreated, response.Code)
	})

	t.Run("create_fail: when the JSON does not contain the required fields, should return code 422", func(t *testing.T) {
		controller := controllers.NewLocalityController(nil)

		r := testutil.SetUpRouter()
		r.POST(EndpointLocality, controller.CreateLocality())
		response := testutil.ExecuteTestRequest(r, http.MethodPost, EndpointLocality, []byte{})

		assert.Equal(t, http.StatusUnprocessableEntity, response.Code)
	})

	t.Run("create_conflict: when the id already exists, should return code 409.", func(t *testing.T) {
		service.
			On("CreateLocality", ctx, &bodyLocality).
			Return(nil, fmt.Errorf("locality with this id alredy exists")).
			Once()

		controller := controllers.NewLocalityController(service)
		requestbodySeller, _ := json.Marshal(bodyLocality)
		r := testutil.SetUpRouter()
		r.POST(EndpointLocality, controller.CreateLocality())

		response := testutil.ExecuteTestRequest(r, http.MethodPost, EndpointLocality, requestbodySeller)

		assert.Equal(t, http.StatusConflict, response.Code)

	})
}

func Test_GetReportLocalitiesController(t *testing.T) {
	service := mocks.NewLocalityService(t)
	controller := controllers.NewLocalityController(service)
	r := testutil.SetUpRouter()
	r.GET(EndpointLocality, controller.GetReportLocalities())

	ctx := context.Background()

	t.Run("report_get_by_id_ok: should return code 200", func(t *testing.T) {
		requestBody, _ := json.Marshal(expectedBodyReportSeller)

		service.
			On("GetByIdReportSeller", ctx, int64(1)).
			Return(&expectedReportSeller, nil).
			Once()

		response := testutil.ExecuteTestRequest(r, http.MethodGet, EndpointLocality+"?id=1", requestBody)

		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("get_by_id_non_exists: should return 404", func(t *testing.T) {
		service.
			On("GetByIdReportSeller", ctx, int64(9999)).
			Return(nil, errors.New("locality not found")).
			Once()

		response := testutil.ExecuteTestRequest(r, http.MethodGet, EndpointLocality+"?id=9999", []byte{})

		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	t.Run("invalid_query_params: when the query params are not valid, should return code 400.", func(t *testing.T) {
		response := testutil.ExecuteTestRequest(r, http.MethodGet, EndpointLocality+"?id=abc", []byte{})

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})

	t.Run("get_all_report_ok: should return code 200", func(t *testing.T) {
		requestBody, _ := json.Marshal(expectedBodyReportSeller)

		service.
			On("GetAllReportSeller", ctx).
			Return(&expectedReportSeller, nil).
			Once()

		response := testutil.ExecuteTestRequest(r, http.MethodGet, EndpointLocality, requestBody)

		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("get_all_internal_server_error: should return code 500", func(t *testing.T) {
		requestBody, _ := json.Marshal(expectedBodyReportSeller)

		service.
			On("GetAllReportSeller", ctx).
			Return(errors.New("request sent to server is invalid or corrupted")).
			Once()

		response := testutil.ExecuteTestRequest(r, http.MethodGet, EndpointLocality, requestBody)

		assert.Equal(t, http.StatusInternalServerError, response.Code)

	})
}
