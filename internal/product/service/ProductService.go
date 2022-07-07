package service

import (
	"context"
	"fmt"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/product/domain"
	productRecordsRepo "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/product_records/domain"
)

type productService struct {
	productRepository        domain.ProductRepository
	productRecordsRepository productRecordsRepo.ProductRecordsRepository
}

func CreateProductService(productRepository domain.ProductRepository, productRecordsRepository productRecordsRepo.ProductRecordsRepository) domain.ProductService {
	return &productService{
		productRepository:        productRepository,
		productRecordsRepository: productRecordsRepository,
	}
}

func (s *productService) GetAll(ctx context.Context) (*[]domain.Product, error) {
	products, err := s.productRepository.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	return products, nil
}

func (s *productService) GetById(ctx context.Context, id int64) (*domain.Product, error) {
	product, err := s.productRepository.GetById(ctx, id)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *productService) Create(ctx context.Context, product *domain.Product) (*domain.Product, error) {

	newProduct, err := s.productRepository.Create(ctx, product)

	if err != nil {
		return nil, err
	}

	return newProduct, nil
}

func (s *productService) UpdateDescription(ctx context.Context, id int64, description string) (*domain.Product, error) {

	productCurrent, err := s.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	productCurrent.Description = description

	productUpdate, err := s.productRepository.UpdateDescription(ctx, productCurrent)
	if err != nil {
		return nil, err
	}

	return productUpdate, nil
}

func (s *productService) Delete(ctx context.Context, id int64) error {

	err := s.productRepository.Delete(ctx, id)

	if err != nil {
		return err
	}

	return nil
}

func (s *productService) GetReportProductRecords(ctx context.Context, id int64) (*[]domain.ProductRecordsReport, error) {
	product, err := s.productRepository.GetById(ctx, id)

	if err != nil {
		return nil, err
	}

	result, err := s.productRecordsRepository.CountByProductId(ctx, product.Id)
	if err != nil {
		return nil, err
	}

	report := make([]domain.ProductRecordsReport, 0)

	productRecordsReport := domain.ProductRecordsReport{
		Id:                  product.Id,
		Description:         product.Description,
		CountProductRecords: result,
	}
	report = append(report, productRecordsReport)

	fmt.Println(result)
	return &report, nil
}

func (s *productService) GetAllReportProductRecords(ctx context.Context) (*[]domain.ProductRecordsReport, error) {

	products, err := s.productRepository.GetAllReportProductRecords(ctx)

	if err != nil {
		return nil, err
	}

	return products, nil
}
