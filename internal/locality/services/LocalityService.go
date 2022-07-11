package services

import (
	"context"

	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/locality/domain"
	locality "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/locality/domain"
)

type service struct {
	repoLocality locality.LocalityRepository
}

func NewLocalityService(repoLocality locality.LocalityRepository) locality.LocalityService {
	return &service{
		repoLocality: repoLocality,
	}
}

func (s service) ReportCarrie(ctx context.Context, locality_id int64) (*[]domain.ReportCarrie, error) {
	localtys, err := s.repoLocality.ReportCarrie(ctx, locality_id)

	if err != nil {
		return nil, err
	}

	return localtys, nil
}

func (s service) Create(ctx context.Context, locality *domain.LocalityModel) (*domain.LocalityModel, error) {
	locality, err := s.repoLocality.Create(ctx, locality)
	if err != nil {
		return nil, err
	}

	return locality, nil
}
