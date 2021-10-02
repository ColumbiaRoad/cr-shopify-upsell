-- +goose Up

-- Loading UUID module IF not already loaded
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
-- Setting up tables IF not already created
CREATE TABLE IF NOT EXISTS "merchants" (
    "id"                SERIAL PRIMARY KEY,
    "created_at"        TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    "updated_at"        TIMESTAMP WITH TIME ZONE,
    "uninstalled_at"    TIMESTAMP WITH TIME ZONE,
    "email"             TEXT                     NULL,
    "name"              TEXT                     NULL,
    "shop_url"          TEXT                     NOT NULL UNIQUE,
    "access_token"      TEXT                     NULL,
    "offset_variant_id" NUMERIC                  NULL
);

-- +goose Down
drop table merchants;