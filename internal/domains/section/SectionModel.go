package section

type Section struct {
	Id                 int64 `json:"id"`
	SectionNumber      int64 `json:"section_number"`
	CurrentTemperature int64 `json:"current_temperature"`
	MinimumTemperature int64 `json:"minimum_temperature"`
	CurrentCapacity    int64 `json:"current_capacity"`
	MinimumCapacity    int64 `json:"minimum_capacity"`
	MaximumCapacity    int64 `json:"maximum_capacity"`
	WarehouseId        int64 `json:"warehouse_id"`
	ProductTypeId      int64 `json:"product_type_id"`
}

type Repository interface {
	Delete(id int64) error
	UpdateCurrentCapacity(id int64, currentCapacity int64) (Section, error)
	GetById(id int64) (Section, error)
	GetAll() ([]Section, error)
	Create(sectionNumber int64, currentTemperature int64, minimumTemperature int64, currentCapacity int64, minimumCapacity int64, maximumCapacity int64, warehouseId int64, productTypeId int64) (Section, error)
	CreateID() int64
}

type Service interface {
	Delete(id int64) error
	UpdateCurrentCapacity(id int64, currentCapacity int64) (Section, error)
	Create(sectionNumber int64, currentTemperature int64, minimumTemperature int64, currentCapacity int64, minimumCapacity int64, maximumCapacity int64, warehouseId int64, productTypeId int64) (Section, error)
	GetById(id int64) (Section, error)
	GetAll() ([]Section, error)
}
