package service

import (
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

func (s *service) Delete(id int64) error {
	return s.repository.Delete(id)
}

func (s *service) UpdateCurrentCapacity(id int64, currentCapacity int64) (domain.SectionModel, error) {
	return s.repository.UpdateCurrentCapacity(id, currentCapacity)
}

func (s *service) Create(sectionNumber int64, currentTemperature int64, minimumTemperature int64, currentCapacity int64, minimumCapacity int64, maximumCapacity int64, warehouseId int64, productTypeId int64) (domain.SectionModel, error) {
	section, err := s.repository.Create(
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

func (s *service) GetById(id int64) (domain.SectionModel, error) {
	section, err := s.repository.GetById(id)
	if err != nil {
		return domain.SectionModel{}, err
	}
	return section, nil
}

func (s *service) GetAll() ([]domain.SectionModel, error) {
	listSection, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return listSection, nil
}
