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

func (m mariaDBProductRepository) GetAll(ctx context.Context) ([]domain.Product, error) {
	products := []domain.Product{}

	rows, err := m.db.QueryContext(ctx, sqlGetAll)

	if err != nil {
		return products, err
	}

	defer rows.Close()

	for rows.Next() {
		var product domain.Product

		err := rows.Scan(&product.Id, &product.ProductCode, &product.Description, &product.Width, &product.Height,
			&product.Length, &product.NetWeight, &product.ExpirationRate, &product.RecommendedFreezingTemperature,
			&product.FreezingRate, &product.ProductTypeId, &product.SellerId)
		if err != nil {
			return products, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (m mariaDBProductRepository) GetById(ctx context.Context, id int64) (*domain.Product, error) {
	row := m.db.QueryRowContext(ctx, sqlGetById, id)

	var product domain.Product

	err := row.Scan(&product.Id, &product.ProductCode, &product.Description, &product.Width, &product.Height,
		&product.Length, &product.NetWeight, &product.ExpirationRate, &product.RecommendedFreezingTemperature,
		&product.FreezingRate, &product.ProductTypeId, &product.SellerId)

	if errors.Is(err, sql.ErrNoRows) {
		return &product, domain.ErrIDNotFound
	}

	if err != nil {
		return &product, err
	}

	return &product, nil
}

func (m mariaDBProductRepository) Create(ctx context.Context, productCode string, description string, width float64, height float64, length float64,
	netWeight float64, expirationRate float64, recommendedFreezingTemperature float64, freezingRate float64, productTypeId int, sellerId int) (*domain.Product, error) {
	product, err := m.db.ExecContext(
		ctx,
		sqlCreate,
		productCode,
		description,
		width,
		height,
		length,
		netWeight,
		expirationRate,
		recommendedFreezingTemperature,
		freezingRate,
		productTypeId,
		sellerId,
	)

	if err != nil {
		return nil, err
	}

	newProductId, err := product.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &domain.Product{
		Id:                             newProductId,
		ProductCode:                    productCode,
		Description:                    description,
		Width:                          width,
		Height:                         height,
		Length:                         length,
		NetWeight:                      netWeight,
		ExpirationRate:                 expirationRate,
		RecommendedFreezingTemperature: recommendedFreezingTemperature,
		FreezingRate:                   freezingRate,
		ProductTypeId:                  productTypeId,
		SellerId:                       sellerId,
	}, nil
}

func (m mariaDBProductRepository) UpdateDescription(ctx context.Context, id int64, description string) (*domain.Product, error) {
	product := domain.Product{
		Id:          id,
		Description: description,
	}
	_, err := m.db.ExecContext(
		ctx,
		sqlUpdateDescription,
		&product.Description,
		&product.Id,
	)
	if err != nil {
		return nil, err
	}

	productUpdate, err := m.GetById(ctx, id)
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
