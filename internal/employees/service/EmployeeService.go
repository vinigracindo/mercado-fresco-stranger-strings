package service

import (
	"context"

	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/employees/domain"
)

type service struct {
	repo domain.EmployeeRepository
}

func (s service) GetAll() ([]domain.Employee, error) {
	employees, err := s.repo.GetAll(context.Background())

	if err != nil {
		return nil, err
	}

	return employees, nil
}

func (s service) GetById(id int64) (*domain.Employee, error) {
	employee, err := s.repo.GetById(context.Background(), id)

	if err != nil {
		return nil, err
	}

	return employee, nil
}

func (s service) UpdateFullname(id int64, firstName string, lastName string) (*domain.Employee, error) {
	employee, err := s.repo.GetById(context.Background(), id)
	if err != nil {
		return nil, err
	}

	employee.SetFullname(firstName, lastName)

	err = s.repo.Update(context.Background(), id, *employee)
	if err != nil {
		return nil, err
	}
	return employee, nil
}

func (s service) Create(cardNumberId string, firstName string, lastName string, warehouseId int64) (domain.Employee, error) {
	employee, err := s.repo.Create(context.Background(), cardNumberId, firstName, lastName, warehouseId)

	if err != nil {
		return domain.Employee{}, err
	}

	return employee, nil
}

func (s service) Delete(id int64) error {
	err := s.repo.Delete(context.Background(), id)

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
