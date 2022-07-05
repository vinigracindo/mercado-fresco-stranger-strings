package domain

import "context"

type Product struct {
	Id                             int64   `json:"id"`
	ProductCode                    string  `json:"product_code"`
	Description                    string  `json:"description"`
	Width                          float64 `json:"width"`
	Height                         float64 `json:"height"`
	Length                         float64 `json:"length"`
	NetWeight                      float64 `json:"net_weight"`
	ExpirationRate                 float64 `json:"expiration_rate"`
	RecommendedFreezingTemperature float64 `json:"recommended_freezing_temperature"`
	FreezingRate                   float64 `json:"freezing_rate"`
	ProductTypeId                  int64   `json:"product_type_id"`
	SellerId                       int64   `json:"seller_id"`
}

type ProductRepository interface {
	GetAll(ctx context.Context) (*[]Product, error)
	GetById(ctx context.Context, id int64) (*Product, error)
	Create(ctx context.Context, product *Product) (*Product, error)
	UpdateDescription(ctx context.Context, product *Product) (*Product, error)
	Delete(ctx context.Context, id int64) error
}

type ProductService interface {
	GetAll(ctx context.Context) (*[]Product, error)
	GetById(ctx context.Context, id int64) (*Product, error)
	Create(ctx context.Context, product *Product) (*Product, error)
	UpdateDescription(ctx context.Context, id int64, description string) (*Product, error)
	Delete(ctx context.Context, id int64) error
}
