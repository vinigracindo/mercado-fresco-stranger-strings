package mariadb_test

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/product/domain"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/product/repository/mariadb"
	"regexp"
	"testing"
)

var expectedProduct = domain.Product{
	Id:                             1,
	ProductCode:                    "PROD01",
	Description:                    "Yogurt",
	Width:                          1.2,
	Height:                         6.4,
	Length:                         4.5,
	NetWeight:                      3.4,
	ExpirationRate:                 1.5,
	RecommendedFreezingTemperature: 1.3,
	FreezingRate:                   2,
	ProductTypeId:                  2,
	SellerId:                       2,
}

var expectedProductList = []domain.Product{
	{
		Id:                             1,
		ProductCode:                    "PROD01",
		Description:                    "Yogurt",
		Width:                          1.2,
		Height:                         6.4,
		Length:                         4.5,
		NetWeight:                      3.4,
		ExpirationRate:                 1.5,
		RecommendedFreezingTemperature: 1.3,
		FreezingRate:                   2,
		ProductTypeId:                  2,
		SellerId:                       2,
	},
	{
		Id:                             2,
		ProductCode:                    "PROD02",
		Description:                    "Yogurt",
		Width:                          1.2,
		Height:                         6.4,
		Length:                         4.5,
		NetWeight:                      3.4,
		ExpirationRate:                 1.5,
		RecommendedFreezingTemperature: 1.3,
		FreezingRate:                   2,
		ProductTypeId:                  2,
		SellerId:                       2,
	},
}

func TestMariaDBProductRepository_GetAll(t *testing.T) {

	t.Run("get_all_ok: should return all products", func(t *testing.T) {

		db, mock, err := sqlmock.New()

		assert.NoError(t, err)
		defer db.Close()

		rows := sqlmock.NewRows([]string{
			"id",
			"product_code",
			"description",
			"width",
			"height",
			"length",
			"net_weight",
			"expiration_rate",
			"recommended_freezing_temperature",
			"freezing_rate",
			"product_type_id",
			"seller_id"})

		for _, product := range expectedProductList {
			rows = rows.AddRow(
				product.Id,
				product.ProductCode,
				product.Description,
				product.Width,
				product.Height,
				product.Length,
				product.NetWeight,
				product.ExpirationRate,
				product.RecommendedFreezingTemperature,
				product.FreezingRate,
				product.ProductTypeId,
				product.SellerId)
		}

		productRepository := mariadb.CreateProductRepository(db)

		mock.
			ExpectQuery(regexp.QuoteMeta(mariadb.SqlGetAll)).
			WillReturnRows(rows)

		result, err := productRepository.GetAll(context.TODO())

		assert.Nil(t, err)
		assert.Equal(t, &expectedProductList, result)
	})

	t.Run("get_all_scan_fails: should return error when scan fail", func(t *testing.T) {

		db, mock, err := sqlmock.New()

		assert.NoError(t, err)
		defer db.Close()

		rows := sqlmock.NewRows([]string{
			"id",
			"product_code",
			"description",
			"width",
			"height",
			"length",
			"net_weight",
			"expiration_rate",
			"recommended_freezing_temperature",
			"freezing_rate",
			"product_type_id",
			"seller_id"}).
			AddRow("", "", "", "", "", "", "", "", "", "", "", "")

		mock.
			ExpectQuery(regexp.QuoteMeta(mariadb.SqlGetAll)).
			WillReturnRows(rows)

		productRepository := mariadb.CreateProductRepository(db)

		_, err = productRepository.GetAll(context.TODO())
		assert.Error(t, err)
	})

	t.Run("get_all_falis: should return error when query fails", func(t *testing.T) {

		db, mock, err := sqlmock.New()

		assert.NoError(t, err)
		defer db.Close()

		productRepository := mariadb.CreateProductRepository(db)

		mock.
			ExpectQuery(regexp.QuoteMeta(mariadb.SqlGetAll)).
			WillReturnError(fmt.Errorf("query error"))

		result, err := productRepository.GetAll(context.TODO())

		assert.Error(t, err)
		assert.Empty(t, result)
	})
}

