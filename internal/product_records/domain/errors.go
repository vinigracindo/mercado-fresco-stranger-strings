package domain

import "errors"

var (
	ErrProductIdNotFound = errors.New("product id not found")

	ErrInvalidDate = errors.New("the product record's date can't be before the current date")
)
