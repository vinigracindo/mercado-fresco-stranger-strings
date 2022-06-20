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

func TestEmployeeService_Store(t *testing.T) {
	expectedEmployee := makeEmployee()

	repo := mocks.NewRepository(t)
	service := employees.NewService(repo)

	t.Run("create_ok: Se contiver os campos necessários, será criado", func(t *testing.T) {
		repo.On("Create", "123456", "John", "Doe", int64(1)).Return(expectedEmployee, nil)
		employee, err := service.Create("123456", "John", "Doe", int64(1))

		assert.Nil(t, err)
		assert.Equal(t, employee, expectedEmployee)
	})

	t.Run("create_conflict: Se o card_number_id já existir, ele não pode ser criado", func(t *testing.T) {
		repo.On("Create", "123456", "First Name", "Last Name", int64(1)).Return(employees.Employee{}, employees.ErrCardNumberMustBeUnique)
		employee, err := service.Create("123456", "First Name", "Last Name", 1)

		assert.NotNil(t, err)
		assert.Empty(t, employee)
	})
}

func TestEmployeeService_GetAll(t *testing.T) {
	repo := mocks.NewRepository(t)
	service := employees.NewService(repo)

	t.Run("find_all: Se a lista tiver \"n\" elementos, retornará uma quantidade do total de elementos", func(t *testing.T) {
		expectedEmployees := []employees.Employee{
			makeEmployee(),
			makeEmployee(),
		}

		repo.On("GetAll").Return(expectedEmployees, nil).Once()

		employees, err := service.GetAll()

		assert.Equal(t, employees, expectedEmployees)
		assert.Nil(t, err)
	})
}

func TestEmployeeService_GetById(t *testing.T) {
	repo := mocks.NewRepository(t)
	service := employees.NewService(repo)

	t.Run("find_by_id_non_existent: Se o elemento procurado por id não existir, retorna null", func(t *testing.T) {
		repo.On("GetById", int64(1)).Return(nil, employees.ErrEmployeeNotFound).Once()

		employee, err := service.GetById(int64(1))

		assert.Nil(t, employee)
		assert.NotNil(t, err)
	})

	t.Run("find_by_id_existent: Se o elemento procurado por id existir, ele retornará as informações do elemento solicitado", func(t *testing.T) {
		expectedEmployee := makeEmployee()

		repo.On("GetById", int64(1)).Return(&expectedEmployee, nil).Once()

		employee, err := service.GetById(int64(1))

		assert.Nil(t, err)
		assert.Equal(t, employee, &expectedEmployee)
	})
}

func TestEmployeeService_UpdateFullname(t *testing.T) {
	repo := mocks.NewRepository(t)
	service := employees.NewService(repo)

	t.Run("update_existent: Quando a atualização dos dados for bem-sucedida, o funcionário será devolvido com as informações atualizadas", func(t *testing.T) {
		updatedEmployee := makeEmployee()
		updatedEmployee.FirstName = "Jane"
		updatedEmployee.LastName = "Doe"

		repo.On("UpdateFullname", int64(1), "Jane", "Doe").Return(&updatedEmployee, nil).Once()

		employee, err := service.UpdateFullname(int64(1), "Jane", "Doe")

		assert.Equal(t, employee, &updatedEmployee)
		assert.Nil(t, err)
	})

	t.Run("update_non_existent: Se o funcionário a ser atualizado não existir, será	retornado null.", func(t *testing.T) {
		repo.On("UpdateFullname", int64(1), "John", "Doe").Return(nil, employees.ErrEmployeeNotFound).Once()

		employee, err := service.UpdateFullname(int64(1), "John", "Doe")
		assert.Nil(t, employee)
		assert.NotNil(t, err)
	})

}

func TestEmployeeService_Delete(t *testing.T) {
	repo := mocks.NewRepository(t)
	service := employees.NewService(repo)

	t.Run("delete_non_existent: Quando o funcionário não existir, será retornado null.", func(t *testing.T) {
		repo.On("Delete", int64(1)).Return(employees.ErrEmployeeNotFound).Once()

		err := service.Delete(int64(1))

		assert.NotNil(t, err)
	})

	t.Run("delete_ok: Se a exclusão for bem-sucedida, o item não aparecerá na lista.", func(t *testing.T) {
		repo.On("Delete", int64(1)).Return(nil).Once()

		err := service.Delete(int64(1))

		assert.Nil(t, err)
	})
}
