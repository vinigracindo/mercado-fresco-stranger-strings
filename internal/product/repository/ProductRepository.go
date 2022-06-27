package repository

import (
	"fmt"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/product/domain"
)

var listProducts []domain.Product
var lastId int64 = 1

type repository struct {
}

func CreateProductRepository() domain.ProductRepository {
	listProducts = []domain.Product{}

	listProducts = append(listProducts)

	return &repository{}
}

func (r *repository) GetAll() ([]domain.Product, error) {
	return listProducts, nil
}

func (r *repository) GetById(id int64) (*domain.Product, error) {
	for _, product := range listProducts {
		if product.Id == id {
			return &product, nil
		}
	}
	return nil, fmt.Errorf("the product with the id %d was not found", id)
}

func (r *repository) Create(productCode string, description string, width float64, height float64, length float64, netWeight float64,
	expirationRate float64, recommendedFreezingTemperature float64, freezingRate int, productTypeId int, sellerId int) (*domain.Product, error) {

	nextId := lastId
	newProduct := domain.Product{
		Id:                             nextId,
		ProductCode:                    productCode,
		Description:                    description,
		Width:                          width,
		Height:                         height,
		Length:                         length,
		NetWeight:                      netWeight,
		ExpirationRate:                 expirationRate,
		RecommendedFreezingTemperature: recommendedFreezingTemperature,
		FreezingRate:                   freezingRate,
		ProductTypeId:                  productTypeId,
		SellerId:                       sellerId}

	for _, product := range listProducts {
		if product.ProductCode == productCode {
			return nil, fmt.Errorf("the product with code %s has already been registered", productCode)
		}
	}
	listProducts = append(listProducts, newProduct)
	lastId += 1

	return &newProduct, nil
}

func (r *repository) LastId() int64 {
	return lastId
}

func (r *repository) UpdateDescription(id int64, description string) (*domain.Product, error) {
	var product domain.Product
	update := false
	for i := range listProducts {
		if listProducts[i].Id == id {
			listProducts[i].Description = description
			update = true
			product = listProducts[i]
			break
		}
	}

	if !update {
		return nil, fmt.Errorf("the product with id %d was not found", id)
	}
	return &product, nil
}

func (r *repository) Delete(id int64) error {
	deleted := false
	var index int
	for i := range listProducts {
		if listProducts[i].Id == id {
			index = i
			deleted = true
		}
	}

	if !deleted {
		return fmt.Errorf("the product with id %d was not found", id)
	}
	listProducts = append(listProducts[:index], listProducts[index+1:]...)

	return nil
}
