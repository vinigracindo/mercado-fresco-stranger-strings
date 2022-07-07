package services

import (
	"context"

	carry "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/carry/domain"
	locality "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/locality/domain"
)

type service struct {
	repoCarry    carry.CarryRepository
	repoLocality locality.LocalityRepository
}

func NewLocalityService(repoLocality locality.LocalityRepository, repoCarry carry.CarryRepository) locality.LocalityService {
	return &service{
		repoCarry:    repoCarry,
		repoLocality: repoLocality,
	}
}

func (s service) ReportCarrie(ctx context.Context, locality_id int64) (any, error) {
	total_locality_cary, err := s.repoCarry.CountLocality(ctx, locality_id)

	if err != nil {
		return nil, err
	}

	localty, err := s.repoLocality.GetById(ctx, locality_id)

	if err != nil {
		return nil, err
	}

	return &struct {
		LocalityId   int64  `json:"locality_id"`
		LocalityName string `json:"locality_name"`
		CarriesCount int64  `json:"carries_count"`
	}{
		LocalityId:   localty.Id,
		LocalityName: localty.LocalityName,
		CarriesCount: total_locality_cary,
	}, nil
}
