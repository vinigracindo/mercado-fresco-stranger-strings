package seller

type Service interface {
	GetAll() ([]Seller, error)
	Get(id float64) (Seller, error)
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

func (s *service) Get(id float64) (Seller, error) {
	seller, err := s.repository.Get(id)

	if err != nil {
		return Seller{}, err
	}
	return seller, nil
}
