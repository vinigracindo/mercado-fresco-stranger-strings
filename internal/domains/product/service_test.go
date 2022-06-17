package product_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/product"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/product/mocks"
	"testing"
)

func TestProductServiceCreate(t *testing.T) {

	t.Run("create_ok: Se contiver os campos necessários, será criado", func(t *testing.T) {

		repo := mocks.NewRepository(t)

		repo.On("Create", "PROD02", "Yogurt", 1.2, 6.4, 4.5, 3.4, 1.5, 1.3, 2, 2, 2).Return(
			&product.Product{
				ProductCode:                    "PROD02",
				Description:                    "Yogurt",
				Width:                          1.2,
				Height:                         6.4,
				Length:                         4.5,
				NetWeight:                      3.4,
				ExpirationRate:                 1.5,
				RecommendedFreezingTemperature: 1.3,
				FreezingRate:                   2,
				ProductTypeId:                  2,
				SellerId:                       2,
			}, nil)
		service := product.CreateService(repo)

		expectedProduct, err := service.Create("PROD02", "Yogurt", 1.2, 6.4,
			4.5, 3.4, 1.5, 1.3, 2, 2, 2)

		assert.Nil(t, err)
		assert.Equal(t, expectedProduct.ProductCode, "PROD02")
		assert.Equal(t, expectedProduct.Description, "Yogurt")
		assert.Equal(t, expectedProduct.Width, 1.2)
		assert.Equal(t, expectedProduct.Height, 6.4)
		assert.Equal(t, expectedProduct.Length, 4.5)
		assert.Equal(t, expectedProduct.NetWeight, 3.4)
		assert.Equal(t, expectedProduct.ExpirationRate, 1.5)
		assert.Equal(t, expectedProduct.RecommendedFreezingTemperature, 1.3)
		assert.Equal(t, expectedProduct.FreezingRate, 2)
		assert.Equal(t, expectedProduct.ProductTypeId, 2)
		assert.Equal(t, expectedProduct.SellerId, 2)

		t.Run("create_conflict: Se o product_code já existir, ele não pode ser criado", func(t *testing.T) {
			repo := mocks.NewRepository(t)
			repo.On("Create", "PROD02", "Yogurt", 1.2, 6.4, 4.5, 3.4, 1.5, 1.3, 2, 2, 2).Return(
				nil, fmt.Errorf("the product with code PROD02 has already been registered"))

			service := product.CreateService(repo)

			expectedProduct, err := service.Create(
				"PROD02", "Yogurt", 1.2, 6.4, 4.5, 3.4, 1.5,
				1.3, 2, 2, 2)

			assert.NotNil(t, err)
			assert.Nil(t, expectedProduct)
			assert.Equal(t, err.Error(), "the product with code PROD02 has already been registered")
		})
	})
}
