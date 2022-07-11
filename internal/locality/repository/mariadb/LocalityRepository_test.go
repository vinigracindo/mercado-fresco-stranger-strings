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
	ProvinceId:   1,
}

var expectedReportSeller = []domain.ReportSeller{
	{
		LocalityId:   1,
		LocalityName: "Salvador",
		SellerCount:  1,
	},
}

func Test_GetByIdRepository(t *testing.T) {
	t.Run("get_by_id_ok: shopuld return locality by id", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		row := sqlmock.NewRows([]string{
			"id",
			"locality_name",
			"province_name",
			"country_name",
			"province_id",
		}).AddRow(
			&expectedLocality.Id,
			&expectedLocality.LocalityName,
			&expectedLocality.ProvinceName,
			&expectedLocality.CountryName,
			&expectedLocality.ProvinceId,
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
	t.Run("creat_ok: Should creat a seller", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.
			ExpectExec(regexp.QuoteMeta(repository.QuerryCreateLocality)).
			WithArgs(
				&expectedLocality.LocalityName,
				&expectedLocality.ProvinceName,
				&expectedLocality.CountryName,
				&expectedLocality.ProvinceId).
			WillReturnResult(sqlmock.NewResult(1, 1))

		localityRepository := repository.NewMariadbLocalityRepository(db)

		result, err := localityRepository.CreateLocality(context.TODO(), &expectedLocality)

		assert.NoError(t, err)
		assert.Equal(t, result, &expectedLocality)
	})

	t.Run("create_query_error: Should return error when query execution fails", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.
			ExpectExec(regexp.QuoteMeta(repository.QuerryCreateLocality)).
			WithArgs(0, 0, 0, 0, 0).
			WillReturnResult(sqlmock.NewResult(0, 0))

		localityRepository := repository.NewMariadbLocalityRepository(db)

		_, err = localityRepository.CreateLocality(context.TODO(), &expectedLocality)

		assert.Error(t, err)

	})
}
