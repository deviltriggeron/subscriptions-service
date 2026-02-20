-- +goose Up
-- +goose StatementBegin
CREATE TABLE subscriptions (
    subscription_id UUID PRIMARY KEY,
    service_name TEXT NOT NULL,
    price price INTEGER NOT NULL CHECK (price > 0),
    user_id UUID NOT NULL,
    start_date TIMESTAMP NOT NULL DEFAULT now(),
    end_date DATE,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE subscriptions;
-- +goose StatementEnd