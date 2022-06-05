package section

type Service interface {
	Delete(id int64) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) Delete(id int64) error {
	return s.repository.Delete(id)
}
