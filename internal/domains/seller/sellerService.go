package seller

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

func (s *service) GetById(id int64) (Seller, error) {
	seller, err := s.repository.GetById(id)

	if err != nil {
		return Seller{}, err
	}
	return seller, nil
}

func (s *service) Create(cid int64, companyName, address, telephone string) (Seller, error) {
	seller, err := s.repository.Create(cid, companyName, address, telephone)
	if err != nil {
		return Seller{}, err
	}
	return seller, nil
}

func (s *service) Update(id int64, address, telephone string) (Seller, error) {
	seller, err := s.repository.Update(id, address, telephone)
	if err != nil {
		return Seller{}, err
	}
	return seller, nil
}

func (s *service) Delete(id int64) error {
	err := s.repository.Delete(id)

	if err != nil {
		return err
	}
	return nil
}
