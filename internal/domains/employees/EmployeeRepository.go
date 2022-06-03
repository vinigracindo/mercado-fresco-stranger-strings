package employees

var employees = []Employee{}

type Repository interface {
	GetAll() ([]Employee, error)
}

type repository struct{}

func (repository) GetAll() ([]Employee, error) {
	return employees, nil
}

func NewRepository() Repository {
	return &repository{}
}
