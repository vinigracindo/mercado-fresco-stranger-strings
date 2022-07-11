package repository_test

import (
	"context"
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/employees/domain"
	repository "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/employees/repository/mariadb"
)

func makeEmployee(id int64) domain.Employee {
	return domain.Employee{
		Id:           id,
		CardNumberId: "123456",
		FirstName:    "John",
		LastName:     "Doe",
		WarehouseId:  1,
	}
}

func makeEmployeeInboundOrdersReport(id int64) domain.EmployeeInboundOrdersReport {
	return domain.EmployeeInboundOrdersReport{
		Employee: makeEmployee(id),
		Count:    10,
	}
}

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

	t.Run("get_all_ok: should return all employees", func(t *testing.T) {
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

	t.Run("get_all_query_fails: should return error when query fails", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		employeeRepository := repository.NewMariaDBEmployeeRepository(db)

		mock.ExpectQuery(repository.SQLFindAllEmployees).WillReturnError(fmt.Errorf("query error"))

		result, err := employeeRepository.GetAll(ctx)
		assert.Error(t, err)
		assert.Empty(t, result)
	})

	t.Run("get_all_scan_fails: should return error when scan fails", func(t *testing.T) {
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

	t.Run("get_by_id_ok: should return employee by id", func(t *testing.T) {
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

	t.Run("get_by_id_query_fails: should return error when query fails", func(t *testing.T) {
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

	t.Run("create_ok: should create employee", func(t *testing.T) {
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

	t.Run("create_query_fails: should return error when query execution fails", func(t *testing.T) {
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

	t.Run("update_ok: should update employee full name", func(t *testing.T) {
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

	t.Run("update_query_fails: should return error when query execution fails", func(t *testing.T) {
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
	t.Run("delete_ok: should delete employee", func(t *testing.T) {
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

	t.Run("delete_query_fails: should return error when query execution fails", func(t *testing.T) {
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

	t.Run("delete_employee_not_found: should return a error when employee not found. ", func(t *testing.T) {
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

func TestEmployeeRepository_GetAllReportInboundOrders(t *testing.T) {
	employeeID := int64(1)

	t.Run("get_all_report_inbound_errors: should return all employees with inbound orders count", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		expectedInboundOrders := []domain.EmployeeInboundOrdersReport{
			makeEmployeeInboundOrdersReport(employeeID),
			makeEmployeeInboundOrdersReport(employeeID + 1),
			makeEmployeeInboundOrdersReport(employeeID + 2),
		}

		employeeRepository := repository.NewMariaDBEmployeeRepository(db)

		mock.
			ExpectQuery(regexp.QuoteMeta(repository.SQLReportInboundOrders)).
			WillReturnRows(
				sqlmock.
					NewRows([]string{"id", "card_number_id", "first_name", "last_name", "warehouse_id", "inbound_orders_count"}).
					AddRow(1, "123456", "John", "Doe", int64(1), 10).
					AddRow(2, "123456", "John", "Doe", int64(1), 10).
					AddRow(3, "123456", "John", "Doe", int64(1), 10),
			)

		result, err := employeeRepository.GetAllReportInboundOrders(context.TODO())

		assert.NoError(t, err)
		assert.Equal(t, expectedInboundOrders, result)
	})

	t.Run("get_all_report_inbound_errors_query_fails: should return error when query execution fails", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		employeeRepository := repository.NewMariaDBEmployeeRepository(db)

		mock.
			ExpectQuery(regexp.QuoteMeta(repository.SQLReportInboundOrders)).
			WillReturnError(fmt.Errorf("query error"))

		result, err := employeeRepository.GetAllReportInboundOrders(context.TODO())

		assert.Error(t, err)
		assert.Empty(t, result)
	})

	t.Run("get_all_report_inbound_errors_scan_fails: should return error when scan fails", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		employeeRepository := repository.NewMariaDBEmployeeRepository(db)

		mock.
			ExpectQuery(regexp.QuoteMeta(repository.SQLReportInboundOrders)).
			WillReturnRows(
				sqlmock.
					NewRows([]string{"id", "card_number_id", "first_name", "last_name", "warehouse_id", "inbound_orders_count"}).
					AddRow("invalid_id", "123456", "John", "Doe", int64(1), 10),
			)

		result, err := employeeRepository.GetAllReportInboundOrders(context.TODO())

		assert.Error(t, err)
		assert.Empty(t, result)
	})
}

func TestEmployeeRepository_GetReportInboundOrdersById(t *testing.T) {
	employeeID := int64(1)

	t.Run("get_by_id_report_inbound: should return employee with inbound orders count", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		expectedInboundOrders := makeEmployeeInboundOrdersReport(employeeID)

		employeeRepository := repository.NewMariaDBEmployeeRepository(db)

		mock.
			ExpectQuery(regexp.QuoteMeta(repository.SQLReportInboundOrders)).
			WithArgs(employeeID).
			WillReturnRows(
				sqlmock.
					NewRows([]string{"id", "card_number_id", "first_name", "last_name", "warehouse_id", "inbound_orders_count"}).
					AddRow(1, "123456", "John", "Doe", int64(1), 10),
			)

		result, err := employeeRepository.GetReportInboundOrdersById(context.TODO(), employeeID)

		assert.NoError(t, err)
		assert.Equal(t, expectedInboundOrders, result)
	})

	t.Run("get_by_id_report_inbound_errors: should return employee not found error when employee not found", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		employeeRepository := repository.NewMariaDBEmployeeRepository(db)

		mock.
			ExpectQuery(regexp.QuoteMeta(repository.SQLFindEmployeeByID)).
			WithArgs(employeeID).
			WillReturnRows(sqlmock.NewRows([]string{}))

		result, err := employeeRepository.GetReportInboundOrdersById(context.TODO(), employeeID)

		assert.Error(t, err)
		assert.Empty(t, result)
	})
}
