package employees

import "fmt"

var employees = []Employee{}

type Repository interface {
	GetAll() ([]Employee, error)
	Get(id int64) (Employee, error)
	Store(cardNumberId string, firstName string, lastName string, warehouseId int64) (Employee, error)
}

type repository struct{}

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

func (repository) Store(cardNumberId string, firstName string, lastName string, warehouseId int64) (Employee, error) {
	employee := Employee{
		Id:           int64(len(employees) + 1),
		CardNumberId: cardNumberId,
		FirstName:    firstName,
		LastName:     lastName,
		WarehouseId:  warehouseId,
	}
	employees = append(employees, employee)
	return employee, nil
}

func NewRepository() Repository {
	return &repository{}
}
