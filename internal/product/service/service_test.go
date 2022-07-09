package service_test

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/product/domain"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/product/domain/mocks"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/product/service"

	"testing"

	mocksProductRecords "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/product_records/domain/mocks"
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

var expectedReportProductRecords = domain.ProductRecordsReport{
	Id:                  1,
	Description:         "Yogurt",
	CountProductRecords: 5,
}

func TestProductService_Create(t *testing.T) {
	mockProductRepository := mocks.NewProductRepository(t)
	mockRepositoryProductRecords := mocksProductRecords.NewProductRecordsRepository(t)

	productService := service.CreateProductService(mockProductRepository, mockRepositoryProductRecords)

	t.Run("create_ok: when it contains the mandatory fields, should create a product", func(t *testing.T) {

		mockProductRepository.
			On("Create", context.TODO(), &expectedProduct).
			Return(&expectedProduct, nil).
			Once()

		prod, err := productService.Create(context.TODO(), &expectedProduct)

		assert.Nil(t, err)
		assert.Equal(t, prod, &expectedProduct)
	})

	t.Run("create_conflict: when product_code already exists, should not create a product", func(t *testing.T) {

		mockProductRepository.
			On("Create", context.TODO(), &expectedProduct).
			Return(nil, fmt.Errorf("the product code has already been registered")).
			Once()

		expectedProduct, err := productService.Create(context.TODO(), &expectedProduct)

		assert.NotNil(t, err)
		assert.Nil(t, expectedProduct)
		assert.Equal(t, err.Error(), "the product code has already been registered")
	})
}

func TestProductService_GetAll(t *testing.T) {
	mockProductRepository := mocks.NewProductRepository(t)
	mockRepositoryProductRecords := mocksProductRecords.NewProductRecordsRepository(t)

	productService := service.CreateProductService(mockProductRepository, mockRepositoryProductRecords)

	t.Run("get_all: when exists products, should return a list", func(t *testing.T) {

		expectedProductList := &[]domain.Product{expectedProduct, expectedProduct}

		mockProductRepository.
			On("GetAll", context.TODO()).
			Return(expectedProductList, nil).
			Once()

		productList, err := productService.GetAll(context.TODO())

		assert.Nil(t, err)
		assert.Equal(t, expectedProductList, productList)

	})

	t.Run("get_all_error: should return any error", func(t *testing.T) {

		mockProductRepository.
			On("GetAll", context.TODO()).
			Return(&[]domain.Product{}, fmt.Errorf("error: products not found")).
			Once()

		_, err := productService.GetAll(context.TODO())

		assert.NotNil(t, err)
	})
}

func TestProductService_GetById(t *testing.T) {
	mockProductRepository := mocks.NewProductRepository(t)
	mockRepositoryProductRecords := mocksProductRecords.NewProductRecordsRepository(t)

	productService := service.CreateProductService(mockProductRepository, mockRepositoryProductRecords)

	t.Run("find_by_id_non_existent: when the element searched for by id does not exist, should return an error", func(t *testing.T) {

		mockProductRepository.
			On("GetById", context.TODO(), int64(1)).
			Return(nil, fmt.Errorf("the product id was not found")).
			Once()

		prod, err := productService.GetById(context.TODO(), int64(1))

		assert.Nil(t, prod)
		assert.NotNil(t, err)
	})

	t.Run("find_by_id_existent: when element searched for by id exists, should return a product", func(t *testing.T) {

		mockProductRepository.
			On("GetById", context.TODO(), int64(1)).
			Return(&expectedProduct, nil).
			Once()

		resultProduct, err := productService.GetById(context.TODO(), 1)

		assert.Nil(t, err)
		assert.Equal(t, &expectedProduct, resultProduct)
	})

}

func TestProductService_UpdateDescription(t *testing.T) {

	mockProductRepository := mocks.NewProductRepository(t)
	mockRepositoryProductRecords := mocksProductRecords.NewProductRecordsRepository(t)

	productService := service.CreateProductService(mockProductRepository, mockRepositoryProductRecords)

	dummyUpdatedProduct := domain.Product{
		Id:          expectedProduct.Id,
		Description: "Strawberry yogurt",
	}

	t.Run("update_existent: when the data update is successful, should return the updated product", func(t *testing.T) {

		mockProductRepository.
			On("GetById", context.TODO(), int64(1)).
			Return(&expectedProduct, nil).
			Once()

		mockProductRepository.
			On("UpdateDescription", context.TODO(), &expectedProduct).
			Return(&expectedProduct, nil).
			Once()

		prod, err := productService.UpdateDescription(context.TODO(), expectedProduct.Id, dummyUpdatedProduct.Description)

		assert.Nil(t, err)
		assert.Equal(t, &expectedProduct, prod)
	})

	t.Run("update_error: should return any error", func(t *testing.T) {

		mockProductRepository.
			On("GetById", context.TODO(), int64(1)).
			Return(&expectedProduct, nil).
			Once()

		mockProductRepository.
			On("UpdateDescription", context.TODO(), &expectedProduct).
			Return(nil, fmt.Errorf("error: any error")).
			Once()

		_, err := productService.UpdateDescription(context.TODO(), expectedProduct.Id, dummyUpdatedProduct.Description)

		assert.Error(t, err)
	})

	t.Run("update_non_existent: when the element searched for by id does not exist, should return an error", func(t *testing.T) {

		mockProductRepository.
			On("GetById", context.TODO(), int64(1)).
			Return(nil, fmt.Errorf("the product id was not found")).
			Once()

		prod, err := productService.UpdateDescription(context.TODO(), expectedProduct.Id, dummyUpdatedProduct.Description)

		assert.Nil(t, prod)
		assert.NotNil(t, err)
	})
}

