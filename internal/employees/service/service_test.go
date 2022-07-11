package service_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/employees/domain"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/employees/domain/mocks"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/employees/service"
)

func makeEmployee() domain.Employee {
	return domain.Employee{
		CardNumberId: "123456",
		FirstName:    "John",
		LastName:     "Doe",
		WarehouseId:  1,
	}
}

func makeEmployeeInboundOrdersReport() domain.EmployeeInboundOrdersReport {
	return domain.EmployeeInboundOrdersReport{
		Employee: makeEmployee(),
		Count:    10,
	}
}

func TestEmployeeService_Create(t *testing.T) {
	expectedEmployee := makeEmployee()

	repo := mocks.NewEmployeeRepository(t)
	service := service.NewEmployeeService(repo) // employees.NewService

	t.Run("create_ok: when it contains the mandatory fields, should create a employee", func(t *testing.T) {
		repo.
			On("Create", mock.Anything, "123456", "John", "Doe", int64(1)).
			Return(expectedEmployee, nil).
			Once()

		employee, err := service.Create(context.TODO(), "123456", "John", "Doe", int64(1))

		assert.Nil(t, err)
		assert.Equal(t, employee, expectedEmployee)
	})

	t.Run("create_conflict: when card_number already exists, should not create a employee", func(t *testing.T) {
		repo.
			On("Create", mock.Anything, "123456", "First Name", "Last Name", int64(1)).
			Return(domain.Employee{}, domain.ErrCardNumberMustBeUnique).
			Once()

		employee, err := service.Create(context.TODO(), "123456", "First Name", "Last Name", 1)

		assert.NotNil(t, err)
		assert.Empty(t, employee)
	})
}

func TestEmployeeService_GetAll(t *testing.T) {
	repo := mocks.NewEmployeeRepository(t)
	service := service.NewEmployeeService(repo)

	t.Run("find_all: when exists employees, should return a list", func(t *testing.T) {
		expectedEmployees := []domain.Employee{
			makeEmployee(),
			makeEmployee(),
		}

		repo.On("GetAll", mock.Anything).Return(expectedEmployees, nil).Once()

		employees, err := service.GetAll(context.TODO())

		assert.Equal(t, employees, expectedEmployees)
		assert.Nil(t, err)
	})

	t.Run("find_all_err: should return error.", func(t *testing.T) {
		repo.
			On("GetAll", mock.Anything).
			Return(nil, domain.ErrEmployeeNotFound).
			Once()

		employees, err := service.GetAll(context.TODO())

		assert.Nil(t, employees)
		assert.NotNil(t, err)
	})
}

func TestEmployeeService_GetById(t *testing.T) {
	repo := mocks.NewEmployeeRepository(t)
	service := service.NewEmployeeService(repo)

	t.Run("find_by_id_non_existent: when element searched for by id exists, should return a employee", func(t *testing.T) {
		repo.
			On("GetById", mock.Anything, int64(1)).
			Return(nil, domain.ErrEmployeeNotFound).
			Once()

		employee, err := service.GetById(context.TODO(), int64(1))

		assert.Nil(t, employee)
		assert.NotNil(t, err)
	})

	t.Run("find_by_id_existent: when the element searched for by id does not exists, should return an error", func(t *testing.T) {
		expectedEmployee := makeEmployee()

		repo.
			On("GetById", mock.Anything, int64(1)).
			Return(&expectedEmployee, nil).
			Once()

		employee, err := service.GetById(context.TODO(), int64(1))

		assert.Nil(t, err)
		assert.Equal(t, employee, &expectedEmployee)
	})
}

