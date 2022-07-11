package services

import (
	"context"
	"errors"

	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/locality/domain"
	locality "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/locality/domain"
	seller "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/seller/domain"
)

type service struct {
	repoLocality locality.LocalityRepository
	repoSeller   seller.RepositorySeller
}

func NewLocalityService(repoLocality locality.LocalityRepository, repoSeller seller.RepositorySeller) locality.LocalityService {
	return &service{
		repoLocality: repoLocality,
		repoSeller:   repoSeller,
	}
}

func (s service) ReportCarrie(ctx context.Context, locality_id int64) (*[]domain.ReportCarrie, error) {
	localities, err := s.repoLocality.ReportCarrie(ctx, locality_id)

	if err != nil {
		return nil, err
	}

	return localities, nil
}

func (s service) CreateLocality(ctx context.Context, locality *domain.LocalityModel) (*domain.LocalityModel, error) {
	locality, err := s.repoLocality.CreateLocality(ctx, locality)
	if err != nil {
		return nil, err
	}

	return locality, nil
}

func (s service) GetByIdReportSeller(ctx context.Context, locality_id int64) (*[]domain.ReportSeller, error) {
	locality, err := s.repoLocality.GetById(ctx, locality_id)

	if err != nil {
		return nil, errors.New("locality not found")
	}

	countSeller, err := s.repoSeller.CountByLocalityId(ctx, locality_id)

	if err != nil {
		return nil, err
	}

	var result []domain.ReportSeller

	reportSeller := domain.ReportSeller{
		LocalityId:   locality_id,
		LocalityName: locality.LocalityName,
		SellerCount:  countSeller,
	}

	result = append(result, reportSeller)

	return &result, nil

}

func (s service) GetAllReportSeller(ctx context.Context) (*[]domain.ReportSeller, error) {
	localities, err := s.repoLocality.GetAllReportSeller(ctx)

	if err != nil {
		return nil, err
	}

	return localities, nil
}
