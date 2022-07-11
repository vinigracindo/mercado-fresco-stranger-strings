package service

import (
	"context"

	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/section/domain"
)

type service struct {
	repository domain.SectionRepository
}

func NewServiceSection(r domain.SectionRepository) domain.SectionService {
	return &service{
		repository: r,
	}
}

func (s *service) Delete(ctx context.Context, id int64) error {
	return s.repository.Delete(ctx, id)
}

func (s *service) UpdateCurrentCapacity(ctx context.Context, id int64, currentCapacity int64) (*domain.SectionModel, error) {
	sectionCurrent, err := s.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	if currentCapacity > 0 {
		sectionCurrent.CurrentCapacity = currentCapacity
	}

	section, err := s.repository.UpdateCurrentCapacity(ctx, &sectionCurrent)
	if err != nil {
		return nil, err
	}

	return section, nil
}

func (s *service) Create(
	ctx context.Context,
	sectionNumber int64,
	currentTemperature float64,
	minimumTemperature float64,
	currentCapacity int64,
	minimumCapacity int64,
	maximumCapacity int64,
	warehouseId int64,
	productTypeId int64) (domain.SectionModel, error) {

	section, err := s.repository.Create(
		ctx,
		sectionNumber,
		currentTemperature,
		minimumTemperature,
		currentCapacity,
		minimumCapacity,
		maximumCapacity,
		warehouseId,
		productTypeId,
	)
	if err != nil {
		return domain.SectionModel{}, err
	}
	return section, nil
}

func (s *service) GetById(ctx context.Context, id int64) (domain.SectionModel, error) {
	section, err := s.repository.GetById(ctx, id)
	if err != nil {
		return domain.SectionModel{}, err
	}
	return section, nil
}

func (s *service) GetAll(ctx context.Context) ([]domain.SectionModel, error) {
	listSection, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return listSection, nil
}

func (s *service) GetAllProductCountBySection(ctx context.Context) (*[]domain.ReportProductsModel, error) {
	productsCount, err := s.repository.GetAllProductCountBySection(ctx)
	if err != nil {
		return nil, err
	}

	return productsCount, nil
}

func (s *service) GetByIdProductCountBySection(ctx context.Context, id int64) (*domain.ReportProductsModel, error) {
	_, err := s.repository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	result, err := s.repository.GetByIdProductCountBySection(ctx, id)
	if err != nil {
		return nil, err
	}

	return result, nil
}
