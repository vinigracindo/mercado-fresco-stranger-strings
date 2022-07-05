package repository_test

import (
	"context"
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/employees/domain"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/employees/repository"
)

func TestEmployeeRepository_GetAll(t *testing.T) {
	ctx := context.Background()

	expectedEmployees := []domain.Employee{
		{
			Id:           1,
			CardNumberId: "123456",
			FirstName:    "John",
			LastName:     "Doe",
			WarehouseId:  1,
		},
		{
			Id:           2,
			CardNumberId: "789012",
			FirstName:    "Jane",
			LastName:     "Doe",
			WarehouseId:  2,
		},
	}

	t.Run("should return all employees", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		rows := sqlmock.NewRows([]string{"id", "card_number_id", "first_name", "last_name", "warehouse_id"})
		for _, employee := range expectedEmployees {
			rows = rows.AddRow(employee.Id, employee.CardNumberId, employee.FirstName, employee.LastName, employee.WarehouseId)
		}

		employeeRepository := repository.NewMariaDBEmployeeRepository(db)

		mock.ExpectQuery(repository.SQLFindAllEmployees).WillReturnRows(rows)

		result, err := employeeRepository.GetAll(ctx)

		assert.Nil(t, err)
		assert.Equal(t, expectedEmployees, result)
	})

	t.Run("should return error when query fails", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		employeeRepository := repository.NewMariaDBEmployeeRepository(db)

		mock.ExpectQuery(repository.SQLFindAllEmployees).WillReturnError(fmt.Errorf("query error"))

		result, err := employeeRepository.GetAll(ctx)
		assert.Error(t, err)
		assert.Empty(t, result)
	})

	t.Run("should return error when scan fails", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		employeeRepository := repository.NewMariaDBEmployeeRepository(db)

		rows := sqlmock.NewRows(
			[]string{"id", "card_number_id", "first_name", "last_name", "warehouse_id"},
		).AddRow("", "", "", "", "")

		mock.ExpectQuery(repository.SQLFindAllEmployees).WillReturnRows(rows)

		result, err := employeeRepository.GetAll(ctx)

		assert.Error(t, err)
		assert.Empty(t, result)
	})
}

func TestEmployeeRepository_GetById(t *testing.T) {
	expectedEmployee := domain.Employee{
		Id:           1,
		CardNumberId: "123456",
		FirstName:    "John",
		LastName:     "Doe",
		WarehouseId:  1,
	}

	t.Run("should return employee by id", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		rows := sqlmock.NewRows([]string{"id", "card_number_id", "first_name", "last_name", "warehouse_id"}).AddRow(expectedEmployee.Id, expectedEmployee.CardNumberId, expectedEmployee.FirstName, expectedEmployee.LastName, expectedEmployee.WarehouseId)

		employeeRepository := repository.NewMariaDBEmployeeRepository(db)

		mock.
			ExpectQuery(repository.SQLFindEmployeeByID).
			WithArgs(expectedEmployee.Id).
			WillReturnRows(rows)

		result, err := employeeRepository.GetById(context.TODO(), expectedEmployee.Id)

		assert.Nil(t, err)
		assert.Equal(t, expectedEmployee, *result)
	})

	t.Run("should return error when query fails", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		employeeRepository := repository.NewMariaDBEmployeeRepository(db)

		mock.
			ExpectQuery(repository.SQLFindEmployeeByID).
			WithArgs(expectedEmployee.Id).
			WillReturnError(fmt.Errorf("query error"))

		result, err := employeeRepository.GetById(context.TODO(), expectedEmployee.Id)

		assert.Error(t, err)
		assert.Empty(t, result)
	})
}

func TestEmployeeRepository_Create(t *testing.T) {
	expectedEmployee := domain.Employee{
		Id:           1,
		CardNumberId: "123456",
		FirstName:    "John",
		LastName:     "Doe",
		WarehouseId:  1,
	}

	t.Run("should create employee", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		employeeRepository := repository.NewMariaDBEmployeeRepository(db)

		mock.
			ExpectExec(regexp.QuoteMeta(repository.SQLCreateEmployee)).
			WithArgs("123456", "John", "Doe", int64(1)).
			WillReturnResult(sqlmock.NewResult(1, 1))

		result, err := employeeRepository.Create(
			context.TODO(),
			"123456", "John", "Doe", int64(1),
		)

		assert.Nil(t, err)
		assert.Equal(t, expectedEmployee, result)
	})

	t.Run("should return error when query execution fails", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		employeeRepository := repository.NewMariaDBEmployeeRepository(db)

		mock.
			ExpectExec(regexp.QuoteMeta(repository.SQLCreateEmployee)).
			WillReturnError(fmt.Errorf("query error"))

		result, err := employeeRepository.Create(
			context.TODO(),
			"123456", "John", "Doe", int64(1),
		)

		assert.Error(t, err)
		assert.Empty(t, result)
	})
}

func TestEmployeeRepository_Update(t *testing.T) {
	expectedEmployee := domain.Employee{
		Id:           int64(1),
		CardNumberId: "123456",
		FirstName:    "John",
		LastName:     "Doe",
		WarehouseId:  int64(1),
	}

	t.Run("should update employee full name", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		employeeRepository := repository.NewMariaDBEmployeeRepository(db)

		mock.
			ExpectExec(regexp.QuoteMeta(repository.SQLUpdateEmployeeFullname)).
			WithArgs("John", "Doe", 1).
			WillReturnResult(sqlmock.NewResult(0, 1))

		err = employeeRepository.Update(
			context.TODO(),
			int64(1),
			expectedEmployee,
		)

		assert.NoError(t, err)
	})

	t.Run("should return error when query execution fails", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		employeeRepository := repository.NewMariaDBEmployeeRepository(db)

		mock.
			ExpectExec(regexp.QuoteMeta(repository.SQLUpdateEmployeeFullname)).
			WillReturnError(fmt.Errorf("query error"))

		err = employeeRepository.Update(
			context.TODO(),
			int64(1),
			expectedEmployee,
		)

		assert.Error(t, err)
	})
}

func TestEmployeeRepository_Delete(t *testing.T) {
	t.Run("should delete employee", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		employeeRepository := repository.NewMariaDBEmployeeRepository(db)

		mock.
			ExpectExec(regexp.QuoteMeta(repository.SQLDeleteEmployee)).
			WithArgs(1).
			WillReturnResult(sqlmock.NewResult(0, 1))

		err = employeeRepository.Delete(context.TODO(), 1)

		assert.NoError(t, err)
	})

	t.Run("should return error when query execution fails", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		employeeRepository := repository.NewMariaDBEmployeeRepository(db)

		mock.
			ExpectExec(regexp.QuoteMeta(repository.SQLDeleteEmployee)).
			WillReturnError(fmt.Errorf("query error"))

		err = employeeRepository.Delete(context.TODO(), 1)

		assert.Error(t, err)
	})

	t.Run("delete_not_found: ", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(regexp.QuoteMeta(repository.SQLDeleteEmployee)).
			WithArgs(int64(1)).
			WillReturnResult(sqlmock.NewResult(0, 0))

		sectionRepository := repository.NewMariaDBEmployeeRepository(db)

		err = sectionRepository.Delete(context.Background(), int64(1))

		assert.Error(t, err)
	})
}