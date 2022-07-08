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

func (m repository) ReportCarrie(ctx context.Context, id int64) (*[]domain.ReportCarrie, error) {
	result, err := m.db.QueryContext(ctx, QueryCarryReport, id, id)

	if err != nil {
		return nil, err
	}

	defer result.Close()

	var listReport []domain.ReportCarrie

	for result.Next() {
		row := domain.ReportCarrie{}

		err = result.Scan(
			&row.LocalityId,
			&row.LocalityName,
			&row.CarriesCount,
		)

		if err != nil {
			return nil, err
		}

		listReport = append(listReport, row)

	}

	return &listReport, nil
}
