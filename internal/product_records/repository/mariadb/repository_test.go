package mariadb_test

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/product_records/domain"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/product_records/repository/mariadb"
	"regexp"
	"testing"
	"time"
)

var expectedProductRecords = domain.ProductRecords{
	Id:             1,
	LastUpdateDate: time.Now(),
	PurchasePrice:  10.5,
	SalePrice:      15.2,
	ProductId:      1,
}

func TestMariaDBProductRecordsRepository_Create(t *testing.T) {

	t.Run("create_ok: should create product records", func(t *testing.T) {

		db, mock, err := sqlmock.New()

		assert.NoError(t, err)
		defer db.Close()

		mock.
			ExpectExec(regexp.QuoteMeta(mariadb.SqlCreate)).
			WithArgs(
				expectedProductRecords.LastUpdateDate,
				expectedProductRecords.PurchasePrice,
				expectedProductRecords.SalePrice,
				expectedProductRecords.ProductId).
			WillReturnResult(sqlmock.NewResult(1, 1))

		productRecordsRepository := mariadb.CreateProductRecordsRepository(db)

		result, err := productRecordsRepository.Create(context.TODO(), &expectedProductRecords)

		assert.NoError(t, err)
		assert.Equal(t, result, &expectedProductRecords)
	})

	t.Run("create_fail_exec: should return error when query execution fails", func(t *testing.T) {

		db, mock, err := sqlmock.New()

		assert.NoError(t, err)
		defer db.Close()

		mock.
			ExpectExec(regexp.QuoteMeta(mariadb.SqlCreate)).
			WithArgs(0, 0, 0, 0).
			WillReturnResult(sqlmock.NewResult(1, 1))

		productRecordsRepository := mariadb.CreateProductRecordsRepository(db)
		_, err = productRecordsRepository.Create(context.TODO(), &expectedProductRecords)

		assert.Error(t, err)
	})
}
