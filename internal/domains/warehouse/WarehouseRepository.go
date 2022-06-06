package warehouse

import (
	"fmt"
)

var wh = []WarehouseModel{}
var id int64 = 0

type repository struct{}

type Repository interface {
	Store(wr *WarehouseModel) (WarehouseModel, error)
	GetAll() ([]WarehouseModel, error)
	GetById(id int64) (WarehouseModel, error)
	Delete(id int64) error
	Update(id int64, wh *WarehouseModel) (WarehouseModel, error)
	createID() int64
}

func NewRepository() Repository {
	return &repository{}
}

func (w repository) createID() int64 {
	id += 1
	return id
}

func (w repository) Store(new *WarehouseModel) (WarehouseModel, error) {
	for i := range wh {
		if wh[i].WarehouseCode == new.WarehouseCode {
			return WarehouseModel{}, fmt.Errorf("error: already a warehouse with the code: %s", new.WarehouseCode)
		}
	}
	new.Id = w.createID()
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
			if patchwh.Address != "" {
				wh[i].Address = patchwh.Address
			}

			if patchwh.Telephone != "" {
				wh[i].Telephone = patchwh.Telephone
			}

			if patchwh.WarehouseCode != "" {
				wh[i].WarehouseCode = patchwh.WarehouseCode
			}

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
