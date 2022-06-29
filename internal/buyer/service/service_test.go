package service_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/buyer/domain"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/buyer/domain/mocks"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/buyer/service"
)

var expectBuyer = &domain.Buyer{
	Id:           1,
	CardNumberId: "402323",
	FirstName:    "FirstNameTest",
	LastName:     "LastNameTest",
}

var expectBuyerList = &[]domain.Buyer{
	{
		CardNumberId: "402323",
		FirstName:    "FirstNameTest",
		LastName:     "LastNameTest",
	},
	{
		CardNumberId: "402300",
		FirstName:    "FirstNameTest 2",
		LastName:     "LastNameTestTest2",
	},
}

var ctx = context.Background()

func TestService_Create(t *testing.T) {
	repo := mocks.NewBuyerRepository(t)
	service := service.NewBuyerService(repo)

	t.Run("crete_ok: when it contains the mandatory fields, should create a buyer", func(t *testing.T) {

		repo.
			On("Create",
				ctx,
				expectBuyer.CardNumberId,
				expectBuyer.FirstName,
				expectBuyer.LastName,
			).
			Return(expectBuyer, nil).
			Once()

		result, err := service.Create(ctx, "402323", "FirstNameTest", "LastNameTest")

		assert.Nil(t, err)
		assert.Equal(t, expectBuyer, result)
	})

	t.Run("create_conflict: when card_number_id already exists, should not create a buyer", func(t *testing.T) {
		errorConflict := fmt.Errorf("Card number id is not unique.")

		repo.
			On("Create",
				ctx,
				expectBuyer.CardNumberId,
				expectBuyer.FirstName,
				expectBuyer.LastName,
			).
			Return(&domain.Buyer{}, errorConflict).
			Once()

		result, err := service.Create(ctx, "402323", "FirstNameTest", "LastNameTest")

		assert.Equal(t, &domain.Buyer{}, result)
		assert.Equal(t, errorConflict, err)
	})
}

func TestService_GetAll(t *testing.T) {
	repo := mocks.NewBuyerRepository(t)
	service := service.NewBuyerService(repo)

	t.Run("find_all: when exists buyers, should return a list", func(t *testing.T) {

		repo.
			On("GetAll", ctx).
			Return(expectBuyerList, nil).
			Once()

		buyerList, _ := service.GetAll(ctx)

		assert.Equal(t, expectBuyerList, buyerList)
	})

	t.Run("get_all_error: should return any error", func(t *testing.T) {
		repo.On("GetAll", ctx).
			Return(expectBuyerList, fmt.Errorf("any error")).
			Once()

		_, err := service.GetAll(ctx)

		assert.NotNil(t, err)

	})
}

func TestService_GetId(t *testing.T) {
	repo := mocks.NewBuyerRepository(t)
	service := service.NewBuyerService(repo)

	t.Run("find_by_id_non_existent: when the element searched for by id does not exist, should return an error", func(t *testing.T) {
		errorNotFound := fmt.Errorf("Buyer not found.")
		repo.
			On("GetId", ctx, int64(3)).
			Return(nil, errorNotFound).
			Once()

		buyer, err := service.GetId(ctx, int64(3))

		assert.Nil(t, buyer)
		assert.NotNil(t, err)
	})

	t.Run("find_by_id_existent: when element searched for by id exists, should return a buyer", func(t *testing.T) {

		repo.
			On("GetId", ctx, int64(1)).
			Return(expectBuyer, nil).
			Once()

		buyer, err := service.GetId(ctx, int64(1))

		assert.Nil(t, err)
		assert.Equal(t, buyer, expectBuyer)

	})
}

func TestService_Update(t *testing.T) {
	repo := mocks.NewBuyerRepository(t)
	service := service.NewBuyerService(repo)

	t.Run("update_existent: when the data update is successful, should return the updated session", func(t *testing.T) {

		repo.
			On("Update", ctx, int64(1), "402300", "LastNameTest 2").
			Return(expectBuyer, nil).
			Once()

		buyer, err := service.Update(ctx, int64(1), "402300", "LastNameTest 2")

		assert.Equal(t, expectBuyer, buyer)
		assert.Nil(t, err)

	})

	t.Run("update_non_existent: when the element searched for by id does not exist, should return an error.", func(t *testing.T) {

		repo.
			On("Update", ctx, int64(1), "402300", "LastNameTest 2").
			Return(nil, fmt.Errorf("Buyer not found.")).
			Once()

		buyer, err := service.Update(ctx, int64(1), "402300", "LastNameTest 2")
		assert.Nil(t, buyer)
		assert.NotNil(t, err)
	})

}

func TestService_Delete(t *testing.T) {
	repo := mocks.NewBuyerRepository(t)
	service := service.NewBuyerService(repo)

	t.Run("delete_non_existent: when the buyer does not exist, should return an error.", func(t *testing.T) {

		repo.
			On("Delete", ctx, int64(1)).
			Return(fmt.Errorf("buyer not found.")).
			Once()

		err := service.Delete(ctx, int64(1))

		assert.NotNil(t, err)
	})

	t.Run("delete_ok: when the buyer exist, should delete a buyer.", func(t *testing.T) {

		repo.
			On("Delete", ctx, int64(1)).
			Return(nil).
			Once()

		err := service.Delete(ctx, int64(1))

		assert.Nil(t, err)
	})
}
