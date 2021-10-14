-- +goose Up

-- Setting up tables IF not already created
ALTER TABLE "merchants" ADD COLUMN IF NOT EXISTS
    "subscription_id"   TEXT NULL;
ALTER TABLE "merchants" ADD COLUMN IF NOT EXISTS
    "accepted_terms_at" TIMESTAMP WITH TIME ZONE;


-- +goose Down
ALTER TABLE "merchants" DROP COLUMN "subscription_id";
ALTER TABLE "merchants" DROP COLUMN "accepted_terms_at";
