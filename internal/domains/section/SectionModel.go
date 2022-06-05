package section

type Section struct {
	Id                 int64 `json:"id"`
	SectionNumber      int64 `json:"section_number" binding:"required"`
	CurrentTemperature int64 `json:"current_temperature" binding:"required"`
	MinimumTemperature int64 `json:"minimum_temperature" binding:"required"`
	CurrentCapacity    int64 `json:"current_capacity" binding:"required"`
	MinimumCapacity    int64 `json:"minimum_capacity" binding:"required"`
	MaximumCapacity    int64 `json:"maximum_capacity" binding:"required"`
	WarehouseId        int64 `json:"warehouse_id" binding:"required"`
	ProductTypeId      int64 `json:"product_type_id" binding:"required"`
}
