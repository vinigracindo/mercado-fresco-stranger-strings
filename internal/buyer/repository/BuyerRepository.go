package repository

import (
	"fmt"

	buyer "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/buyer/domain"
)

var buyers = []buyer.Buyer{}
var id int64 = 0

type repository struct{}

func (r *repository) CreateId() int64 {
	id += 1
	return id
}

func (r *repository) Create(cardNumberId int64, firstName string, lastName string) (*buyer.Buyer, error) {

	for i := range buyers {
		if buyers[i].CardNumberId == cardNumberId {
			return nil, fmt.Errorf("buyer already registered: %d", cardNumberId)
		}
	}
	newBuyer := buyer.Buyer{
		Id:           r.CreateId(),
		CardNumberId: cardNumberId,
		FirstName:    firstName,
		LastName:     lastName,
	}
	buyers = append(buyers, newBuyer)
	return &newBuyer, nil
}

func (r *repository) GetAll() ([]buyer.Buyer, error) {
	return buyers, nil
}

func (r *repository) GetId(id int64) (*buyer.Buyer, error) {
	for _, buyer := range buyers {
		if buyer.Id == id {
			return &buyer, nil
		}
	}
	return nil, fmt.Errorf("buyer with id %d not found", id)
}

func (r *repository) Update(id int64, cardNumberId int64, lastName string) (*buyer.Buyer, error) {
	for i, buyer := range buyers {
		if buyer.Id == id {
			buyers[i].CardNumberId = cardNumberId
			buyers[i].LastName = lastName
			return &buyers[i], nil
		}
	}
	return nil, fmt.Errorf("buyer with id %d not found", id)
}

func (r *repository) Delete(id int64) error {
	deleted := false
	var index int
	for i := range buyers {
		if buyers[i].Id == id {
			index = i
			deleted = true
		}
	}
	if !deleted {
		return fmt.Errorf("buyer with id %d not found", id)
	}
	buyers = append(buyers[:index], buyers[index+1:]...)
	return nil
}

func NewBuyerRepository() buyer.BuyerRepository {
	return &repository{}
}
