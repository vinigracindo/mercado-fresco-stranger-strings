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

func (repo *mariaDBEmployeerepository) cardNumberIsUnique(cardNumberId string) bool {
	row := repo.db.QueryRow("select card_number_id from employees where card_number_id=?")

	if err := row.Scan(&cardNumberId); err != nil {
		return true
	}
	return false
}

func (repo *mariaDBEmployeerepository) GetAll(ctx context.Context) ([]domain.Employee, error) {
	employees := []domain.Employee{}

	rows, err := repo.db.QueryContext(ctx, "select id, card_number_id, first_name, last_name, warehouse_id from employees")

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

func (repo *mariaDBEmployeerepository) GetById(id int64) (*domain.Employee, error) {
	var employee domain.Employee

	row := repo.db.QueryRowContext(
		context.Background(),
		"select id, card_number_id, first_name, last_name, warehouse_id from employees where id=?",
		id,
	)
	err := row.Scan(&employee.Id, &employee.CardNumberId, &employee.FirstName, &employee.LastName, &employee.WarehouseId)
	if err != nil {
		return nil, domain.ErrEmployeeNotFound
	}

	return &employee, nil
}

func (repo mariaDBEmployeerepository) Create(cardNumberId string, firstName string, lastName string, warehouseId int64) (domain.Employee, error) {
	nextId := lastId
	employee := domain.Employee{
		Id:           nextId,
		CardNumberId: cardNumberId,
		FirstName:    firstName,
		LastName:     lastName,
		WarehouseId:  warehouseId,
	}

	if !repo.cardNumberIsUnique(cardNumberId) {
		return domain.Employee{}, domain.ErrCardNumberMustBeUnique
	}

	employees = append(employees, employee)
	lastId += 1
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
