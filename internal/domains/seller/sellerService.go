package seller

type Service interface {
	GetAll() ([]Seller, error)
	Get(id int64) (Seller, error)
	CreateSeller(cid int64, companyName, address, telephone string) (Seller, error)
	UpdateSeller(id int64, cid int64, companyName, address, telephone string) (Seller, error)
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

func (s *service) Get(id int64) (Seller, error) {
	seller, err := s.repository.Get(id)

	if err != nil {
		return Seller{}, err
	}
	return seller, nil
}

func (s *service) CreateSeller(cid int64, companyName, address, telephone string) (Seller, error) {
	seller, err := s.repository.CreateSeller(cid, companyName, address, telephone)
	if err != nil {
		return Seller{}, err
	}
	return seller, nil
}

func (s *service) UpdateSeller(id int64, cid int64, companyName, address, telephone string) (Seller, error) {
	seller, err := s.repository.UpdateSeller(id, cid, companyName, address, telephone)
	if err != nil {
		return Seller{}, err
	}
	return seller, nil
}
