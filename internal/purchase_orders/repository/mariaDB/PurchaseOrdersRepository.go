package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/purchase_orders/domain"
)

type mariadbPurchaseOrdersRepository struct {
	db *sql.DB
}

func NewMariadbPurchaseOrdersRepository(db *sql.DB) domain.PurchaseOrdersRepository {
	return &mariadbPurchaseOrdersRepository{db: db}
}

func (repository *mariadbPurchaseOrdersRepository) Create(ctx context.Context, orderNumber string, orderDate time.Time, trackingCode string, buyerId, productRecordId, orderStatusId int64) (*domain.PurchaseOrders, error) {

	purchaseOrders := domain.PurchaseOrders{
		OrderNumber:     orderNumber,
		OrderDate:       orderDate,
		TrackingCode:    trackingCode,
		BuyerId:         buyerId,
		ProductRecordId: productRecordId,
		OrderStatusId:   orderStatusId,
	}

	newPurchaseOrders, err := repository.db.ExecContext(
		ctx,
		SQLCreatePurchaseOrders,
		purchaseOrders.OrderNumber,
		purchaseOrders.OrderDate,
		purchaseOrders.TrackingCode,
		purchaseOrders.BuyerId,
		purchaseOrders.ProductRecordId,
		purchaseOrders.OrderStatusId,
	)
	if err != nil {
		return nil, err
	}

	id, _ := newPurchaseOrders.LastInsertId()
	purchaseOrders.Id = id

	return &purchaseOrders, nil
}

func (repository *mariadbPurchaseOrdersRepository) ContByBuyerId(ctx context.Context, buyerId int64) (int64, error) {

	row := repository.db.QueryRowContext(
		ctx,
		SQLContByBuyerId,
		buyerId)

	var purchaseOrdersCont int64

	err := row.Scan(
		&purchaseOrdersCont)
	if err != nil {
		return 0, errors.New("sql: no rows in result set")
	}
	return purchaseOrdersCont, nil

}
