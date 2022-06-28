package service

import (
	"context"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/product/domain"
)

type service struct {
	repository domain.ProductRepository
}

func CreateProductService(r domain.ProductRepository) domain.ProductService {
	return &service{repository: r}
}

func (s *service) GetAll(ctx context.Context) ([]domain.Product, error) {
	products, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *service) GetById(ctx context.Context, id int64) (*domain.Product, error) {
	product, err := s.repository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *service) Create(ctx context.Context, productCode string, description string, width float64, height float64, length float64, netWeight float64,
	expirationRate float64, recommendedFreezingTemperature float64, freezingRate float64, productTypeId int, sellerId int) (*domain.Product, error) {

	newProduct, err := s.repository.
		Create(ctx, productCode, description, width, height, length, netWeight, expirationRate,
			recommendedFreezingTemperature, freezingRate, productTypeId, sellerId)

	if err != nil {
		return nil, err
	}

	return newProduct, nil
}

func (s *service) UpdateDescription(ctx context.Context, id int64, description string) (*domain.Product, error) {
	productUpdate, err := s.repository.UpdateDescription(ctx, id, description)

	if err != nil {
		return nil, err
	}

	return productUpdate, nil
}

func (s *service) Delete(ctx context.Context, id int64) error {
	err := s.repository.Delete(ctx, id)

	if err != nil {
		return err
	}

	return nil
}
