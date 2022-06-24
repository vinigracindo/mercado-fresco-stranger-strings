package domain

type SectionModel struct {
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

type SectionRepository interface {
	Delete(id int64) error
	UpdateCurrentCapacity(id int64, currentCapacity int64) (SectionModel, error)
	GetById(id int64) (SectionModel, error)
	GetAll() ([]SectionModel, error)
	Create(
		sectionNumber int64,
		currentTemperature int64,
		minimumTemperature int64,
		currentCapacity int64,
		minimumCapacity int64,
		maximumCapacity int64,
		warehouseId int64,
		productTypeId int64) (SectionModel, error)
	CreateID() int64
}

type SectionService interface {
	Delete(id int64) error
	UpdateCurrentCapacity(id int64, currentCapacity int64) (SectionModel, error)
	Create(
		sectionNumber int64,
		currentTemperature int64,
		minimumTemperature int64,
		currentCapacity int64,
		minimumCapacity int64,
		maximumCapacity int64,
		warehouseId int64,
		productTypeId int64) (SectionModel, error)
	GetById(id int64) (SectionModel, error)
	GetAll() ([]SectionModel, error)
}
