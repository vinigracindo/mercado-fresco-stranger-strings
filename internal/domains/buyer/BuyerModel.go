package buyer

type Buyer struct {
	Id           int64  `json:"id"`
	CardNumberId int64  `json:"card_number_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
}

type Repository interface {
	Create(cardNumberId int64, firstName string, lastName string) (*Buyer, error)
	GetAll() ([]Buyer, error)
	GetId(id int64) (*Buyer, error)
	Update(id int64, cardNumberId int64, lastName string) (*Buyer, error)
	Delete(id int64) error
	CreateId() int64
}

type Service interface {
	Create(cardNumberId int64, firstName string, lastName string) (*Buyer, error)
	GetAll() ([]Buyer, error)
	GetId(id int64) (*Buyer, error)
	Update(id int64, cardNumberId int64, lastName string) (*Buyer, error)
	Delete(id int64) error
}
