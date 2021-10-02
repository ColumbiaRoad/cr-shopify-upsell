-- +goose Up

-- Loading UUID module IF not already loaded
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
-- Setting up tables IF not already created
CREATE TABLE IF NOT EXISTS "users"
(
    "id"                SERIAL PRIMARY KEY,
    "created_at"        TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    "updated_at"        TIMESTAMP WITH TIME ZONE,
    "uninstalled_at"    TIMESTAMP WITH TIME ZONE,
    "email"             TEXT                     NOT NULL UNIQUE,
    "name"              TEXT                     NOT NULL,
    "shop_url"          TEXT                     NOT NULL UNIQUE,
    "access_token"      TEXT                     NULL,
    "offset_variant_id" NUMERIC                  NULL
)
-- Insert dummy data

-- +goose Down
DROP TABLE users;
