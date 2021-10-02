-- +goose Up

-- Setting up tables IF not already created
CREATE TABLE IF NOT EXISTS "webhooks" (
    "id"                    SERIAL PRIMARY KEY,
    "created_at"            TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    "updated_at"            TIMESTAMP WITH TIME ZONE,
    "shop_url"              TEXT                     NULL,
    "compensation_quantity" NUMERIC                  NULL,
    "order_id"              TEXT                     NULL,
    "total_price"           TEXT                     NULL,
    "billing_status"        TEXT                     NULL
);

-- +goose Down
drop table webhooks;