package services

import (
	"context"

	warehouse "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/warehouse/domain"
)

type service struct {
	repository warehouse.WarehouseRepository
}

func NewWarehouseService(r warehouse.WarehouseRepository) warehouse.WarehouseService {
	return &service{
		repository: r,
	}
}

func (s service) Create(ctx context.Context, adress, tel, code string, mintemp float64, mincap int64, locality int64) (warehouse.WarehouseModel, error) {
	new := warehouse.WarehouseModel{
		Address:            adress,
		Telephone:          tel,
		WarehouseCode:      code,
		MinimunCapacity:    mincap,
		MinimunTemperature: mintemp,
		LocalityID:         locality,
	}

	wh, err := s.repository.Create(ctx, &new)

	if err != nil {
		return warehouse.WarehouseModel{}, err
	}

	return wh, nil
}

func (s service) GetAll(ctx context.Context) ([]warehouse.WarehouseModel, error) {
	swh, err := s.repository.GetAll(ctx)

	if err != nil {
		return []warehouse.WarehouseModel{}, err
	}

	return swh, nil
}

func (s service) GetById(ctx context.Context, id int64) (warehouse.WarehouseModel, error) {
	hw, err := s.repository.GetById(ctx, id)

	if err != nil {
		return warehouse.WarehouseModel{}, err
	}

	return hw, nil
}

func (s service) Delete(ctx context.Context, id int64) error {
	err := s.repository.Delete(ctx, id)

	if err != nil {
		return err
	}

	return nil
}

func (s service) UpdateTempAndCap(ctx context.Context, id int64, mintemp float64, mincap int64) (warehouse.WarehouseModel, error) {
	wh := warehouse.WarehouseModel{
		MinimunCapacity:    mincap,
		MinimunTemperature: mintemp,
	}

	parchWh, err := s.repository.Update(ctx, id, &wh)

	if err != nil {
		return warehouse.WarehouseModel{}, err
	}

	return parchWh, nil
}
