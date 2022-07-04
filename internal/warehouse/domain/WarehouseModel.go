package domain

import "context"

type WarehouseModel struct {
	Id                 int64   `json:"id"`
	Address            string  `json:"address"`
	Telephone          string  `json:"telephone"`
	WarehouseCode      string  `json:"warehouse_code"`
	MinimunCapacity    int64   `json:"minimun_capacity"`
	MinimunTemperature float64 `json:"minimun_temperature"`
	LocalityID         int64   `json:"locality_id"`
}

type WarehouseRepository interface {
	Create(ctx context.Context, wr *WarehouseModel) (WarehouseModel, error)
	GetAll(ctx context.Context) ([]WarehouseModel, error)
	GetById(ctx context.Context, id int64) (WarehouseModel, error)
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, id int64, wh *WarehouseModel) (WarehouseModel, error)
}

type WarehouseService interface {
	GetAll(ctx context.Context) ([]WarehouseModel, error)
	GetById(ctx context.Context, id int64) (WarehouseModel, error)
	Delete(ctx context.Context, id int64) error
	UpdateTempAndCap(ctx context.Context, id int64, mintemp float64, mincap int64) (WarehouseModel, error)
	Create(ctx context.Context, adress, tel, code string, mintemp float64, mincap int64, locality int64) (WarehouseModel, error)
}
