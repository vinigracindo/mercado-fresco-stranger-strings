package section

import "fmt"

var listSection []Section
var id int64 = 0

type Repository interface {
	CreateSection(section *Section) (Section, error)
	CreateID() int64
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) CreateID() int64 {
	id += 1
	return id
}

func (r *repository) CreateSection(section *Section) (Section, error) {
	for i := range listSection {
		if listSection[i].SectionNumber == section.SectionNumber {
			return Section{}, fmt.Errorf("Already a secton with the code: %d", section.SectionNumber)
		}
	}

	section.Id = r.CreateID()

	listSection = append(listSection, *section)

	return *section, nil
}