func TestProductService_Delete(t *testing.T) {
	mockProductRepository := mocks.NewProductRepository(t)
	mockRepositoryProductRecords := mocksProductRecords.NewProductRecordsRepository(t)

	productService := service.CreateProductService(mockProductRepository, mockRepositoryProductRecords)

	t.Run("delete_non_existent: when the product does not exist, should return an error", func(t *testing.T) {

		mockProductRepository.
			On("Delete", context.TODO(), int64(1)).
			Return(fmt.Errorf("product was not found")).
			Once()

		err := productService.Delete(context.TODO(), int64(1))

		assert.NotNil(t, err)
	})

	t.Run("delete_ok: when the section exist, should delete a product", func(t *testing.T) {

		mockProductRepository.
			On("Delete", context.TODO(), int64(1)).
			Return(nil).
			Once()

		err := productService.Delete(context.TODO(), int64(1))

		assert.Nil(t, err)
	})
}

func TestProductService_GetReportProductRecordsById(t *testing.T) {

	mockProductRepository := mocks.NewProductRepository(t)
	mockRepositoryProductRecords := mocksProductRecords.NewProductRecordsRepository(t)

	productService := service.CreateProductService(mockProductRepository, mockRepositoryProductRecords)

	var expectedReportProductRecordsList = []domain.ProductRecordsReport{
		{
			Id:                  1,
			Description:         "Yogurt",
			CountProductRecords: 5,
		},
	}
	product := domain.Product{
		Id:          1,
		Description: "Yogurt",
	}

	t.Run("get_by_id_report_existent: when element searched for by id exists, should return number of records for a particular product", func(t *testing.T) {

		mockProductRepository.
			On("GetById", context.TODO(), product.Id).
			Return(&product, nil).
			Once()

		mockRepositoryProductRecords.On("CountByProductId", context.TODO(), product.Id).
			Return(int64(5), nil).
			Once()

		resultProductRecords, err := productService.GetReportProductRecordsById(context.TODO(), product.Id)

		assert.Nil(t, err)
		assert.Equal(t, resultProductRecords, &expectedReportProductRecordsList)
	})

	t.Run("find_by_id_non_existent: when the element searched for by id does not exist, should return an error", func(t *testing.T) {

		mockProductRepository.
			On("GetById", context.TODO(), product.Id).
			Return(nil, fmt.Errorf("the product id was not found")).
			Once()

		resultProductRecords, err := productService.GetReportProductRecordsById(context.TODO(), product.Id)

		assert.Nil(t, resultProductRecords)
		assert.NotNil(t, err)
	})

	t.Run("get_by_id_report_error: return an error when the service fails", func(t *testing.T) {

		mockProductRepository.
			On("GetById", context.TODO(), product.Id).
			Return(&product, nil).
			Once()

		mockRepositoryProductRecords.On("CountByProductId", context.TODO(), product.Id).
			Return(int64(-1), fmt.Errorf("error")).
			Once()

		resultProductRecords, err := productService.GetReportProductRecordsById(context.TODO(), product.Id)

		assert.NotNil(t, err)
		assert.Empty(t, &resultProductRecords)
	})

}

func TestProductService_GetAllReportProductRecords(t *testing.T) {

	mockProductRepository := mocks.NewProductRepository(t)
	mockRepositoryProductRecords := mocksProductRecords.NewProductRecordsRepository(t)

	productService := service.CreateProductService(mockProductRepository, mockRepositoryProductRecords)

	t.Run("get_report_ok: should return a list of number of records of each product", func(t *testing.T) {

		expectedReportProductRecordsList := &[]domain.ProductRecordsReport{expectedReportProductRecords, expectedReportProductRecords}

		mockProductRepository.
			On("GetAllReportProductRecords", context.TODO()).
			Return(expectedReportProductRecordsList, nil).
			Once()

		productRecordsList, err := productService.GetAllReportProductRecords(context.TODO())

		assert.Nil(t, err)
		assert.Equal(t, expectedReportProductRecordsList, productRecordsList)
	})

	t.Run("get_report_error: should return any error", func(t *testing.T) {

		mockProductRepository.
			On("GetAllReportProductRecords", context.TODO()).
			Return(nil, fmt.Errorf("error")).
			Once()

		_, err := productService.GetAllReportProductRecords(context.TODO())

		assert.NotNil(t, err)
	})
}
