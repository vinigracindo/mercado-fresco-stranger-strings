package service

import (
	"context"

	DomainBuyer "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/buyer/domain"
	DomainPurchaseOrders "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/purchase_orders/domain"
)

type service struct {
	repository      DomainPurchaseOrders.PurchaseOrdersRepository
	repositoryBuyer DomainBuyer.BuyerRepository
}

func NewPurchaseOrdersService(repository DomainPurchaseOrders.PurchaseOrdersRepository, repositoryBuyer DomainBuyer.BuyerRepository) DomainPurchaseOrders.PurchaseOrdersService {
	return &service{
		repository:      repository,
		repositoryBuyer: repositoryBuyer,
	}
}

func (s service) Create(ctx context.Context, orderNumber, orderDate, trackingCode string, buyerId, productRecordId, orderStatusId int64) (*DomainPurchaseOrders.PurchaseOrders, error) {

	_, err := s.repositoryBuyer.GetId(ctx, buyerId)

	if err != nil {
		return nil, DomainBuyer.ErrIDNotFound
	}

	newPurchaseOrders, err := s.repository.Create(ctx, orderNumber, orderDate, trackingCode, buyerId, productRecordId, orderStatusId)

	if err != nil {
		return nil, err
	}
	return newPurchaseOrders, nil

}
