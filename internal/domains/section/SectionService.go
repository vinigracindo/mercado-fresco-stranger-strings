package section

type Service interface {
	GetById(id int64) (Section, error)
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
