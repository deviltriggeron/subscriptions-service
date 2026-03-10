package service

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"

	"subscriptions-service/internal/domain"
	"subscriptions-service/internal/dto"
	"subscriptions-service/internal/mapper"
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
	sub, err := mapper.ReqToDomain(req)
	if err != nil {
		return nil, err
	}

	if strings.TrimSpace(sub.ServiceName) == "" {
		return nil, domain.ErrInvalidServiceName
	}

	if sub.Price < 0 {
		return nil, domain.ErrInvalidPrice
	}

	sub.EndDate = getEndDate(sub.StartDate)

	err = s.repo.Create(ctx, *sub)
	if err != nil {
		return nil, err
	}

	return sub, nil
}

func (s *service) GetByID(ctx context.Context, id uuid.UUID) (*domain.Subscription, error) {
	sub, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if sub == nil {
		return nil, fmt.Errorf("subscription not found")
	}

	return sub, nil
}

func (s *service) List(ctx context.Context, req dto.CreateSubscriptionFilterReq) ([]domain.Subscription, error) {
	filter, err := mapper.ReqFilterToDomain(req)
	if err != nil {
		return nil, err
	}

	return s.repo.List(ctx, *filter)
}

func (s *service) Update(ctx context.Context, id uuid.UUID, req dto.UpdateSubscriptionReq) (*domain.Subscription, error) {
	sub, err := mapper.ReqUpdateToDomain(req)
	if err != nil {
		return nil, err
	}

	err = s.repo.Update(ctx, *sub)
	if err != nil {
		return nil, err
	}

	return sub, nil
}

func (s *service) Delete(ctx context.Context, id uuid.UUID) error {
	err := s.repo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) TotalCost(ctx context.Context, req dto.CreateSubscriptionFilterReq) (int64, error) {
	filter, err := mapper.ReqFilterToDomain(req)
	if err != nil {
		return 0, err
	}

	return s.repo.TotalCost(ctx, *filter)
}

func getEndDate(startDate time.Time) time.Time {
	return startDate.AddDate(0, 1, 0)
}
