package repository_test

import (
	"context"
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/purchase_orders/domain"
	repository "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/purchase_orders/repository/mariaDB"
)

var expectedPurchaseOrders = domain.PurchaseOrders{
	Id:              1,
	OrderNumber:     "order#1",
	OrderDate:       "2021-04-04",
	TrackingCode:    "abscf123",
	BuyerId:         1,
	ProductRecordId: 1,
	OrderStatusId:   1,
}

var mockPurchaseOrders = &domain.PurchaseOrders{
	Id:              1,
	OrderNumber:     "order#1",
	OrderDate:       "2021-04-04",
	TrackingCode:    "abscf123",
	BuyerId:         1,
	ProductRecordId: 1,
	OrderStatusId:   1,
}

func TestBuyerRepository_Create(t *testing.T) {
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
