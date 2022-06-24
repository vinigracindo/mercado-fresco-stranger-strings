package service

import "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/employees/domain"

type service struct {
	repo domain.EmployeeRepository
}

func (s service) GetAll() ([]domain.Employee, error) {
	employees, err := s.repo.GetAll()

	if err != nil {
		return nil, err
	}

	return employees, nil
}

func (s service) GetById(id int64) (*domain.Employee, error) {
	employee, err := s.repo.GetById(id)

	if err != nil {
		return nil, err
	}

	return employee, nil
}

func (s service) UpdateFullname(id int64, firstName string, lastName string) (*domain.Employee, error) {
	employee, err := s.repo.UpdateFullname(id, firstName, lastName)
	if err != nil {
		return nil, err
	}
	return employee, nil
}

func (s service) Create(cardNumberId string, firstName string, lastName string, warehouseId int64) (domain.Employee, error) {
	employee, err := s.repo.Create(cardNumberId, firstName, lastName, warehouseId)

	if err != nil {
		return domain.Employee{}, err
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

func NewEmployeeService(r domain.EmployeeRepository) domain.EmployeeService {
	return &service{
		repo: r,
	}
}
