package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/inboud_orders/domain"
)

type mariaDBInboundOrdersRepository struct {
	db *sql.DB
}

func NewMariaDBInboundRepositoryRepository(db *sql.DB) domain.InboundOrdersRepository {
	return &mariaDBInboundOrdersRepository{db: db}
}

func (repo *mariaDBInboundOrdersRepository) Create(
	ctx context.Context,
	orderDate time.Time,
	orderNumber string,
	employeeId int64,
	productBatchId int64,
	warehouseId int64,
) (domain.InboundOrders, error) {
	inboundOrders := domain.InboundOrders{
		OrderDate:      orderDate,
		OrderNumber:    orderNumber,
		EmployeeId:     employeeId,
		ProductBatchId: productBatchId,
		WarehouseId:    warehouseId,
	}

	res, err := repo.db.ExecContext(
		ctx,
		SQLCreateInboundOrder,
		&inboundOrders.OrderDate,
		&inboundOrders.OrderNumber,
		&inboundOrders.EmployeeId,
		&inboundOrders.ProductBatchId,
		&inboundOrders.WarehouseId,
	)
	if err != nil {
		return domain.InboundOrders{}, err
	}

	id, _ := res.LastInsertId()
	inboundOrders.Id = id

	return inboundOrders, nil
}
