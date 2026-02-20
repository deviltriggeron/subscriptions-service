package ports

import (
	"context"
	"subscriptions-service/internal/domain"

	"github.com/google/uuid"
)

type Repository interface {
	Create(ctx context.Context, sub domain.Subscription) error
	GetByID(ctx context.Context, id uuid.UUID) (*domain.Subscription, error)
	List(ctx context.Context, filter domain.SubscriptionFilter) (int64, error)
	Update(ctx context.Context, sub domain.Subscription) error
	Delete(ctx context.Context, id uuid.UUID) error
}
