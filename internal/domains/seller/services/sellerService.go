package services

import "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/seller/domain"

type service struct {
	repository domain.RepositorySeller
}

func NewSellerService(r domain.RepositorySeller) domain.ServiceSeller {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]domain.Seller, error) {
	listSeller, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return listSeller, nil
}

func (s *service) GetById(id int64) (domain.Seller, error) {
	seller, err := s.repository.GetById(id)

	if err != nil {
		return domain.Seller{}, err
	}
	return seller, nil
}

func (s *service) Create(cid int64, companyName, address, telephone string) (domain.Seller, error) {
	seller, err := s.repository.Create(cid, companyName, address, telephone)
	if err != nil {
		return domain.Seller{}, err
	}
	return seller, nil
}

func (s *service) Update(id int64, address, telephone string) (domain.Seller, error) {
	seller, err := s.repository.Update(id, address, telephone)
	if err != nil {
		return domain.Seller{}, err
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
