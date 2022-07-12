package repository

import (
	"context"
	"database/sql"
	"log"

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
		&locality.CountryName,
		&locality.ProvinceName,
		&locality.LocalityName,
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

func getOrCreateCountry(ctx context.Context, db *sql.DB, countryName string) (int64, error) {
	result := db.QueryRowContext(ctx, QueryGetCountryByName, countryName)

	var id int64

	if err := result.Scan(&id); err != nil {
		if err == sql.ErrNoRows {
			result = db.QueryRowContext(ctx, QueryCreateCountry, countryName)

			if err := result.Scan(&id); err != nil {
				return 0, err
			}
		} else {
			return 0, err
		}
	}

	return id, nil
}

func getOrCreateProvince(ctx context.Context, db *sql.DB, countryId int64, provinceName string) (int64, error) {
	result := db.QueryRowContext(ctx, QueryGetProvinceByName, provinceName)

	var id int64

	if err := result.Scan(&id); err != nil {
		if err == sql.ErrNoRows {
			result = db.QueryRowContext(ctx, QueryCreateProvince, provinceName, countryId)

			if err := result.Scan(&id); err != nil {
				return 0, err
			}
		} else {
			return 0, err
		}
	}

	return id, nil
}

func (m repository) CreateLocality(ctx context.Context, locality *domain.LocalityModel) (*domain.LocalityModel, error) {
	transaction, _ := m.db.BeginTx(ctx, nil)
	stmt, _ := transaction.Prepare(QueryCreateLocality)

	country_id, _ := getOrCreateCountry(ctx, m.db, locality.CountryName)
	province_id, _ := getOrCreateProvince(ctx, m.db, country_id, locality.ProvinceName)

	localityResult, err := stmt.ExecContext(
		ctx,
		&locality.LocalityName,
		&province_id)

	if err != nil {
		transaction.Rollback()
		log.Print(err)
	}

	lastId, _ := localityResult.LastInsertId()

	locality.Id = lastId

	transaction.Commit()

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
