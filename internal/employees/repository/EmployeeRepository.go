package repository

import (
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/employees/domain"
)

var employees = []domain.Employee{}
var lastId int64 = 1

type mariaDBEmployeerepository struct{}

func NewMariaDBEmployeeRepository() domain.EmployeeRepository {
	return &mariaDBEmployeerepository{}
}

func (mariaDBEmployeerepository) cardNumberIsUnique(cardNumberId string) bool {
	for _, employee := range employees {
		if employee.CardNumberId == cardNumberId {
			return false
		}
	}
	return true
}

func (mariaDBEmployeerepository) GetAll() ([]domain.Employee, error) {
	return employees, nil
}

func (mariaDBEmployeerepository) GetById(id int64) (*domain.Employee, error) {
	for _, employee := range employees {
		if employee.Id == id {
			return &employee, nil
		}
	}

	return nil, domain.ErrEmployeeNotFound
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
