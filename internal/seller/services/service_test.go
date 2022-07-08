package services_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/seller/domain"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/seller/domain/mocks"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/seller/services"
)

var expectedSeller = domain.Seller{
	Id:          1,
	Cid:         123,
	CompanyName: "Mercado Livre",
	Address:     "Osasco, SP",
	Telephone:   "11 99999999",
	LocalityId:  1,
}

func Test_Service_Creat(t *testing.T) {
	repo := mocks.NewRepositorySeller(t)
	ctx := context.Background()

	t.Run("create_ok: when it contains the mandatory fields, should create a seller", func(t *testing.T) {
		repo.
			On("Create", ctx, &expectedSeller).
			Return(&expectedSeller, nil).
			Once()

		service := services.NewSellerService(repo)

		result, err := service.Create(ctx, &expectedSeller)

		assert.Nil(t, err)
		assert.Equal(t, &expectedSeller, result)

	})

	t.Run("create_conflict: when cid already exists, should not create a seller", func(t *testing.T) {
		errMsg := fmt.Errorf("The seller whith cid %d has already been registered", expectedSeller.Cid)

		repo.
			On("Create", ctx, &expectedSeller).
			Return(nil, errMsg).
			Once()
		service := services.NewSellerService(repo)

		result, err := service.Create(ctx, &expectedSeller)

		assert.NotNil(t, err)
		assert.Nil(t, result)
		assert.Equal(t, err, errMsg)

	})

}

func Test_Service_GetAll(t *testing.T) {

	t.Run("find_all: when exists sellers, should return a list", func(t *testing.T) {
		repo := mocks.NewRepositorySeller(t)
		ctx := context.Background()
		expectedListSeller := &[]domain.Seller{expectedSeller, expectedSeller}

		repo.
			On("GetAll", ctx).
			Return(expectedListSeller, nil).
			Once()

		service := services.NewSellerService(repo)

		sellers, err := service.GetAll(ctx)

		assert.Equal(t, sellers, expectedListSeller)
		assert.Nil(t, err)

	})

}

func Test_Service_GetById(t *testing.T) {
	repo := mocks.NewRepositorySeller(t)
	ctx := context.Background()

	t.Run("find_by_id_non_existent: when element searched for by id non exists, should return nil", func(t *testing.T) {
		repo.
			On("GetById", ctx, int64(1)).
			Return(nil, fmt.Errorf("Seller not found.")).
			Once()
		service := services.NewSellerService(repo)

		result, err := service.GetById(ctx, int64(1))

		assert.Nil(t, result)
		assert.NotNil(t, err)
	})

	t.Run("find_by_id_existent: when element searched for by id exists, should return a seller", func(t *testing.T) {

		repo.
			On("GetById", ctx, int64(1)).
			Return(&expectedSeller, nil).
			Once()
		service := services.NewSellerService(repo)

		result, err := service.GetById(ctx, int64(1))

		assert.Nil(t, err)
		assert.Equal(t, &expectedSeller, result)
	})
}

func Test_Service_Update(t *testing.T) {
	repo := mocks.NewRepositorySeller(t)
	ctx := context.Background()

	t.Run("update_ok: when the data update is successful, should return the updated seller", func(t *testing.T) {
		repo.
			On("GetById", ctx, int64(1)).
			Return(&expectedSeller, nil).
			Once()

		repo.
			On("Update", ctx, &expectedSeller).
			Return(&expectedSeller, nil).
			Once()

		service := services.NewSellerService(repo)

		result, err := service.Update(ctx, int64(1), "Salvador, BA", "11 98989898")

		assert.Equal(t, result, &expectedSeller)
		assert.Nil(t, err)

	})

	t.Run("update_non_existent: when element searched for by id non exists, should return nil", func(t *testing.T) {
		repo.
			On("GetById", ctx, int64(3)).
			Return(nil, fmt.Errorf("Seller not found.")).
			Once()

		service := services.NewSellerService(repo)

		result, err := service.Update(ctx, int64(3), "Salvador, BA", "11 98989898")

		assert.Nil(t, result)
		assert.NotNil(t, err)

	})
}

func Test_Service_Delete(t *testing.T) {
	repo := mocks.NewRepositorySeller(t)
	ctx := context.Background()

	t.Run("delete_ok: when the seller exists, should delete", func(t *testing.T) {
		repo.
			On("Delete", ctx, int64(1)).
			Return(nil).
			Once()
		service := services.NewSellerService(repo)

		err := service.Delete(ctx, int64(1))

		assert.Nil(t, err)

	})

	t.Run("delete_non_existent: when seller does not exist, should return nil", func(t *testing.T) {
		repo.
			On("Delete", ctx, int64(3)).
			Return(fmt.Errorf("Seller not found.")).
			Once()
		service := services.NewSellerService(repo)

		err := service.Delete(ctx, int64(3))

		assert.NotNil(t, err)

	})
}
