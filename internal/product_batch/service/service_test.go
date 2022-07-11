package service_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	productMocks "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/product/domain/mocks"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/product_batch/domain"
	productBatch "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/product_batch/domain/mocks"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/product_batch/service"
	section "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/section/domain"
	sectionMocks "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/section/domain/mocks"
)

var timeNow = time.Now()

var expectedProductBatch = domain.ProductBatch{
	Id:                 1,
	BatchNumber:        1,
	CurrentQuantity:    1,
	CurrentTemperature: 1.0,
	DueDate:            timeNow,
	InitialQuantity:    1.0,
	ManufacturingDate:  timeNow,
	ManufacturingHour:  1,
	MinumumTemperature: 1.0,
	ProductId:          1,
	SectionId:          1,
}

func TestProductBatchService_Create(t *testing.T) {
	errorNotFound := fmt.Errorf("not found")
	errorAny := fmt.Errorf("any error")

	mockRepositoryProductBatch := productBatch.NewProductBatchRepository(t)
	mockRepositoryProduct := productMocks.NewProductRepository(t)
	mockRepositorySection := sectionMocks.NewSectionRepository(t)

	service := service.NewProductBatchService(mockRepositoryProductBatch, mockRepositoryProduct, mockRepositorySection)

	t.Run("create_ok: when it contains the mandatory fields, should create a product batch", func(t *testing.T) {
		mockRepositorySection.
			On("GetById", context.TODO(), int64(1)).
			Return(section.SectionModel{}, nil).
			Once()

		mockRepositoryProduct.
			On("GetById", context.TODO(), int64(1)).
			Return(nil, nil).
			Once()

		mockRepositoryProductBatch.
			On("Create", context.TODO(), &expectedProductBatch).
			Return(&expectedProductBatch, nil).
			Once()

		productBatch, err := service.Create(context.TODO(), &expectedProductBatch)
		assert.Nil(t, err)
		assert.Equal(t, productBatch, &expectedProductBatch)
	})

	t.Run("create_product_does_not_exist: when product does not exist, should not create a product batch", func(t *testing.T) {
		mockRepositoryProduct.
			On("GetById", context.TODO(), int64(1)).
			Return(nil, errorNotFound).
			Once()

		_, err := service.Create(context.TODO(), &expectedProductBatch)

		assert.Equal(t, errorNotFound, err)
		assert.Equal(t, nil, nil)
	})

	t.Run("create_section_does_not_exist: when product does not exist, should not create a product batch", func(t *testing.T) {
		mockRepositorySection.
			On("GetById", context.TODO(), int64(1)).
			Return(section.SectionModel{}, errorNotFound).
			Once()

		mockRepositoryProduct.
			On("GetById", context.TODO(), int64(1)).
			Return(nil, nil).
			Once()

		_, err := service.Create(context.TODO(), &expectedProductBatch)

		assert.Equal(t, errorNotFound, err)
	})

	t.Run("create_error: when create product batch fails, should return error", func(t *testing.T) {
		mockRepositorySection.
			On("GetById", context.TODO(), int64(1)).
			Return(section.SectionModel{}, nil).
			Once()

		mockRepositoryProduct.
			On("GetById", context.TODO(), int64(1)).
			Return(nil, nil).
			Once()

		mockRepositoryProductBatch.
			On("Create", context.TODO(), &expectedProductBatch).
			Return(nil, errorAny).
			Once()

		_, err := service.Create(context.TODO(), &expectedProductBatch)

		assert.Equal(t, errorAny, err)
		assert.Equal(t, nil, nil)
	})

}
