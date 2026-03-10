package domain

import "errors"

var (
	ErrInvalidServiceName   = errors.New("service_name cannot be empty")
	ErrInvalidPrice         = errors.New("price cannot be negative")
	ErrSubscriptionNotFound = errors.New("subscriber not found")
)
