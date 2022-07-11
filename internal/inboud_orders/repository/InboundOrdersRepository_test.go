package repository_test

import (
	"context"
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/inboud_orders/domain"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/inboud_orders/repository"
)

func TestInboundOrdersRepository_Create(t *testing.T) {
	ctx := context.Background()

	now := time.Now()

	expectedInboundOrders := domain.InboundOrders{
		Id:             1,
		OrderDate:      now,
		OrderNumber:    "order#1",
		EmployeeId:     1,
		ProductBatchId: 1,
		WarehouseId:    1,
	}

	t.Run("create_ok: should create inbound order", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.
			ExpectExec(regexp.QuoteMeta(repository.SQLCreateInboundOrder)).
			WithArgs(now, "order#1", int64(1), int64(1), int64(1)).
			WillReturnResult(sqlmock.NewResult(1, 1))

		inboundOrdersRepository := repository.NewMariaDBInboundRepositoryRepository(db)

		result, err := inboundOrdersRepository.Create(
			ctx,
			now,
			"order#1",
			int64(1),
			int64(1),
			int64(1),
		)

		assert.Nil(t, err)
		assert.Equal(t, expectedInboundOrders, result)
	})

	t.Run("create_fail_exec: should return error when query fails", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.
			ExpectExec(regexp.QuoteMeta(repository.SQLCreateInboundOrder)).
			WillReturnError(fmt.Errorf("query error"))

		inboundOrdersRepository := repository.NewMariaDBInboundRepositoryRepository(db)

		_, err = inboundOrdersRepository.Create(
			ctx,
			now,
			"order#1",
			int64(1),
			int64(1),
			int64(1),
		)

		assert.Error(t, err)
	})
}
