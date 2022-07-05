package respository

import (
	"context"
	"database/sql"

	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/carry/domain"
)

type mariadbCarry struct {
	db *sql.DB
}

func NewMariadbCarryRepository(db *sql.DB) domain.CarryRepository {
	return &mariadbCarry{
		db: db,
	}
}

// (cid, company_name, address, telephone, locality_id)
func (m mariadbCarry) Create(ctx context.Context, carry *domain.CarryModel) (*domain.CarryModel, error) {

	result, err := m.db.ExecContext(
		ctx,
		QueryCreateCarry,
		carry.Cid,
		carry.CompanyName,
		carry.Address,
		carry.Telephone,
		carry.LocalityID,
	)

	if err != nil {
		return nil, err
	}

	newCarryId, _ := result.LastInsertId()

	carry.Id = newCarryId

	return carry, nil
}

func (m mariadbCarry) GetById(ctx context.Context, id int64) (*domain.CarryModel, error) {

	result := m.db.QueryRowContext(ctx, QueryGetCarry, id)

	var carry domain.CarryModel

	if err := result.Scan(
		&carry.Id,
		&carry.Cid,
		&carry.CompanyName,
		&carry.Address,
		&carry.Telephone,
		&carry.LocalityID,
	); err != nil {
		return nil, err
	}

	return &carry, nil

}

func (m mariadbCarry) CountLocality(ctx context.Context, locality_id int64) (int64, error) {
	result := m.db.QueryRowContext(ctx, QueryCountLocality, locality_id)

	var count int64

	if err := result.Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}
