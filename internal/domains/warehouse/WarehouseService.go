package warehouse

type service struct {
	repository Repository
}
type Service interface {
	GetAll() ([]WarehouseModel, error)
	GetById(id int64) (WarehouseModel, error)
	Delete(id int64) error
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

func (s service) GetAll() ([]WarehouseModel, error) {
	swh, err := s.repository.GetAll()

	if err != nil {
		return []WarehouseModel{}, err
	}

	return swh, nil
}

func (s service) GetById(id int64) (WarehouseModel, error) {
	hw, err := s.repository.GetById(id)

	if err != nil {
		return WarehouseModel{}, err
	}

	return hw, nil
}

func (s service) Delete(id int64) error {
	err := s.repository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
