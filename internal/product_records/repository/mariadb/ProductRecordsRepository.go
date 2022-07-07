package mariadb

import (
	"context"
	"database/sql"
	"errors"
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

func (m mariaDBProductRecordsRepository) CountByProductId(ctx context.Context, productId int64) (int64, error) {

	rows := m.db.QueryRowContext(
		ctx,
		SqlCountByProductId,
		productId,
	)

	var productRecordsCount int64

	err := rows.Scan(&productRecordsCount)

	if err != nil {
		return 0, errors.New("sql: no rows in result set")
	}

	return productRecordsCount, nil
}
