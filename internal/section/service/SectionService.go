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

func (s *service) UpdateCurrentCapacity(ctx context.Context, id int64, currentCapacity int64) (domain.SectionModel, error) {
	return s.repository.UpdateCurrentCapacity(ctx, id, currentCapacity)
}

func (s *service) Create(
	ctx context.Context,
	sectionNumber string,
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
