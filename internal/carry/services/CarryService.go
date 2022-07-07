package services

import (
	"context"

	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/carry/domain"
)

type service struct {
	repository domain.CarryRepository
}

func NewCarryService(repo domain.CarryRepository) domain.CarryService {
	return &service{
		repository: repo,
	}
}

func (s service) GetById(ctx context.Context, id int64) (*domain.CarryModel, error) {

	carry, err := s.repository.GetById(ctx, id)

	if err != nil {
		return nil, err
	}

	return carry, nil
}

func (s service) Create(ctx context.Context, carry *domain.CarryModel) (*domain.CarryModel, error) {

	newCarry, err := s.repository.Create(ctx, carry)

	if err != nil {
		return nil, err
	}

	return newCarry, nil
}
