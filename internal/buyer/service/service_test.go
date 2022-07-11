package service_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/buyer/domain"
	buyerRepositoryMock "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/buyer/domain/mocks"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/buyer/service"
	mockPurchaseOrder "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/purchase_orders/domain/mocks"
)

var expectedBuyer = &domain.Buyer{
	Id:           1,
	CardNumberId: "402323",
	FirstName:    "FirstNameTest",
	LastName:     "LastNameTest",
}

var expectedBuyerList = &[]domain.Buyer{
	{
		CardNumberId: "402323",
		FirstName:    "FirstNameTest",
		LastName:     "LastNameTest",
	},
	{
		CardNumberId: "402300",
		FirstName:    "FirstNameTest 2",
		LastName:     "LastNameTestTest2",
	},
}

var expectedPurchaseOrders = []domain.PurchaseOrdersReport{
	{
		Id:                 1,
		CardNumberId:       "402323",
		FirstName:          "FirstNameTest",
		LastName:           "LastNameTest",
		CountBuyersRecords: 1,
	},
}

var expectedPurchaseOrdersList = &[]domain.PurchaseOrdersReport{
	{
		Id:                 3,
		CardNumberId:       "40232212",
		FirstName:          "FirstNameTest",
		LastName:           "LastNameTest",
		CountBuyersRecords: 2,
	},

	{
		Id:                 4,
		CardNumberId:       "40232213",
		FirstName:          "FirstNameTest",
		LastName:           "LastNameTest",
		CountBuyersRecords: 3,
	},
}

var ctx = context.Background()

func TestService_Create(t *testing.T) {

	purchaseOrdersRepo := mockPurchaseOrder.NewPurchaseOrdersRepository(t)
	buyerRepo := buyerRepositoryMock.NewBuyerRepository(t)

	service := service.NewBuyerService(buyerRepo, purchaseOrdersRepo)

	t.Run("crete_ok: when it contains the mandatory fields, should create a buyer", func(t *testing.T) {

		buyerRepo.
			On("Create",
				ctx,
				expectedBuyer.CardNumberId,
				expectedBuyer.FirstName,
				expectedBuyer.LastName,
			).
			Return(expectedBuyer, nil).
			Once()

		result, err := service.Create(ctx, "402323", "FirstNameTest", "LastNameTest")

		assert.Nil(t, err)
		assert.Equal(t, expectedBuyer, result)
	})

	t.Run("create_conflict: when card_number_id already exists, should not create a buyer", func(t *testing.T) {
		errorConflict := fmt.Errorf("Card number id is not unique.")

		buyerRepo.
			On("Create",
				ctx,
				expectedBuyer.CardNumberId,
				expectedBuyer.FirstName,
				expectedBuyer.LastName,
			).
			Return(&domain.Buyer{}, errorConflict).
			Once()

		result, err := service.Create(ctx, "402323", "FirstNameTest", "LastNameTest")

		assert.Equal(t, &domain.Buyer{}, result)
		assert.Equal(t, errorConflict, err)
	})
}

func TestService_GetAll(t *testing.T) {
	purchaseOrdersRepo := mockPurchaseOrder.NewPurchaseOrdersRepository(t)
	buyerRepo := buyerRepositoryMock.NewBuyerRepository(t)

	service := service.NewBuyerService(buyerRepo, purchaseOrdersRepo)

	t.Run("find_all: when exists buyers, should return a list", func(t *testing.T) {

		buyerRepo.
			On("GetAll", ctx).
			Return(expectedBuyerList, nil).
			Once()

		buyerList, _ := service.GetAll(ctx)

		assert.Equal(t, expectedBuyerList, buyerList)
	})

	t.Run("get_all_error: should return any error", func(t *testing.T) {
		buyerRepo.On("GetAll", ctx).
			Return(expectedBuyerList, fmt.Errorf("any error")).
			Once()

		_, err := service.GetAll(ctx)

		assert.NotNil(t, err)

	})
}

func TestService_GetId(t *testing.T) {
	purchaseOrdersRepo := mockPurchaseOrder.NewPurchaseOrdersRepository(t)
	buyerRepo := buyerRepositoryMock.NewBuyerRepository(t)

	service := service.NewBuyerService(buyerRepo, purchaseOrdersRepo)

	t.Run("find_by_id_non_existent: when the element searched for by id does not exist, should return an error", func(t *testing.T) {
		errorNotFound := fmt.Errorf("Buyer not found.")
		buyerRepo.
			On("GetId", ctx, int64(3)).
			Return(nil, errorNotFound).
			Once()

		buyer, err := service.GetId(ctx, int64(3))

		assert.Nil(t, buyer)
		assert.NotNil(t, err)
	})

	t.Run("find_by_id_existent: when element searched for by id exists, should return a buyer", func(t *testing.T) {

		buyerRepo.
			On("GetId", ctx, int64(1)).
			Return(expectedBuyer, nil).
			Once()

		buyer, err := service.GetId(ctx, int64(1))

		assert.Nil(t, err)
		assert.Equal(t, buyer, expectedBuyer)

	})
}

