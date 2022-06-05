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
			wh = append(wh[:id], wh[id+1:]...)
			return nil
		}
	}

	return fmt.Errorf("erros: no warehouse was found with id %d", id)
}
