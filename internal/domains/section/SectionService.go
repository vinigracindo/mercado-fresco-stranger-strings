package section

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) Delete(id int64) error {
	return s.repository.Delete(id)
}

func (s *service) UpdateCurrentCapacity(id int64, currentCapacity int64) (Section, error) {
	return s.repository.UpdateCurrentCapacity(id, currentCapacity)
}

func (s *service) CreateSection(sectionNumber int64, currentTemperature int64, minimumTemperature int64, currentCapacity int64, minimumCapacity int64, maximumCapacity int64, warehouseId int64, productTypeId int64) (Section, error) {
	section, err := s.repository.CreateSection(
		sectionNumber,
		currentTemperature,
		minimumTemperature,
		currentCapacity,
		minimumCapacity,
		maximumCapacity,
		warehouseId,
		productTypeId,
	)
	if err != nil {
		return Section{}, err
	}
	return section, nil
}

func (s *service) GetById(id int64) (Section, error) {
	section, err := s.repository.GetById(id)
	if err != nil {
		return Section{}, err
	}
	return section, nil
}

func (s *service) GetAll() ([]Section, error) {
	listSection, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return listSection, nil
}
