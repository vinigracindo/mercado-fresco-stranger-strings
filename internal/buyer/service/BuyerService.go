package service

import (
	"context"

	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/buyer/domain"
)

type service struct {
	repository domain.BuyerRepository
}

func (s service) Create(ctx context.Context, cardNumberId, firstName, lastName string) (*domain.Buyer, error) {

	buyer, err := s.repository.Create(ctx, cardNumberId, firstName, lastName)
	if err != nil {
		return &domain.Buyer{}, err
	}
	return buyer, nil
}

func (s service) GetAll(ctx context.Context) (*[]domain.Buyer, error) {
	buyers, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return buyers, nil
}

func (s service) GetId(ctx context.Context, id int64) (*domain.Buyer, error) {
	buyer, err := s.repository.GetId(ctx, id)
	if err != nil {
		return &domain.Buyer{}, err
	}
	return buyer, nil
}

func (s service) Update(ctx context.Context, id int64, cardNumberId, lastName string) (*domain.Buyer, error) {
	buyer, err := s.repository.Update(ctx, id, cardNumberId, lastName)
	if err != nil {
		return &domain.Buyer{}, err
	}
	return buyer, nil
}

func (s service) Delete(ctx context.Context, id int64) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func NewBuyerService(r domain.BuyerRepository) domain.BuyerService {
	return &service{
		repository: r,
	}
}
