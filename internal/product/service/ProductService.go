package service

import (
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/product/domain"
)

type service struct {
	repository domain.ProductRepository
}

func CreateProductService(r domain.ProductRepository) domain.ProductService {
	return &service{repository: r}
}

func (s *service) GetAll() ([]domain.Product, error) {
	products, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *service) GetById(id int64) (*domain.Product, error) {
	product, err := s.repository.GetById(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *service) Create(productCode string, description string, width float64, height float64, length float64, netWeight float64,
	expirationRate float64, recommendedFreezingTemperature float64, freezingRate int, productTypeId int, sellerId int) (*domain.Product, error) {

	newProduct, err := s.repository.
		Create(productCode, description, width, height, length, netWeight, expirationRate,
			recommendedFreezingTemperature, freezingRate, productTypeId, sellerId)

	if err != nil {
		return nil, err
	}

	return newProduct, nil
}

func (s *service) UpdateDescription(id int64, description string) (*domain.Product, error) {
	productUpdate, err := s.repository.UpdateDescription(id, description)

	if err != nil {
		return nil, err
	}

	return productUpdate, nil
}

func (s *service) Delete(id int64) error {
	err := s.repository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
