package domain

import (
	"context"
	"time"
)

type PurchaseOrders struct {
	Id              int64     `json:"id"`
	OrderNumber     string    `json:"order_number"`
	OrderDate       time.Time `json:"order_date"`
	TrackingCode    string    `json:"tracking_code"`
	BuyerId         int64     `json:"buyer_id"`
	ProductRecordId int64     `json:"product_record_id"`
	OrderStatusId   int64     `json:"order_status_id"`
}

type PurchaseOrdersRepository interface {
	Create(ctx context.Context, OrderNumber string, OrderDate time.Time, TrackingCode string, BuyerId, ProductRecordId, OrderStatusId int64) (*PurchaseOrders, error)
	ContByBuyerId(ctx context.Context, buyerId int64) (int64, error)
}

type PurchaseOrdersService interface {
	Create(ctx context.Context, OrderNumber string, OrderDate time.Time, TrackingCode string, BuyerId, ProductRecordId, OrderStatusId int64) (*PurchaseOrders, error)
}
