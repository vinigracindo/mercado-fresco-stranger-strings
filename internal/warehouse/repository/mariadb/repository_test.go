package respository

import (
	"context"
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	warehouse "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/warehouse/domain"
)

var expectedWarehouse warehouse.WarehouseModel = warehouse.WarehouseModel{
	Id:                 1,
	Address:            "Avenida Teste",
	Telephone:          "31 999999999",
	WarehouseCode:      "30",
	MinimunCapacity:    10,
	MinimunTemperature: 9,
	LocalityID:         1,
}

var mockWarehouse warehouse.WarehouseModel = warehouse.WarehouseModel{
	Address:            "Avenida Teste",
	Telephone:          "31 999999999",
	WarehouseCode:      "30",
	MinimunCapacity:    10,
	MinimunTemperature: 9,
	LocalityID:         1,
}

var updateWarehouse warehouse.WarehouseModel = warehouse.WarehouseModel{
	MinimunCapacity:    77777,
	MinimunTemperature: 888888,
}

func Test_repository_create(t *testing.T) {

	t.Run("sucess: if all the fields are correct database will create new warehouse and return it", func(t *testing.T) {

		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		defer db.Close()

		mock.ExpectExec(regexp.QuoteMeta(CreateWarehouse)).WithArgs(
			mockWarehouse.Address,
			mockWarehouse.Telephone,
			mockWarehouse.WarehouseCode,
			mockWarehouse.MinimunCapacity,
			mockWarehouse.MinimunTemperature,
			mockWarehouse.LocalityID,
		).WillReturnResult(sqlmock.NewResult(1, 1))

		mariadbWarehouse := NewMariadbWarehouseRepository(db)

		newWarehouse, err := mariadbWarehouse.Create(context.TODO(), &mockWarehouse)

		assert.NoError(t, err)
		assert.Equal(t, expectedWarehouse, newWarehouse)
	})

	t.Run("error: when ExecContext return an error", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		defer db.Close()

		mock.ExpectExec(regexp.QuoteMeta(CreateWarehouse)).WithArgs(
			mockWarehouse.Address,
			mockWarehouse.Telephone,
			mockWarehouse.WarehouseCode,
			mockWarehouse.MinimunCapacity,
			mockWarehouse.MinimunTemperature,
			mockWarehouse.LocalityID,
		).WillReturnError(fmt.Errorf("some problem on the query"))

		mariadbWarehouse := NewMariadbWarehouseRepository(db)

		newWarehouse, err := mariadbWarehouse.Create(context.TODO(), &mockWarehouse)

		assert.Error(t, err)
		assert.Empty(t, newWarehouse)
	})

}

func Test_repository_update(t *testing.T) {
	t.Run("update_not_found", func(t *testing.T) {

		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		defer db.Close()

		mock.ExpectExec(regexp.QuoteMeta(UpdateWarehouse)).WithArgs(
			updateWarehouse.MinimunCapacity,
			updateWarehouse.MinimunTemperature,
			expectedWarehouse.Id,
		).WillReturnResult(sqlmock.NewResult(0, 0))

		mariadbWarehouse := NewMariadbWarehouseRepository(db)

		_, err = mariadbWarehouse.Update(context.TODO(), expectedWarehouse.Id, &updateWarehouse)

		assert.Error(t, err)

	})

	t.Run("update_sucess", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(regexp.QuoteMeta(UpdateWarehouse)).WithArgs(
			updateWarehouse.MinimunCapacity,
			updateWarehouse.MinimunTemperature,
			expectedWarehouse.Id,
		).WillReturnResult(sqlmock.NewResult(0, 1))

		mariadbWarehouse := NewMariadbWarehouseRepository(db)

		expectReturn, err := mariadbWarehouse.Update(context.Background(), expectedWarehouse.Id, &updateWarehouse)

		assert.NoError(t, err)
		assert.Equal(t, expectReturn, updateWarehouse)
	})
}
