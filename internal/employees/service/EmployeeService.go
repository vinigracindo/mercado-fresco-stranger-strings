package service

import (
	"context"

	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/employees/domain"
)

type service struct {
	repo domain.EmployeeRepository
}

func NewEmployeeService(r domain.EmployeeRepository) domain.EmployeeService {
	return &service{
		repo: r,
	}
}

func (s service) GetAll(ctx context.Context) ([]domain.Employee, error) {
	employees, err := s.repo.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	return employees, nil
}

func (s service) GetById(ctx context.Context, id int64) (*domain.Employee, error) {
	employee, err := s.repo.GetById(ctx, id)

	if err != nil {
		return nil, err
	}

	return employee, nil
}

func (s service) UpdateFullname(ctx context.Context, id int64, firstName string, lastName string) (*domain.Employee, error) {
	employee, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	employee.SetFullname(firstName, lastName)

	err = s.repo.Update(ctx, id, *employee)
	if err != nil {
		return nil, err
	}
	return employee, nil
}

func (s service) Create(ctx context.Context, cardNumberId string, firstName string, lastName string, warehouseId int64) (domain.Employee, error) {
	employee, err := s.repo.Create(ctx, cardNumberId, firstName, lastName, warehouseId)

	if err != nil {
		return domain.Employee{}, err
	}

	return employee, nil
}

func (s service) Delete(ctx context.Context, id int64) error {
	err := s.repo.Delete(ctx, id)

	if err != nil {
		return err
	}

	return nil
}

func (s service) GetAllReportInboundOrders(ctx context.Context) ([]domain.EmployeeInboundOrdersReport, error) {
	result, err := s.repo.GetAllReportInboundOrders(ctx)

	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s service) GetReportInboundOrdersById(ctx context.Context, employeeID int64) (domain.EmployeeInboundOrdersReport, error) {
	_, err := s.repo.GetById(ctx, employeeID)
	if err != nil {
		return domain.EmployeeInboundOrdersReport{}, domain.ErrEmployeeNotFound
	}

	result, err := s.repo.GetReportInboundOrdersById(ctx, employeeID)

	if err != nil {
		return domain.EmployeeInboundOrdersReport{}, err
	}
	return result, nil
}
