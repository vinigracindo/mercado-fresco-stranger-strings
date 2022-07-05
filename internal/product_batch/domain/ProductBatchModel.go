package domain

import (
	"context"
	"time"
)

type ProductBatch struct {
	Id                 int64     `json:"id"`
	BatchNumber        int64     `json:"batch_number"`
	CurrentQuantity    int64     `json:"current_quantity"`
	CurrentTemperature float64   `json:"current_temperature"`
	DueDate            time.Time `json:"due_date"`
	InitialQuantity    int64     `json:"initial_quantity"`
	ManufacturingDate  time.Time `json:"manufacturing_date"`
	ManufacturingHour  int64     `json:"manufacturing_hour"`
	MinumumTemperature float64   `json:"minumum_temperature"`
	ProductId          int64     `json:"product_id"`
	SectionId          int64     `json:"section_id"`
}

type ProductBatchRepository interface {
	Create(ctx context.Context, productBatch *ProductBatch) (*ProductBatch, error)
}

type ProductBatchService interface {
	Create(ctx context.Context, productBatch *ProductBatch) (*ProductBatch, error)
}
