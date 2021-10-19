-- +goose Up

-- Setting up tables IF not already created
ALTER TABLE "merchants" ADD COLUMN IF NOT EXISTS
    "subscription_id"   BIGINT DEFAULT 0;
ALTER TABLE "merchants" ADD COLUMN IF NOT EXISTS
    "accepted_terms_at" TIMESTAMP WITH TIME ZONE;
ALTER TABLE "merchants" ADD COLUMN IF NOT EXISTS
    "should_render" BOOLEAN DEFAULT FALSE;


-- +goose Down
ALTER TABLE "merchants" DROP COLUMN "subscription_id";
ALTER TABLE "merchants" DROP COLUMN "accepted_terms_at";
ALTER TABLE "merchants" DROP COLUMN "should_render";
