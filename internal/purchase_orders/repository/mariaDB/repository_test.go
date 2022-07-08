package repository_test

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/purchase_orders/domain"
	repository "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/purchase_orders/repository/mariaDB"
)

var expectedPurchaseOrders = domain.PurchaseOrders{
	Id:              1,
	OrderNumber:     "order#1",
	OrderDate:       time.Now(),
	TrackingCode:    "abscf123",
	BuyerId:         1,
	ProductRecordId: 1,
	OrderStatusId:   1,
}

var mockPurchaseOrders = &domain.PurchaseOrders{
	Id:              1,
	OrderNumber:     "order#1",
	OrderDate:       time.Now(),
	TrackingCode:    "abscf123",
	BuyerId:         1,
	ProductRecordId: 1,
	OrderStatusId:   1,
}

func TestPurchaseOrderRepository_Create(t *testing.T) {
	t.Run("create_ok: should create purchase orders", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(regexp.QuoteMeta(repository.SQLCreatePurchaseOrders)).WithArgs(
			mockPurchaseOrders.OrderNumber,
			mockPurchaseOrders.OrderDate,
			mockPurchaseOrders.TrackingCode,
			mockPurchaseOrders.BuyerId,
			mockPurchaseOrders.ProductRecordId,
			mockPurchaseOrders.OrderStatusId,
		).WillReturnResult(sqlmock.NewResult(1, 1))

		purchaseOrdersRepository := repository.NewMariadbPurchaseOrdersRepository(db)

		newPurchaseOrders, err := purchaseOrdersRepository.Create(context.Background(),
			mockPurchaseOrders.OrderNumber,
			mockPurchaseOrders.OrderDate,
			mockPurchaseOrders.TrackingCode,
			mockPurchaseOrders.BuyerId,
			mockPurchaseOrders.ProductRecordId,
			mockPurchaseOrders.OrderStatusId,
		)

		assert.NoError(t, err)
		assert.Equal(t, newPurchaseOrders, mockPurchaseOrders)
	})

	t.Run("create_error: should return error when query execution fails", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(regexp.QuoteMeta(repository.SQLCreatePurchaseOrders)).WithArgs(
			mockPurchaseOrders.OrderNumber,
			mockPurchaseOrders.OrderDate,
			mockPurchaseOrders.TrackingCode,
			mockPurchaseOrders.BuyerId,
			mockPurchaseOrders.ProductRecordId,
			mockPurchaseOrders.OrderStatusId,
		).WillReturnError(fmt.Errorf("erro"))

		purchaseOrdersRepository := repository.NewMariadbPurchaseOrdersRepository(db)

		newPurchaseOrders, err := purchaseOrdersRepository.Create(context.Background(),
			mockPurchaseOrders.OrderNumber,
			mockPurchaseOrders.OrderDate,
			mockPurchaseOrders.TrackingCode,
			mockPurchaseOrders.BuyerId,
			mockPurchaseOrders.ProductRecordId,
			mockPurchaseOrders.OrderStatusId,
		)

		assert.Error(t, err)
		assert.Empty(t, newPurchaseOrders)

	})

}

func TestPurchaseOrderRepository_ContByBuyerId(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		var mockPurchaseOrdersCont int64

		row := sqlmock.NewRows([]string{"conts"}).AddRow(mockPurchaseOrdersCont)

		mock.ExpectQuery(regexp.QuoteMeta(repository.SQLContByBuyerId)).
			WithArgs(1).
			WillReturnRows(row)

		purchaseOrdersRepository := repository.NewMariadbPurchaseOrdersRepository(db)

		result, err := purchaseOrdersRepository.ContByBuyerId(context.TODO(), 1)

		assert.NoError(t, err)
		assert.Equal(t, mockPurchaseOrdersCont, result)
	})

	t.Run("error - not found", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectQuery(regexp.QuoteMeta(repository.SQLContByBuyerId)).
			WithArgs(1).
			WillReturnError(errors.New("sql: no rows in result set"))

		purchaseOrdersRepository := repository.NewMariadbPurchaseOrdersRepository(db)

		_, err = purchaseOrdersRepository.ContByBuyerId(context.TODO(), 1)

		assert.Error(t, err)
	})
}
