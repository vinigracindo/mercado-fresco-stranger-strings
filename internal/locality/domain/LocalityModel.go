package domain

import "context"

type LocalityModel struct {
	Id           int64
	LocalityName string
	ProvinceName string
	CountryName  string
}

type ReportCarrie struct {
	LocalityId   int64  `json:"locality_id"`
	LocalityName string `json:"locality_name"`
	CarriesCount int64  `json:"carries_count"`
}

type LocalityService interface {
	ReportCarrie(ctx context.Context, locality_id int64) (*[]ReportCarrie, error)
}
type LocalityRepository interface {
	GetById(ctx context.Context, id int64) (*LocalityModel, error)
	ReportCarrie(ctx context.Context, id int64) (*[]ReportCarrie, error)
}
