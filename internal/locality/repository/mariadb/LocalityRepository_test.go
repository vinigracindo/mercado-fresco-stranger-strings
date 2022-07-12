package repository_test

import (
	"context"
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/locality/domain"
	repository "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/locality/repository/mariadb"
)

var expectedLocality = domain.LocalityModel{
	Id:           1,
	LocalityName: "Salvador",
	ProvinceName: "Bahia",
	CountryName:  "Brasil",
}

var expectedReportSeller = []domain.ReportSeller{
	{
		LocalityId:   1,
		LocalityName: "Salvador",
		SellerCount:  1,
	},
}

var expectedCarries = []domain.ReportCarrie{
	{
		LocalityId:   1,
		LocalityName: "Salvador",
		CarriesCount: 1,
	},
	{
		LocalityId:   2,
		LocalityName: "Belo Horizonte",
		CarriesCount: 10000,
	},
}

func Test_GetByIdRepository(t *testing.T) {
	t.Run("get_by_id_ok: shopuld return locality by id", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		row := sqlmock.NewRows([]string{
			"id",
			"country_name",
			"province_name",
			"locality_name",
		}).AddRow(
			&expectedLocality.Id,
			&expectedLocality.CountryName,
			&expectedLocality.ProvinceName,
			&expectedLocality.LocalityName,
		)

		localityRepository := repository.NewMariadbLocalityRepository(db)

		mock.
			ExpectQuery(regexp.QuoteMeta(repository.QueryGetById)).
			WithArgs(int64(1)).
			WillReturnRows(row)

		result, err := localityRepository.GetById(context.TODO(), int64(1))

		assert.NoError(t, err)
		assert.Equal(t, result, &expectedLocality)

	})

	t.Run("get_by_id_scan_fails: should return an error when sacan fail", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		row := sqlmock.NewRows([]string{
			"id",
			"locality_name",
			"province_name",
			"country_name",
			"province_id",
		}).AddRow(nil, nil, nil, nil, nil)

		localityRepository := repository.NewMariadbLocalityRepository(db)

		mock.
			ExpectQuery(regexp.QuoteMeta(repository.QueryGetById)).
			WithArgs(int64(9999)).
			WillReturnRows(row)

		result, err := localityRepository.GetById(context.TODO(), int64(9999))

		assert.Error(t, err)
		assert.Empty(t, result)
	})
}

func Test_GetAllReportSellerRepository(t *testing.T) {
	t.Run("get_all_ok: should return all sellers", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		rows := sqlmock.NewRows([]string{
			"locality_id",
			"locality_name",
			"seller_count",
		})

		for _, locality := range expectedReportSeller {
			rows = rows.AddRow(
				&locality.LocalityId,
				&locality.LocalityName,
				&locality.SellerCount,
			)
		}

		localityRepository := repository.NewMariadbLocalityRepository(db)

		mock.
			ExpectQuery(regexp.QuoteMeta(repository.QueryGetAllLocality)).
			WillReturnRows(rows)

		result, err := localityRepository.GetAllReportSeller(context.TODO())

		assert.Nil(t, err)
		assert.Equal(t, &expectedReportSeller, result)
	})

	t.Run("error_scan_get_all: return error when try scan query", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		rows := sqlmock.NewRows([]string{
			"locality_id",
			"locality_name",
			"seller_count",
		}).AddRow(nil, nil, nil)

		localityRepository := repository.NewMariadbLocalityRepository(db)
		mock.
			ExpectQuery(regexp.QuoteMeta(repository.QueryGetAllLocality)).
			WillReturnRows(rows)

		result, err := localityRepository.GetAllReportSeller(context.TODO())

		assert.Empty(t, result)
		assert.Error(t, err)
	})

	t.Run("error_query_get_all: return error when try exec query", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		localityRepository := repository.NewMariadbLocalityRepository(db)

		mock.
			ExpectQuery(regexp.QuoteMeta(repository.QueryGetAllLocality)).
			WillReturnError(fmt.Errorf("error: invalid query"))

		result, err := localityRepository.GetAllReportSeller(context.TODO())

		assert.Empty(t, result)
		assert.Error(t, err)
	})
}

