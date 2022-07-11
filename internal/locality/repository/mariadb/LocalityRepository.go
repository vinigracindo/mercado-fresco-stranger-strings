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

	var locality domain.LocalityModel

	if err := result.Scan(
		&locality.Id,
		&locality.LocalityName,
		&locality.ProvinceName,
		&locality.CountryName,
		&locality.ProvinceId,
	); err != nil {
		return nil, err
	}

	return &locality, nil
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

func (m repository) GetAllReportSeller(ctx context.Context) (*[]domain.ReportSeller, error) {
	rows, err := m.db.QueryContext(ctx, QueryGetAllLocality)

	if err != nil {
		return nil, err
	}

	var result []domain.ReportSeller

	for rows.Next() {
		report := domain.ReportSeller{}

		err := rows.Scan(
			&report.LocalityId,
			&report.LocalityName,
			&report.SellerCount)

		if err != nil {
			return nil, err
		}

		result = append(result, report)
	}

	return &result, nil
}
