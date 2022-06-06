package employees

import "fmt"

var employees = []Employee{}

type Repository interface {
	GetAll() ([]Employee, error)
	Get(id int64) (Employee, error)
	Store(cardNumberId string, firstName string, lastName string, warehouseId int64) (Employee, error)
}

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

func (repository) Get(id int64) (Employee, error) {
	for _, employee := range employees {
		if employee.Id == id {
			return employee, nil
		}
	}

	return Employee{}, fmt.Errorf("employee with id %d not found", id)
}

func (repo repository) Store(cardNumberId string, firstName string, lastName string, warehouseId int64) (Employee, error) {
	nextId := employees[len(employees)-1].Id + 1
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
	return employee, nil
}

func NewRepository() Repository {
	return &repository{}
}
