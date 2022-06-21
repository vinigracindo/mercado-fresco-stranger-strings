package employees

import "errors"

var (
	ErrCardNumberMustBeUnique = errors.New("card number must be unique")
	ErrEmployeeNotFound       = errors.New("employee not found")
)
