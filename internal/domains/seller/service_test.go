package seller_test

import (
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

		result, _ := service.Create(123, "Mercado Livre", "Osasco, SP", "11 99999999")

		assert.Equal(t, expectedSeller, result)

	})
}
