package mapper

import "errors"

var (
	ErrInvalidUUID      = errors.New("invalid UUID format")
	ErrInvalidStartDate = errors.New("invalid start_date format, expected MM-YYYY")
	ErrInvalidEndDate   = errors.New("invalid end_date format, expected MM-YYYY")
)