func TestEmployeeService_UpdateFullname(t *testing.T) {
	repo := mocks.NewEmployeeRepository(t)
	service := service.NewEmployeeService(repo)

	employee := makeEmployee()

	t.Run("update_existent: when the data update is successful, should return the updated employee", func(t *testing.T) {
		idWillBeUpdated := int64(1)

		repo.
			On("GetById", context.TODO(), int64(1)).
			Return(&employee, nil).Once()

		updatedEmployee := employee
		updatedEmployee.SetFullname("Jane", "Doe")

		repo.
			On("Update", context.TODO(), idWillBeUpdated, updatedEmployee).
			Return(nil).Once()

		emp, err := service.UpdateFullname(context.TODO(), idWillBeUpdated, "Jane", "Doe")

		assert.Equal(t, emp, &updatedEmployee)
		assert.Nil(t, err)
	})

	t.Run("update_non_existent: when the element searched for by id does not exist, should return an error", func(t *testing.T) {
		repo.
			On("GetById", context.TODO(), int64(32)).
			Return(nil, fmt.Errorf("Employee not found.")).
			Once()

		res, err := service.UpdateFullname(context.TODO(), 32, "Jane", "Doe")

		assert.Nil(t, res)
		assert.Error(t, err)
	})

	t.Run("update_invalid_data: when the data update is not successful, should return an error", func(t *testing.T) {
		idWillBeUpdated := int64(1)

		repo.
			On("GetById", context.TODO(), int64(1)).
			Return(&employee, nil).Once()

		updatedEmployee := employee
		updatedEmployee.SetFullname("Jane", "Doe")

		repo.
			On("Update", context.TODO(), idWillBeUpdated, updatedEmployee).
			Return(fmt.Errorf("error")).Once()

		emp, err := service.UpdateFullname(context.TODO(), idWillBeUpdated, "Jane", "Doe")

		assert.Nil(t, emp)
		assert.Error(t, err)

	})

}
func TestEmployeeService_Delete(t *testing.T) {
	repo := mocks.NewEmployeeRepository(t)
	service := service.NewEmployeeService(repo)

	t.Run("delete_non_existent: when the section does not exist, should return an error", func(t *testing.T) {
		repo.
			On("Delete", mock.Anything, int64(1)).
			Return(domain.ErrEmployeeNotFound).
			Once()

		err := service.Delete(context.TODO(), int64(1))

		assert.NotNil(t, err)
	})

	t.Run("delete_ok: when the section exists, should delete a employee", func(t *testing.T) {
		repo.
			On("Delete", mock.Anything, int64(1)).
			Return(nil).
			Once()

		err := service.Delete(context.TODO(), int64(1))

		assert.Nil(t, err)
	})
}

func TestEmployeeService_GetAllReportInboundOrders(t *testing.T) {
	repo := mocks.NewEmployeeRepository(t)
	service := service.NewEmployeeService(repo)

	t.Run("report_ok: Return a list of employees with the total amount of inbound orders", func(t *testing.T) {
		expectedInboundOrders := []domain.EmployeeInboundOrdersReport{
			makeEmployeeInboundOrdersReport(),
			makeEmployeeInboundOrdersReport(),
			makeEmployeeInboundOrdersReport(),
		}

		repo.
			On("GetAllReportInboundOrders", context.TODO()).
			Return(expectedInboundOrders, nil).
			Once()

		result, err := service.GetAllReportInboundOrders(context.TODO())
		assert.Nil(t, err)
		assert.Equal(t, result, expectedInboundOrders)
	})

	t.Run("report_error: Return an error when the service fails", func(t *testing.T) {
		repo.
			On("GetAllReportInboundOrders", context.TODO()).
			Return(nil, fmt.Errorf("error")).
			Once()

		result, err := service.GetAllReportInboundOrders(context.TODO())
		assert.NotNil(t, err)
		assert.Nil(t, result)
	})
}

func TestEmployeeService_GetReportInboundOrdersById(t *testing.T) {
	repo := mocks.NewEmployeeRepository(t)
	service := service.NewEmployeeService(repo)

	t.Run("report_ok: Return a list of employees with the total amount of inbound orders", func(t *testing.T) {
		expectedInboundOrders := makeEmployeeInboundOrdersReport()
		employee := makeEmployee()

		repo.
			On("GetById", mock.Anything, int64(1)).
			Return(&employee, nil).
			Once()

		repo.
			On("GetReportInboundOrdersById", context.TODO(), int64(1)).
			Return(expectedInboundOrders, nil).
			Once()

		result, err := service.GetReportInboundOrdersById(context.TODO(), int64(1))
		assert.Nil(t, err)
		assert.Equal(t, result, expectedInboundOrders)
	})

	t.Run("report_error: Return an error when the service fails", func(t *testing.T) {
		employee := makeEmployee()

		repo.
			On("GetById", mock.Anything, int64(1)).
			Return(&employee, nil).
			Once()

		repo.
			On("GetReportInboundOrdersById", context.TODO(), int64(1)).
			Return(domain.EmployeeInboundOrdersReport{}, fmt.Errorf("error")).
			Once()

		result, err := service.GetReportInboundOrdersById(context.TODO(), int64(1))
		assert.NotNil(t, err)
		assert.Empty(t, result)
	})

	t.Run("employee_does_not_exist: return an error when the employee does not exist", func(t *testing.T) {
		repo.
			On("GetById", mock.Anything, int64(1)).
			Return(nil, fmt.Errorf("error")).
			Once()

		result, err := service.GetReportInboundOrdersById(context.TODO(), int64(1))
		assert.NotNil(t, err)
		assert.Empty(t, result)
	})

}
