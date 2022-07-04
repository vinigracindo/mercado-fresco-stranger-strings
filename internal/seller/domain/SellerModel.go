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
	GetAll(ctx context.Context) ([]Seller, error)
	GetById(ctx context.Context, id int64) (Seller, error)
	Create(ctx context.Context, cid int64, companyName, address, telephone string) (Seller, error)
	Update(ctx context.Context, id int64, address, telephone string) (Seller, error)
	Delete(ctx context.Context, id int64) error
}

type RepositorySeller interface {
	GetAll(ctx context.Context) ([]Seller, error)
	GetById(ctx context.Context, id int64) (Seller, error)
	Create(ctx context.Context, cid int64, companyName, address, telephone string) (Seller, error)
	Update(ctx context.Context, id int64, address, telephone string) (Seller, error)
	CreatID() int64
	Delete(ctx context.Context, id int64) error
}
