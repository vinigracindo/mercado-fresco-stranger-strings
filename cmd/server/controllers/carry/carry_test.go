package controllers_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	controllers "github.com/vinigracindo/mercado-fresco-stranger-strings/cmd/server/controllers/carry"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/carry/domain"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/carry/domain/mocks"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/pkg/testutil"
)

const endPointCarries = "/api/v1/carries"

func Test_Controller_Warehouse_CreateWarehouse(t *testing.T) {

	body := &controllers.RequestCarryPost{
		Cid:         1,
		CompanyName: "Mercado Livre",
		Address:     "Avenida Teste",
		Telephone:   "31 999999999",
		LocalityID:  1,
	}

	mockCarry := &domain.CarryModel{
		Cid:         1,
		CompanyName: "Mercado Livre",
		Address:     "Avenida Teste",
		Telephone:   "31 999999999",
		LocalityID:  1,
	}

	expect := &domain.CarryModel{
		Id:          1,
		Cid:         1,
		CompanyName: "Mercado Livre",
		Address:     "Avenida Teste",
		Telephone:   "31 999999999",
		LocalityID:  1,
	}

	service := mocks.NewCarryService(t)

	requestBody, _ := json.Marshal(body)

	t.Run("create_ok: if carry was successfully created", func(t *testing.T) {

		service.On("Create",
			context.TODO(),
			mockCarry,
		).Return(expect, nil).Once()

		controller := controllers.NewCarryController(service)

		r := testutil.SetUpRouter()

		r.POST(endPointCarries, controller.CreateCarry())

		response := testutil.ExecuteTestRequest(r, http.MethodPost, endPointCarries, requestBody)

		expect := map[string]interface{}{
			"data": expect,
		}

		assert.Equal(t, http.StatusCreated, response.Code)
		assert.JSONEq(t, testutil.StringJSON(expect), response.Body.String())
	})

	t.Run("create_fail: return 409, because the is already an warehouse with that code", func(t *testing.T) {

		service.On("Create",
			context.TODO(),
			mockCarry,
		).Return(nil, fmt.Errorf("error: already a warehouse with the code: %d", mockCarry.Cid)).Once()

		controller := controllers.NewCarryController(service)

		r := testutil.SetUpRouter()

		r.POST(endPointCarries, controller.CreateCarry())

		response := testutil.ExecuteTestRequest(r, http.MethodPost, endPointCarries, requestBody)

		assert.Equal(t, http.StatusConflict, response.Code)
	})

	t.Run("create_fail: when json object do not have all necessary fields, return 422 code", func(t *testing.T) {

		controller := controllers.NewCarryController(nil)

		r := testutil.SetUpRouter()

		r.POST(endPointCarries, controller.CreateCarry())

		response := testutil.ExecuteTestRequest(r, http.MethodPost, endPointCarries, []byte{})

		assert.Equal(t, http.StatusUnprocessableEntity, response.Code)
	})
}
