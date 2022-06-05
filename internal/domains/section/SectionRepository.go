package section

import "fmt"

var listSection []Section

type Repository interface {
	Delete(id int64) error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) Delete(id int64) error {
	listSection = append(listSection, Section{1, 1, 1, 1, 1, 1, 1, 1, 1})
	listSection = append(listSection, Section{2, 2, 1, 1, 1, 1, 1, 1, 1})
	listSection = append(listSection, Section{4, 2, 1, 1, 1, 1, 1, 1, 1})

	deleted := false
	var index int

	for i := range listSection {
		if listSection[i].Id == id {
			index = i
			deleted = true
		}
	}

	if !deleted {
		return fmt.Errorf("Section %d not found", id)
	}

	listSection = append(listSection[:index], listSection[index+1:]...)
	return nil
}
