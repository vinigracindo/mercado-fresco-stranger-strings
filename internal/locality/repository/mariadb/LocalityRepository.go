package repository

import (
	"context"
	"database/sql"

	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/locality/domain"
)

type repository struct {
	db *sql.DB
}

func NewMariadbLocalityRepository(db *sql.DB) domain.LocalityRepository {
	return &repository{
		db: db,
	}
}

func (m repository) GetById(ctx context.Context, id int64) (*domain.LocalityModel, error) {
	result := m.db.QueryRowContext(ctx, QueryGetById, id)

	var carry domain.LocalityModel

	if err := result.Scan(
		&carry.Id,
		&carry.CountryName,
		&carry.ProvinceName,
		&carry.LocalityName,
	); err != nil {
		return nil, err
	}

	return &carry, nil
}
