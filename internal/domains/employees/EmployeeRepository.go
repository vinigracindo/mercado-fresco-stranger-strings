package employees

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

func (repository) GetById(id int64) (*Employee, error) {
	for _, employee := range employees {
		if employee.Id == id {
			return &employee, nil
		}
	}

	return nil, ErrEmployeeNotFound
}

func (repo repository) Create(cardNumberId string, firstName string, lastName string, warehouseId int64) (Employee, error) {
	nextId := lastId
	employee := Employee{
		Id:           nextId,
		CardNumberId: cardNumberId,
		FirstName:    firstName,
		LastName:     lastName,
		WarehouseId:  warehouseId,
	}

	if !repo.cardNumberIsUnique(cardNumberId) {
		return Employee{}, ErrCardNumberMustBeUnique
	}

	employees = append(employees, employee)
	lastId += 1
	return employee, nil
}

func (repo repository) UpdateFullname(id int64, firstName string, lastName string) (*Employee, error) {
	for i, employee := range employees {
		if employee.Id == id {
			employees[i].FirstName = firstName
			employees[i].LastName = lastName
			return &employees[i], nil
		}
	}
	return nil, ErrEmployeeNotFound
}

func (repo repository) Delete(id int64) error {
	for i, employee := range employees {
		if employee.Id == id {
			employees = append(employees[:i], employees[i+1:]...)
			return nil
		}
	}
	return ErrEmployeeNotFound
}

func NewRepository() Repository {
	return &repository{}
}
