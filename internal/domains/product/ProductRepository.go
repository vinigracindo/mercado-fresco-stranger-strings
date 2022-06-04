package product

import (
	"fmt"
)

var listProducts []Product

type Repository interface {
	GetAll() ([]Product, error)
	GetById(id int64) (*Product, error)
}

type repository struct {
}

func (r repository) GetAll() ([]Product, error) {
	return listProducts, nil
}

func (r repository) GetById(id int64) (*Product, error) {
	for _, product := range listProducts {
		if product.Id == id {
			return &product, nil
		}
	}
	return nil, fmt.Errorf("O produto com o id %d não foi encontrado", id)
}

func NewRepository() Repository {
	listProducts = []Product{}

	// TODO: para testes. remover depois
	prod1 := Product{Id: 1, ProductCode: "XX", Description: "Product 1", Width: 1.5, Height: 2.2, NetWeight: 4.52, ExpirationRate: 15.1,
		RecommendedFreezingTemperature: 32.5, FreezingRate: 5, ProductTypeId: 2, SellerId: 5}
	prod2 := Product{Id: 2, ProductCode: "YY", Description: "Product 2", Width: 1.5, Height: 2.2, NetWeight: 4.52, ExpirationRate: 15.1,
		RecommendedFreezingTemperature: 32.5, FreezingRate: 5, ProductTypeId: 2, SellerId: 5}

	listProducts = append(listProducts, prod1, prod2)
	return &repository{}
}
