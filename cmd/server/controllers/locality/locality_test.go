package locality_test

import (
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/locality/domain"
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
	Id:           1,
	LocalityName: "Salvador",
	ProvinceName: "Bahia",
	CountryName:  "Brasil",
	ProvinceId:   1,
}

/*func Test_CreateLocalityController(t *testing.T) {
	service := mocks.NewLocalityService(t)
	ctx := context.Background()

	t.Run("create_ok: when data entry is successful, should return code 201.", func(t *testing.T) {
		service.
			On("CreateLocality", ctx, &expectedLocality).
			Return(&expectedLocality, nil).
			Once()

		controller := controllers.NewLocalityController(service)

		requestbodyLocality, _ := json.Marshal(expectedLocality)

		r := testutil.SetUpRouter()

		r.POST(EndpointLocality, controller.CreateLocality())

		response := testutil.ExecuteTestRequest(r, http.MethodPost, EndpointLocality, requestbodyLocality)

		assert.Equal(t, http.StatusCreated, response.Code)
	})
}*/
