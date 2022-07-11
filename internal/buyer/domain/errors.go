package domain

import "errors"

var (
	ErrBuyerNotFound = errors.New("buyer not found")
	ErrIDNotFound    = errors.New("buyer id not found")
)
