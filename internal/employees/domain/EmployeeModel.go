package domain

import "context"

type Employee struct {
	Id           int64  `json:"id"`
	CardNumberId string `json:"card_number_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	WarehouseId  int64  `json:"warehouse_id"`
}

func (e *Employee) SetFullname(firstName string, lastName string) {
	e.FirstName = firstName
	e.LastName = lastName
}

type EmployeeInboundOrdersReport struct {
	Employee
	Count int64 `json:"inbound_orders_count"`
}

type EmployeeService interface {
	GetAll(ctx context.Context) ([]Employee, error)
	GetById(ctx context.Context, id int64) (*Employee, error)
	Create(ctx context.Context, cardNumberId string, firstName string, lastName string, warehouseId int64) (Employee, error)
	UpdateFullname(ctx context.Context, id int64, firstName string, lastName string) (*Employee, error)
	Delete(ctx context.Context, id int64) error
	GetAllReportInboundOrders(ctx context.Context) ([]EmployeeInboundOrdersReport, error)
	GetReportInboundOrdersById(ctx context.Context, employeeID int64) (EmployeeInboundOrdersReport, error)
}

type EmployeeRepository interface {
	GetAll(ctx context.Context) ([]Employee, error)
	GetById(ctx context.Context, id int64) (*Employee, error)
	Create(ctx context.Context, cardNumberId string, firstName string, lastName string, warehouseId int64) (Employee, error)
	Update(ctx context.Context, employeeID int64, updatedEmployee Employee) error
	Delete(ctx context.Context, id int64) error
	GetAllReportInboundOrders(ctx context.Context) ([]EmployeeInboundOrdersReport, error)
	GetReportInboundOrdersById(ctx context.Context, employeeID int64) (EmployeeInboundOrdersReport, error)
}
