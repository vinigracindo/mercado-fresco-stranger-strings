package domain

import "context"

type LocalityModel struct {
	Id           int64
	LocalityName string
	ProvinceName string
	CountryName  string
}

type LocalityService interface {
	ReportCarrie(ctx context.Context, locality_id int64) (any, error)
}
type LocalityRepository interface {
	GetById(ctx context.Context, id int64) (*LocalityModel, error)
}
