package seller_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/seller"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/seller/mocks"
)

func Test_Service_Creat(t *testing.T) {
	expectedSeller := seller.Seller{
		Id:          1,
		Cid:         123,
		CompanyName: "Mercado Livre",
		Address:     "Osasco, SP",
		Telephone:   "11 99999999",
	}

	t.Run("create_ok: Se contiver os campos necessários, será criado", func(t *testing.T) {
		repo := mocks.NewRepository(t)
		repo.On("Create", expectedSeller.Cid, expectedSeller.CompanyName, expectedSeller.Address, expectedSeller.Telephone).Return(expectedSeller, nil)
		service := seller.NewService(repo)

		result, err := service.Create(123, "Mercado Livre", "Osasco, SP", "11 99999999")

		assert.Nil(t, err)
		assert.Equal(t, expectedSeller, result)

	})

	t.Run("create_conflict: Se o cid já existir ele não pode ser criado", func(t *testing.T) {
		errMsg := fmt.Errorf("The seller whith cid %d has already been registered", expectedSeller.Cid)
		repo := mocks.NewRepository(t)
		repo.On("Create", expectedSeller.Cid, expectedSeller.CompanyName, expectedSeller.Address, expectedSeller.Telephone).Return(seller.Seller{}, errMsg)
		service := seller.NewService(repo)

		result, err := service.Create(123, "Mercado Livre", "Osasco, SP", "11 99999999")

		assert.Equal(t, seller.Seller{}, result)
		assert.Equal(t, errMsg, err)

	})
}

func Test_Service_GetAll(t *testing.T) {
	expectedListSeller := []seller.Seller{
		{
			Id:          1,
			Cid:         123,
			CompanyName: "Mercado Livre",
			Address:     "Osasco, SP",
			Telephone:   "11 99999999",
		},
		{
			Id:          2,
			Cid:         1234,
			CompanyName: "Mercado Pago",
			Address:     "Salvador, BA",
			Telephone:   "11 88888888",
		},
	}
	t.Run("find_all: Se a lista tiver 'n' elementos, retornará uma quantidade do total de elementos", func(t *testing.T) {
		repo := mocks.NewRepository(t)
		repo.On("GetAll").Return(expectedListSeller, nil)
		service := seller.NewService(repo)

		seller, err := service.GetAll()

		assert.Equal(t, seller, expectedListSeller)
		assert.Nil(t, err)

	})

}

func Test_Service_GetById(t *testing.T) {
	repo := mocks.NewRepository(t)

	t.Run("find_by_id_non_existent: se o elemento procurado por id não existir, retorna nil", func(t *testing.T) {
		repo.On("GetById", int64(1)).Return(seller.Seller{}, fmt.Errorf("Seller not found."))
		service := seller.NewService(repo)

		result, err := service.GetById(int64(1))

		assert.Error(t, err)
		assert.Equal(t, seller.Seller{}, result)
	})

	t.Run("find_by_id_existent: se o elemento procurado por id existir, ele retornará as informações do elemento solicitado", func(t *testing.T) {
		expectedListSeller := []seller.Seller{
			{
				Id:          1,
				Cid:         123,
				CompanyName: "Mercado Livre",
				Address:     "Osasco, SP",
				Telephone:   "11 99999999",
			},
			{
				Id:          2,
				Cid:         1234,
				CompanyName: "Mercado Pago",
				Address:     "Salvador, BA",
				Telephone:   "11 88888888",
			},
		}

		repo.On("GetById", int64(2)).Return(expectedListSeller[1], nil)
		service := seller.NewService(repo)

		result, err := service.GetById(int64(2))

		assert.Nil(t, err)
		assert.Equal(t, expectedListSeller[1], result)
	})
}

func Test_Service_Update(t *testing.T) {
	repo := mocks.NewRepository(t)

	expectedSeller := seller.Seller{
		Id:          1,
		Cid:         123,
		CompanyName: "Mercado Livre",
		Address:     "Osasco, SP",
		Telephone:   "11 99999999",
	}

	t.Run("update_ok: Se os campos forem atualizados com sucesso retornará a informação do elemento atualizado", func(t *testing.T) {
		repo.On("Update", int64(1), "Salvador, BA", "11 98989898").Return(expectedSeller, nil)
		service := seller.NewService(repo)

		result, err := service.Update(int64(1), "Salvador, BA", "11 98989898")

		assert.Equal(t, result, expectedSeller)
		assert.Nil(t, err)

	})

	t.Run("update_non_existent: Se o elemento a ser atualizado não existir, retornar nil", func(t *testing.T) {
		repo.On("Update", int64(3), "Salvador, BA", "11 98989898").Return(seller.Seller{}, fmt.Errorf("Seller not found"))
		service := seller.NewService(repo)

		result, err := service.Update(int64(3), "Salvador, BA", "11 98989898")

		assert.Equal(t, seller.Seller{}, result)
		assert.NotNil(t, err)

	})

}

func Test_Service_Delete(t *testing.T) {
	repo := mocks.NewRepository(t)

	t.Run("delete_ok: Se a exclusão for bem sucedida, o item não aparecerá na lista", func(t *testing.T) {
		repo.On("Delete", int64(1)).Return(nil)
		service := seller.NewService(repo)

		err := service.Delete(int64(1))

		assert.Nil(t, err)

	})

	t.Run("delete_non_existent: Se o elemento a ser removido não existir, retornar nil", func(t *testing.T) {
		repo.On("Delete", int64(3)).Return(fmt.Errorf("Seller not found."))
		service := seller.NewService(repo)

		err := service.Delete(int64(3))

		assert.NotNil(t, err)

	})

}
