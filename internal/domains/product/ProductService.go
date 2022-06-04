package product

type Service interface {
	GetAll() ([]Product, error)
}

type service struct {
	repository Repository
}

func (s service) GetAll() ([]Product, error) {
	product, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return product, nil
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}
