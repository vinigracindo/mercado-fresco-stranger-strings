package domain

import "context"

type Seller struct {
	Id          int64  `json:"id"`
	Cid         int64  `json:"cid"`
	CompanyName string `json:"company_name"`
	Address     string `json:"address"`
	Telephone   string `json:"telephone"`
}

type ServiceSeller interface {
	GetAll(ctx context.Context) (*[]Seller, error)
	GetById(ctx context.Context, id int64) (*Seller, error)
	Create(ctx context.Context, seller *Seller) (*Seller, error)
	Update(ctx context.Context, id int64, adress, telephone string) (*Seller, error)
	Delete(ctx context.Context, id int64) error
}

type RepositorySeller interface {
	GetAll(ctx context.Context) (*[]Seller, error)
	GetById(ctx context.Context, id int64) (*Seller, error)
	Create(ctx context.Context, seller *Seller) (*Seller, error)
	Update(ctx context.Context, seller *Seller) (*Seller, error)
	Delete(ctx context.Context, id int64) error
}
