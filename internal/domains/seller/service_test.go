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
