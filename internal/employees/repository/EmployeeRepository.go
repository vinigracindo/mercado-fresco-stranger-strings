package repository

import (
	"context"
	"database/sql"

	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/employees/domain"
)

var employees = []domain.Employee{}
var lastId int64 = 1

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
			return employees, err
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

func (repo *mariaDBEmployeerepository) Create(cardNumberId string, firstName string, lastName string, warehouseId int64) (domain.Employee, error) {
	employee := domain.Employee{
		CardNumberId: cardNumberId,
		FirstName:    firstName,
		LastName:     lastName,
		WarehouseId:  warehouseId,
	}
	res, err := repo.db.ExecContext(
		context.Background(),
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

func (repo mariaDBEmployeerepository) UpdateFullname(id int64, firstName string, lastName string) (*domain.Employee, error) {
	for i, employee := range employees {
		if employee.Id == id {
			employees[i].FirstName = firstName
			employees[i].LastName = lastName
			return &employees[i], nil
		}
	}
	return nil, domain.ErrEmployeeNotFound
}

func (repo mariaDBEmployeerepository) Delete(id int64) error {
	for i, employee := range employees {
		if employee.Id == id {
			employees = append(employees[:i], employees[i+1:]...)
			return nil
		}
	}
	return domain.ErrEmployeeNotFound
}
