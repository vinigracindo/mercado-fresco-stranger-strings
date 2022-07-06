package repository

import (
	"context"
	"database/sql"

	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/employees/domain"
)

type mariaDBEmployeerepository struct {
	db *sql.DB
}

func NewMariaDBEmployeeRepository(db *sql.DB) domain.EmployeeRepository {
	return &mariaDBEmployeerepository{db: db}
}

func (repo *mariaDBEmployeerepository) GetAll(ctx context.Context) ([]domain.Employee, error) {
	employees := []domain.Employee{}

	rows, err := repo.db.QueryContext(ctx, SQLFindAllEmployees)

	if err != nil {
		return employees, err
	}
	defer rows.Close()

	for rows.Next() {
		var employee domain.Employee
		if err := rows.Scan(&employee.Id, &employee.CardNumberId, &employee.FirstName, &employee.LastName, &employee.WarehouseId); err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}

	return employees, nil
}

func (repo *mariaDBEmployeerepository) GetById(ctx context.Context, id int64) (*domain.Employee, error) {
	var employee domain.Employee

	row := repo.db.QueryRowContext(ctx, SQLFindEmployeeByID, id)
	err := row.Scan(&employee.Id, &employee.CardNumberId, &employee.FirstName, &employee.LastName, &employee.WarehouseId)
	if err != nil {
		return nil, domain.ErrEmployeeNotFound
	}

	return &employee, nil
}

func (repo *mariaDBEmployeerepository) Create(ctx context.Context, cardNumberId string, firstName string, lastName string, warehouseId int64) (domain.Employee, error) {
	employee := domain.Employee{
		CardNumberId: cardNumberId,
		FirstName:    firstName,
		LastName:     lastName,
		WarehouseId:  warehouseId,
	}
	res, err := repo.db.ExecContext(
		ctx,
		SQLCreateEmployee,
		&employee.CardNumberId, &employee.FirstName, &employee.LastName, &employee.WarehouseId,
	)
	if err != nil {
		return domain.Employee{}, err
	}

	id, _ := res.LastInsertId()
	employee.Id = id

	return employee, nil
}

func (repo mariaDBEmployeerepository) Update(ctx context.Context, employeeID int64, updatedEmployee domain.Employee) error {
	res, err := repo.db.ExecContext(
		ctx,
		SQLUpdateEmployeeFullname,
		updatedEmployee.FirstName, updatedEmployee.LastName, employeeID,
	)

	if err != nil {
		return err
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return domain.ErrEmployeeNotFound
	}

	return nil
}

func (repo mariaDBEmployeerepository) Delete(ctx context.Context, id int64) error {
	result, err := repo.db.ExecContext(ctx, SQLDeleteEmployee, id)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return domain.ErrEmployeeNotFound
	}

	return nil
}

func (repo mariaDBEmployeerepository) ReportInboundOrders(ctx context.Context, employeeID *int64) ([]domain.EmployeeInboundOrdersReport, error) {
	result := []domain.EmployeeInboundOrdersReport{}

	rows, err := repo.db.QueryContext(ctx, SQLReportInboundOrders, employeeID)

	if err != nil {
		return result, err
	}

	for rows.Next() {
		res := domain.EmployeeInboundOrdersReport{}

		err := rows.Scan(
			&res.Id,
			&res.CardNumberId,
			&res.FirstName,
			&res.LastName,
			&res.WarehouseId,
			&res.Count,
		)
		if err != nil {
			return nil, err
		}

		result = append(result, res)
	}

	return result, nil
}
