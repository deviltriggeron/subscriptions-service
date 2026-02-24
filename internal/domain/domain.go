package domain

import (
	"time"

	"github.com/google/uuid"
)

type Subscription struct {
	ID          uuid.UUID
	Servicename string
	Price       int
	UserID      uuid.UUID
	StartDate   time.Time
	EndDate     time.Time
}
