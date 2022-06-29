package repository

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

	rows, err := m.db.QueryContext(ctx, sqlGetAll)

	if err != nil {
		return &products, err
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
	row := m.db.QueryRowContext(ctx, sqlGetById, id)

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
		return nil, domain.ErrIDNotFound
	}

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (m mariaDBProductRepository) Create(ctx context.Context, product *domain.Product) (*domain.Product, error) {

	productResult, err := m.db.ExecContext(
		ctx,
		sqlCreate,
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

	lastId, err := productResult.LastInsertId()
	if err != nil {
		return nil, err
	}

	product.Id = lastId

	return product, nil
}

func (m mariaDBProductRepository) UpdateDescription(ctx context.Context, product *domain.Product) (*domain.Product, error) {

	productResult, err := m.db.ExecContext(
		ctx,
		sqlUpdateDescription,
		&product.Description,
		&product.Id,
	)
	if err != nil {
		return nil, err
	}

	affectedRows, err := productResult.RowsAffected()
	if affectedRows == 0 {
		return nil, domain.ErrIDNotFound
	}

	if err != nil {
		return nil, err
	}

	productUpdate, err := m.GetById(ctx, product.Id)
	if err != nil {
		return nil, err
	}

	return productUpdate, nil
	
}

func (m mariaDBProductRepository) Delete(ctx context.Context, id int64) error {
	result, err := m.db.ExecContext(ctx, sqlDelete, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return domain.ErrIDNotFound
	}

	return nil
}
