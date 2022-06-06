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
