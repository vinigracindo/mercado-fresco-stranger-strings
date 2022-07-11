package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/buyer/domain"
)

type mariadbBuyerRepository struct {
	db *sql.DB
}

func NewmariadbBuyerRepository(db *sql.DB) domain.BuyerRepository {
	return &mariadbBuyerRepository{db: db}
}

func (repo *mariadbBuyerRepository) Create(ctx context.Context, cardNumberId, firstName, lastName string) (*domain.Buyer, error) {
	var newBuyer domain.Buyer

	result, err := repo.db.ExecContext(
		ctx,
		SQLCreateBuyer,
		cardNumberId,
		firstName,
		lastName,
	)

	if err != nil {
		return nil, err
	}

	lastId, _ := result.LastInsertId()

	newBuyer.Id = lastId

	return &domain.Buyer{
		Id:           newBuyer.Id,
		CardNumberId: cardNumberId,
		FirstName:    firstName,
		LastName:     lastName,
	}, nil
}

func (repo *mariadbBuyerRepository) GetAll(ctx context.Context) (*[]domain.Buyer, error) {

	buyers := []domain.Buyer{}

	rows, err := repo.db.QueryContext(ctx, SQLGetAllBuyer)

	if err != nil {
		return &buyers, err
	}
	defer rows.Close()

	for rows.Next() {
		var buyer domain.Buyer
		if err := rows.Scan(
			&buyer.Id,
			&buyer.CardNumberId,
			&buyer.FirstName,
			&buyer.LastName,
		); err != nil {
			return &buyers, err
		}
		buyers = append(buyers, buyer)
	}
	return &buyers, nil
}

func (repo *mariadbBuyerRepository) GetId(ctx context.Context, id int64) (*domain.Buyer, error) {

	row := repo.db.QueryRowContext(ctx, SQLGetByIdBuyer, id)

	var buyer domain.Buyer

	err := row.Scan(
		&buyer.Id,
		&buyer.CardNumberId,
		&buyer.FirstName,
		&buyer.LastName,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return &buyer, domain.ErrBuyerNotFound
	}

	if err != nil {
		return &buyer, err
	}
	return &buyer, nil
}

func (repo *mariadbBuyerRepository) Update(ctx context.Context, id int64, cardNumberId, lastName string) (*domain.Buyer, error) {

	result, err := repo.db.ExecContext(
		ctx,
		SQLUpdateBuyer,
		cardNumberId,
		lastName,
		id,
	)

	if err != nil {
		return nil, err
	}

	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0 {
		return nil, domain.ErrBuyerNotFound
	}
	return &domain.Buyer{
		Id:           id,
		CardNumberId: cardNumberId,
		LastName:     lastName,
	}, nil
}

func (repo *mariadbBuyerRepository) Delete(ctx context.Context, id int64) error {
	result, err := repo.db.ExecContext(ctx, SQLDeleteBuyer, id)
	if err != nil {
		return err
	}

	affectRows, _ := result.RowsAffected()

	if affectRows == 0 {
		return domain.ErrBuyerNotFound
	}
	return nil
}

func (repo *mariadbBuyerRepository) GetAllPurchaseOrdersReports(ctx context.Context) (*[]domain.PurchaseOrdersReport, error) {
	var result []domain.PurchaseOrdersReport

	rows, err := repo.db.QueryContext(ctx, SQLGetAllPurchaseOrdersReports)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		res := domain.PurchaseOrdersReport{}

		err := rows.Scan(
			&res.Id,
			&res.CardNumberId,
			&res.FirstName,
			&res.LastName,
			&res.CountBuyersRecords,
		)
		if err != nil {
			return nil, err
		}

		result = append(result, res)
	}

	return &result, nil
}
