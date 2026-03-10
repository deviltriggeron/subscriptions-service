package mapper

import (
	"time"

	"github.com/google/uuid"

	"subscriptions-service/internal/db"
	"subscriptions-service/internal/domain"
	"subscriptions-service/internal/dto"
)

func DomainToModel(sub domain.Subscription) db.SubscriptionModel {
	return db.SubscriptionModel{
		ID:          sub.ID,
		ServiceName: sub.ServiceName,
		Price:       sub.Price,
		UserID:      sub.UserID,
		StartDate:   sub.StartDate,
		EndDate:     sub.EndDate,
	}
}

func ModelToDomain(sub db.SubscriptionModel) domain.Subscription {
	return domain.Subscription{
		ID:          sub.ID,
		ServiceName: sub.ServiceName,
		Price:       sub.Price,
		UserID:      sub.UserID,
		StartDate:   sub.StartDate,
		EndDate:     sub.EndDate,
	}
}

func ReqToDomain(req dto.CreateSubscriptionReq) (*domain.Subscription, error) {
	userID, err := uuid.Parse(req.UserID)
	if err != nil {
		return nil, ErrInvalidUUID
	}

	startDate, err := time.Parse("01-2006", req.StartDate)
	if err != nil {
		return nil, ErrInvalidStartDate
	}

	return &domain.Subscription{
		ServiceName: req.ServiceName,
		Price:       req.Price,
		UserID:      userID,
		StartDate:   startDate,
	}, nil
}

func ReqUpdateToDomain(req dto.UpdateSubscriptionReq) (*domain.Subscription, error) {
	var sub domain.Subscription
	if req.ServiceName != nil {
		sub.ServiceName = *req.ServiceName
	}

	if req.Price != nil {
		sub.Price = *req.Price
	}

	if req.StartDate != nil {
		t, err := time.Parse("01-2006", *req.StartDate)
		if err != nil {
			return nil, ErrInvalidStartDate
		}
		sub.StartDate = t
	}

	if req.EndDate != nil {
		t, err := time.Parse("01-2006", *req.EndDate)
		if err != nil {
			return nil, ErrInvalidEndDate
		}
		sub.EndDate = t
	}

	return &sub, nil
}

func ReqFilterToDomain(req dto.CreateSubscriptionFilterReq) (*domain.SubscriptionFilter, error) {
	var filter domain.SubscriptionFilter

	if req.ServiceName != nil {
		filter.ServiceName = req.ServiceName
	}

	if req.UserID != nil {
		userID, err := uuid.Parse(*req.UserID)
		if err != nil {
			return nil, ErrInvalidUUID
		}

		filter.UserID = &userID
	}

	if req.From != nil {
		t, err := time.Parse("01-2006", *req.From)
		if err != nil {
			return nil, ErrInvalidStartDate
		}

		filter.From = &t
	}

	if req.To != nil {
		t, err := time.Parse("01-2006", *req.To)
		if err != nil {
			return nil, ErrInvalidEndDate
		}

		filter.To = &t
	}

	return &filter, nil
}

func DomainToResp(sub domain.Subscription) dto.SubscriptionResp {
	return dto.SubscriptionResp{
		ServiceName: sub.ServiceName,
		Price:       sub.Price,
		UserID:      sub.UserID.String(),
		StartDate:   sub.StartDate.String(),
		EndDate:     sub.EndDate.String(),
	}
}
