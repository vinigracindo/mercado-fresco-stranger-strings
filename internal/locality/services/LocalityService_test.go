package services_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/locality/domain"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/locality/domain/mocks"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/locality/services"
	mocksSeller "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/seller/domain/mocks"
)

var expectedLocality = domain.LocalityModel{
	Id:           1,
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

var expectedReportCarrie = []domain.ReportCarrie{
	{
		LocalityId:   1,
		LocalityName: "Salvador",
		CarriesCount: 1,
	},
}

func Test_CreateLocalityService(t *testing.T) {
	repoLocality := mocks.NewLocalityRepository(t)
	repoSeller := mocksSeller.NewRepositorySeller(t)
	ctx := context.Background()

	t.Run("create_ok: when it contains the mandatory fields, should create a seller", func(t *testing.T) {
		repoLocality.
			On("CreateLocality", ctx, &expectedLocality).
			Return(&expectedLocality, nil).
			Once()

		service := services.NewLocalityService(repoLocality, repoSeller)

		result, err := service.CreateLocality(ctx, &expectedLocality)

		assert.Nil(t, err)
		assert.Equal(t, &expectedLocality, result)

	})

	t.Run("create_conflict: when id already exists, should not create a locality", func(t *testing.T) {
		errMsg := fmt.Errorf("The locality whith id %d has already been registered", expectedLocality.Id)

		repoLocality.
			On("CreateLocality", ctx, &expectedLocality).
			Return(nil, errMsg).
			Once()
		service := services.NewLocalityService(repoLocality, repoSeller)

		result, err := service.CreateLocality(ctx, &expectedLocality)

		assert.NotNil(t, err)
		assert.Nil(t, result)
		assert.Equal(t, err, errMsg)

	})
}

func Test_GetByIdReportSeller(t *testing.T) {
	repoLocality := mocks.NewLocalityRepository(t)
	repoSeller := mocksSeller.NewRepositorySeller(t)
	ctx := context.Background()

	t.Run("GetByIdReportSeller_ok: should return reportSeller", func(t *testing.T) {
		repoLocality.
			On("GetById", ctx, int64(1)).
			Return(&expectedLocality, nil).
			Once()

		repoSeller.
			On("CountByLocalityId", ctx, int64(1)).
			Return(int64(1), nil).
			Once()

		service := services.NewLocalityService(repoLocality, repoSeller)

		result, err := service.GetByIdReportSeller(ctx, int64(1))

		assert.NoError(t, err)
		assert.Equal(t, result, &expectedReportSeller)
	})

	t.Run("GetByIdReportSeller_id_non_exists: should return an error", func(t *testing.T) {
		repoLocality.
			On("GetById", ctx, int64(1)).
			Return(nil, errors.New("locality not found")).
			Once()

		service := services.NewLocalityService(repoLocality, repoSeller)

		result, err := service.GetByIdReportSeller(ctx, int64(1))

		assert.Error(t, err)
		assert.Nil(t, result)

	})

	t.Run("GetByIdReportSeller_CountByLocality_error: should return an error", func(t *testing.T) {
		repoLocality.
			On("GetById", ctx, int64(1)).
			Return(&expectedLocality, nil).
			Once()

		repoSeller.
			On("CountByLocalityId", ctx, int64(1)).
			Return(int64(0), errors.New("there ins't sellers in this locality")).
			Once()

		service := services.NewLocalityService(repoLocality, repoSeller)

		_, err := service.GetByIdReportSeller(ctx, int64(1))

		assert.Error(t, err)
	})
}

func Test_GetAllReportSeller(t *testing.T) {
	repoLocality := mocks.NewLocalityRepository(t)
	repoSeller := mocksSeller.NewRepositorySeller(t)
	ctx := context.Background()

	t.Run("GetAll_ok: should return reportSeller list", func(t *testing.T) {
		repoLocality.
			On("GetAllReportSeller", ctx).
			Return(&expectedReportSeller, nil).
			Once()

		service := services.NewLocalityService(repoLocality, repoSeller)

		result, err := service.GetAllReportSeller(ctx)

		assert.NoError(t, err)
		assert.Equal(t, result, &expectedReportSeller)
	})

	t.Run("GetAll_error: should return an error", func(t *testing.T) {
		repoLocality.
			On("GetAllReportSeller", ctx).
			Return(nil, errors.New("error")).
			Once()

		service := services.NewLocalityService(repoLocality, repoSeller)

		result, err := service.GetAllReportSeller(ctx)

		assert.Error(t, err)
		assert.Nil(t, result)
	})
}

func Test_ReportCarrie(t *testing.T) {
	repoLocality := mocks.NewLocalityRepository(t)
	repoSeller := mocksSeller.NewRepositorySeller(t)
	ctx := context.TODO()

	t.Run("ReportCarrie_ok: should return reportSeller list", func(t *testing.T) {
		repoLocality.
			On("ReportCarrie", ctx, int64(1)).
			Return(&expectedReportCarrie, nil).
			Once()

		service := services.NewLocalityService(repoLocality, repoSeller)

		result, err := service.ReportCarrie(ctx, int64(1))

		assert.NoError(t, err)
		assert.Equal(t, result, &expectedReportCarrie)
	})

	t.Run("ReportCarrie_error: should return an error", func(t *testing.T) {
		repoLocality.
			On("ReportCarrie", ctx, int64(1)).
			Return(nil, errors.New("error")).
			Once()

		service := services.NewLocalityService(repoLocality, repoSeller)

		result, err := service.ReportCarrie(ctx, int64(1))

		assert.Error(t, err)
		assert.Nil(t, result)
	})
}
