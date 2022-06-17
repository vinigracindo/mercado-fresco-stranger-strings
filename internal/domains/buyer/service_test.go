package buyer_test

import (
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/buyer"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/buyer/mocks"
)

func Test_Service_Create(t *testing.T) {
	expectBuyer := buyer.Buyer{
		//Id:           1,
		CardNumberId: 123,
		FirstName:    "Jessica",
		LastName:     "Cruz",
	}

	t.Run("crete_ok: Se contiver os campos necessários, será criado", func(t *testing.T) {
		repo := mocks.NewRepository(t)

		repo.On("Create", int64(123), "Jessica", "Cruz").Return(expectBuyer, nil)

		service := buyer.NewService(repo)

		result, _ := service.Create(int64(123), "Jessica", "Cruz")

		assert.Equal(t, expectBuyer, result)
	})
}
