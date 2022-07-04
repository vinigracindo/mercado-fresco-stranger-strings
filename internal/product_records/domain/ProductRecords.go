package domain

import (
	"context"
	"time"
)

type ProductRecords struct {
	Id             int64     `json:"id"`
	LastUpdateDate time.Time `json:"last_update_date"`
	PurchasePrice  float64   `json:"purchase_price"`
	SalePrice      float64   `json:"sale_price"`
	ProductId      int64     `json:"product_id"`
}

type ProductRecordsRepository interface {
	Create(ctx context.Context, productRecords *ProductRecords) (*ProductRecords, error)
}

type ProductRecordsService interface {
	Create(ctx context.Context, productRecords *ProductRecords) (*ProductRecords, error)
}
