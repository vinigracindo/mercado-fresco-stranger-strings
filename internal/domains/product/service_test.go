package product_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/product"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/product/mocks"
	"testing"
)

func TestProductService_Create(t *testing.T) {
	mockRepo := mocks.NewProductRepository(t)

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

		mockRepo.
			On("Create", "PROD02", "Yogurt", 1.2, 6.4, 4.5, 3.4, 1.5, 1.3, 2, 2, 2).
			Return(expectedProduct, nil).
			Once()

		service := product.CreateService(mockRepo)

		prod, err := service.Create("PROD02", "Yogurt", 1.2, 6.4, 4.5, 3.4, 1.5, 1.3, 2, 2, 2)

		assert.Nil(t, err)
		assert.Equal(t, prod, expectedProduct)
	})

	t.Run("create_conflict: Se o product_code já existir, ele não pode ser criado", func(t *testing.T) {

		mockRepo.
			On("Create", "PROD02", "Yogurt", 1.2, 6.4, 4.5, 3.4, 1.5, 1.3, 2, 2, 2).
			Return(nil, fmt.Errorf("the product code has already been registered")).
			Once()

		service := product.CreateService(mockRepo)

		expectedProduct, err := service.Create("PROD02", "Yogurt", 1.2, 6.4, 4.5,
			3.4, 1.5, 1.3, 2, 2, 2)

		assert.NotNil(t, err)
		assert.Nil(t, expectedProduct)
		assert.Equal(t, err.Error(), "the product code has already been registered")
	})
}

func TestProductService_GetAll(t *testing.T) {
	mockRepo := mocks.NewProductRepository(t)

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

		mockRepo.
			On("GetAll").
			Return(expectedProduct, nil).
			Once()
		service := product.CreateService(mockRepo)

		productList, err := service.GetAll()

		assert.Nil(t, err)
		assert.Equal(t, expectedProduct, productList)

	})

	t.Run("find_all_err: quando não encontrar todos os produtos, retornará um erro", func(t *testing.T) {

		mockRepo.
			On("GetAll").
			Return([]product.Product{}, fmt.Errorf("error: products not found")).
			Once()

		service := product.CreateService(mockRepo)

		_, err := service.GetAll()

		assert.Error(t, err)
	})
}

func TestProductService_GetById(t *testing.T) {
	mockRepo := mocks.NewProductRepository(t)

	t.Run("find_by_id_non_existent: Se o elemento procurado por id não existir, retorna null", func(t *testing.T) {

		mockRepo.
			On("GetById", int64(1)).
			Return(nil, fmt.Errorf("the product id was not found")).
			Once()

		service := product.CreateService(mockRepo)

		prod, err := service.GetById(int64(1))

		assert.Nil(t, prod)
		assert.NotNil(t, err)
	})

	t.Run("find_by_id_existent: Se o elemento procurado por id existir, ele retornará as informações do elemento solicitado", func(t *testing.T) {
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

		mockRepo.
			On("GetById", int64(1)).
			Return(expectedProduct[1], nil).
			Once()

		service := product.CreateService(mockRepo)

		resultProduct, err := service.GetById(int64(1))

		assert.Nil(t, err)
		assert.Equal(t, expectedProduct[1], resultProduct)

	})
}

func TestProductService_UpdateDescription(t *testing.T) {
	mockRepo := mocks.NewProductRepository(t)

	t.Run("update_existent: Quando a atualização dos dados for bem sucedida, o produto será devolvido com as informações atualizadas", func(t *testing.T) {
		expectedProduct := &product.Product{
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

		mockRepo.
			On("UpdateDescription", int64(1), "Strawberry yogurt").
			Return(expectedProduct, nil).
			Once()

		service := product.CreateService(mockRepo)

		prod, err := service.UpdateDescription(int64(1), "Strawberry yogurt")

		assert.Nil(t, err)
		assert.Equal(t, prod, expectedProduct)
	})

	t.Run("update_non_existent: Se o produto a ser atualizado não existir, será retornado null.", func(t *testing.T) {

		mockRepo.
			On("UpdateDescription", int64(1), "Strawberry yogurt").
			Return(nil, fmt.Errorf("product was not found")).
			Once()

		service := product.CreateService(mockRepo)

		prod, err := service.UpdateDescription(int64(1), "Strawberry yogurt")

		assert.Nil(t, prod)
		assert.NotNil(t, err)
	})
}

func TestProductService_Delete(t *testing.T) {
	mockRepo := mocks.NewProductRepository(t)

	t.Run("delete_non_existent: Quando o produto não existe, null será retornado.", func(t *testing.T) {

		mockRepo.
			On("Delete", int64(1)).
			Return(fmt.Errorf("product was not found")).
			Once()

		service := product.CreateService(mockRepo)

		err := service.Delete(int64(1))

		assert.NotNil(t, err)
	})

	t.Run("delete_ok: Se a exclusão for bem-sucedida, o item não aparecerá na lista.", func(t *testing.T) {

		mockRepo.
			On("Delete", int64(1)).
			Return(nil).
			Once()

		service := product.CreateService(mockRepo)

		err := service.Delete(int64(1))

		assert.Nil(t, err)
	})
}
