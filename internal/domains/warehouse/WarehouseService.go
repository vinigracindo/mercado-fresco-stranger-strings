package warehouse

type service struct {
	repository Repository
}
type Service interface {
	Create(wr *WarehouseModel) (WarehouseModel, error)
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s service) Create(new *WarehouseModel) (WarehouseModel, error) {
	wh, err := s.repository.Store(new)

	if err != nil {
		return WarehouseModel{}, err
	}

	return wh, nil
}
