package buyer

type Buyer struct {
	Id           int64  `json:"id"`
	CardNumberId int64  `json:"card_number_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
}
