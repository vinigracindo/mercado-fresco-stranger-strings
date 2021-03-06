package mariadb

import (
	"context"
	"database/sql"
	"errors"

	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/product/domain"
)

type mariaDBProductRepository struct {
	db *sql.DB
}

func CreateProductRepository(db *sql.DB) domain.ProductRepository {
	return &mariaDBProductRepository{db: db}
}

func (m mariaDBProductRepository) GetAll(ctx context.Context) (*[]domain.Product, error) {
	var products []domain.Product

	rows, err := m.db.QueryContext(ctx, SqlGetAll)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var product domain.Product

		err := rows.Scan(
			&product.Id,
			&product.ProductCode,
			&product.Description,
			&product.Width,
			&product.Height,
			&product.Length,
			&product.NetWeight,
			&product.ExpirationRate,
			&product.RecommendedFreezingTemperature,
			&product.FreezingRate,
			&product.ProductTypeId,
			&product.SellerId)

		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}
	return &products, nil
}

func (m mariaDBProductRepository) GetById(ctx context.Context, id int64) (*domain.Product, error) {
	row := m.db.QueryRowContext(ctx, SqlGetById, id)

	var product domain.Product

	err := row.Scan(
		&product.Id,
		&product.ProductCode,
		&product.Description,
		&product.Width,
		&product.Height,
		&product.Length,
		&product.NetWeight,
		&product.ExpirationRate,
		&product.RecommendedFreezingTemperature,
		&product.FreezingRate,
		&product.ProductTypeId,
		&product.SellerId)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, domain.ErrProductIdNotFound
	}

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (m mariaDBProductRepository) Create(ctx context.Context, product *domain.Product) (*domain.Product, error) {

	productResult, err := m.db.ExecContext(
		ctx,
		SqlCreate,
		&product.ProductCode,
		&product.Description,
		&product.Width,
		&product.Height,
		&product.Length,
		&product.NetWeight,
		&product.ExpirationRate,
		&product.RecommendedFreezingTemperature,
		&product.FreezingRate,
		&product.ProductTypeId,
		&product.SellerId,
	)

	if err != nil {
		return nil, err
	}

	lastId, _ := productResult.LastInsertId()

	product.Id = lastId

	return product, nil
}

func (m mariaDBProductRepository) UpdateDescription(ctx context.Context, product *domain.Product) (*domain.Product, error) {

	_, err := m.db.ExecContext(
		ctx,
		SqlUpdateDescription,
		&product.Description,
		&product.Id,
	)

	if err != nil {
		return nil, err
	}

	return product, nil

}

func (m mariaDBProductRepository) Delete(ctx context.Context, id int64) error {
	result, err := m.db.ExecContext(ctx, SqlDelete, id)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0 {
		return domain.ErrProductIdNotFound
	}

	return nil
}

func (m mariaDBProductRepository) GetAllReportProductRecords(ctx context.Context) (*[]domain.ProductRecordsReport, error) {
	var result []domain.ProductRecordsReport

	rows, err := m.db.QueryContext(ctx, SqlGetAllReportProductRecords)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		res := domain.ProductRecordsReport{}

		err := rows.Scan(
			&res.Id,
			&res.Description,
			&res.CountProductRecords,
		)
		if err != nil {
			return nil, err
		}

		result = append(result, res)
	}

	return &result, nil
}
