package domain

import (
	"context"
	"time"
)

type InboundOrders struct {
	Id             int64     `json:"id"`
	OrderDate      time.Time `json:"order_date"`
	OrderNumber    string    `json:"order_number"`
	EmployeeId     int64     `json:"employee_id"`
	ProductBatchId int64     `json:"product_batch_id"`
	WarehouseId    int64     `json:"warehouse_id"`
}

type InboundOrdersService interface {
	Create(
		ctx context.Context,
		orderDate time.Time,
		orderNumber string,
		employeeId int64,
		productBatchId int64,
		warehouseId int64,
	) (InboundOrders, error)
}

type InboundOrdersRepository interface {
	Create(
		ctx context.Context,
		orderDate time.Time,
		orderNumber string,
		employeeId int64,
		productBatchId int64,
		warehouseId int64,
	) (InboundOrders, error)
}
