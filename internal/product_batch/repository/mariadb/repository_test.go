package repository_test

import (
	"context"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/product_batch/domain"
	repository "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/product_batch/repository/mariadb"
)

var timeNow = time.Now()

var expectedProductBatch = domain.ProductBatch{
	Id:                 1,
	BatchNumber:        1,
	CurrentQuantity:    1,
	CurrentTemperature: 1.0,
	DueDate:            timeNow,
	InitialQuantity:    1.0,
	ManufacturingDate:  timeNow,
	ManufacturingHour:  1,
	MinumumTemperature: 1.0,
	ProductId:          1,
	SectionId:          1,
}

func TestMariaDBProductRecordsRepository_Create(t *testing.T) {
	t.Run("create_ok: should create product batch", func(t *testing.T) {

		db, mock, err := sqlmock.New()

		assert.NoError(t, err)
		defer db.Close()

		mock.
			ExpectExec(regexp.QuoteMeta(repository.SQLCreate)).
			WithArgs(
				expectedProductBatch.BatchNumber,
				expectedProductBatch.CurrentQuantity,
				expectedProductBatch.CurrentTemperature,
				expectedProductBatch.DueDate,
				expectedProductBatch.InitialQuantity,
				expectedProductBatch.ManufacturingDate,
				expectedProductBatch.ManufacturingHour,
				expectedProductBatch.MinumumTemperature,
				expectedProductBatch.ProductId,
				expectedProductBatch.SectionId,
			).
			WillReturnResult(sqlmock.NewResult(1, 1))

		productBatchRepository := repository.NewMariadbProductBatchRepository(db)

		result, err := productBatchRepository.Create(context.TODO(), &expectedProductBatch)

		assert.NoError(t, err)
		assert.Equal(t, result, &expectedProductBatch)
	})

	t.Run("create_fail_exec: should return error when query execution fails", func(t *testing.T) {

		db, mock, err := sqlmock.New()

		assert.NoError(t, err)
		defer db.Close()

		mock.
			ExpectExec(regexp.QuoteMeta(repository.SQLCreate)).
			WithArgs(0, 0, 0, 0).
			WillReturnResult(sqlmock.NewResult(1, 1))

		productBatchRepository := repository.NewMariadbProductBatchRepository(db)
		_, err = productBatchRepository.Create(context.TODO(), &expectedProductBatch)

		assert.Error(t, err)
	})
}
