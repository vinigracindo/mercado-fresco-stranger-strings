package employees

type Service interface {
	GetAll() ([]Employee, error)
}

type service struct{}

func (service) GetAll() ([]Employee, error) {
	return []Employee{}, nil
}

func NewService() Service {
	return &service{}
}
