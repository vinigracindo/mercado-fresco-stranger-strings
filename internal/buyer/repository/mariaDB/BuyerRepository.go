package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/buyer/domain"
)

//var buyers = []domain.Buyer{}
//var lastId int64 = 0

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
		return &domain.Buyer{}, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return &newBuyer, err
	}

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

	_, err := repo.db.ExecContext(
		ctx,
		SQLUpdateAwardBuyer,
		cardNumberId,
		lastName,
		id,
	)
	if err != nil {
		return &domain.Buyer{}, err
	}

	buyerUpdated, err := repo.GetId(ctx, id)
	if err != nil {
		return &domain.Buyer{}, err

	}

	return buyerUpdated, nil
}

func (repo *mariadbBuyerRepository) Delete(ctx context.Context, id int64) error {
	result, err := repo.db.ExecContext(ctx, SQLDeleteBuyer, id)
	if err != nil {
		return err
	}

	affectRows, err := result.RowsAffected()

	if affectRows == 0 {
		return domain.ErrBuyerNotFound
	}

	if err != nil {
		return err
	}
	return nil
}
