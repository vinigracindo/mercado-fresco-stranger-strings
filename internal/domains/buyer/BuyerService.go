package buyer

type Service interface {
	Store(id int64, cardNumberId int64, firstName string, lastName string) (Buyer, error)
	GetAll() ([]Buyer, error)
	GetId(id int64) (*Buyer, error)
}

type service struct {
	repository Repository
}

func (s service) Store(id int64, cardNumberId int64, firstName string, lastName string) (Buyer, error) {

	buyer, err := s.repository.Store(id, cardNumberId, firstName, lastName)

	if err != nil {
		return Buyer{}, err
	}

	return buyer, nil
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s service) GetAll() ([]Buyer, error) {
	buyers, err := s.repository.GetAll()

	if err != nil {
		return []Buyer{}, err
	}

	return buyers, nil
}

func (s service) GetId(id int64) (*Buyer, error) {
	buyer, err := s.repository.GetId(id)
	if err != nil {
		return nil, err
	}
	return buyer, nil
}