package service

import (
	"context"
	"time"

	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/inboud_orders/domain"
)

type service struct {
	repo domain.InboundOrderRepository
}

func NewInboundOrderService(repo domain.InboundOrderRepository) domain.InboundOrderService {
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
) (domain.InboundOrder, error) {
	inboundOrder, err := s.repo.Create(ctx, orderDate, orderType, employeeId, productBatchId, warehouseId)

	if err != nil {
		return domain.InboundOrder{}, err
	}

	return inboundOrder, nil
}
