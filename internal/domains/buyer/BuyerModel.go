package buyer

type Buyer struct {
	Id           int64  `json:"id"`
	CardNumberId int64  `json:"cardNumberId"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
}
