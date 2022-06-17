package warehouse_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/warehouse"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/warehouse/mocks"
)

func Test_Service_Create(t *testing.T) {

	expectedWarehouse := warehouse.WarehouseModel{
		Id:                 0,
		Address:            "Avenida Teste",
		Telephone:          "31 999999999",
		WarehouseCode:      "30",
		MinimunCapacity:    10,
		MinimunTemperature: 9,
	}

	t.Run("create_ok: Se contiver os campos necessários, será criado", func(t *testing.T) {

		repo := mocks.NewRepository(t)

		repo.On("Create", &expectedWarehouse).Return(expectedWarehouse, nil)

		service := warehouse.NewService(repo)

		result, _ := service.Create("Avenida Teste", "31 999999999", "30", 9, 10)

		assert.Equal(t, expectedWarehouse, result)
	})

	t.Run("create_conflict: warehouse_code duplicado", func(t *testing.T) {

		errMsg := fmt.Errorf("the product with code %d has already been registered", expectedWarehouse.Id)

		repo := mocks.NewRepository(t)

		repo.On("Create", &expectedWarehouse).Return(warehouse.WarehouseModel{}, errMsg)

		service := warehouse.NewService(repo)

		_, err := service.Create("Avenida Teste", "31 999999999", "30", 9, 10)

		assert.Error(t, err)
	})

}
