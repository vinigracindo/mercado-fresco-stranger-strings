package domain

import "context"

type SectionModel struct {
	Id                 int64   `json:"id"`
	SectionNumber      int64   `json:"section_number"`
	CurrentTemperature float64 `json:"current_temperature"`
	MinimumTemperature float64 `json:"minimum_temperature"`
	CurrentCapacity    int64   `json:"current_capacity"`
	MinimumCapacity    int64   `json:"minimum_capacity"`
	MaximumCapacity    int64   `json:"maximum_capacity"`
	WarehouseId        int64   `json:"warehouse_id"`
	ProductTypeId      int64   `json:"product_type_id"`
}

type ReportProductsModel struct {
	Id            int64 `json:"section_id"`
	SectionNumber int64 `json:"section_number"`
	ProductsCount int64 `json:"products_count"`
}

type SectionRepository interface {
	Delete(ctx context.Context, id int64) error
	UpdateCurrentCapacity(ctx context.Context, section *SectionModel) (*SectionModel, error)
	GetById(ctx context.Context, id int64) (SectionModel, error)
	GetAll(ctx context.Context) ([]SectionModel, error)
	Create(
		ctx context.Context,
		sectionNumber int64,
		currentTemperature float64,
		minimumTemperature float64,
		currentCapacity int64,
		minimumCapacity int64,
		maximumCapacity int64,
		warehouseId int64,
		productTypeId int64) (SectionModel, error)
	GetAllProductCountBySection(ctx context.Context) (*[]ReportProductsModel, error)
	GetByIdProductCountBySection(ctx context.Context, id int64) (*ReportProductsModel, error)
}

type SectionService interface {
	Delete(ctx context.Context, id int64) error
	UpdateCurrentCapacity(ctx context.Context, id int64, currentCapacity int64) (*SectionModel, error)
	Create(
		ctx context.Context,
		sectionNumber int64,
		currentTemperature float64,
		minimumTemperature float64,
		currentCapacity int64,
		minimumCapacity int64,
		maximumCapacity int64,
		warehouseId int64,
		productTypeId int64) (SectionModel, error)
	GetById(ctx context.Context, id int64) (SectionModel, error)
	GetAll(ctx context.Context) ([]SectionModel, error)
	GetAllProductCountBySection(ctx context.Context) (*[]ReportProductsModel, error)
	GetByIdProductCountBySection(ctx context.Context, id int64) (*ReportProductsModel, error)
}
