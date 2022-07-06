package repository_test

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/seller/domain"
	repository "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/seller/repository/mariadb"
)

var expectedSeller = domain.Seller{
	Id:          1,
	Cid:         123,
	CompanyName: "Mercado Livre",
	Address:     "Osasco, SP",
	Telephone:   "11 99999999",
}

var expectedListSeller = []domain.Seller{
	{
		Id:          1,
		Cid:         123,
		CompanyName: "Mercado Livre",
		Address:     "Osasco, SP",
		Telephone:   "11 99999999",
	},
	{
		Id:          2,
		Cid:         456,
		CompanyName: "Mercado Pago",
		Address:     "Salvador, BA",
		Telephone:   "71 88888888",
	},
}

func TestSellerRepository_GetAll(t *testing.T) {
	t.Run("get_all_ok: should return all sellers", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		rows := sqlmock.NewRows([]string{
			"id",
			"cid",
			"company_name",
			"adress",
			"telephone",
		})

		for _, seller := range expectedListSeller {
			rows = rows.AddRow(
				&seller.Id,
				&seller.Cid,
				&seller.CompanyName,
				&seller.Address,
				&seller.Telephone,
			)
		}

		sellerRepository := repository.NewMariaDBSellerRepository(db)

		mock.
			ExpectQuery(regexp.QuoteMeta(repository.SqlGetAllSeller)).
			WillReturnRows(rows)

		result, err := sellerRepository.GetAll(context.TODO())

		assert.Nil(t, err)
		assert.Equal(t, &expectedListSeller, result)
	})

	t.Run("error_query_get_all: return error when try exec query", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		sellerRepository := repository.NewMariaDBSellerRepository(db)

		mock.
			ExpectQuery(regexp.QuoteMeta(repository.SqlGetAllSeller)).
			WillReturnError(fmt.Errorf("error: invalid query"))

		result, err := sellerRepository.GetAll(context.TODO())

		assert.Empty(t, result)
		assert.Error(t, err)
	})

	t.Run("error_scan_get_all: return error when try scan query", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		rows := sqlmock.NewRows([]string{
			"id",
			"cid",
			"company_name",
			"adress",
			"telephone",
		}).AddRow(nil, nil, nil, nil, nil)

		sellerRepository := repository.NewMariaDBSellerRepository(db)
		mock.
			ExpectQuery(regexp.QuoteMeta(repository.SqlGetAllSeller)).
			WillReturnRows(rows)

		result, err := sellerRepository.GetAll(context.TODO())

		assert.Empty(t, result)
		assert.Error(t, err)
	})
}

func TestSellerRepository_GetById(t *testing.T) {
	t.Run("get_by_id_ok: should return seller by id", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		row := sqlmock.NewRows([]string{
			"id",
			"cid",
			"company_name",
			"adress",
			"telephone",
		}).AddRow(
			&expectedSeller.Id,
			&expectedSeller.Cid,
			&expectedSeller.CompanyName,
			&expectedSeller.Address,
			&expectedSeller.Telephone,
		)

		sellerRepository := repository.NewMariaDBSellerRepository(db)

		mock.
			ExpectQuery(regexp.QuoteMeta(repository.SqlGetByIdSeller)).
			WithArgs(int64(1)).
			WillReturnRows(row)

		result, err := sellerRepository.GetById(context.TODO(), int64(1))

		assert.NoError(t, err)
		assert.Equal(t, &expectedSeller, result)
	})

	t.Run("get_by_id_scan_fails: should return an error when scan fails", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		row := sqlmock.NewRows([]string{
			"id",
			"cid",
			"company_name",
			"adress",
			"telephone",
		}).AddRow(nil, nil, nil, nil, nil)

		sellerRepository := repository.NewMariaDBSellerRepository(db)

		mock.ExpectQuery(repository.SqlGetByIdSeller).
			WithArgs(int64(9999)).
			WillReturnRows(row)

		result, err := sellerRepository.GetById(context.TODO(), 9999)

		assert.Empty(t, result)
		assert.Error(t, err)
	})

	t.Run("get_by_id_query_fails: should return error when query fails", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		sellerRepository := repository.NewMariaDBSellerRepository(db)

		mock.
			ExpectQuery(repository.SqlGetByIdSeller).
			WithArgs(int64(9999)).
			WillReturnError(fmt.Errorf("querry error"))

		result, err := sellerRepository.GetById(context.TODO(), int64(1))

		assert.Empty(t, result)
		assert.Error(t, err)
	})

}

