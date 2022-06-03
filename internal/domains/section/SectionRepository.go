package section

var listSection []Section

type Repository interface {
	GetAll() ([]Section, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]Section, error) {
	listSection = []Section{}
	return listSection, nil
}
