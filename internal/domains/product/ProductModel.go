package product

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
	FreezingRate                   int     `json:"freezing_rate"`
	ProductTypeId                  int     `json:"product_type_id"`
	SellerId                       int     `json:"seller_id"`
}

type ProductRepository interface {
	GetAll() ([]Product, error)
	GetById(id int64) (*Product, error)
	Create(productCode string, description string, width float64, height float64, length float64, netWeight float64,
		expirationRate float64, recommendedFreezingTemperature float64, freezingRate int, productTypeId int, sellerId int) (*Product, error)
	UpdateDescription(id int64, description string) (*Product, error)
	Delete(id int64) error
}

type ProductService interface {
	GetAll() ([]Product, error)
	GetById(id int64) (*Product, error)
	Create(productCode string, description string, width float64, height float64, length float64, netWeight float64,
		expirationRate float64, recommendedFreezingTemperature float64, freezingRate int, productTypeId int, sellerId int) (*Product, error)
	UpdateDescription(id int64, description string) (*Product, error)
	Delete(id int64) error
}
