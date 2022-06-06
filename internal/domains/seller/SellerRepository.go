package seller

import "fmt"

var listSeller []Seller

type Repository interface {
	GetAll() ([]Seller, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]Seller, error) {
	listSeller = []Seller{}
	return listSeller, nil
}

func (r *repository) Get(id float64) (Seller, error) {
	for _, seller := range listSeller {
		if seller.Id == id {
			return seller, nil
		}
	}

	return Seller{}, fmt.Errorf("seller id %f not found", id)
}
