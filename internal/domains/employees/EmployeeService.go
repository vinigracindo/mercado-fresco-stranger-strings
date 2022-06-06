package employees

type Service interface {
	GetAll() ([]Employee, error)
	Get(id int64) (Employee, error)
	Store(cardNumberId string, firstName string, lastName string, warehouseId int64) (Employee, error)
	Update(id int64, cardNumberId string, firstName string, lastName string, warehouseId int64) (Employee, error)
	Delete(id int64) error
}

type service struct {
	repo Repository
}

func (s service) GetAll() ([]Employee, error) {
	employees, err := s.repo.GetAll()

	if err != nil {
		return nil, err
	}

	return employees, nil
}

func (s service) Get(id int64) (Employee, error) {
	employee, err := s.repo.Get(id)

	if err != nil {
		return Employee{}, err
	}

	return employee, nil
}

func (s service) Update(id int64, cardNumberId string, firstName string, lastName string, warehouseId int64) (Employee, error) {
	employee, err := s.repo.Update(id, cardNumberId, firstName, lastName, warehouseId)
	if err != nil {
		return Employee{}, err
	}
	return employee, nil
}

func (s service) Store(cardNumberId string, firstName string, lastName string, warehouseId int64) (Employee, error) {
	employee, err := s.repo.Store(cardNumberId, firstName, lastName, warehouseId)

	if err != nil {
		return Employee{}, err
	}

	return employee, nil
}

func (s service) Delete(id int64) error {
	err := s.repo.Delete(id)

	if err != nil {
		return err
	}

	return nil
}

func NewService(r Repository) Service {
	return &service{
		repo: r,
	}
}
