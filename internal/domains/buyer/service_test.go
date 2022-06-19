package buyer_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/buyer"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/buyer/mocks"
)

func TestService_Create(t *testing.T) {
	expectBuyer := &buyer.Buyer{
		CardNumberId: 402323,
		FirstName:    "FirstNameTest",
		LastName:     "LastNameTest",
	}

	t.Run("crete_ok: Se contiver os campos necessários, será criado", func(t *testing.T) {
		repo := mocks.NewRepository(t)

		repo.On("Create", int64(402323), "FirstNameTest", "LastNameTest").Return(expectBuyer, nil)

		service := buyer.NewService(repo)

		result, _ := service.Create(int64(402323), "FirstNameTest", "LastNameTest")

		assert.Equal(t, expectBuyer, result)
	})

	t.Run("create_conflict: Se o card_number_id já existir, ele não pode ser criado", func(t *testing.T) {

		repo := mocks.NewRepository(t)

		repo.On("Create", int64(402323), "FirstNameTest", "LastNameTest").Return(expectBuyer, fmt.Errorf("Card number id is not unique."))

		service := buyer.NewService(repo)

		buyer, err := service.Create(int64(402323), "FirstNameTest", "LastNameTest")

		assert.NotNil(t, err)
		assert.Empty(t, buyer)
		assert.Equal(t, err.Error(), "Card number id is not unique.")

	})
}
