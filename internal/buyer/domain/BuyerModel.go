package domain

import "context"

type Buyer struct {
	Id           int64  `json:"id"`
	CardNumberId string `json:"card_number_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
}

type PurchaseOrdersReport struct {
	Id                 int64  `json:"id"`
	CardNumberId       string `json:"card_number_id"`
	FirstName          string `json:"first_name"`
	LastName           string `json:"last_name"`
	CountBuyersRecords int64  `json:"purchase_orders_count"`
}

type BuyerRepository interface {
	Create(ctx context.Context, cardNumberId, firstName, lastName string) (*Buyer, error)
	GetAll(ctx context.Context) (*[]Buyer, error)
	GetId(ctx context.Context, id int64) (*Buyer, error)
	Update(ctx context.Context, id int64, cardNumberId, lastName string) (*Buyer, error)
	Delete(ctx context.Context, id int64) error
	GetAllPurchaseOrdersReports(ctx context.Context) (*[]PurchaseOrdersReport, error)
}

type BuyerService interface {
	Create(ctx context.Context, cardNumberId, firstName string, lastName string) (*Buyer, error)
	GetAll(ctx context.Context) (*[]Buyer, error)
	GetId(ctx context.Context, id int64) (*Buyer, error)
	Update(ctx context.Context, id int64, cardNumberId, lastName string) (*Buyer, error)
	Delete(ctx context.Context, id int64) error
	GetPurchaseOrdersReports(ctx context.Context, id int64) (*[]PurchaseOrdersReport, error)
	GetAllPurchaseOrdersReports(ctx context.Context) (*[]PurchaseOrdersReport, error)
}
