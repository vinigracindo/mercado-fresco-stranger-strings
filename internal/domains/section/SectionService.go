package section

type Service interface {
	UpdateCurrentCapacity(id int64, currentCapacity int64) (Section, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) UpdateCurrentCapacity(id int64, currentCapacity int64) (Section, error) {
	return s.repository.UpdateCurrentCapacity(id, currentCapacity)
}
