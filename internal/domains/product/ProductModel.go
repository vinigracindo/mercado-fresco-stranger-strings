package product

type Product struct {
	Id                             int64   `json:"id"`
	ProductCode                    string  `json:"productCode"`
	Description                    string  `json:"description"`
	Width                          float64 `json:"width"`
	Height                         float64 `json:"height"`
	Length                         float64 `json:"length"`
	NetWeight                      float64 `json:"netWeight"`
	ExpirationRate                 int     `json:"expirationRate"`
	RecommendedFreezingTemperature float64 `json:"recommendedFreezingTemperature"`
	FreezingRate                   int     `json:"freezingRate"`
	ProductTypeId                  int     `json:"productTypeId"`
	SellerId                       int     `json:"sellerId"`
}
