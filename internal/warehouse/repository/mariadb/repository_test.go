package respository

import (
	"context"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	warehouse "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/warehouse/domain"
)

func Test_repository_create(t *testing.T) {
	mockWarehouse := warehouse.WarehouseModel{
		Address:            "Avenida Teste",
		Telephone:          "31 999999999",
		WarehouseCode:      "30",
		MinimunCapacity:    10,
		MinimunTemperature: 9,
		LocalityID:         1,
	}

	t.Run("sucess: if all the fields are correct database will create new warehouse and return it", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		defer db.Close()

		mock.ExpectExec(regexp.QuoteMeta(CreateWarehouse)).WithArgs(
			1,
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

		expectedWarehouse := warehouse.WarehouseModel{
			Id:                 1,
			Address:            "Avenida Teste",
			Telephone:          "31 999999999",
			WarehouseCode:      "30",
			MinimunCapacity:    10,
			MinimunTemperature: 9,
			LocalityID:         1,
		}

		assert.Equal(t, expectedWarehouse, newWarehouse)
	})
}