func TestSellerRepository_Create(t *testing.T) {

	t.Run("creat_ok: Should creat a seller", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.
			ExpectExec(regexp.QuoteMeta(repository.SqlCreateSeller)).
			WithArgs(
				&expectedSeller.Cid,
				&expectedSeller.CompanyName,
				&expectedSeller.Address,
				&expectedSeller.Telephone).
			WillReturnResult(sqlmock.NewResult(1, 1))

		sellerRepository := repository.NewMariaDBSellerRepository(db)

		result, err := sellerRepository.Create(context.TODO(), &expectedSeller)

		assert.NoError(t, err)
		assert.Equal(t, result, &expectedSeller)
	})

	t.Run("create_query_error: Should return error when query execution fails", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.
			ExpectExec(regexp.QuoteMeta(repository.SqlCreateSeller)).
			WithArgs(0, 0, 0, 0).
			WillReturnResult(sqlmock.NewResult(1, 1))

		sellerRepository := repository.NewMariaDBSellerRepository(db)

		_, err = sellerRepository.Create(context.TODO(), &expectedSeller)

		assert.Error(t, err)

	})

}

func TestSellerRepository_Update(t *testing.T) {

	t.Run("Update_ok: should update seller adress and telephone", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.
			ExpectExec(regexp.QuoteMeta(repository.SqlUpdateSeller)).
			WithArgs(
				&expectedSeller.Address,
				&expectedSeller.Telephone,
				&expectedSeller.Id).
			WillReturnResult(sqlmock.NewResult(0, 1))

		sellerRepository := repository.NewMariaDBSellerRepository(db)

		result, err := sellerRepository.Update(context.TODO(), &expectedSeller)

		assert.NoError(t, err)
		assert.Equal(t, &expectedSeller, result)
	})

	t.Run("update_not_ok: return error when rows not affected", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.
			ExpectExec(regexp.QuoteMeta(repository.SqlUpdateSeller)).
			WithArgs(
				&expectedSeller.Address,
				&expectedSeller.Telephone,
				&expectedSeller.Id).
			WillReturnResult(sqlmock.NewResult(0, 0))

		sellerRepository := repository.NewMariaDBSellerRepository(db)

		_, err = sellerRepository.Update(context.TODO(), &expectedSeller)

		assert.Error(t, err)

	})

	t.Run("update_not_found: should return error when seller not found", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.
			ExpectExec(regexp.QuoteMeta(repository.SqlUpdateSeller)).
			WithArgs(
				&expectedSeller.Id,
				&expectedSeller.Address,
				&expectedSeller.Telephone).
			WillReturnResult(sqlmock.NewResult(0, 0))

		sellerRepository := repository.NewMariaDBSellerRepository(db)

		_, err = sellerRepository.Update(context.TODO(), &expectedSeller)

		assert.Error(t, err)
	})

	t.Run("update_query_fails: Should return error when query execution fails", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.
			ExpectExec(regexp.QuoteMeta(repository.SqlUpdateSeller)).
			WithArgs(&expectedSeller.Telephone).
			WillReturnError(errors.New("query executions fails"))

		sellerRepository := repository.NewMariaDBSellerRepository(db)

		_, err = sellerRepository.Update(context.TODO(), &expectedSeller)

		assert.Error(t, err)
	})
}

func TestSellerRepository_Delete(t *testing.T) {
	t.Run("delete_ok: should delete seller", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.
			ExpectExec(regexp.QuoteMeta(repository.SqlDeleteSeller)).
			WithArgs(int64(1)).
			WillReturnResult(sqlmock.NewResult(1, 1))

		sellerRepository := repository.NewMariaDBSellerRepository(db)

		err = sellerRepository.Delete(context.TODO(), int64(1))

		assert.NoError(t, err)
	})

	t.Run("delete_not_found: should return an error when seller is not found", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.
			ExpectExec(regexp.QuoteMeta(repository.SqlDeleteSeller)).
			WithArgs(int64(1)).
			WillReturnResult(sqlmock.NewResult(0, 0))

		sellerRepository := repository.NewMariaDBSellerRepository(db)

		err = sellerRepository.Delete(context.TODO(), int64(1))

		assert.Error(t, err)
	})

	t.Run("delete_query_fail: Should return error when query execution fails", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.
			ExpectExec(regexp.QuoteMeta(repository.SqlDeleteSeller)).
			WithArgs(int64(1)).
			WillReturnError(errors.New("query executions fails"))

		sellerRepository := repository.NewMariaDBSellerRepository(db)

		err = sellerRepository.Delete(context.TODO(), int64(1))

		assert.Error(t, err)
	})
}
