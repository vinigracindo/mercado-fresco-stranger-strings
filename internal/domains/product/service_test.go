package product_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/product"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/product/mocks"
	"testing"
)

func TestProductService_Create(t *testing.T) {

	t.Run("create_ok: Se contiver os campos necessários, será criado", func(t *testing.T) {
		expectedProduct := &product.Product{
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
		}

		repo := mocks.NewRepository(t)
		repo.On("Create", "PROD02", "Yogurt", 1.2, 6.4, 4.5, 3.4, 1.5, 1.3, 2, 2, 2).Return(expectedProduct, nil)
		service := product.CreateService(repo)

		prod, err := service.Create("PROD02", "Yogurt", 1.2, 6.4,
			4.5, 3.4, 1.5, 1.3, 2, 2, 2)

		assert.Nil(t, err)
		assert.Equal(t, prod, expectedProduct)

		t.Run("create_conflict: Se o product_code já existir, ele não pode ser criado", func(t *testing.T) {
			repo := mocks.NewRepository(t)
			repo.On("Create", "PROD02", "Yogurt", 1.2, 6.4, 4.5, 3.4, 1.5, 1.3, 2, 2, 2).
				Return(nil, fmt.Errorf("the product with code PROD02 has already been registered"))
			service := product.CreateService(repo)

			expectedProduct, err := service.Create("PROD02", "Yogurt", 1.2, 6.4, 4.5,
				3.4, 1.5, 1.3, 2, 2, 2)

			assert.NotNil(t, err)
			assert.Nil(t, expectedProduct)
			assert.Equal(t, err.Error(), "the product with code PROD02 has already been registered")
		})
	})
}

func TestProductService_GetAll(t *testing.T) {
	t.Run("find_all: Se a lista tiver 'n' elementos, retornará uma quantidade do total de elementos", func(t *testing.T) {
		expectedProduct := []product.Product{
			{
				Id:                             1,
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
			},
			{
				Id:                             2,
				ProductCode:                    "PROD03",
				Description:                    "Yogurt light",
				Width:                          1.5,
				Height:                         5.4,
				Length:                         3.5,
				NetWeight:                      4.4,
				ExpirationRate:                 1.8,
				RecommendedFreezingTemperature: 1.2,
				FreezingRate:                   2,
				ProductTypeId:                  3,
				SellerId:                       3,
			},
		}

		repo := mocks.NewRepository(t)
		repo.On("GetAll").Return(expectedProduct, nil)
		service := product.CreateService(repo)

		productList, err := service.GetAll()

		assert.Nil(t, err)
		assert.Equal(t, expectedProduct, productList)

	})

	t.Run("find_all_err: quando não encontrar todos os produtos, retornará um erro", func(t *testing.T) {
		repo := mocks.NewRepository(t)
		repo.On("GetAll").Return([]product.Product{}, fmt.Errorf("error: produtos não encontrados"))
		service := product.CreateService(repo)

		_, err := service.GetAll()

		assert.Error(t, err)
	})
}

func TestProductService_GetById(t *testing.T) {
	t.Run("find_by_id_non_existent: Se o elemento procurado por id não existir, retorna null", func(t *testing.T) {
		expectedProduct := []*product.Product{
			{
				Id:                             1,
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
			},
			{
				Id:                             2,
				ProductCode:                    "PROD03",
				Description:                    "Yogurt light",
				Width:                          1.5,
				Height:                         5.4,
				Length:                         3.5,
				NetWeight:                      4.4,
				ExpirationRate:                 1.8,
				RecommendedFreezingTemperature: 1.2,
				FreezingRate:                   2,
				ProductTypeId:                  3,
				SellerId:                       3,
			},
		}

		t.Run("find_by_id_existent: Se o elemento procurado por id existir, ele retornará as informações do elemento solicitado", func(t *testing.T) {
			repo := mocks.NewRepository(t)
			repo.On("GetById", int64(1)).Return(expectedProduct[1], nil)
			service := product.CreateService(repo)

			resultProduct, err := service.GetById(int64(1))

			assert.Nil(t, err)
			assert.Equal(t, expectedProduct[1], resultProduct)

		})

		repo := mocks.NewRepository(t)
		repo.On("GetById", int64(1)).Return(nil, fmt.Errorf("the product with the id was not found"))
		service := product.CreateService(repo)

		prod, err := service.GetById(int64(1))

		assert.Nil(t, prod)
		assert.NotNil(t, err)
	})
}

func TestProductService_UpdateDescription(t *testing.T) {

	t.Run("update_existent: Quando a atualização dos dados for bem sucedida, o produto será devolvido com as informações atualizadas", func(t *testing.T) {
		expectedProduct := product.Product{
			Id:                             1,
			ProductCode:                    "PROD02",
			Description:                    "Strawberry yogurt",
			Width:                          1.2,
			Height:                         6.4,
			Length:                         4.5,
			NetWeight:                      3.4,
			ExpirationRate:                 1.5,
			RecommendedFreezingTemperature: 1.3,
			FreezingRate:                   2,
			ProductTypeId:                  2,
			SellerId:                       2,
		}

		repo := mocks.NewRepository(t)
		repo.On("UpdateDescription", int64(1), "Strawberry yogurt").Return(expectedProduct, nil)
		service := product.CreateService(repo)

		prod, err := service.UpdateDescription(int64(1), "Strawberry yogurt")

		assert.Nil(t, err)
		assert.Equal(t, prod, expectedProduct)
	})

	t.Run("update_non_existent: Se o produto a ser atualizado não existir, será retornado null.", func(t *testing.T) {
		repo := mocks.NewRepository(t)
		repo.On("UpdateDescription", int64(1), "Strawberry yogurt").Return(nil, fmt.Errorf("product was not found"))
		service := product.CreateService(repo)

		prod, err := service.UpdateDescription(int64(1), "Strawberry yogurt")

		assert.Nil(t, prod)
		assert.NotNil(t, err)
	})
}
