package mariadb_test

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/product/domain"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/product/repository/mariadb"
	"testing"
)

var expectedProduct = domain.Product{
	Id:                             1,
	ProductCode:                    "PROD01",
	Description:                    "Yogurt",
	Width:                          1.2,
	Height:                         6.4,
	Length:                         4.5,
	NetWeight:                      3.4,
	ExpirationRate:                 1.5,
	RecommendedFreezingTemperature: 1.3,
	FreezingRate:                   2,
	ProductTypeId:                  2,
	SellerId:                       2,
}

var expectedProductList = []domain.Product{
	{
		Id:                             1,
		ProductCode:                    "PROD01",
		Description:                    "Yogurt",
		Width:                          1.2,
		Height:                         6.4,
		Length:                         4.5,
		NetWeight:                      3.4,
		ExpirationRate:                 1.5,
		RecommendedFreezingTemperature: 1.3,
		FreezingRate:                   2,
		ProductTypeId:                  2,
		SellerId:                       2,
	},
	{
		Id:                             2,
		ProductCode:                    "PROD02",
		Description:                    "Yogurt",
		Width:                          1.2,
		Height:                         6.4,
		Length:                         4.5,
		NetWeight:                      3.4,
		ExpirationRate:                 1.5,
		RecommendedFreezingTemperature: 1.3,
		FreezingRate:                   2,
		ProductTypeId:                  2,
		SellerId:                       2,
	},
}

func TestMariaDBProductRepository_GetAll(t *testing.T) {

	ctx := context.Background()

	t.Run("should return all employees", func(t *testing.T) {

		db, mock, err := sqlmock.New()

		assert.NoError(t, err)
		defer db.Close()

		rows := sqlmock.NewRows([]string{"id", "product_code", "description", "width", "height", "length", "net_weight",
			"expiration_rate", "recommended_freezing_temperature", "freezing_rate", "product_type_id", "seller_id"})

		for _, product := range expectedProductList {
			rows = rows.AddRow(
				product.Id,
				product.ProductCode,
				product.Description,
				product.Width,
				product.Height,
				product.Length,
				product.NetWeight,
				product.ExpirationRate,
				product.RecommendedFreezingTemperature,
				product.FreezingRate,
				product.ProductTypeId,
				product.SellerId)
		}

		productRepository := mariadb.CreateProductRepository(db)

		mock.ExpectQuery(mariadb.SqlGetAll).WillReturnRows(rows)

		result, err := productRepository.GetAll(ctx)

		assert.Nil(t, err)
		assert.Equal(t, expectedProductList, result)
	})
}
