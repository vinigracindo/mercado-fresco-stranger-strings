package section

import "fmt"

var listSection []Section

type Repository interface {
	GetById(id int64) (Section, error)
	GetAll() ([]Section, error)
	CreateSection(sectionNumber int64, currentTemperature int64, minimumTemperature int64, currentCapacity int64, minimumCapacity int64, maximumCapacity int64, warehouseId int64, productTypeId int64) (Section, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) CreateSection(sectionNumber int64, currentTemperature int64, minimumTemperature int64, currentCapacity int64, minimumCapacity int64, maximumCapacity int64, warehouseId int64, productTypeId int64) (Section, error) {
	section := Section{
		Id:                 int64(len(listSection) + 1),
		SectionNumber:      sectionNumber,
		CurrentTemperature: currentTemperature,
		MinimumTemperature: minimumTemperature,
		CurrentCapacity:    currentCapacity,
		MinimumCapacity:    minimumCapacity,
		MaximumCapacity:    maximumCapacity,
		WarehouseId:        warehouseId,
		ProductTypeId:      productTypeId,
	}

	for i := range listSection {
		if listSection[i].SectionNumber == section.SectionNumber {
			return Section{}, fmt.Errorf("Already a secton with the code: %d", section.SectionNumber)
		}
	}

	listSection = append(listSection, section)
	return section, nil
}

func (r *repository) GetById(id int64) (Section, error) {
	section := Section{}
	found := false

	for i := range listSection {
		if listSection[i].Id == id {
			section = listSection[i]
			found = true
		}
	}

	if !found {
		return Section{}, fmt.Errorf("Sessão %d não encontrada", id)
	}

	return section, nil
}

func (r *repository) GetAll() ([]Section, error) {
	listSection = []Section{}
	return listSection, nil
}
