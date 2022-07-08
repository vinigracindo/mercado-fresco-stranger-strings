package service

import (
	"context"
	"fmt"

	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/buyer/domain"
	purchaseOrdersRepo "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/purchase_orders/domain"
)

type buyerService struct {
	buyerRepository          domain.BuyerRepository
	purchaseOrdersRepository purchaseOrdersRepo.PurchaseOrdersRepository
}

func (s buyerService) Create(ctx context.Context, cardNumberId, firstName, lastName string) (*domain.Buyer, error) {

	buyer, err := s.buyerRepository.Create(ctx, cardNumberId, firstName, lastName)
	if err != nil {
		return &domain.Buyer{}, err
	}
	return buyer, nil
}

func (s buyerService) GetAll(ctx context.Context) (*[]domain.Buyer, error) {
	buyers, err := s.buyerRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return buyers, nil
}

func (s buyerService) GetId(ctx context.Context, id int64) (*domain.Buyer, error) {
	buyer, err := s.buyerRepository.GetId(ctx, id)
	if err != nil {
		return nil, err
	}
	return buyer, nil
}

func (s buyerService) Update(ctx context.Context, id int64, cardNumberId, lastName string) (*domain.Buyer, error) {

	_, err := s.buyerRepository.Update(ctx, id, cardNumberId, lastName)

	if err != nil {
		return &domain.Buyer{}, err
	}

	buyerUpdate, err := s.buyerRepository.GetId(ctx, id)

	if err != nil {
		return &domain.Buyer{}, err
	}

	return buyerUpdate, nil

}

func (s buyerService) Delete(ctx context.Context, id int64) error {
	err := s.buyerRepository.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (s buyerService) GetPurchaseOrdersReports(ctx context.Context, id int64) (*[]domain.PurchaseOrdersReport, error) {
	buyer, err := s.buyerRepository.GetId(ctx, id)

	if err != nil {
		return nil, err
	}
	result, err := s.purchaseOrdersRepository.ContByBuyerId(ctx, buyer.Id)
	if err != nil {
		return nil, err
	}

	report := make([]domain.PurchaseOrdersReport, 0)

	purchaseOrdersReport := domain.PurchaseOrdersReport{
		Id:                 buyer.Id,
		CardNumberId:       buyer.CardNumberId,
		FirstName:          buyer.FirstName,
		LastName:           buyer.LastName,
		CountBuyersRecords: result,
	}
	report = append(report, purchaseOrdersReport)

	fmt.Println(0)
	return &report, nil
}

func (s buyerService) GetAllPurchaseOrdersReports(ctx context.Context) (*[]domain.PurchaseOrdersReport, error) {

	buyers, err := s.buyerRepository.GetAllPurchaseOrdersReports(ctx)

	if err != nil {
		return nil, err
	}

	return buyers, nil
}

func NewBuyerService(buyerRepository domain.BuyerRepository, purchaseOrdersRepository purchaseOrdersRepo.PurchaseOrdersRepository) domain.BuyerService {
	return &buyerService{
		buyerRepository:          buyerRepository,
		purchaseOrdersRepository: purchaseOrdersRepository,
	}
}
