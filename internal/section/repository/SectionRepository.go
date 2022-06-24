package repository

import (
	"fmt"

	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/section/domain"
)

var listSection = []domain.SectionModel{}
var id int64 = 0

type repository struct{}

func NewRepositorySection() domain.SectionRepository {
	return &repository{}
}

func (r repository) CreateID() int64 {
	id += 1
	return id
}

func (r *repository) Delete(id int64) error {
	deleted := false
	var index int

	for i := range listSection {
		if listSection[i].Id == id {
			index = i
			deleted = true
		}
	}

	if !deleted {
		return fmt.Errorf("section %d not found", id)
	}

	listSection = append(listSection[:index], listSection[index+1:]...)
	return nil
}

func (r *repository) UpdateCurrentCapacity(id int64, currentCapacity int64) (domain.SectionModel, error) {
	var section domain.SectionModel
	updated := false

	for i := range listSection {
		if listSection[i].Id == id {
			listSection[i].CurrentCapacity = currentCapacity
			section = listSection[i]
			updated = true
		}
	}

	if !updated {
		return domain.SectionModel{}, fmt.Errorf("section %d not found", id)
	}
	return section, nil
}

func (r *repository) Create(sectionNumber int64, currentTemperature int64, minimumTemperature int64, currentCapacity int64, minimumCapacity int64, maximumCapacity int64, warehouseId int64, productTypeId int64) (domain.SectionModel, error) {
	for i := range listSection {
		if listSection[i].SectionNumber == sectionNumber {
			return domain.SectionModel{}, fmt.Errorf("already a secton with the code: %d", sectionNumber)
		}
	}

	section := domain.SectionModel{
		Id:                 r.CreateID(),
		SectionNumber:      sectionNumber,
		CurrentTemperature: currentTemperature,
		MinimumTemperature: minimumTemperature,
		CurrentCapacity:    currentCapacity,
		MinimumCapacity:    minimumCapacity,
		MaximumCapacity:    maximumCapacity,
		WarehouseId:        warehouseId,
		ProductTypeId:      productTypeId,
	}

	listSection = append(listSection, section)
	return section, nil
}

func (r *repository) GetById(id int64) (domain.SectionModel, error) {
	section := domain.SectionModel{}
	found := false

	for i := range listSection {
		if listSection[i].Id == id {
			section = listSection[i]
			found = true
		}
	}

	if !found {
		return domain.SectionModel{}, fmt.Errorf("section %d not found", id)
	}

	return section, nil
}

func (r *repository) GetAll() ([]domain.SectionModel, error) {
	return listSection, nil
}
