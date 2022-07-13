package repository

import (
	"context"
	"database/sql"

	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/seller/domain"
)

type mariaDBSellerRepository struct {
	db *sql.DB
}

func NewMariaDBSellerRepository(db *sql.DB) domain.RepositorySeller {
	return &mariaDBSellerRepository{db: db}
}

func (m *mariaDBSellerRepository) GetAll(ctx context.Context) (*[]domain.Seller, error) {
	listSeller := []domain.Seller{}

	rows, err := m.db.QueryContext(ctx, SqlGetAllSeller)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var seller domain.Seller

		if err := rows.Scan(
			&seller.Id,
			&seller.Cid,
			&seller.CompanyName,
			&seller.Address,
			&seller.Telephone,
			&seller.LocalityId,
		); err != nil {
			return nil, err
		}

		listSeller = append(listSeller, seller)
	}

	return &listSeller, nil
}

func (m *mariaDBSellerRepository) GetById(ctx context.Context, id int64) (*domain.Seller, error) {
	row := m.db.QueryRowContext(ctx, SqlGetByIdSeller, id)

	var seller domain.Seller

	err := row.Scan(
		&seller.Id,
		&seller.Cid,
		&seller.CompanyName,
		&seller.Address,
		&seller.Telephone,
		&seller.LocalityId,
	)

	if err != nil {
		return nil, domain.ErrIDNotFound
	}

	return &seller, nil
}

func (m *mariaDBSellerRepository) Create(ctx context.Context, seller *domain.Seller) (*domain.Seller, error) {
	sellerResult, err := m.db.ExecContext(
		ctx,
		SqlCreateSeller,
		&seller.Cid,
		&seller.CompanyName,
		&seller.Address,
		&seller.Telephone,
		&seller.LocalityId,
	)

	if err != nil {
		return nil, err
	}

	lastId, _ := sellerResult.LastInsertId()

	seller.Id = lastId

	return seller, nil
}

func (m *mariaDBSellerRepository) Update(ctx context.Context, seller *domain.Seller) (*domain.Seller, error) {
	_, err := m.db.ExecContext(
		ctx,
		SqlUpdateSeller,
		&seller.Address,
		&seller.Telephone,
		&seller.Id,
	)

	if err != nil {
		return nil, err
	}

	return seller, nil

}

func (m *mariaDBSellerRepository) Delete(ctx context.Context, id int64) error {
	sellerResult, err := m.db.ExecContext(ctx, SqlDeleteSeller, id)

	if err != nil {
		return err
	}

	affectedRows, _ := sellerResult.RowsAffected()

	if affectedRows == 0 {
		return domain.ErrIDNotFound
	}

	return nil
}

func (m *mariaDBSellerRepository) CountByLocalityId(ctx context.Context, localityId int64) (int64, error) {
	rows := m.db.QueryRowContext(
		ctx,
		QueryCountByLocalityId,
		localityId,
	)

	var countSellersInLocalityId int64

	err := rows.Scan(&countSellersInLocalityId)

	if err != nil {
		return 0, err
	}

	return countSellersInLocalityId, nil
}
