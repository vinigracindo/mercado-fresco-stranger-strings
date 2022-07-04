package service

import (
	"context"
	"time"

	EmployeesDomain "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/employees/domain"
	InboundOrdersDomain "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/inboud_orders/domain"
)

type service struct {
	repoInbound  InboundOrdersDomain.InboundOrdersRepository
	repoEmployee EmployeesDomain.EmployeeRepository
}

func NewInboundOrderService(repoInbound InboundOrdersDomain.InboundOrdersRepository, repoEmployee EmployeesDomain.EmployeeRepository) InboundOrdersDomain.InboundOrdersService {
	return &service{
		repoInbound:  repoInbound,
		repoEmployee: repoEmployee,
	}
}

func (s service) Create(
	ctx context.Context,
	orderDate time.Time,
	orderType string,
	employeeId int64,
	productBatchId int64,
	warehouseId int64,
) (InboundOrdersDomain.InboundOrders, error) {
	_, err := s.repoEmployee.GetById(ctx, employeeId)
	if err != nil {
		return InboundOrdersDomain.InboundOrders{}, EmployeesDomain.ErrEmployeeNotFound
	}

	inboundOrder, err := s.repoInbound.Create(ctx, orderDate, orderType, employeeId, productBatchId, warehouseId)

	if err != nil {
		return InboundOrdersDomain.InboundOrders{}, err
	}

	return inboundOrder, nil
}
