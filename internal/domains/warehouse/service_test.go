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

func Test_Service_GetAll(t *testing.T) {
	expectedWarehouseList := []warehouse.WarehouseModel{
		{
			Id:                 0,
			Address:            "Avenida Teste",
			Telephone:          "31 999999999",
			WarehouseCode:      "hg312",
			MinimunCapacity:    1111111,
			MinimunTemperature: 9999999,
		},
		{
			Id:                 1,
			Address:            "Avenida Teste Segunda",
			Telephone:          "31 77777777",
			WarehouseCode:      "od78",
			MinimunCapacity:    5555555,
			MinimunTemperature: 444444,
		},
	}

	t.Run("find_all: retonar uma list com varios warehouses", func(t *testing.T) {
		repo := mocks.NewRepository(t)

		repo.On("GetAll").Return(expectedWarehouseList, nil)

		service := warehouse.NewService(repo)

		resultList, _ := service.GetAll()

		assert.Equal(t, expectedWarehouseList, resultList)
	})

	t.Run("find_all_err: retonar um erro ao tentar buscar pro todas as warehouses", func(t *testing.T) {

		errMsg := fmt.Errorf("error: database não found")

		repo := mocks.NewRepository(t)

		repo.On("GetAll").Return([]warehouse.WarehouseModel{}, errMsg)

		service := warehouse.NewService(repo)

		_, err := service.GetAll()

		assert.Error(t, err)
	})
}

func Test_Service_GetByID(t *testing.T) {
	expectedWarehouseList := []warehouse.WarehouseModel{
		{
			Id:                 0,
			Address:            "Avenida Teste",
			Telephone:          "31 999999999",
			WarehouseCode:      "hg312",
			MinimunCapacity:    1111111,
			MinimunTemperature: 9999999,
		},
		{
			Id:                 1,
			Address:            "Avenida Teste Segunda",
			Telephone:          "31 77777777",
			WarehouseCode:      "od78",
			MinimunCapacity:    5555555,
			MinimunTemperature: 444444,
		},
	}

	t.Run("find_by_id_existent: procura um warehouse pelo id valida e retornar", func(t *testing.T) {
		repo := mocks.NewRepository(t)

		repo.On("GetById", int64(1)).Return(expectedWarehouseList[1], nil)

		service := warehouse.NewService(repo)

		result, _ := service.GetById(int64(1))

		assert.Equal(t, expectedWarehouseList[1], result)
	})

	t.Run("find_by_id_non_existent: procura um warehouse por um ID invalida e retonar um erro", func(t *testing.T) {

		var fakeID int64 = 9999
		errMsg := fmt.Errorf("erros: no warehouse was found with id %d", fakeID)

		repo := mocks.NewRepository(t)

		repo.On("GetById", int64(fakeID)).Return(warehouse.WarehouseModel{}, errMsg)

		service := warehouse.NewService(repo)

		_, err := service.GetById(int64(fakeID))

		assert.Error(t, err)
	})
}
