package service_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	employeeDomain "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/employees/domain"
	employeesMock "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/employees/domain/mocks"
	InboundOrdersDomain "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/inboud_orders/domain"
	inboundOrdersMock "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/inboud_orders/domain/mocks"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/inboud_orders/service"
)

var inboundTimeNow = time.Now()

func makeInboundOrder() InboundOrdersDomain.InboundOrders {
	return InboundOrdersDomain.InboundOrders{
		OrderDate:      inboundTimeNow,
		OrderNumber:    "order#1",
		EmployeeId:     1,
		ProductBatchId: 1,
		WarehouseId:    1,
	}
}

func TestInboundOrdersService_Create(t *testing.T) {
	mockRepoInbound := inboundOrdersMock.NewInboundOrdersRepository(t)
	mockRepoEmployee := employeesMock.NewEmployeeRepository(t)
	service := service.NewInboundOrderService(mockRepoInbound, mockRepoEmployee)

	t.Run("create_ok: when it contains the mandatory fields, should create a inbound order", func(t *testing.T) {
		expectedInboundOrders := makeInboundOrder()

		mockRepoEmployee.
			On("GetById", context.TODO(), int64(1)).
			Return(&employeeDomain.Employee{}, nil).
			Once()

		mockRepoInbound.
			On("Create", context.TODO(), inboundTimeNow, "order#1", int64(1), int64(1), int64(1)).
			Return(expectedInboundOrders, nil).
			Once()

		inboundOrders, err := service.Create(context.TODO(), inboundTimeNow, "order#1", int64(1), int64(1), int64(1))

		assert.Nil(t, err)
		assert.Equal(t, expectedInboundOrders, inboundOrders)
	})

	t.Run("create_employee_does_not_exist: when employee does not exist, should not create a inbound order", func(t *testing.T) {
		mockRepoEmployee.
			On("GetById", context.TODO(), int64(1)).
			Return(nil, employeeDomain.ErrEmployeeNotFound).
			Once()

		inboundOrders, err := service.Create(context.TODO(), inboundTimeNow, "order#1", int64(1), int64(1), int64(1))

		assert.Equal(t, employeeDomain.ErrEmployeeNotFound, err)
		assert.Equal(t, InboundOrdersDomain.InboundOrders{}, inboundOrders)
	})

	t.Run("create_error: when create inbound order fails, should return error", func(t *testing.T) {
		mockRepoEmployee.
			On("GetById", context.TODO(), int64(1)).
			Return(&employeeDomain.Employee{}, nil).
			Once()

		mockRepoInbound.
			On("Create", context.TODO(), inboundTimeNow, "order#1", int64(1), int64(1), int64(1)).
			Return(InboundOrdersDomain.InboundOrders{}, InboundOrdersDomain.ErrCreateInboundOrder).
			Once()

		inboundOrders, err := service.Create(context.TODO(), inboundTimeNow, "order#1", int64(1), int64(1), int64(1))

		assert.Equal(t, InboundOrdersDomain.ErrCreateInboundOrder, err)
		assert.Equal(t, InboundOrdersDomain.InboundOrders{}, inboundOrders)
	})
}
