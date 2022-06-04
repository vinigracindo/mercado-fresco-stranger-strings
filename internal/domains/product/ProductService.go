package product

type Service interface {
	GetAll() ([]Product, error)
	GetById(id int64) (*Product, error)
}

type service struct {
	repository Repository
}

func (s service) GetAll() ([]Product, error) {
	products, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s service) GetById(id int64) (*Product, error) {
	product, err := s.repository.GetById(id)
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
