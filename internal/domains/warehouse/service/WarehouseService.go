package service

import "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/warehouse"

type service struct {
	repository warehouse.Repository
}

func NewService(r warehouse.Repository) warehouse.Service {
	return &service{
		repository: r,
	}
}

func (s service) Create(adress, tel, code string, mintemp float64, mincap int64) (warehouse.WarehouseModel, error) {
	new := warehouse.WarehouseModel{
		Address:            adress,
		Telephone:          tel,
		WarehouseCode:      code,
		MinimunCapacity:    mincap,
		MinimunTemperature: mintemp,
	}

	wh, err := s.repository.Create(&new)

	if err != nil {
		return warehouse.WarehouseModel{}, err
	}

	return wh, nil
}

func (s service) GetAll() ([]warehouse.WarehouseModel, error) {
	swh, err := s.repository.GetAll()

	if err != nil {
		return []warehouse.WarehouseModel{}, err
	}

	return swh, nil
}

func (s service) GetById(id int64) (warehouse.WarehouseModel, error) {
	hw, err := s.repository.GetById(id)

	if err != nil {
		return warehouse.WarehouseModel{}, err
	}

	return hw, nil
}

func (s service) Delete(id int64) error {
	err := s.repository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}

func (s service) UpdateTempAndCap(id int64, mintemp float64, mincap int64) (warehouse.WarehouseModel, error) {
	wh := warehouse.WarehouseModel{
		MinimunCapacity:    mincap,
		MinimunTemperature: mintemp,
	}

	parchWh, err := s.repository.Update(id, &wh)

	if err != nil {
		return warehouse.WarehouseModel{}, err
	}

	return parchWh, nil
}
