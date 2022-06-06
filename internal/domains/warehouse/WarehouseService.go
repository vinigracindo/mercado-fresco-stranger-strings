package warehouse

type service struct {
	repository Repository
}
type Service interface {
	GetAll() ([]WarehouseModel, error)
	GetById(id int64) (WarehouseModel, error)
	Delete(id int64) error
	UpdateTempAndCap(id int64, adress, tel, code string, mintemp float64, mincap int64) (WarehouseModel, error)
	Create(adress, tel, code string, mintemp float64, mincap int64) (WarehouseModel, error)
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s service) Create(adress, tel, code string, mintemp float64, mincap int64) (WarehouseModel, error) {
	new := WarehouseModel{
		Address:            adress,
		Telephone:          tel,
		WarehouseCode:      code,
		MinimunCapacity:    mincap,
		MinimunTemperature: mintemp,
	}

	wh, err := s.repository.Store(&new)

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

func (s service) UpdateTempAndCap(id int64, adress, tel, code string, mintemp float64, mincap int64) (WarehouseModel, error) {
	wh := WarehouseModel{
		Address:            adress,
		Telephone:          tel,
		WarehouseCode:      code,
		MinimunCapacity:    mincap,
		MinimunTemperature: mintemp,
	}

	parchWh, err := s.repository.Update(id, &wh)

	if err != nil {
		return WarehouseModel{}, err
	}

	return parchWh, nil
}
