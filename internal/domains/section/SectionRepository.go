package section

import "fmt"

var listSection []Section

type Repository interface {
	UpdateCurrentCapacity(id int64, currentCapacity int64) (Section, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) UpdateCurrentCapacity(id int64, currentCapacity int64) (Section, error) {
	var section Section
	updated := false

	for i := range listSection {
		if listSection[i].Id == id {
			listSection[i].CurrentCapacity = currentCapacity
			section = listSection[i]
			updated = true
		}
	}

	if !updated {
		return Section{}, fmt.Errorf("Section %d not found", id)
	}
	return section, nil
}
