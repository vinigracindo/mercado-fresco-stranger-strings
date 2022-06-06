package warehouse

type WarehouseModel struct {
	Id                 int64   `json:"id"`
	Address            string  `json:"address"`
	Telephone          string  `json:"telephone"`
	WarehouseCode      string  `json:"warehouse_code,require"`
	MinimunCapacity    int64   `json:"minimun_capacity"`
	MinimunTemperature float64 `json:"minimun_temperature"`
}