func Test_CreateLocalityRepository(t *testing.T) {
	t.Run("create_ok: Should create a locality when there are country and province records in database.", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectBegin()

		row_country := sqlmock.NewRows([]string{
			"id",
		}).AddRow(expectedLocality.Id)

		mock.ExpectPrepare(regexp.QuoteMeta(repository.QueryCreateLocality))

		mock.
			ExpectQuery(regexp.QuoteMeta(repository.QueryGetCountryByName)).
			WithArgs(&expectedLocality.CountryName).
			WillReturnRows(row_country)

		row_province := sqlmock.NewRows([]string{
			"id",
		}).AddRow(expectedLocality.Id)

		mock.
			ExpectQuery(regexp.QuoteMeta(repository.QueryGetProvinceByName)).
			WithArgs(&expectedLocality.ProvinceName).
			WillReturnRows(row_province)

		mock.
			ExpectExec(regexp.QuoteMeta(repository.QueryCreateLocality)).
			WithArgs(
				&expectedLocality.LocalityName,
				&expectedLocality.Id,
			).
			WillReturnResult(sqlmock.NewResult(1, 1))

		mock.ExpectCommit()

		localityRepository := repository.NewMariadbLocalityRepository(db)

		result, err := localityRepository.CreateLocality(context.TODO(), &expectedLocality)

		assert.NoError(t, err)
		assert.Equal(t, result, &expectedLocality)
	})

	t.Run("create_ok: Should create a locality when there is a country record in database but there is not a province.", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectBegin()

		row_country := sqlmock.NewRows([]string{
			"id",
		}).AddRow(expectedLocality.Id)

		mock.ExpectPrepare(regexp.QuoteMeta(repository.QueryCreateLocality))

		mock.
			ExpectQuery(regexp.QuoteMeta(repository.QueryGetCountryByName)).
			WithArgs(&expectedLocality.CountryName).
			WillReturnRows(row_country)

		mock.
			ExpectExec(regexp.QuoteMeta(repository.QueryCreateProvince)).
			WithArgs(
				&expectedLocality.ProvinceName,
				int64(1),
			).
			WillReturnResult(sqlmock.NewResult(1, 1))

		mock.
			ExpectExec(regexp.QuoteMeta(repository.QueryCreateLocality)).
			WithArgs(
				&expectedLocality.LocalityName,
				&expectedLocality.Id,
			).
			WillReturnResult(sqlmock.NewResult(1, 1))

		mock.ExpectCommit()

		localityRepository := repository.NewMariadbLocalityRepository(db)

		result, err := localityRepository.CreateLocality(context.TODO(), &expectedLocality)

		assert.NoError(t, err)
		assert.Equal(t, result, &expectedLocality)
	})

	t.Run("create_ok: Should create a locality when there is a province record in database but there is not a country.", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectBegin()

		mock.ExpectPrepare(regexp.QuoteMeta(repository.QueryCreateLocality))

		mock.
			ExpectExec(regexp.QuoteMeta(repository.QueryCreateCountry)).
			WithArgs(
				&expectedLocality.CountryName,
			).
			WillReturnResult(sqlmock.NewResult(1, 1))

		row_province := sqlmock.NewRows([]string{
			"id",
		}).AddRow(expectedLocality.Id)

		mock.
			ExpectQuery(regexp.QuoteMeta(repository.QueryGetProvinceByName)).
			WithArgs(&expectedLocality.ProvinceName).
			WillReturnRows(row_province)

		mock.
			ExpectExec(regexp.QuoteMeta(repository.QueryCreateLocality)).
			WithArgs(
				&expectedLocality.LocalityName,
				&expectedLocality.Id,
			).
			WillReturnResult(sqlmock.NewResult(1, 1))

		mock.ExpectCommit()

		localityRepository := repository.NewMariadbLocalityRepository(db)

		result, err := localityRepository.CreateLocality(context.TODO(), &expectedLocality)

		assert.NoError(t, err)
		assert.Equal(t, result, &expectedLocality)
	})

}

func Test_report_carry(t *testing.T) {
	t.Run("success_get_report", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		rows := sqlmock.NewRows([]string{
			"locality_id",
			"locality_name",
			"total_locality",
		}).AddRow(
			expectedCarries[0].LocalityId,
			expectedCarries[0].LocalityName,
			expectedCarries[0].CarriesCount,
		).AddRow(
			expectedCarries[1].LocalityId,
			expectedCarries[1].LocalityName,
			expectedCarries[1].CarriesCount)

		mock.ExpectQuery(regexp.QuoteMeta(repository.QueryCarryReport)).WillReturnRows(rows)

		localityRepository := repository.NewMariadbLocalityRepository(db)

		result, err := localityRepository.ReportCarrie(context.TODO(), 0)

		assert.NoError(t, err)
		assert.Equal(t, result, &expectedCarries)
	})

	t.Run("error_exec_query", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectQuery(regexp.QuoteMeta(repository.QueryCarryReport)).WithArgs("a").WillReturnError(fmt.Errorf("error: invalid id"))

		localityRepository := repository.NewMariadbLocalityRepository(db)

		result, err := localityRepository.ReportCarrie(context.TODO(), 0)

		assert.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("error_when_scan_query_result", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		rows := sqlmock.NewRows([]string{
			"locality_id",
			"locality_name",
		}).AddRow(
			expectedCarries[0].LocalityId,
			expectedCarries[0].LocalityName,
		).AddRow(
			expectedCarries[1].LocalityId,
			expectedCarries[1].LocalityName,
		)

		mock.ExpectQuery(regexp.QuoteMeta(repository.QueryCarryReport)).WillReturnRows(rows)

		localityRepository := repository.NewMariadbLocalityRepository(db)

		result, err := localityRepository.ReportCarrie(context.TODO(), 1)

		assert.Error(t, err)
		assert.Nil(t, result)
	})
}
