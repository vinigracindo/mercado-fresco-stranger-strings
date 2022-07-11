package service_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	buyerDomain "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/buyer/domain"
	buyerRepositoryMock "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/buyer/domain/mocks"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/purchase_orders/domain"
	purchaseOrdersRepositoryMock "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/purchase_orders/domain/mocks"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/purchase_orders/service"
)

var orderDateNow = time.Now()

var expectedPurchaseOrders = &domain.PurchaseOrders{
	Id:              1,
	OrderNumber:     "order#1",
	OrderDate:       orderDateNow,
	TrackingCode:    "abscf123",
	BuyerId:         1,
	ProductRecordId: 1,
	OrderStatusId:   1,
}

var ctx = context.Background()

var expectedBuyer = &buyerDomain.Buyer{
	Id:           1,
	CardNumberId: "402323",
	FirstName:    "FirstNameTest",
	LastName:     "LastNameTest",
}

func TestService_Create(t *testing.T) {
	repo := purchaseOrdersRepositoryMock.NewPurchaseOrdersRepository(t)
	buyerRepo := buyerRepositoryMock.NewBuyerRepository(t)
	service := service.NewPurchaseOrdersService(repo, buyerRepo)

	t.Run("crete_ok: when it contains the mandatory fields, should create a purchase orders", func(t *testing.T) {

		repo.On("Create",
			ctx,
			expectedPurchaseOrders.OrderNumber,
			expectedPurchaseOrders.OrderDate,
			expectedPurchaseOrders.TrackingCode,
			expectedPurchaseOrders.BuyerId,
			expectedPurchaseOrders.ProductRecordId,
			expectedPurchaseOrders.OrderStatusId,
		).
			Return(expectedPurchaseOrders, nil).
			Once()

		buyerRepo.
			On("GetId", ctx, int64(1)).
			Return(expectedBuyer, nil).
			Once()

		result, err := service.Create(ctx, "order#1", orderDateNow, "abscf123", 1, 1, 1)

		assert.Nil(t, err)
		assert.Equal(t, expectedPurchaseOrders, result)
	})

	t.Run("create_error: when create buyer records fails, should return error", func(t *testing.T) {

		repo.On("Create",
			context.TODO(),
			expectedPurchaseOrders.OrderNumber,
			expectedPurchaseOrders.OrderDate,
			expectedPurchaseOrders.TrackingCode,
			expectedPurchaseOrders.BuyerId,
			expectedPurchaseOrders.ProductRecordId,
			expectedPurchaseOrders.OrderStatusId,
		).
			Return(nil, buyerDomain.ErrIDNotFound).
			Once()

		buyerRepo.
			On("GetId", ctx, int64(1)).
			Return(nil, nil).
			Once()

		_, err := service.Create(ctx, "order#1", orderDateNow, "abscf123", 1, 1, 1)

		assert.Equal(t, buyerDomain.ErrIDNotFound, err)
		assert.Equal(t, nil, nil)

	})

	t.Run("create_error: when create buyer records fails, should return error", func(t *testing.T) {

		buyerRepo.
			On("GetId", ctx, int64(1)).
			Return(nil, buyerDomain.ErrIDNotFound).
			Once()

		_, err := service.Create(ctx, "order#1", orderDateNow, "abscf123", 1, 1, 1)

		assert.Equal(t, buyerDomain.ErrIDNotFound, err)
		assert.Equal(t, nil, nil)

	})

	t.Run("create_error: when the purchase order date is earlier than the current date should return an error", func(t *testing.T) {

		buyerRepo.
			On("GetId", ctx, int64(1)).
			Return(nil, nil).
			Once()

		date := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)

		purchaseOrdersIinvalidDate := domain.PurchaseOrders{
			Id:              1,
			OrderNumber:     "order#1",
			OrderDate:       date,
			TrackingCode:    "abscf123",
			BuyerId:         1,
			ProductRecordId: 1,
			OrderStatusId:   1,
		}

		_, err := service.Create(ctx,
			purchaseOrdersIinvalidDate.OrderNumber,
			purchaseOrdersIinvalidDate.OrderDate,
			purchaseOrdersIinvalidDate.TrackingCode,
			purchaseOrdersIinvalidDate.BuyerId,
			purchaseOrdersIinvalidDate.ProductRecordId,
			purchaseOrdersIinvalidDate.OrderStatusId,
		)
		assert.Equal(t, domain.ErrInvalidDate, err)
		assert.Equal(t, nil, nil)
	})

}
