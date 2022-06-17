package warehouse

type WarehouseModel struct {
	Id                 int64   `json:"id"`
	Address            string  `json:"address"`
	Telephone          string  `json:"telephone"`
	WarehouseCode      string  `json:"warehouse_code,require"`
	MinimunCapacity    int64   `json:"minimun_capacity"`
	MinimunTemperature float64 `json:"minimun_temperature"`
}

type Repository interface {
	Create(wr *WarehouseModel) (WarehouseModel, error)
	GetAll() ([]WarehouseModel, error)
	GetById(id int64) (WarehouseModel, error)
	Delete(id int64) error
	Update(id int64, wh *WarehouseModel) (WarehouseModel, error)
	CreateID() int64
}

type Service interface {
	GetAll() ([]WarehouseModel, error)
	GetById(id int64) (WarehouseModel, error)
	Delete(id int64) error
	UpdateTempAndCap(id int64, mintemp float64, mincap int64) (WarehouseModel, error)
	Create(adress, tel, code string, mintemp float64, mincap int64) (WarehouseModel, error)
}
