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

func (m repository) CreateLocality(ctx context.Context, locality *domain.LocalityModel) (*domain.LocalityModel, error) {
	localityResult, err := m.db.ExecContext(
		ctx,
		QuerryCreateLocality,
		&locality.LocalityName,
		&locality.ProvinceName,
		&locality.CountryName,
		&locality.ProvinceId,
	)

	if err != nil {
		return nil, err
	}

	lastId, _ := localityResult.LastInsertId()

	locality.Id = lastId

	return locality, nil
}

func (m repository) CountByLocalityId(ctx context.Context, localityId int64) (int64, error) {
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
