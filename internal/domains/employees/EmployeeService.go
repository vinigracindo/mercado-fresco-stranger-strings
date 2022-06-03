package employees

type Service interface {
	GetAll() ([]Employee, error)
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

func NewService(r Repository) Service {
	return &service{
		repo: r,
	}
}
