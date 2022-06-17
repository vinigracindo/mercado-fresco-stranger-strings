package employees_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/employees"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/employees/mocks"
)

func TestEmployeeService_Store(t *testing.T) {
	t.Run("create_ok: Se contiver os campos necessários, será criado", func(t *testing.T) {
		expectedEmployee := employees.Employee{
			CardNumberId: "123456",
			FirstName:    "John",
			LastName:     "Doe",
			WarehouseId:  1,
		}

		repo := mocks.NewRepository(t)
		repo.On("Store", "123456", "John", "Doe", int64(1)).Return(expectedEmployee, nil)
		service := employees.NewService(repo)

		employee, err := service.Store("123456", "John", "Doe", int64(1))

		assert.Nil(t, err)
		assert.Equal(t, employee, expectedEmployee)
	})

	t.Run("create_conflict: Se o card_number_id já existir, ele não pode ser criado", func(t *testing.T) {
		repo := mocks.NewRepository(t)
		repo.On("Store", "123456", "First Name", "Last Name", int64(1)).Return(employees.Employee{}, fmt.Errorf("Card number id is not unique."))
		service := employees.NewService(repo)

		employee, err := service.Store("123456", "First Name", "Last Name", 1)

		assert.NotNil(t, err)
		assert.Empty(t, employee)
		assert.Equal(t, err.Error(), "Card number id is not unique.")
	})
}

func TestEmployeeService_GetAll(t *testing.T) {
	t.Run("find_all: Se a lista tiver \"n\" elementos, retornará uma quantidade do total de elementos", func(t *testing.T) {
		expectedEmployees := []employees.Employee{
			{
				CardNumberId: "123456",
				FirstName:    "John",
				LastName:     "Doe",
				WarehouseId:  1,
			},
			{
				CardNumberId: "789012",
				FirstName:    "Jane",
				LastName:     "Doe",
				WarehouseId:  2,
			},
		}

		repo := mocks.NewRepository(t)
		repo.On("GetAll").Return(expectedEmployees, nil)
		service := employees.NewService(repo)

		employees, err := service.GetAll()

		assert.Equal(t, employees, expectedEmployees)
		assert.Nil(t, err)
	})
}

func TestEmployeeService_GetById(t *testing.T) {
	t.Run("find_by_id_non_existent: Se o elemento procurado por id não existir, retorna null", func(t *testing.T) {
		repo := mocks.NewRepository(t)
		repo.On("GetById", int64(1)).Return(nil, fmt.Errorf("Employee not found."))
		service := employees.NewService(repo)

		employee, err := service.GetById(int64(1))

		assert.Nil(t, employee)
		assert.NotNil(t, err)
	})
}
