package domain

import (
	"context"
	"time"
)

type InboundOrder struct {
	Id             int64     `json:"id"`
	OrderDate      time.Time `json:"order_date"`
	OrderType      string    `json:"order_type"`
	EmployeeId     int64     `json:"employee_id"`
	ProductBatchId int64     `json:"product_batch_id"`
	WarehouseId    int64     `json:"warehouse_id"`
}

type InboundOrderService interface {
	Create(
		ctx context.Context,
		orderDate time.Time,
		orderType string,
		employeeId int64,
		productBatchId int64,
		warehouseId int64,
	) (InboundOrder, error)
}

type InboundOrderRepository interface {
	Create(
		ctx context.Context,
		orderDate time.Time,
		orderType string,
		employeeId int64,
		productBatchId int64,
		warehouseId int64,
	) (InboundOrder, error)
}
