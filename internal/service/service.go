package service

import (
	"context"

	"github.com/google/uuid"

	"subscriptions-service/internal/domain"
	"subscriptions-service/internal/dto"
	"subscriptions-service/internal/ports"
)

type service struct {
	repo ports.Repository
}

func NewService(repo ports.Repository) ports.Service {
	return &service{
		repo: repo,
	}
}

func (s *service) Create(ctx context.Context, req dto.CreateSubscriptionReq) (*domain.Subscription, error) {
	return nil, nil
}

func (s *service) GetByID(ctx context.Context, id uuid.UUID) (*domain.Subscription, error) {
	return nil, nil
}

func (s *service) List(ctx context.Context, filter domain.SubscriptionFilter) ([]domain.Subscription, error) {
	return nil, nil
}

func (s *service) Update(ctx context.Context, id uuid.UUID, req dto.UpdateSubscriptionReq) (*domain.Subscription, error) {
	return nil, nil
}

func (s *service) Delete(ctx context.Context, id uuid.UUID) error {
	return nil
}

func (s *service) TotalCost(ctx context.Context, filter domain.TotalCostFilter) (int64, error) {
	return 0, nil
}