func TestService_Update(t *testing.T) {
	purchaseOrdersRepo := mockPurchaseOrder.NewPurchaseOrdersRepository(t)
	buyerRepo := buyerRepositoryMock.NewBuyerRepository(t)

	service := service.NewBuyerService(buyerRepo, purchaseOrdersRepo)

	buyerUpdated := domain.Buyer{
		CardNumberId: "402300",
		LastName:     "LastNameTest 2",
	}

	t.Run("update_existent: when the data update is successful, should return the updated session", func(t *testing.T) {

		buyerRepo.
			On("Update", ctx, expectedBuyer.Id, buyerUpdated.CardNumberId, buyerUpdated.LastName).
			Return(expectedBuyer, nil).
			Once()

		buyerRepo.
			On("GetId", ctx, expectedBuyer.Id).
			Return(expectedBuyer, nil).
			Once()

		result, err := service.Update(ctx, expectedBuyer.Id, buyerUpdated.CardNumberId, buyerUpdated.LastName)

		assert.Equal(t, expectedBuyer, result)
		assert.Nil(t, err)

	})

	t.Run("update_non_existent: when the element searched for by id does not exist, should return an error.", func(t *testing.T) {

		buyerRepo.On("Update", ctx, expectedBuyer.Id, buyerUpdated.CardNumberId, buyerUpdated.LastName).
			Return(expectedBuyer, nil).
			Once()

		buyerRepo.
			On("GetId", ctx, expectedBuyer.Id).
			Return(nil, fmt.Errorf("Buyer not found.")).
			Once()

		_, err := service.Update(ctx, expectedBuyer.Id, buyerUpdated.CardNumberId, buyerUpdated.LastName)

		assert.Error(t, err)
	})

	t.Run("update_non_existent: when the element searched for by id does not exist, should return an error.", func(t *testing.T) {

		buyerRepo.On("Update", ctx, expectedBuyer.Id, buyerUpdated.CardNumberId, buyerUpdated.LastName).
			Return(nil, fmt.Errorf("Buyer not found.")).
			Once()

		_, err := service.Update(ctx, expectedBuyer.Id, buyerUpdated.CardNumberId, buyerUpdated.LastName)

		assert.Error(t, err)
	})

}

func TestService_Delete(t *testing.T) {
	purchaseOrdersRepo := mockPurchaseOrder.NewPurchaseOrdersRepository(t)
	buyerRepo := buyerRepositoryMock.NewBuyerRepository(t)

	service := service.NewBuyerService(buyerRepo, purchaseOrdersRepo)

	t.Run("delete_non_existent: when the buyer does not exist, should return an error.", func(t *testing.T) {

		buyerRepo.
			On("Delete", ctx, int64(1)).
			Return(fmt.Errorf("buyer not found.")).
			Once()

		err := service.Delete(ctx, int64(1))

		assert.NotNil(t, err)
	})

	t.Run("delete_ok: when the buyer exist, should delete a buyer.", func(t *testing.T) {

		buyerRepo.
			On("Delete", ctx, int64(1)).
			Return(nil).
			Once()

		err := service.Delete(ctx, int64(1))

		assert.Nil(t, err)
	})
}

func TestService_GetPurchaseOrdersReportsById(t *testing.T) {
	purchaseOrdersRepo := mockPurchaseOrder.NewPurchaseOrdersRepository(t)
	buyerRepo := buyerRepositoryMock.NewBuyerRepository(t)

	service := service.NewBuyerService(buyerRepo, purchaseOrdersRepo)

	t.Run("GetById_ok: should returns a report with the number of purchase orders sent to the buyer", func(t *testing.T) {

		buyerRepo.
			On("GetId", context.TODO(), expectedBuyer.Id).
			Return(expectedBuyer, nil).
			Once()

		purchaseOrdersRepo.
			On("ContByBuyerId", ctx, expectedBuyer.Id).
			Return(int64(1), nil).
			Once()

		result, err := service.GetPurchaseOrdersReports(context.TODO(), expectedBuyer.Id)

		assert.Nil(t, err)
		assert.Equal(t, result, &expectedPurchaseOrders)
	})

	t.Run("GetById_err: return an error when the service fails", func(t *testing.T) {
		buyerRepo.
			On("GetId", ctx, expectedBuyer.Id).
			Return(expectedBuyer, nil).
			Once()

		purchaseOrdersRepo.
			On("ContByBuyerId", ctx, expectedBuyer.Id).
			Return(int64(-1), fmt.Errorf("error")).
			Once()

		result, err := service.GetPurchaseOrdersReports(context.TODO(), expectedBuyer.Id)

		assert.NotNil(t, err)
		assert.Empty(t, result)
	})

	t.Run("GetById_err: when the element searched for by id does not exist, should return an error", func(t *testing.T) {

		errorNotFound := fmt.Errorf("the buyer id was not found")

		buyerRepo.
			On("GetId", context.TODO(), expectedBuyer.Id).
			Return(nil, errorNotFound).
			Once()

		result, err := service.GetPurchaseOrdersReports(context.TODO(), int64(1))

		assert.Nil(t, result)
		assert.NotNil(t, err)
	})
}

func TestPurchaseOrderService_GetAllPurchaseOrdersReports(t *testing.T) {
	purchaseOrdersRepo := mockPurchaseOrder.NewPurchaseOrdersRepository(t)
	buyerRepo := buyerRepositoryMock.NewBuyerRepository(t)

	service := service.NewBuyerService(buyerRepo, purchaseOrdersRepo)

	t.Run("getAll_ok: ", func(t *testing.T) {

		buyerRepo.
			On("GetAllPurchaseOrdersReports", ctx).
			Return(expectedPurchaseOrdersList, nil).
			Once()

		result, _ := service.GetAllPurchaseOrdersReports(ctx)

		assert.Equal(t, expectedPurchaseOrdersList, result)

	})

	t.Run("GetAll_err: should return any error", func(t *testing.T) {

		buyerRepo.
			On("GetAllPurchaseOrdersReports", context.TODO()).
			Return(nil, fmt.Errorf("error")).
			Once()

		_, err := service.GetAllPurchaseOrdersReports(context.TODO())

		assert.NotNil(t, err)
	})
}
