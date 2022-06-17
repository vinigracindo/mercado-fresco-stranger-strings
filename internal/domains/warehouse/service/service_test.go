package service_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/warehouse"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/warehouse/mocks"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/warehouse/service"
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

		service := service.NewService(repo)

		result, _ := service.Create("Avenida Teste", "31 999999999", "30", 9, 10)

		assert.Equal(t, expectedWarehouse, result)
	})

	t.Run("create_conflict: warehouse_code duplicado", func(t *testing.T) {

		errMsg := fmt.Errorf("the product with code %d has already been registered", expectedWarehouse.Id)

		repo := mocks.NewRepository(t)

		repo.On("Create", &expectedWarehouse).Return(warehouse.WarehouseModel{}, errMsg)

		service := service.NewService(repo)

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

		service := service.NewService(repo)

		resultList, _ := service.GetAll()

		assert.Equal(t, expectedWarehouseList, resultList)
	})

	t.Run("find_all_err: retonar um erro ao tentar buscar pro todas as warehouses", func(t *testing.T) {

		errMsg := fmt.Errorf("error: database não found")

		repo := mocks.NewRepository(t)

		repo.On("GetAll").Return([]warehouse.WarehouseModel{}, errMsg)

		service := service.NewService(repo)

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

		service := service.NewService(repo)

		result, _ := service.GetById(int64(1))

		assert.Equal(t, expectedWarehouseList[1], result)
	})

	t.Run("find_by_id_non_existent: procura um warehouse por um ID invalida e retonar um erro", func(t *testing.T) {

		var Id int64 = 9999
		errMsg := fmt.Errorf("erros: no warehouse was found with id %d", Id)

		repo := mocks.NewRepository(t)

		repo.On("GetById", int64(Id)).Return(warehouse.WarehouseModel{}, errMsg)

		service := service.NewService(repo)

		_, err := service.GetById(int64(Id))

		assert.Error(t, err)
	})
}

func Test_Service_UpdateTempAndCap(t *testing.T) {
	expectedWarehouse := warehouse.WarehouseModel{
		Id:                 1,
		Address:            "Avenida Teste Segunda",
		Telephone:          "31 77777777",
		WarehouseCode:      "od78",
		MinimunCapacity:    999.0,
		MinimunTemperature: 999.0,
	}

	updateWarehouse := warehouse.WarehouseModel{
		MinimunCapacity:    999.0,
		MinimunTemperature: 777.0,
	}

	t.Run("update_existent: Se os campos forem atualizados com sucesso retornará a informação do elemento atualizado", func(t *testing.T) {
		repo := mocks.NewRepository(t)

		repo.On("Update", expectedWarehouse.Id, &updateWarehouse).Return(expectedWarehouse, nil)

		service := service.NewService(repo)

		result, _ := service.UpdateTempAndCap(expectedWarehouse.Id, updateWarehouse.MinimunTemperature, updateWarehouse.MinimunCapacity)

		assert.Equal(t, expectedWarehouse, result)
	})

	t.Run("update_non_existent: Se não for encontrado um warehouse com o ID retornar um error informando", func(t *testing.T) {

		errMsg := fmt.Errorf("erros: no warehouse was found with id %d", expectedWarehouse.Id)

		repo := mocks.NewRepository(t)

		repo.On("Update", expectedWarehouse.Id, &updateWarehouse).Return(warehouse.WarehouseModel{}, errMsg)

		service := service.NewService(repo)

		_, err := service.UpdateTempAndCap(expectedWarehouse.Id, updateWarehouse.MinimunTemperature, updateWarehouse.MinimunCapacity)

		assert.Error(t, err)
	})
}

func Test_Service_Delete(t *testing.T) {

	t.Run("delete_non_existent: Se o elemento a ser removido não existir, ele retornará erro.", func(t *testing.T) {
		var id int64 = 1
		errMsg := fmt.Errorf("erros: no warehouse was found with id %d", id)

		repo := mocks.NewRepository(t)

		repo.On("Delete", id).Return(errMsg)

		service := service.NewService(repo)

		err := service.Delete(id)

		assert.NotNil(t, err)
	})

	t.Run("delete_ok: Se a exclusão for bem-sucedida, retorna um erro vazio.", func(t *testing.T) {
		var id int64 = 1

		repo := mocks.NewRepository(t)

		repo.On("Delete", id).Return(nil)

		service := service.NewService(repo)

		err := service.Delete(id)

		assert.Nil(t, err)
	})
}
