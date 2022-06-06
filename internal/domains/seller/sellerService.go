package seller

type Service interface {
	GetAll() ([]Seller, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]Seller, error) {
	listSeller, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return listSeller, nil
}
