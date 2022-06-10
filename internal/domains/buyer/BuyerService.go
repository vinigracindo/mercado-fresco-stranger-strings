package buyer

type Service interface {
	Store(cardNumberId int64, firstName string, lastName string) (Buyer, error)
	GetAll() ([]Buyer, error)
	GetId(id int64) (*Buyer, error)
	Update(id int64, cardNumberId int64, lastName string) (Buyer, error)
	Delete(id int64) error
}

type service struct {
	repository Repository
}

func (s service) Store(cardNumberId int64, firstName string, lastName string) (Buyer, error) {
	buyer, err := s.repository.Store(cardNumberId, firstName, lastName)
	if err != nil {
		return Buyer{}, err
	}
	return buyer, nil
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

func (s service) Update(id int64, cardNumberId int64, lastName string) (Buyer, error) {
	buyer, err := s.repository.Update(id, cardNumberId, lastName)
	if err != nil {
		return Buyer{}, err
	}
	return buyer, nil
}

func (s service) Delete(id int64) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}
