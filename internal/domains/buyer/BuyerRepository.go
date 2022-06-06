package buyer

import "fmt"

var buyers []Buyer = []Buyer{}

type Repository interface {
	Store(id int64, cardNumberId int64, firstName string, lastName string) (Buyer, error)
	GetAll() ([]Buyer, error)
	GetId(id int64) (*Buyer, error)
	Update(id int64, cardNumberId int64, firstName string, lastName string) (*Buyer, error)
	Delete(id int64) error
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

func (repository) GetAll() ([]Buyer, error) {
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

func (repository) Update(id int64, cardNumberId int64, firstName string, lastName string) (*Buyer, error) {
	buyerUpdate := Buyer{id, cardNumberId, firstName, lastName}
	for i, buyer := range buyers {
		if buyer.Id == id {
			buyers[i] = buyerUpdate
			return &buyerUpdate, nil
		}
	}

	return nil, fmt.Errorf("não foi possível alterar, id %d não encontrado", id)
}

func (repository) Delete(id int64) error {
	deleted := false
	var index int
	for i := range buyers {
		if buyers[i].Id == id {
			index = i
			deleted = true
		}
	}
	if !deleted {
		return fmt.Errorf("comprador %d nao encontrado", id)
	}

	buyers = append(buyers[:index], buyers[index+1:]...)
	return nil
}

func NewRepository() Repository {
	return &repository{}
}
