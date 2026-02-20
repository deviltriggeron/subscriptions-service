package mapper

import (
	"subscriptions-service/internal/db"
	"subscriptions-service/internal/domain"
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
