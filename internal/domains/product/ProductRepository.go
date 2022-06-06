package product

import (
	"fmt"
)

var listProducts []Product
var lastId int64

type Repository interface {
	GetAll() ([]Product, error)
	GetById(id int64) (*Product, error)
	Create(id int64, productCode string, description string, width float64, height float64, length float64, netWeight float64,
		expirationRate float64, recommendedFreezingTemperature float64, freezingRate int, productTypeId int, sellerId int) (Product, error)
	LastId() int64
	UpdateDescription(id int64, description string) (Product, error)
}

type repository struct {
}

func (r *repository) GetAll() ([]Product, error) {
	return listProducts, nil
}

func (r *repository) GetById(id int64) (*Product, error) {
	for _, product := range listProducts {
		if product.Id == id {
			return &product, nil
		}
	}
	return nil, fmt.Errorf("O produto com o id %d não foi encontrado", id)
}

func (r *repository) Create(id int64, productCode string, description string, width float64, height float64, length float64, netWeight float64,
	expirationRate float64, recommendedFreezingTemperature float64, freezingRate int, productTypeId int, sellerId int) (Product, error) {
	newProduct := Product{id, productCode, description, width, height, length, netWeight,
		expirationRate, recommendedFreezingTemperature, freezingRate, productTypeId, sellerId}

	for _, product := range listProducts {
		if product.ProductCode == productCode {
			return product, fmt.Errorf("O produto com o código %s já foi cadastrado", productCode)
		}
	}
	listProducts = append(listProducts, newProduct)
	lastId = newProduct.Id

	return newProduct, nil
}

func (r *repository) LastId() int64 {
	return lastId
}

func (r *repository) UpdateDescription(id int64, description string) (Product, error) {
	var product Product
	update := false
	for i := range listProducts {
		currentProd := listProducts[i]
		if currentProd.Id == id {
			currentProd.Description = description
			update = true
			product = currentProd
			break
		}
	}

	if !update {
		return Product{}, fmt.Errorf("produto %d não foi encontrado", id)
	}
	return product, nil
}

func CreateRepository() Repository {
	listProducts = []Product{}

	// TODO: para testes. remover depois
	prod1 := Product{Id: 1, ProductCode: "XX", Description: "Product 1", Width: 1.5, Height: 2.2, NetWeight: 4.52, ExpirationRate: 15.1,
		RecommendedFreezingTemperature: 32.5, FreezingRate: 5, ProductTypeId: 2, SellerId: 5}
	prod2 := Product{Id: 2, ProductCode: "YY", Description: "Product 2", Width: 1.5, Height: 2.2, NetWeight: 4.52, ExpirationRate: 15.1,
		RecommendedFreezingTemperature: 32.5, FreezingRate: 5, ProductTypeId: 2, SellerId: 5}

	lastId = 2

	listProducts = append(listProducts, prod1, prod2)

	return &repository{}
}
