package section

type Service interface {
	GetAll() ([]Section, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]Section, error) {
	listSection, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return listSection, nil
}
