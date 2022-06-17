package employees

type Employee struct {
	Id           int64  `json:"id"`
	CardNumberId string `json:"card_number_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	WarehouseId  int64  `json:"warehouse_id"`
}

type Service interface {
	GetAll() ([]Employee, error)
	GetById(id int64) (Employee, error)
	Store(cardNumberId string, firstName string, lastName string, warehouseId int64) (Employee, error)
	UpdateFullname(id int64, firstName string, lastName string) (Employee, error)
	Delete(id int64) error
}

type Repository interface {
	GetAll() ([]Employee, error)
	GetById(id int64) (Employee, error)
	Store(cardNumberId string, firstName string, lastName string, warehouseId int64) (Employee, error)
	UpdateFullname(id int64, firstName string, lastName string) (Employee, error)
	Delete(id int64) error
}
