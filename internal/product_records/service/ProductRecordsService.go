package service

import (
	"context"
	product "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/product/domain"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/product_records/domain"
)

type productRecordsService struct {
	repositoryProductRecords domain.ProductRecordsRepository
	repositoryProduct        product.ProductRepository
}

func CreateProductRecordsService(repositoryProductRecords domain.ProductRecordsRepository, repositoryProduct product.ProductRepository) domain.ProductRecordsService {
	return &productRecordsService{
		repositoryProductRecords: repositoryProductRecords,
		repositoryProduct:        repositoryProduct}
}

func (s productRecordsService) Create(ctx context.Context, productRecords *domain.ProductRecords) (*domain.ProductRecords, error) {

	_, err := s.repositoryProduct.GetById(ctx, productRecords.ProductId)

	if err != nil {
		return nil, product.ErrIDNotFound
	}

	newProductRecords, err := s.repositoryProductRecords.Create(ctx, productRecords)

	if err != nil {
		return nil, err
	}

	return newProductRecords, nil
}
