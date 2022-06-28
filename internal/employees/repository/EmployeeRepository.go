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
		return employee, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return employee, err
	}
	employee.Id = int64(id)

	return employee, nil
}

func (repo mariaDBEmployeerepository) UpdateFullname(ctx context.Context, id int64, firstName string, lastName string) (*domain.Employee, error) {
	employee, err := repo.GetById(context.Background(), id)
	if err != nil {
		return nil, domain.ErrEmployeeNotFound
	}

	employee.FirstName = firstName
	employee.LastName = lastName

	_, err = repo.db.ExecContext(
		ctx,
		SQLUpdateEmployeeFullname,
		&employee.FirstName, &employee.LastName, &employee.Id,
	)

	if err != nil {
		return nil, err
	}

	return employee, nil
}

func (repo mariaDBEmployeerepository) Delete(ctx context.Context, id int64) error {
	_, err := repo.db.ExecContext(ctx, SQLDeleteEmployee, id)
	if err != nil {
		return domain.ErrEmployeeNotFound
	}
	return nil
}
