package mariadb

import (
	"context"
	"database/sql"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/product_records/domain"
)

type mariaDBProductRecordsRepository struct {
	db *sql.DB
}

func CreateProductRecordsRepository(db *sql.DB) domain.ProductRecordsRepository {
	return &mariaDBProductRecordsRepository{db: db}
}

func (m mariaDBProductRecordsRepository) Create(ctx context.Context, productRecords *domain.ProductRecords) (*domain.ProductRecords, error) {
	productRecordsResult, err := m.db.ExecContext(
		ctx,
		SqlCreate,
		&productRecords.LastUpdateDate,
		&productRecords.PurchasePrice,
		&productRecords.SalePrice,
		&productRecords.ProductId,
	)

	if err != nil {
		return nil, err
	}

	lastId, _ := productRecordsResult.LastInsertId()

	productRecords.Id = lastId

	return productRecords, nil
}
