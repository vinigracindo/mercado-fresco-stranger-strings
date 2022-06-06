package employees

import "fmt"

var employees = []Employee{}

type Repository interface {
	GetAll() ([]Employee, error)
	GetById(id int64) (Employee, error)
	Store(cardNumberId string, firstName string, lastName string, warehouseId int64) (Employee, error)
	Update(id int64, cardNumberId string, firstName string, lastName string, warehouseId int64) (Employee, error)
	Delete(id int64) error
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

func (repository) GetById(id int64) (Employee, error) {
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

func (repo repository) Update(id int64, cardNumberId string, firstName string, lastName string, warehouseId int64) (Employee, error) {
	for i, employee := range employees {
		if employee.Id == id {
			employees[i].CardNumberId = cardNumberId
			employees[i].FirstName = firstName
			employees[i].LastName = lastName
			employees[i].WarehouseId = warehouseId
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
