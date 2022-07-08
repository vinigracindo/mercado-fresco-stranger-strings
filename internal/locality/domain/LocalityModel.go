package domain

import "context"

type LocalityModel struct {
	Id           int64  `json:"id"`
	LocalityName string `json:"locality_name"`
	ProvinceName string `json:"province_name"`
	CountryName  string `json:"country_name"`
	ProvinceId   int64  `json:"province_id"`
}

type ReportSeller struct {
	LocalityId   int64  `json:"locality_id"`
	LocalityName string `json:"locality_name"`
	SellerCount  int64  `json:"seller_count"`
}

type ReportCarrie struct {
	LocalityId   int64  `json:"locality_id"`
	LocalityName string `json:"locality_name"`
	CarriesCount int64  `json:"carries_count"`
}

type LocalityService interface {
	ReportCarrie(ctx context.Context, locality_id int64) (*[]ReportCarrie, error)
	Create(ctx context.Context, locality *LocalityModel) (*LocalityModel, error)
	ReportSeller(ctx context.Context, id int64) (*[]ReportSeller, error)
}

type LocalityRepository interface {
	GetById(ctx context.Context, id int64) (*LocalityModel, error)
	ReportCarrie(ctx context.Context, id int64) (*[]ReportCarrie, error)
	Create(ctx context.Context, locality *LocalityModel) (*LocalityModel, error)
	ReportSeller(ctx context.Context, id int64) (*[]ReportSeller, error)
}
