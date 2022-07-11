package domain

import "errors"

var (
	ErrIDNotFound  = errors.New("PurchaseOrder id not found")
	ErrInvalidDate = errors.New("the purchase order date can't be before the current date")
)
