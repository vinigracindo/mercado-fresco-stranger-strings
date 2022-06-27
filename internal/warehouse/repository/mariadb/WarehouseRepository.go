package respository

import (
	"context"
	"database/sql"

	warehouse "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/warehouse/domain"
)

type mariadbWarehouse struct {
	db *sql.DB
}

func NewMariadbWarehouseRepository(connection *sql.DB) warehouse.WarehouseRepository {
	return &mariadbWarehouse{
		db: connection,
	}
}

func (r *mariadbWarehouse) Create(ctx context.Context, wr *warehouse.WarehouseModel) (warehouse.WarehouseModel, error) {
	result, err := r.db.ExecContext(
		ctx,
		CreateWarehouse,
		wr.Address,
		wr.Telephone,
		wr.WarehouseCode,
		wr.MinimunCapacity,
		wr.MinimunTemperature,
		wr.LocalityID,
	)

	if err != nil {
		return warehouse.WarehouseModel{}, err
	}

	newWarehouseId, err := result.LastInsertId()

	if err != nil {
		return warehouse.WarehouseModel{}, err
	}

	return warehouse.WarehouseModel{
		Id:                 newWarehouseId,
		Address:            wr.Address,
		Telephone:          wr.Telephone,
		WarehouseCode:      wr.WarehouseCode,
		MinimunCapacity:    wr.MinimunCapacity,
		MinimunTemperature: wr.MinimunTemperature,
	}, nil
}

func (r *mariadbWarehouse) GetAll(ctx context.Context) ([]warehouse.WarehouseModel, error) {
	return []warehouse.WarehouseModel{}, nil
}
func (r *mariadbWarehouse) GetById(ctx context.Context, id int64) (warehouse.WarehouseModel, error) {
	return warehouse.WarehouseModel{}, nil
}
func (r *mariadbWarehouse) Delete(ctx context.Context, id int64) error {
	return nil
}
func (r *mariadbWarehouse) Update(ctx context.Context, id int64, wh *warehouse.WarehouseModel) (warehouse.WarehouseModel, error) {
	return warehouse.WarehouseModel{}, nil
}
