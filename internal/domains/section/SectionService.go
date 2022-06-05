package section

type Service interface {
	CreateSection(section *Section) (Section, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) CreateSection(section *Section) (Section, error) {

	sectionNew, err := s.repository.CreateSection(section)
	if err != nil {
		return Section{}, err
	}
	return sectionNew, nil
}
