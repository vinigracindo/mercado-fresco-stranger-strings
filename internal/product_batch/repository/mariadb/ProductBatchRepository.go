package repository

import (
	"context"
	"database/sql"

	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/product_batch/domain"
)

type mariaDbProductBatchRepository struct {
	db *sql.DB
}

func NewMariadbProductBatchRepository(db *sql.DB) domain.ProductBatchRepository {
	return &mariaDbProductBatchRepository{db: db}
}

func (m mariaDbProductBatchRepository) Create(ctx context.Context, productBatch *domain.ProductBatch) (*domain.ProductBatch, error) {

	productBatchResult, err := m.db.ExecContext(
		ctx,
		SQLCreate,
		&productBatch.BatchNumber,
		&productBatch.CurrentQuantity,
		&productBatch.CurrentTemperature,
		&productBatch.DueDate,
		&productBatch.InitialQuantity,
		&productBatch.ManufacturingDate,
		&productBatch.ManufacturingHour,
		&productBatch.MinumumTemperature,
		&productBatch.ProductId,
		&productBatch.SectionId,
	)

	if err != nil {
		return nil, err
	}

	lastId, _ := productBatchResult.LastInsertId()

	productBatch.Id = lastId

	return productBatch, nil
}
