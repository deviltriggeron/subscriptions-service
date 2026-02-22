package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"

	"subscriptions-service/internal/db"
	"subscriptions-service/internal/domain"
	"subscriptions-service/internal/mapper"
	"subscriptions-service/internal/ports"
)

type repo struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) ports.Repository {
	return &repo{
		db: db,
	}
}

func (r *repo) Create(ctx context.Context, sub domain.Subscription) error {
	model := mapper.DomainToModel(sub)

	q := `
		INSERT INTO subscriptions (subscription_id, service_name, price, user_id, start_date, end_date)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := r.db.ExecContext(ctx, q, model.ID, model.ServiceName, model.Price, model.UserID, model.StartDate, model.EndDate)
	if err != nil {
		return fmt.Errorf("error insert subscription to DB: %v", err)
	}

	return nil
}
func (r *repo) GetByID(ctx context.Context, id uuid.UUID) (*domain.Subscription, error) {
	var model db.SubscriptionModel
	q := `
		SELECT *
		FROM subscriptions
		WHERE subscription_id = $1;
	`

	err := r.db.QueryRowContext(ctx, q, id).Scan(&model.ID, &model.ServiceName, &model.Price, &model.UserID, &model.StartDate, &model.EndDate)
	if err != nil {
		return nil, fmt.Errorf("error select subscription(%s): %v", id, err)
	}

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	sub := mapper.ModelToDomain(model)

	return &sub, nil
}

func (r *repo) List(ctx context.Context, filter domain.SubscriptionFilter) ([]domain.Subscription, error) {
	query := `
		SELECT
			subscription_id,
			service_name,
			price,
			user_id,
			start_date,
			end_date
		FROM subscriptions
		WHERE start_date < $1
		  AND (end_date IS NULL OR end_date > $2)
	`

	args := []any{filter.To, filter.From}

	if filter.UserID != nil {
		query += fmt.Sprintf(" AND user_id = $%d", len(args)+1)
		args = append(args, *filter.UserID)
	}

	if filter.ServiceName != nil {
		query += fmt.Sprintf(" AND service_name ILIKE $%d", len(args)+1)
		args = append(args, "%"+*filter.ServiceName+"%")
	}

	query += " ORDER BY start_date DESC"

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("list subscriptions: %w", err)
	}
	defer rows.Close()

	var subs []domain.Subscription

	for rows.Next() {
		var sub domain.Subscription
		var endDate sql.NullTime

		if err := rows.Scan(
			&sub.ID,
			&sub.ServiceName,
			&sub.Price,
			&sub.UserID,
			&sub.StartDate,
			&endDate,
		); err != nil {
			return nil, fmt.Errorf("scan subscription: %w", err)
		}

		if endDate.Valid {
			sub.EndDate = endDate.Time
		}

		subs = append(subs, sub)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return subs, nil
}

func (r *repo) TotalCost(ctx context.Context, filter domain.TotalCostFilter) (int64, error) {
	query := `
		SELECT COALESCE(SUM(price), 0)
		FROM subscriptions
		WHERE start_date < $1
		  AND (end_date IS NULL OR end_date > $2)
	`

	args := []any{filter.To, filter.From}

	if filter.UserID != nil {
		query += fmt.Sprintf(" AND user_id = $%d", len(args)+1)
		args = append(args, *filter.UserID)
	}

	if filter.ServiceName != nil {
		query += fmt.Sprintf(" AND service_name ILIKE $%d", len(args)+1)
		args = append(args, "%"+*filter.ServiceName+"%")
	}

	var total int64

	err := r.db.QueryRowContext(ctx, query, args...).Scan(&total)
	if err != nil {
		return 0, fmt.Errorf("total cost: %w", err)
	}

	return total, nil
}

func (r *repo) Update(ctx context.Context, sub domain.Subscription) error {
	q := `
		UPDATE subscriptions
		SET service_name = $1,
			price = $2,
			user_id = $3,
			start_date = $4,
			end_date = $5
		WHERE id = $6
	`

	res, err := r.db.ExecContext(
		ctx, q, sub.ServiceName, sub.Price, sub.UserID, sub.StartDate, sub.EndDate, sub.ID,
	)
	if err != nil {
		return fmt.Errorf("error update subscription: %w", err)
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("subscription not found")
	}

	return nil
}

func (r *repo) Delete(ctx context.Context, id uuid.UUID) error {
	q := `
		DELETE FROM subscriptions
		WHERE subscription_id = $1; 
	`

	res, err := r.db.ExecContext(ctx, q, id)
	if err != nil {
		return fmt.Errorf("error delete subscription: %v", err)
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("subscription %s not found", id)
	}

	return nil
}
