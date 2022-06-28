package repository

import (
	"context"
	"database/sql"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/product/domain"
)

type mariaDBProductRepository struct {
	db *sql.DB
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
	//TODO implement me
	panic("implement me")
}

func (m mariaDBProductRepository) Create(ctx context.Context, productCode string, description string, width float64, height float64, length float64, netWeight float64, expirationRate float64, recommendedFreezingTemperature float64, freezingRate float64, productTypeId int, sellerId int) (*domain.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (m mariaDBProductRepository) UpdateDescription(ctx context.Context, id int64, description string) (*domain.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (m mariaDBProductRepository) Delete(ctx context.Context, id int64) error {
	//TODO implement me
	panic("implement me")
}

func CreateProductRepository(db *sql.DB) domain.ProductRepository {
	return &mariaDBProductRepository{db: db}
}
