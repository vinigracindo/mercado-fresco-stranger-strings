package services_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	warehouse "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/warehouse/domain"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/warehouse/domain/mocks"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/warehouse/services"
)

var expectedWarehouse = warehouse.WarehouseModel{
	Id:                 0,
	Address:            "Avenida Teste",
	Telephone:          "31 999999999",
	WarehouseCode:      "30",
	MinimunCapacity:    10,
	MinimunTemperature: 9,
	LocalityID:         1,
}

var expectedWarehouseList = []warehouse.WarehouseModel{
	{
		Id:                 0,
		Address:            "Avenida Teste",
		Telephone:          "31 999999999",
		WarehouseCode:      "hg312",
		MinimunCapacity:    1111111,
		MinimunTemperature: 9999999,
		LocalityID:         1,
	},
	{
		Id:                 1,
		Address:            "Avenida Teste Segunda",
		Telephone:          "31 77777777",
		WarehouseCode:      "od78",
		MinimunCapacity:    5555555,
		MinimunTemperature: 444444,
		LocalityID:         2,
	},
}

func Test_Service_Create(t *testing.T) {

	t.Run("create_ok: if all the fields are correct warehouse will be created", func(t *testing.T) {
		ctx := context.Background()

		repo := mocks.NewWarehouseRepository(t)

		repo.On("Create", ctx, &expectedWarehouse).Return(expectedWarehouse, nil)

		service := services.NewWarehouseService(repo)

		result, _ := service.Create(ctx, "Avenida Teste", "31 999999999", "30", 9, 10, 1)

		assert.Equal(t, expectedWarehouse, result)
	})

	t.Run("create_conflict: return erro when try to register a warehouse with code the already exist", func(t *testing.T) {
		ctx := context.Background()

		errMsg := fmt.Errorf("the product with code %d has already been registered", expectedWarehouse.Id)

		repo := mocks.NewWarehouseRepository(t)

		repo.On("Create", ctx, &expectedWarehouse).Return(warehouse.WarehouseModel{}, errMsg)

		service := services.NewWarehouseService(repo)

		_, err := service.Create(ctx, "Avenida Teste", "31 999999999", "30", 9, 10, 1)

		assert.Error(t, err)
	})

}

func Test_Service_GetAll(t *testing.T) {

	t.Run("find_all: return list of warehouses", func(t *testing.T) {
		ctx := context.Background()

		repo := mocks.NewWarehouseRepository(t)

		repo.On("GetAll", ctx).Return(expectedWarehouseList, nil)

		service := services.NewWarehouseService(repo)

		resultList, _ := service.GetAll(ctx)

		assert.Equal(t, expectedWarehouseList, resultList)
	})

	t.Run("find_all_err: error ocorrency on when try to get a list of warehouses", func(t *testing.T) {
		ctx := context.Background()

		errMsg := fmt.Errorf("error: database não found")

		repo := mocks.NewWarehouseRepository(t)

		repo.On("GetAll", ctx).Return([]warehouse.WarehouseModel{}, errMsg)

		service := services.NewWarehouseService(repo)

		_, err := service.GetAll(ctx)

		assert.Error(t, err)
	})
}

func Test_Service_GetByID(t *testing.T) {

	t.Run("find_by_id_existent: search warehouses by id and return", func(t *testing.T) {
		ctx := context.Background()

		repo := mocks.NewWarehouseRepository(t)

		repo.On("GetById", ctx, int64(1)).Return(expectedWarehouseList[1], nil)

		service := services.NewWarehouseService(repo)

		result, _ := service.GetById(ctx, int64(1))

		assert.Equal(t, expectedWarehouseList[1], result)
	})

	t.Run("find_by_id_non_existent: search warehouses by an invalid id and return an error", func(t *testing.T) {
		ctx := context.Background()

		var Id int64 = 9999
		errMsg := fmt.Errorf("erros: no warehouse was found with id %d", Id)

		repo := mocks.NewWarehouseRepository(t)

		repo.On("GetById", ctx, int64(Id)).Return(warehouse.WarehouseModel{}, errMsg)

		service := services.NewWarehouseService(repo)

		_, err := service.GetById(ctx, int64(Id))

		assert.Error(t, err)
	})
}

func Test_Service_UpdateTempAndCap(t *testing.T) {

	updateWarehouse := warehouse.WarehouseModel{
		MinimunCapacity:    999.0,
		MinimunTemperature: 777.0,
	}

	t.Run("update_existent: Se os campos forem atualizados com sucesso retornará a informação do elemento atualizado", func(t *testing.T) {
		ctx := context.Background()

		repo := mocks.NewWarehouseRepository(t)

		repo.On("Update", ctx, expectedWarehouse.Id, &updateWarehouse).Return(expectedWarehouse, nil)

		service := services.NewWarehouseService(repo)

		result, _ := service.UpdateTempAndCap(ctx, expectedWarehouse.Id, updateWarehouse.MinimunTemperature, updateWarehouse.MinimunCapacity)

		assert.Equal(t, expectedWarehouse, result)
	})

	t.Run("update_non_existent: Se não for encontrado um warehouse com o ID retornar um error informando", func(t *testing.T) {
		ctx := context.Background()

		errMsg := fmt.Errorf("erros: no warehouse was found with id %d", expectedWarehouse.Id)

		repo := mocks.NewWarehouseRepository(t)

		repo.On("Update", ctx, expectedWarehouse.Id, &updateWarehouse).Return(warehouse.WarehouseModel{}, errMsg)

		service := services.NewWarehouseService(repo)

		_, err := service.UpdateTempAndCap(ctx, expectedWarehouse.Id, updateWarehouse.MinimunTemperature, updateWarehouse.MinimunCapacity)

		assert.Error(t, err)
	})
}

func Test_Service_Delete(t *testing.T) {

	t.Run("delete_non_existent: return error when try to remove an element the do not exist", func(t *testing.T) {
		ctx := context.Background()

		var id int64 = 1
		errMsg := fmt.Errorf("erros: no warehouse was found with id %d", id)

		repo := mocks.NewWarehouseRepository(t)

		repo.On("Delete", ctx, id).Return(errMsg)

		service := services.NewWarehouseService(repo)

		err := service.Delete(ctx, id)

		assert.NotNil(t, err)
	})

	t.Run("delete_ok: if warehouse was successfully deleted, return empty struct", func(t *testing.T) {
		ctx := context.Background()

		var id int64 = 1

		repo := mocks.NewWarehouseRepository(t)

		repo.On("Delete", ctx, id).Return(nil)

		service := services.NewWarehouseService(repo)

		err := service.Delete(ctx, id)

		assert.Nil(t, err)
	})
}
