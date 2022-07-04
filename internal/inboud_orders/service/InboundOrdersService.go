package service

import (
	"context"
	"time"

	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/inboud_orders/domain"
)

type service struct {
	repo domain.InboundOrdersRepository
}

func NewInboundOrderService(repo domain.InboundOrdersRepository) domain.InboundOrdersService {
	return &service{
		repo: repo,
	}
}

func (s service) Create(
	ctx context.Context,
	orderDate time.Time,
	orderType string,
	employeeId int64,
	productBatchId int64,
	warehouseId int64,
) (domain.InboundOrders, error) {
	inboundOrder, err := s.repo.Create(ctx, orderDate, orderType, employeeId, productBatchId, warehouseId)

	if err != nil {
		return domain.InboundOrders{}, err
	}

	return inboundOrder, nil
}
