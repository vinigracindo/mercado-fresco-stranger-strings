package section

type Service interface {
	GetById(id int64) (Section, error)
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

func (s *service) GetById(id int64) (Section, error) {
	section, err := s.repository.GetById(id)
	if err != nil {
		return Section{}, err
	}
	return section, nil
}

func (s *service) GetAll() ([]Section, error) {
	listSection, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return listSection, nil
}
