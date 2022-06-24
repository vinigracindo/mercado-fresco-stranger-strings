package repository

import (
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/employees/domain"
)

var employees = []domain.Employee{}
var lastId int64 = 1

type repository struct{}

func NewEmployeeRepository() domain.EmployeeRepository {
	return &repository{}
}

func (repository) cardNumberIsUnique(cardNumberId string) bool {
	for _, employee := range employees {
		if employee.CardNumberId == cardNumberId {
			return false
		}
	}
	return true
}

func (repository) GetAll() ([]domain.Employee, error) {
	return employees, nil
}

func (repository) GetById(id int64) (*domain.Employee, error) {
	for _, employee := range employees {
		if employee.Id == id {
			return &employee, nil
		}
	}

	return nil, domain.ErrEmployeeNotFound
}

func (repo repository) Create(cardNumberId string, firstName string, lastName string, warehouseId int64) (domain.Employee, error) {
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

func (repo repository) UpdateFullname(id int64, firstName string, lastName string) (*domain.Employee, error) {
	for i, employee := range employees {
		if employee.Id == id {
			employees[i].FirstName = firstName
			employees[i].LastName = lastName
			return &employees[i], nil
		}
	}
	return nil, domain.ErrEmployeeNotFound
}

func (repo repository) Delete(id int64) error {
	for i, employee := range employees {
		if employee.Id == id {
			employees = append(employees[:i], employees[i+1:]...)
			return nil
		}
	}
	return domain.ErrEmployeeNotFound
}
