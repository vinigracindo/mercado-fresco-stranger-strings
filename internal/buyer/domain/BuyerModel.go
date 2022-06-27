package domain

type Buyer struct {
	Id           int64  `json:"id"`
	CardNumberId string `json:"card_number_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
}

type BuyerRepository interface {
	Create(cardNumberId, firstName, lastName string) (*Buyer, error)
	GetAll() ([]Buyer, error)
	GetId(id int64) (*Buyer, error)
	Update(id int64, cardNumberId, lastName string) (*Buyer, error)
	Delete(id int64) error
	CreateId() int64
}

type BuyerService interface {
	Create(cardNumberId, firstName string, lastName string) (*Buyer, error)
	GetAll() ([]Buyer, error)
	GetId(id int64) (*Buyer, error)
	Update(id int64, cardNumberId, lastName string) (*Buyer, error)
	Delete(id int64) error
}
