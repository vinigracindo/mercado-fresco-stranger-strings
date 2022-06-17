package employees

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

func (s service) GetById(id int64) (*Employee, error) {
	employee, err := s.repo.GetById(id)

	if err != nil {
		return nil, err
	}

	return employee, nil
}

func (s service) UpdateFullname(id int64, firstName string, lastName string) (*Employee, error) {
	employee, err := s.repo.UpdateFullname(id, firstName, lastName)
	if err != nil {
		return nil, err
	}
	return employee, nil
}

func (s service) Create(cardNumberId string, firstName string, lastName string, warehouseId int64) (Employee, error) {
	employee, err := s.repo.Create(cardNumberId, firstName, lastName, warehouseId)

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
