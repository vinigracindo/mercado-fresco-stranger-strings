package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/seller/domain"
)

type mariadbRepository struct {
	db *sql.DB
}

func NewMariaDBRepository(db *sql.DB) domain.RepositorySeller {
	return mariadbRepository{db: db}
}

func (m *mariadbRepository) GetAll(ctx context.Context) ([]domain.Seller, error) {
	listSeller := []domain.Seller{}

	rows, err := m.db.QueryContext(ctx, SqlGetAllSeller)
	if err != nil {
		return []domain.Seller{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var seller domain.Seller

		if err := rows.Scan(
			seller.Id,
			seller.Cid,
			seller.Address,
			seller.Telephone,
		); err != nil {
			return []domain.Seller{}, err
		}

		listSeller = append(listSeller, seller)
	}

	return listSeller, nil
}

func (m *mariadbRepository) GetById(ctx context.Context, id int64) (domain.Seller, error) {
	row := m.db.QueryRowContext(ctx, SqlGetByIdSeller)

	seller := domain.Seller{}

	err := row.Scan(
		seller.Id,
		seller.Cid,
		seller.Address,
		seller.Telephone,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return domain.Seller{}, err
	}

	if err != nil {
		return seller, err
	}

	return seller, nil
}

func (m *mariadbRepository) Create(ctx context.Context, seller *domain.Seller) (*domain.Seller, error) {
	sellerResult, err := m.db.ExecContext(
		ctx,
		SqlCreateSeller,
		&seller.Cid,
		&seller.CompanyName,
		&seller.Address,
		&seller.Telephone,
	)

	if err != nil {
		return nil, err
	}

	lastId, _ := sellerResult.LastInsertId()

	seller.Id = lastId

	return seller, nil
}

func (m *mariadbRepository) Update(ctx context.Context, seller *domain.Seller) (*domain.Seller, error) {
	sellerResult, err := m.db.ExecContext(
		ctx,
		SqlUpdateSeller,
		&seller.Address,
		&seller.Telephone,
		&seller.Id,
	)

	if err != nil {
		return nil, err
	}

	affectedRows, err := sellerResult.RowsAffected()

	if affectedRows == 0 {
		return nil, domain.ErrIDNotFound
	}

	if err != nil {
		return nil, err
	}

	return seller, nil

}

func (m *mariadbRepository) Delete(ctx context.Context, id int64) error {
	sellerResult, err := m.db.ExecContext(ctx, SqlDeleteSeller, id)

	if err != nil {
		return err
	}

	affectedRows, _ := sellerResult.RowsAffected()

	if affectedRows == 0 {
		return fmt.Errorf("seller with id %d not found", id)
	}

	return nil
}
