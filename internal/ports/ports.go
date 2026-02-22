package ports

import (
	"context"

	"github.com/google/uuid"

	"subscriptions-service/internal/domain"
	"subscriptions-service/internal/dto"
)

type Repository interface {
	Create(ctx context.Context, sub domain.Subscription) error
	GetByID(ctx context.Context, id uuid.UUID) (*domain.Subscription, error)
	List(ctx context.Context, filter domain.SubscriptionFilter) ([]domain.Subscription, error)
	TotalCost(ctx context.Context, filter domain.TotalCostFilter) (int64, error)
	Update(ctx context.Context, sub domain.Subscription) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type Service interface {
	Create(ctx context.Context, req dto.CreateSubscriptionReq) (*domain.Subscription, error)
	GetByID(ctx context.Context, id uuid.UUID) (*domain.Subscription, error)
	List(ctx context.Context, filter domain.SubscriptionFilter) ([]domain.Subscription, error)
	Update(ctx context.Context, id uuid.UUID, req dto.UpdateSubscriptionReq) (*domain.Subscription, error)
	Delete(ctx context.Context, id uuid.UUID) error
	TotalCost(ctx context.Context, filter domain.TotalCostFilter) (int64, error)
}
