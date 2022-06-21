package warehouse

import (
	"fmt"
)

var wh = []WarehouseModel{}
var id int64 = 0

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (w repository) CreateID() int64 {
	id += 1
	return id
}

func (w repository) Create(new *WarehouseModel) (WarehouseModel, error) {
	for i := range wh {
		if wh[i].WarehouseCode == new.WarehouseCode {
			return WarehouseModel{}, fmt.Errorf("error: already a warehouse with the code: %s", new.WarehouseCode)
		}
	}
	new.Id = w.CreateID()
	wh = append(wh, *new)

	return *new, nil
}

func (w repository) GetAll() ([]WarehouseModel, error) {
	return wh, nil
}

func (w repository) GetById(id int64) (WarehouseModel, error) {
	for i := range wh {
		if wh[i].Id == id {
			return wh[i], nil
		}
	}

	return WarehouseModel{}, fmt.Errorf("erros: no warehouse was found with id %d", id)
}

func (w repository) Delete(id int64) error {
	for i := range wh {
		if wh[i].Id == id {
			wh = append(wh[:i], wh[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("erros: no warehouse was found with id %d", id)
}

func (w repository) Update(id int64, patchwh *WarehouseModel) (WarehouseModel, error) {
	for i := range wh {
		if wh[i].WarehouseCode == patchwh.WarehouseCode {
			return WarehouseModel{}, fmt.Errorf("error: already a warehouse with the code: %s", patchwh.WarehouseCode)
		}

		if wh[i].Id == id {
			if patchwh.MinimunCapacity != 0 {
				wh[i].MinimunCapacity = patchwh.MinimunCapacity
			}

			if patchwh.MinimunTemperature != 0 {
				wh[i].MinimunTemperature = patchwh.MinimunTemperature
			}

			return wh[i], nil
		}
	}
	return WarehouseModel{}, fmt.Errorf("erros: no warehouse was found with id %d", id)
}
