package domain

import "context"

type CarryModel struct {
	Id          int64  `json:"id"`
	Cid         int64  `json:"cid" binding:"required"`
	CompanyName string `json:"company_name" binding:"required"`
	Address     string `json:"address" binding:"required"`
	Telephone   string `json:"telephone" binding:"required"`
	LocalityID  int64  `json:"locality_id" binding:"required"`
}

type CarryRepository interface {
	Create(ctx context.Context, carry *CarryModel) (*CarryModel, error)
	GetById(ctx context.Context, id int64) (*CarryModel, error)
	CountLocality(ctx context.Context, locality_id int64) (int64, error)
	// GetAll(ctx context.Context) (*[]CarryModel, error)
	// Delete(ctx context.Context, id int64) error
	// Update(ctx context.Context, id int64, wh *CarryModel) (*CarryModel, error)
}

type CarryService interface {
	GetById(ctx context.Context, id int64) (*CarryModel, error)
	Create(ctx context.Context, carry *CarryModel) (*CarryModel, error)
	// GetAll(ctx context.Context) (*[]CarryModel, error)
	// Delete(ctx context.Context, id int64) error
	// Update(ctx context.Context, carry *CarryModel) (*CarryModel, error)
}
