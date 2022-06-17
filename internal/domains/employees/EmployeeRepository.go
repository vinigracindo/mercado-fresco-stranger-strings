package employees

import "fmt"

var employees = []Employee{}
var lastId int64 = 1

type repository struct{}

func (repository) cardNumberIsUnique(cardNumberId string) bool {
	for _, employee := range employees {
		if employee.CardNumberId == cardNumberId {
			return false
		}
	}
	return true
}

func (repository) GetAll() ([]Employee, error) {
	return employees, nil
}

func (repository) GetById(id int64) (Employee, error) {
	for _, employee := range employees {
		if employee.Id == id {
			return employee, nil
		}
	}

	return Employee{}, fmt.Errorf("employee with id %d not found", id)
}

func (repo repository) Store(cardNumberId string, firstName string, lastName string, warehouseId int64) (Employee, error) {
	nextId := lastId
	employee := Employee{
		Id:           nextId,
		CardNumberId: cardNumberId,
		FirstName:    firstName,
		LastName:     lastName,
		WarehouseId:  warehouseId,
	}

	if !repo.cardNumberIsUnique(cardNumberId) {
		return Employee{}, fmt.Errorf("card number %s is already in use", cardNumberId)
	}

	employees = append(employees, employee)
	lastId += 1
	return employee, nil
}

func (repo repository) UpdateFullname(id int64, firstName string, lastName string) (Employee, error) {
	for i, employee := range employees {
		if employee.Id == id {
			employees[i].FirstName = firstName
			employees[i].LastName = lastName
			return employees[i], nil
		}
	}
	return Employee{}, fmt.Errorf("employee with id %d not found", id)
}

func (repo repository) Delete(id int64) error {
	for i, employee := range employees {
		if employee.Id == id {
			employees = append(employees[:i], employees[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("employee with id %d not found", id)
}

func NewRepository() Repository {
	return &repository{}
}
