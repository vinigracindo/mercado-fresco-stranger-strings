package employees

type Service interface {
	GetAll() ([]Employee, error)
	Get(id int64) (Employee, error)
}

type service struct {
	repo Repository
}

func (s service) GetAll() ([]Employee, error) {
	employees, err := s.repo.GetAll()

	if err != nil {
		return nil, err
	}

	return employees, nil
}

func (s service) Get(id int64) (Employee, error) {
	employee, err := s.repo.Get(id)

	if err != nil {
		return Employee{}, err
	}

	return employee, nil
}

func NewService(r Repository) Service {
	return &service{
		repo: r,
	}
}
