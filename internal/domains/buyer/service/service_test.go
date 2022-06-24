package service_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/buyer/domain"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/buyer/domain/mocks"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/buyer/service"
)

var expectBuyer = &domain.Buyer{
	Id:           1,
	CardNumberId: 402323,
	FirstName:    "FirstNameTest",
	LastName:     "LastNameTest",
}

var expectBuyerList = []domain.Buyer{
	{
		CardNumberId: 402323,
		FirstName:    "FirstNameTest",
		LastName:     "LastNameTest",
	},
	{
		CardNumberId: 402300,
		FirstName:    "FirstNameTest 2",
		LastName:     "LastNameTestTest2",
	},
}

func TestService_Create(t *testing.T) {
	repo := mocks.NewBuyerRepository(t)

	t.Run("crete_ok: when it contains the mandatory fields, should create a buyer", func(t *testing.T) {

		repo.
			On("Create", int64(402323), "FirstNameTest", "LastNameTest").
			Return(expectBuyer, nil).
			Once()

		service := service.NewBuyerService(repo)

		result, _ := service.Create(int64(402323), "FirstNameTest", "LastNameTest")

		assert.Equal(t, expectBuyer, result)
	})

	t.Run("create_conflict: when card_number_id already exists, should not create a buyer", func(t *testing.T) {

		repo.
			On("Create", int64(402323), "FirstNameTest", "LastNameTest").
			Return(expectBuyer, fmt.Errorf("Card number id is not unique.")).
			Once()

		service := service.NewBuyerService(repo)

		buyer, err := service.Create(int64(402323), "FirstNameTest", "LastNameTest")

		assert.NotNil(t, err)
		assert.Empty(t, buyer)
		assert.Equal(t, err.Error(), "Card number id is not unique.")

	})
}

func TestService_GetAll(t *testing.T) {
	repo := mocks.NewBuyerRepository(t)
	t.Run("find_all: when exists buyers, should return a list", func(t *testing.T) {

		repo.
			On("GetAll").
			Return(expectBuyerList, nil).
			Once()

		service := service.NewBuyerService(repo)

		buyerList, _ := service.GetAll()

		assert.Equal(t, expectBuyerList, buyerList)
	})

	t.Run("get_all_error: should return any error", func(t *testing.T) {
		repo.On("GetAll").
			Return(expectBuyerList, fmt.Errorf("any error")).
			Once()

		service := service.NewBuyerService(repo)

		_, err := service.GetAll()

		assert.NotNil(t, err)

	})
}

func TestService_GetId(t *testing.T) {
	repo := mocks.NewBuyerRepository(t)

	t.Run("find_by_id_non_existent: when the element searched for by id does not exist, should return an error", func(t *testing.T) {

		repo.
			On("GetId", int64(4)).
			Return(nil, fmt.Errorf("Buyer not found.")).
			Once()

		service := service.NewBuyerService(repo)

		buyer, err := service.GetId(int64(4))

		assert.Nil(t, buyer)
		assert.NotNil(t, err)
	})

	t.Run("find_by_id_existent: when element searched for by id exists, should return a buyer", func(t *testing.T) {

		repo.
			On("GetId", int64(1)).
			Return(expectBuyer, nil).
			Once()

		service := service.NewBuyerService(repo)

		buyer, err := service.GetId(int64(1))

		assert.Nil(t, err)
		assert.Equal(t, buyer, expectBuyer)

	})
}

func TestService_Update(t *testing.T) {
	repo := mocks.NewBuyerRepository(t)

	t.Run("update_existent: when the data update is successful, should return the updated session", func(t *testing.T) {

		repo.
			On("Update", int64(1), int64(456), "LastNameTest 2").
			Return(expectBuyer, nil).
			Once()

		service := service.NewBuyerService(repo)

		buyer, err := service.Update(int64(1), int64(456), "LastNameTest 2")

		assert.Equal(t, expectBuyer, buyer)
		assert.Nil(t, err)

	})

	t.Run("update_non_existent: when the element searched for by id does not exist, should return an error.", func(t *testing.T) {

		repo.
			On("Update", int64(1), int64(456), "LastNameTest 2").
			Return(nil, fmt.Errorf("Buyer not found.")).
			Once()

		service := service.NewBuyerService(repo)

		buyer, err := service.Update(int64(1), int64(456), "LastNameTest 2")
		assert.Nil(t, buyer)
		assert.NotNil(t, err)
	})

}

func TestService_Delete(t *testing.T) {
	repo := mocks.NewBuyerRepository(t)

	t.Run("delete_non_existent: when the buyer does not exist, should return an error.", func(t *testing.T) {

		repo.
			On("Delete", int64(1)).
			Return(fmt.Errorf("buyer not found.")).
			Once()

		service := service.NewBuyerService(repo)

		err := service.Delete(int64(1))

		assert.NotNil(t, err)
	})

	t.Run("delete_ok: when the buyer exist, should delete a buyer.", func(t *testing.T) {

		repo.
			On("Delete", int64(1)).
			Return(nil).
			Once()

		service := service.NewBuyerService(repo)

		err := service.Delete(int64(1))

		assert.Nil(t, err)
	})
}
