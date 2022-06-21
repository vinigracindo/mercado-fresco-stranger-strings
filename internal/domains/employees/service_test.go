package employees_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/employees"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/employees/mocks"
)

func makeEmployee() employees.Employee {
	return employees.Employee{
		CardNumberId: "123456",
		FirstName:    "John",
		LastName:     "Doe",
		WarehouseId:  1,
	}
}

func TestEmployeeService_Create(t *testing.T) {
	expectedEmployee := makeEmployee()

	repo := mocks.NewRepository(t)
	service := employees.NewService(repo)

	t.Run("create_ok: when it contains the mandatory fields, should create a employee", func(t *testing.T) {
		repo.
			On("Create", "123456", "John", "Doe", int64(1)).
			Return(expectedEmployee, nil).
			Once()

		employee, err := service.Create("123456", "John", "Doe", int64(1))

		assert.Nil(t, err)
		assert.Equal(t, employee, expectedEmployee)
	})

	t.Run("create_conflict: when card_number already exists, should not create a employee", func(t *testing.T) {
		repo.
			On("Create", "123456", "First Name", "Last Name", int64(1)).
			Return(employees.Employee{}, employees.ErrCardNumberMustBeUnique).
			Once()

		employee, err := service.Create("123456", "First Name", "Last Name", 1)

		assert.NotNil(t, err)
		assert.Empty(t, employee)
	})
}

func TestEmployeeService_GetAll(t *testing.T) {
	repo := mocks.NewRepository(t)
	service := employees.NewService(repo)

	t.Run("find_all: when exists employees, should return a list", func(t *testing.T) {
		expectedEmployees := []employees.Employee{
			makeEmployee(),
			makeEmployee(),
		}

		repo.On("GetAll").Return(expectedEmployees, nil).Once()

		employees, err := service.GetAll()

		assert.Equal(t, employees, expectedEmployees)
		assert.Nil(t, err)
	})

	t.Run("find_all_err: should return error.", func(t *testing.T) {
		repo.
			On("GetAll").
			Return(nil, employees.ErrEmployeeNotFound).
			Once()

		employees, err := service.GetAll()

		assert.Nil(t, employees)
		assert.NotNil(t, err)
	})
}

func TestEmployeeService_GetById(t *testing.T) {
	repo := mocks.NewRepository(t)
	service := employees.NewService(repo)

	t.Run("find_by_id_non_existent: when element searched for by id exists, should return a employee", func(t *testing.T) {
		repo.
			On("GetById", int64(1)).
			Return(nil, employees.ErrEmployeeNotFound).
			Once()

		employee, err := service.GetById(int64(1))

		assert.Nil(t, employee)
		assert.NotNil(t, err)
	})

	t.Run("find_by_id_existent: when the element searched for by id does not exists, should return an error", func(t *testing.T) {
		expectedEmployee := makeEmployee()

		repo.
			On("GetById", int64(1)).
			Return(&expectedEmployee, nil).
			Once()

		employee, err := service.GetById(int64(1))

		assert.Nil(t, err)
		assert.Equal(t, employee, &expectedEmployee)
	})
}

func TestEmployeeService_UpdateFullname(t *testing.T) {
	repo := mocks.NewRepository(t)
	service := employees.NewService(repo)

	t.Run("update_existent: when the data update is successful, should return the updated employee", func(t *testing.T) {
		updatedEmployee := makeEmployee()
		updatedEmployee.FirstName = "Jane"
		updatedEmployee.LastName = "Doe"

		repo.
			On("UpdateFullname", int64(1), "Jane", "Doe").
			Return(&updatedEmployee, nil).
			Once()

		employee, err := service.UpdateFullname(int64(1), "Jane", "Doe")

		assert.Equal(t, employee, &updatedEmployee)
		assert.Nil(t, err)
	})

	t.Run("update_non_existent: when the element searched for by id does not exist, should return an error", func(t *testing.T) {
		repo.
			On("UpdateFullname", int64(1), "John", "Doe").
			Return(nil, employees.ErrEmployeeNotFound).
			Once()

		employee, err := service.UpdateFullname(int64(1), "John", "Doe")
		assert.Nil(t, employee)
		assert.NotNil(t, err)
	})

}

func TestEmployeeService_Delete(t *testing.T) {
	repo := mocks.NewRepository(t)
	service := employees.NewService(repo)

	t.Run("delete_non_existent: when the section does not exist, should return an error", func(t *testing.T) {
		repo.
			On("Delete", int64(1)).
			Return(employees.ErrEmployeeNotFound).
			Once()

		err := service.Delete(int64(1))

		assert.NotNil(t, err)
	})

	t.Run("delete_ok: when the section exists, should delete a employee", func(t *testing.T) {
		repo.
			On("Delete", int64(1)).
			Return(nil).
			Once()

		err := service.Delete(int64(1))

		assert.Nil(t, err)
	})
}
