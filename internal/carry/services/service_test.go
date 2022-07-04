package services_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/carry/domain"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/carry/domain/mocks"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/carry/services"
)

var mockCarry *domain.CarryModel = &domain.CarryModel{
	Id:          1,
	Cid:         1,
	CompanyName: "Mercado Livre",
	Address:     "Avenida Teste",
	Telephone:   "31 999999999",
	LocalityID:  1,
}

func Test_service_create(t *testing.T) {
	repo := mocks.NewCarryRepository(t)

	t.Run("create_success: if all the fields are correct carry will be created", func(t *testing.T) {
		repo.On("Create",
			context.TODO(),
			mockCarry,
		).Return(mockCarry, nil).Once()

		service := services.NewCarryService(repo)

		result, err := service.Create(context.TODO(), mockCarry)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, result, mockCarry)
	})

	t.Run("error: repository return error", func(t *testing.T) {
		repo.On("Create",
			context.TODO(),
			mockCarry,
		).Return(nil, fmt.Errorf("error: invalid id")).Once()

		service := services.NewCarryService(repo)

		result, err := service.Create(context.TODO(), mockCarry)

		assert.Error(t, err)
		assert.Nil(t, result)
	})
}

func Test_service_getByID(t *testing.T) {
	repo := mocks.NewCarryRepository(t)

	t.Run("find_by_id_existent: search warehouses by id and return", func(t *testing.T) {
		repo.On(
			"GetById",
			context.TODO(),
			mockCarry.Id).
			Return(mockCarry, nil).
			Once()

		service := services.NewCarryService(repo)

		result, err := service.GetById(context.TODO(), mockCarry.Id)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, result, mockCarry)

	})

	t.Run("search_by_id_error: repository will retorn some error", func(t *testing.T) {
		repo.On(
			"GetById",
			context.TODO(),
			mockCarry.Id).
			Return(nil, fmt.Errorf("error: something when try to get the carry")).
			Once()

		service := services.NewCarryService(repo)

		result, err := service.GetById(context.TODO(), mockCarry.Id)

		assert.Error(t, err)
		assert.Nil(t, result)
	})
}
