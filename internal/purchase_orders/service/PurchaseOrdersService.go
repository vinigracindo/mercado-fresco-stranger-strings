package service

import (
	"context"
	"time"

	DomainBuyer "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/buyer/domain"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/purchase_orders/domain"
)

type service struct {
	repository      domain.PurchaseOrdersRepository
	repositoryBuyer DomainBuyer.BuyerRepository
}

func NewPurchaseOrdersService(repository domain.PurchaseOrdersRepository, repositoryBuyer DomainBuyer.BuyerRepository) domain.PurchaseOrdersService {
	return &service{
		repository:      repository,
		repositoryBuyer: repositoryBuyer,
	}
}

func (s service) Create(ctx context.Context, orderNumber string, orderDate time.Time, trackingCode string, buyerId, productRecordId, orderStatusId int64) (*domain.PurchaseOrders, error) {

	_, err := s.repositoryBuyer.GetId(ctx, buyerId)
	if err != nil {
		return nil, DomainBuyer.ErrIDNotFound
	}

	dateInput := orderDate
	currentDate := time.Now().UTC().Truncate(24 * time.Hour)

	if dateInput.Before(currentDate) {
		return nil, domain.ErrInvalidDate
	}

	newPurchaseOrders, err := s.repository.Create(ctx, orderNumber, orderDate, trackingCode, buyerId, productRecordId, orderStatusId)

	if err != nil {
		return nil, err
	}
	return newPurchaseOrders, nil
}
