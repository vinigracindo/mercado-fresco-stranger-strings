package services

import (
	"context"

	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/seller/domain"
)

type service struct {
	repository domain.RepositorySeller
}

func NewSellerService(r domain.RepositorySeller) domain.ServiceSeller {
	return &service{repository: r}
}

func (s *service) GetAll(ctx context.Context) (*[]domain.Seller, error) {
	listSeller, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return listSeller, nil
}

func (s *service) GetById(ctx context.Context, id int64) (*domain.Seller, error) {
	seller, err := s.repository.GetById(ctx, id)

	if err != nil {
		return nil, err
	}
	return seller, nil
}

func (s *service) Create(ctx context.Context, seller *domain.Seller) (*domain.Seller, error) {
	seller, err := s.repository.Create(ctx, seller)
	if err != nil {
		return nil, err
	}
	return seller, nil
}

func (s *service) Update(ctx context.Context, id int64, adress, telephone string) (*domain.Seller, error) {
	seller, err := s.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	seller.Address = adress
	seller.Telephone = telephone

	sellerResult, err := s.repository.Update(ctx, seller)
	if err != nil {
		return nil, err
	}

	return sellerResult, nil
}

func (s *service) Delete(ctx context.Context, id int64) error {
	err := s.repository.Delete(ctx, id)

	if err != nil {
		return err
	}
	return nil
}
