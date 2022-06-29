package domain

import "context"

//passar o context dentro dos metodos

type Buyer struct {
	Id           int64  `json:"id"`
	CardNumberId string `json:"card_number_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
}

type BuyerRepository interface {
	Create(ctx context.Context, cardNumberId, firstName, lastName string) (*Buyer, error)
	GetAll(ctx context.Context) (*[]Buyer, error)
	GetId(ctx context.Context, id int64) (*Buyer, error)
	Update(ctx context.Context, id int64, cardNumberId, lastName string) (*Buyer, error)
	Delete(ctx context.Context, id int64) error
}

type BuyerService interface {
	Create(ctx context.Context, cardNumberId, firstName string, lastName string) (*Buyer, error)
	GetAll(ctx context.Context) (*[]Buyer, error)
	GetId(ctx context.Context, id int64) (*Buyer, error)
	Update(ctx context.Context, id int64, cardNumberId, lastName string) (*Buyer, error)
	Delete(ctx context.Context, id int64) error
}