func TestMariaDBProductRepository_GetById(t *testing.T) {

	t.Run("get_by_id_ok: should return product by id ", func(t *testing.T) {

		db, mock, err := sqlmock.New()

		assert.NoError(t, err)
		defer db.Close()

		row := sqlmock.NewRows([]string{
			"id",
			"product_code",
			"description",
			"width",
			"height",
			"length",
			"net_weight",
			"expiration_rate",
			"recommended_freezing_temperature",
			"freezing_rate",
			"product_type_id",
			"seller_id"}).
			AddRow(
				expectedProduct.Id,
				expectedProduct.ProductCode,
				expectedProduct.Description,
				expectedProduct.Width,
				expectedProduct.Height,
				expectedProduct.Length,
				expectedProduct.NetWeight,
				expectedProduct.ExpirationRate,
				expectedProduct.RecommendedFreezingTemperature,
				expectedProduct.FreezingRate,
				expectedProduct.ProductTypeId,
				expectedProduct.SellerId)

		mock.
			ExpectQuery(regexp.QuoteMeta(mariadb.SqlGetById)).
			WithArgs(expectedProduct.Id).
			WillReturnRows(row)

		productRepository := mariadb.CreateProductRepository(db)

		result, err := productRepository.GetById(context.TODO(), expectedProduct.Id)

		assert.NoError(t, err)
		assert.Equal(t, result.Id, int64(1))
	})

	t.Run("get_by_id_scan_fails: should return error when scan fail", func(t *testing.T) {

		db, mock, err := sqlmock.New()

		assert.NoError(t, err)
		defer db.Close()

		row := sqlmock.NewRows([]string{
			"id",
			"product_code",
			"description",
			"width",
			"height",
			"length",
			"net_weight",
			"expiration_rate",
			"recommended_freezing_temperature",
			"freezing_rate",
			"product_type_id",
			"seller_id"}).
			AddRow("", "", "", "", "", "", "", "", "", "", "", "")

		mock.
			ExpectQuery(regexp.QuoteMeta(mariadb.SqlGetById)).WillReturnRows(row)

		productRepository := mariadb.CreateProductRepository(db)

		_, err = productRepository.GetById(context.TODO(), 1)
		assert.Error(t, err)
	})

	t.Run("get_by_id_not_found: should return error when product not found", func(t *testing.T) {
		db, mock, err := sqlmock.New()

		assert.NoError(t, err)
		defer db.Close()

		mock.
			ExpectQuery(regexp.QuoteMeta(mariadb.SqlGetById)).
			WillReturnError(sql.ErrNoRows)

		productRepository := mariadb.CreateProductRepository(db)

		_, err = productRepository.GetById(context.TODO(), 1)
		assert.Error(t, err)
	})
}

func TestMariaDBProductRepository_Create(t *testing.T) {

	t.Run("create_ok: should create product", func(t *testing.T) {

		db, mock, err := sqlmock.New()

		assert.NoError(t, err)
		defer db.Close()

		mock.
			ExpectExec(regexp.QuoteMeta(mariadb.SqlCreate)).
			WithArgs(
				expectedProduct.ProductCode,
				expectedProduct.Description,
				expectedProduct.Width,
				expectedProduct.Height,
				expectedProduct.Length,
				expectedProduct.NetWeight,
				expectedProduct.ExpirationRate,
				expectedProduct.RecommendedFreezingTemperature,
				expectedProduct.FreezingRate,
				expectedProduct.ProductTypeId,
				expectedProduct.SellerId).
			WillReturnResult(sqlmock.NewResult(1, 1))

		productRepository := mariadb.CreateProductRepository(db)

		result, err := productRepository.Create(context.TODO(), &expectedProduct)

		assert.NoError(t, err)
		assert.Equal(t, result, &expectedProduct)
	})

	t.Run("create_fail_exec: should return error when query execution fails", func(t *testing.T) {

		db, mock, err := sqlmock.New()

		assert.NoError(t, err)
		defer db.Close()

		mock.
			ExpectExec(regexp.QuoteMeta(mariadb.SqlCreate)).
			WithArgs(0, 0, 0, 0, 0, 0, 0, 0).
			WillReturnResult(sqlmock.NewResult(1, 1))

		productRepository := mariadb.CreateProductRepository(db)
		_, err = productRepository.Create(context.TODO(), &expectedProduct)

		assert.Error(t, err)
	})
}

