package buyer

import "fmt"

var buyers []Buyer = []Buyer{}

type Repository interface {
	Store(id int64, cardNumberId int64, firstName string, lastName string) (Buyer, error)
	GetAll() ([]Buyer, error)
	GetId(id int64) (*Buyer, error)
}

type repository struct {
}

func createId() int64 {
	return int64(len(buyers) + 1)
}

func (repository) Store(id int64, cardNumberId int64, firstName string, lastName string) (Buyer, error) {

	for _, buyer := range buyers {
		if buyer.CardNumberId == cardNumberId {
			return Buyer{}, fmt.Errorf("comprador já está cadastrado")
		}
	}
	newBuyer := Buyer{createId(), cardNumberId, firstName, lastName}
	buyers = append(buyers, newBuyer)

	return newBuyer, nil
}

func (w repository) GetAll() ([]Buyer, error) {
	return buyers, nil
}

func (repository) GetId(id int64) (*Buyer, error) {
	for _, buyer := range buyers {
		if buyer.Id == id {
			return &buyer, nil
		}
	}
	return nil, fmt.Errorf("o comprador do id %d não foi encontrado", id)
}

func NewRepository() Repository {
	return &repository{}
}