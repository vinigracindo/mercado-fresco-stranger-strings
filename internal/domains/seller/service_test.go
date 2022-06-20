package seller_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/seller"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/seller/mocks"
)

func Test_Service_Creat(t *testing.T) {
	expectedSeller := seller.Seller{
		Id:          1,
		Cid:         123,
		CompanyName: "Mercado Livre",
		Address:     "Osasco, SP",
		Telephone:   "11 99999999",
	}

	t.Run("create_ok: Se contiver os campos necessários, será criado", func(t *testing.T) {
		repo := mocks.NewRepository(t)
		repo.On("Create", expectedSeller.Cid, expectedSeller.CompanyName, expectedSeller.Address, expectedSeller.Telephone).Return(expectedSeller, nil)
		service := seller.NewService(repo)

		result, err := service.Create(123, "Mercado Livre", "Osasco, SP", "11 99999999")

		assert.Nil(t, err)
		assert.Equal(t, expectedSeller, result)

	})

	t.Run("create_conflict: Se o cid já existir ele não pode ser criado", func(t *testing.T) {
		errMsg := fmt.Errorf("The seller whith cid %d has already been registered", expectedSeller.Cid)
		repo := mocks.NewRepository(t)
		repo.On("Create", expectedSeller.Cid, expectedSeller.CompanyName, expectedSeller.Address, expectedSeller.Telephone).Return(seller.Seller{}, errMsg)
		service := seller.NewService(repo)

		result, err := service.Create(123, "Mercado Livre", "Osasco, SP", "11 99999999")

		assert.Equal(t, seller.Seller{}, result)
		assert.Equal(t, errMsg, err)

	})
}

func Test_Service_GetAll(t *testing.T) {
	t.Run("find_all: Se a lista tiver 'n' elementos, retornará uma quantidade do total de elementos", func(t *testing.T) {
		expectedListSeller := []seller.Seller{
			{
				Id:          1,
				Cid:         123,
				CompanyName: "Mercado Livre",
				Address:     "Osasco, SP",
				Telephone:   "11 99999999",
			},
			{
				Id:          2,
				Cid:         1234,
				CompanyName: "Mercado Pago",
				Address:     "Salvador, BA",
				Telephone:   "11 88888888",
			},
		}
		repo := mocks.NewRepository(t)
		repo.On("GetAll").Return(expectedListSeller, nil)
		service := seller.NewService(repo)

		seller, err := service.GetAll()

		assert.Equal(t, seller, expectedListSeller)
		assert.Nil(t, err)

	})

}
