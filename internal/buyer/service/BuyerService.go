package service

import buyer "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/buyer/domain"

type service struct {
	repository buyer.BuyerRepository
}

func (s service) Create(cardNumberId int64, firstName string, lastName string) (*buyer.Buyer, error) {
	buyer, err := s.repository.Create(cardNumberId, firstName, lastName)
	if err != nil {
		return nil, err
	}
	return buyer, nil
}

func (s service) GetAll() ([]buyer.Buyer, error) {
	buyers, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return buyers, nil
}

func (s service) GetId(id int64) (*buyer.Buyer, error) {
	buyer, err := s.repository.GetId(id)
	if err != nil {
		return nil, err
	}
	return buyer, nil
}

func (s service) Update(id int64, cardNumberId int64, lastName string) (*buyer.Buyer, error) {
	buyer, err := s.repository.Update(id, cardNumberId, lastName)
	if err != nil {
		return nil, err
	}
	return buyer, nil
}

func (s service) Delete(id int64) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func NewBuyerService(r buyer.BuyerRepository) buyer.BuyerService {
	return &service{
		repository: r,
	}
}
