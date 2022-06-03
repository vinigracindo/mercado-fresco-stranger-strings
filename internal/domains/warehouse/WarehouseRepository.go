package warehouse

import (
	"fmt"
)

var wh = make([]WarehouseModel, 100)
var id int64 = 0

type repository struct{}

type Repository interface {
	Store(wr *WarehouseModel) (WarehouseModel, error)
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
