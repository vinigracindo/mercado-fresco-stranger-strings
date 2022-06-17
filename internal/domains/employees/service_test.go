package employees_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/employees"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/employees/mocks"
)

func TestEmployeeService_Store(t *testing.T) {
	expectedEmployee := employees.Employee{
		CardNumberId: "123456",
		FirstName:    "John",
		LastName:     "Doe",
		WarehouseId:  1,
	}

	t.Run("create_ok: Se contiver os campos necessários, será criado", func(t *testing.T) {
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