func TestSectionRepository_Update(t *testing.T) {

	dummyUpdatedProduct := domain.Product{
		Id:          1,
		Description: "Strawberry yogurt",
	}

	t.Run("update_not_found: should return error when product not found", func(t *testing.T) {

		db, mock, err := sqlmock.New()

		assert.NoError(t, err)
		defer db.Close()

		mock.
			ExpectExec(regexp.QuoteMeta(mariadb.SqlUpdateDescription)).
			WithArgs(expectedProduct.Id, dummyUpdatedProduct.Description).
			WillReturnResult(sqlmock.NewResult(0, 0))

		productRepository := mariadb.CreateProductRepository(db)

		_, err = productRepository.UpdateDescription(context.TODO(), &dummyUpdatedProduct)

		assert.Error(t, err)
	})

	t.Run("update_fail: should return error when query execution fails", func(t *testing.T) {

		db, mock, err := sqlmock.New()

		assert.NoError(t, err)
		defer db.Close()

		mock.
			ExpectExec(regexp.QuoteMeta(mariadb.SqlUpdateDescription)).
			WithArgs(expectedProduct.Id).
			WillReturnError(errors.New("any error"))

		productRepository := mariadb.CreateProductRepository(db)

		_, err = productRepository.UpdateDescription(context.TODO(), &dummyUpdatedProduct)

		assert.Error(t, err)
	})

	t.Run("update_not_ok: return error when rows not affected", func(t *testing.T) {

		db, mock, err := sqlmock.New()

		assert.NoError(t, err)
		defer db.Close()

		mock.
			ExpectExec(regexp.QuoteMeta(mariadb.SqlUpdateDescription)).
			WithArgs(dummyUpdatedProduct.Description, expectedProduct.Id).
			WillReturnResult(sqlmock.NewResult(0, 0))

		productRepository := mariadb.CreateProductRepository(db)

		_, err = productRepository.UpdateDescription(context.TODO(), &dummyUpdatedProduct)

		assert.Error(t, err)
	})

	t.Run("update_ok: should update product code", func(t *testing.T) {

		db, mock, err := sqlmock.New()

		assert.NoError(t, err)
		defer db.Close()

		mock.
			ExpectExec(regexp.QuoteMeta(mariadb.SqlUpdateDescription)).
			WithArgs(dummyUpdatedProduct.Description, expectedProduct.Id).
			WillReturnResult(sqlmock.NewResult(0, 1))

		productRepository := mariadb.CreateProductRepository(db)

		prod, err := productRepository.UpdateDescription(context.TODO(), &dummyUpdatedProduct)

		assert.NoError(t, err)
		assert.Equal(t, dummyUpdatedProduct, *prod)
	})

}

func TestSectionRepository_Delete(t *testing.T) {

	id := int64(1)

	t.Run("delete_ok: should delete product", func(t *testing.T) {

		db, mock, err := sqlmock.New()

		assert.NoError(t, err)
		defer db.Close()

		mock.
			ExpectExec(regexp.QuoteMeta(mariadb.SqlDelete)).
			WithArgs(id).
			WillReturnResult(sqlmock.NewResult(0, 1))

		productRepository := mariadb.CreateProductRepository(db)

		err = productRepository.Delete(context.TODO(), id)

		assert.NoError(t, err)
	})

	t.Run("delete_affect_rows_0: return error when rows not affected", func(t *testing.T) {

		db, mock, err := sqlmock.New()

		assert.NoError(t, err)
		defer db.Close()

		mock.
			ExpectExec(regexp.QuoteMeta(mariadb.SqlDelete)).
			WithArgs(id).
			WillReturnResult(sqlmock.NewResult(0, 0))

		productRepository := mariadb.CreateProductRepository(db)

		err = productRepository.Delete(context.TODO(), id)

		assert.Error(t, err)
	})

	t.Run("delete_not_found: should return error when product not found", func(t *testing.T) {

		db, mock, err := sqlmock.New()

		assert.NoError(t, err)
		defer db.Close()

		mock.
			ExpectExec(regexp.QuoteMeta(mariadb.SqlDelete)).
			WithArgs(id).
			WillReturnResult(sqlmock.NewResult(0, 0))

		productRepository := mariadb.CreateProductRepository(db)

		err = productRepository.Delete(context.TODO(), id)

		assert.Error(t, err)
	})

	t.Run("delete_fail: should return error when query execution fails", func(t *testing.T) {

		db, mock, err := sqlmock.New()

		assert.NoError(t, err)
		defer db.Close()

		mock.
			ExpectExec(regexp.QuoteMeta(mariadb.SqlDelete)).
			WithArgs(id).
			WillReturnError(errors.New("any error"))

		productRepository := mariadb.CreateProductRepository(db)

		err = productRepository.Delete(context.TODO(), id)

		assert.Error(t, err)
	})
}
