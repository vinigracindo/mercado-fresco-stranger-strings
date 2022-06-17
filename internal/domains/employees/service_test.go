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
		repo.On("Create", "123456", "John", "Doe", int64(1)).Return(expectedEmployee, nil)
		service := employees.NewService(repo)

		employee, err := service.Create("123456", "John", "Doe", int64(1))

		assert.Nil(t, err)
		assert.Equal(t, employee, expectedEmployee)
	})

	t.Run("create_conflict: Se o card_number_id já existir, ele não pode ser criado", func(t *testing.T) {
		repo := mocks.NewRepository(t)
		repo.On("Create", "123456", "First Name", "Last Name", int64(1)).Return(employees.Employee{}, fmt.Errorf("Card number id is not unique."))
		service := employees.NewService(repo)

		employee, err := service.Create("123456", "First Name", "Last Name", 1)

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

	t.Run("find_by_id_existent: Se o elemento procurado por id existir, ele retornará as informações do elemento solicitado", func(t *testing.T) {
		expectedEmployee := employees.Employee{
			CardNumberId: "123456",
			FirstName:    "John",
			LastName:     "Doe",
			WarehouseId:  1,
		}

		repo := mocks.NewRepository(t)
		repo.On("GetById", int64(1)).Return(&expectedEmployee, nil)
		service := employees.NewService(repo)

		employee, err := service.GetById(int64(1))

		assert.Nil(t, err)
		assert.Equal(t, employee, &expectedEmployee)
	})
}

func TestEmployeeService_UpdateFullname(t *testing.T) {
	t.Run("update_existent: Quando a atualização dos dados for bem-sucedida, o funcionário será devolvido com as informações atualizadas", func(t *testing.T) {
		expectedEmployee := employees.Employee{
			Id:           1,
			CardNumberId: "123456",
			FirstName:    "John",
			LastName:     "Doe",
			WarehouseId:  1,
		}

		repo := mocks.NewRepository(t)
		repo.On("UpdateFullname", int64(1), "John", "Doe").Return(&expectedEmployee, nil)
		service := employees.NewService(repo)

		employee, err := service.UpdateFullname(int64(1), "John", "Doe")

		assert.Equal(t, employee, &expectedEmployee)
		assert.Nil(t, err)
	})

	t.Run("update_non_existent: Se o funcionário a ser atualizado não existir, será	retornado null.", func(t *testing.T) {
		repo := mocks.NewRepository(t)
		repo.On("UpdateFullname", int64(1), "John", "Doe").Return(nil, fmt.Errorf("Employee not found."))
		service := employees.NewService(repo)

		employee, err := service.UpdateFullname(int64(1), "John", "Doe")
		assert.Nil(t, employee)
		assert.NotNil(t, err)
	})

}

func TestEmployeeService_Delete(t *testing.T) {
	t.Run("delete_non_existent: Quando o funcionário não existir, será retornado null.", func(t *testing.T) {
		repo := mocks.NewRepository(t)
		repo.On("Delete", int64(1)).Return(fmt.Errorf("Employee not found."))
		service := employees.NewService(repo)

		err := service.Delete(int64(1))

		assert.NotNil(t, err)
	})

	t.Run("delete_ok: Se a exclusão for bem-sucedida, o item não aparecerá na lista.", func(t *testing.T) {
		repo := mocks.NewRepository(t)
		repo.On("Delete", int64(1)).Return(nil)
		service := employees.NewService(repo)

		err := service.Delete(int64(1))

		assert.Nil(t, err)
	})
}
