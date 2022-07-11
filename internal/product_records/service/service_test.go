package service

import (
	"context"
	"github.com/stretchr/testify/assert"
	mocksProduct "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/product/domain/mocks"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/product_records/domain"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/product_records/domain/mocks"

	"time"

	"testing"
)

var expectedProductRecords = domain.ProductRecords{
	Id:             1,
	LastUpdateDate: time.Now(),
	PurchasePrice:  10.5,
	SalePrice:      15.2,
	ProductId:      1,
}

func TestProductService_Create(t *testing.T) {

	mockRepositoryProductRecords := mocks.NewProductRecordsRepository(t)
	mockRepositoryProduct := mocksProduct.NewProductRepository(t)

	service := CreateProductRecordsService(mockRepositoryProductRecords, mockRepositoryProduct)

	t.Run("create_ok: when it contains the mandatory fields, should create a product records", func(t *testing.T) {

		mockRepositoryProduct.
			On("GetById", context.TODO(), int64(1)).
			Return(nil, nil).
			Once()

		mockRepositoryProductRecords.
			On("Create", context.TODO(), &expectedProductRecords).
			Return(&expectedProductRecords, nil).
			Once()

		productRecords, err := service.Create(context.TODO(), &expectedProductRecords)

		assert.Nil(t, err)
		assert.Equal(t, productRecords, &expectedProductRecords)
	})

	t.Run("create_product_does_not_exist: when product does not exist, should not create a product records", func(t *testing.T) {

		mockRepositoryProduct.
			On("GetById", context.TODO(), int64(1)).
			Return(nil, domain.ErrProductIdNotFound).
			Once()

		_, err := service.Create(context.TODO(), &expectedProductRecords)

		assert.Equal(t, domain.ErrProductIdNotFound, err)
		assert.Equal(t, nil, nil)
	})

	t.Run("create_error: when create product records fails, should return error", func(t *testing.T) {

		mockRepositoryProduct.
			On("GetById", context.TODO(), int64(1)).
			Return(nil, nil).
			Once()

		mockRepositoryProductRecords.
			On("Create", context.TODO(), &expectedProductRecords).
			Return(nil, domain.ErrProductIdNotFound).
			Once()

		_, err := service.Create(context.TODO(), &expectedProductRecords)

		assert.Equal(t, domain.ErrProductIdNotFound, err)
		assert.Equal(t, nil, nil)
	})

	t.Run("create_date_error: when the date is incorrect product records creation fails, should return error", func(t *testing.T) {

		mockRepositoryProduct.
			On("GetById", context.TODO(), int64(1)).
			Return(nil, nil).
			Once()

		date, err := time.Parse("2006-01-02", "2021-05-03")

		expectedDateProductRecords := domain.ProductRecords{
			Id:             1,
			LastUpdateDate: date,
			PurchasePrice:  10.5,
			SalePrice:      15.2,
			ProductId:      1,
		}
		_, err = service.Create(context.TODO(), &expectedDateProductRecords)

		assert.Equal(t, domain.ErrInvalidDate, err)
		assert.Equal(t, nil, nil)
	})

}
