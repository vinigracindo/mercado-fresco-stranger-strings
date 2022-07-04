package service_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/inboud_orders/domain"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/inboud_orders/domain/mocks"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/inboud_orders/service"
)

var inboundTimeNow = time.Now()

func makeInboundOrder() domain.InboundOrders {
	return domain.InboundOrders{
		OrderDate:      inboundTimeNow,
		OrderNumber:    "order#1",
		EmployeeId:     1,
		ProductBatchId: 1,
		WarehouseId:    1,
	}
}

func TestInboundOrdersService_Create(t *testing.T) {
	t.Run("create_ok: when it contains the mandatory fields, should create a inbound order", func(t *testing.T) {
		expectedInboundOrders := makeInboundOrder()

		repo := mocks.NewInboundOrdersRepository(t)
		service := service.NewInboundOrderService(repo)

		repo.
			On("Create", context.TODO(), inboundTimeNow, "order#1", int64(1), int64(1), int64(1)).
			Return(expectedInboundOrders, nil).
			Once()

		inboundOrders, err := service.Create(context.TODO(), inboundTimeNow, "order#1", int64(1), int64(1), int64(1))

		assert.Nil(t, err)
		assert.Equal(t, expectedInboundOrders, inboundOrders)
	})
}
