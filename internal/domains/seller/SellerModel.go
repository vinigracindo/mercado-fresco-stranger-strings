package seller

type Seller struct {
	Id          float64 `json:"id"`
	Cid         float64 `json:"cid"`
	CompanyName string  `json:"company_name"`
	Address     string  `json:"address"`
	Telephone   string  `json:"telephone"`
}
