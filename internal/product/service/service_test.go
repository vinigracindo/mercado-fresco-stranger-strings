package service_test

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/product/domain"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/product/domain/mocks"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/product/service"
	"testing"
)

var expectedProduct = domain.Product{
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
}

func TestProductService_Create(t *testing.T) {
	mockRepository := mocks.NewProductRepository(t)

	ctx, _ := context.WithCancel(context.Background())


	t.Run("create_ok: when it contains the mandatory fields, should create a product", func(t *testing.T) {
		mockRepository.
			On("Create", "PROD02", "Yogurt", 1.2, 6.4, 4.5, 3.4, 1.5, 1.3, 2, 2, 2).
			Return(&expectedProduct, nil).
			Once()

		service := service.CreateProductService(mockRepository)

		prod, err := service.Create(ctx, "PROD02", "Yogurt", , 6.4, 4.5, 3.4, 1.5, 1.3, 2, 2, 2, 2)

		assert.Nil(t, err)
		assert.Equal(t, prod, &expectedProduct)
	})

	t.Run("create_conflict: when product_code already exists, should not create a product", func(t *testing.T) {

		mockRepository.
			On("Create", "PROD02", "Yogurt", 1.2, 6.4, 4.5, 3.4, 1.5, 1.3, 2, 2, 2).
			Return(nil, fmt.Errorf("the product code has already been registered")).
			Once()

		service := service.CreateProductService(mockRepository)

		expectedProduct, err := service.Create(ctx, "PROD02", "Yogurt", 1.2, 6.4, 4.5,
			3.4, 1.5, 1.3, 2, 2, 2)

		assert.NotNil(t, err)
		assert.Nil(t, expectedProduct)
		assert.Equal(t, err.Error(), "the product code has already been registered")
	})
}

func TestProductService_GetAll(t *testing.T) {
	mockRepository := mocks.NewProductRepository(t)

	ctx, _ := context.WithCancel(context.Background())

	t.Run("get_all: when exists products, should return a list", func(t *testing.T) {
		expectedProductList := []domain.Product{expectedProduct, expectedProduct}

		mockRepository.
			On("GetAll").
			Return(expectedProductList, nil).
			Once()
		service := service.CreateProductService(mockRepository)

		productList, err := service.GetAll(ctx)

		assert.Nil(t, err)
		assert.Equal(t, expectedProductList, productList)

	})

	t.Run("get_all_error: should return any error", func(t *testing.T) {

		mockRepository.
			On("GetAll").
			Return([]domain.Product{}, fmt.Errorf("error: products not found")).
			Once()

		service := service.CreateProductService(mockRepository)

		_, err := service.GetAll(ctx)

		assert.NotNil(t, err)
	})
}

func TestProductService_GetById(t *testing.T) {
	mockRepository := mocks.NewProductRepository(t)
	ctx, _ := context.WithCancel(context.Background())

	t.Run("find_by_id_non_existent: when the element searched for by id does not exist, should return an error", func(t *testing.T) {

		mockRepository.
			On("GetById", int64(1)).
			Return(nil, fmt.Errorf("the product id was not found")).
			Once()

		service := service.CreateProductService(mockRepository)

		prod, err := service.GetById(ctx, int64(1))

		assert.Nil(t, prod)
		assert.NotNil(t, err)
	})

	t.Run("find_by_id_existent: when element searched for by id exists, should return a product", func(t *testing.T) {
		mockRepository.
			On("GetById", int64(1)).
			Return(&expectedProduct, nil).
			Once()

		service := service.CreateProductService(mockRepository)
		resultProduct, err := service.GetById(ctx, 1)

		assert.Nil(t, err)
		assert.Equal(t, &expectedProduct, resultProduct)
	})

}

func TestProductService_UpdateDescription(t *testing.T) {
	mockRepo := mocks.NewProductRepository(t)

	ctx, _ := context.WithCancel(context.Background())

	t.Run("update_existent: when the data update is successful, should return the updated product", func(t *testing.T) {
		expectedProduct := &domain.Product{
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

		service := service.CreateProductService(mockRepo)

		prod, err := service.UpdateDescription(ctx, int64(1), "Strawberry yogurt")

		assert.Nil(t, err)
		assert.Equal(t, prod, expectedProduct)
	})

	t.Run("update_non_existent: when the element searched for by id does not exist, should return an error", func(t *testing.T) {

		mockRepo.
			On("UpdateDescription", int64(1), "Strawberry yogurt").
			Return(nil, fmt.Errorf("product was not found")).
			Once()

		service := service.CreateProductService(mockRepo)

		prod, err := service.UpdateDescription(ctx, int64(1), "Strawberry yogurt")

		assert.Nil(t, prod)
		assert.NotNil(t, err)
	})
}

func TestProductService_Delete(t *testing.T) {
	mockRepo := mocks.NewProductRepository(t)

	ctx, _ := context.WithCancel(context.Background())

	t.Run("delete_non_existent: when the product does not exist, should return an error", func(t *testing.T) {

		mockRepo.
			On("Delete", int64(1)).
			Return(fmt.Errorf("product was not found")).
			Once()

		service := service.CreateProductService(mockRepo)

		err := service.Delete(ctx, int64(1))

		assert.NotNil(t, err)
	})

	t.Run("delete_ok: when the section exist, should delete a product", func(t *testing.T) {

		mockRepo.
			On("Delete", int64(1)).
			Return(nil).
			Once()

		service := service.CreateProductService(mockRepo)

		err := service.Delete(ctx, int64(1))

		assert.Nil(t, err)
	})
}
