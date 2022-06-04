package section

import "fmt"

var listSection []Section

type Repository interface {
	GetById(id int64) (Section, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
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
