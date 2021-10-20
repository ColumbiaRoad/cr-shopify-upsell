package db

import (
	"context"

	"github.com/ColumbiaRoad/cr-shopify-upsell/backend/app/merchant"
)

func (db *Database) PersistToken(ctx context.Context, shopURL, accessToken string) (merchantID int64, err error) {
	conn, err := db.Conn(ctx)
	if err != nil {
		return merchantID, err
	}
	defer conn.Release()
	row := conn.QueryRow(ctx,
		`INSERT INTO merchants(created_at, updated_at, shop_url, access_token) 
				VALUES (current_timestamp, current_timestamp, $1, $2)
			RETURNING id`, shopURL, accessToken)

	err = row.Scan(&merchantID)
	if err != nil {
		return merchantID, err
	}
	return merchantID, err
}
func (db *Database) CheckMerchantByShopURL(ctx context.Context, shopURL string) (int64, error) {
	var merchantID int64
	conn, err := db.Conn(ctx)
	if err != nil {
		return 0, err
	}
	defer conn.Release()

	row := conn.QueryRow(ctx,
		`SELECT id FROM merchants WHERE shop_url = $1`, shopURL)
	err = row.Scan(&merchantID)
	if err != nil {
		return 0, err
	}
	return merchantID, err
}

func (db *Database) GetProfileByURL(ctx context.Context, shopURL string) (merchant.Profile, error) {
	profile := merchant.Profile{}
	conn, err := db.Conn(ctx)
	if err != nil {
		return profile, err
	}
	defer conn.Release()

	row := conn.QueryRow(ctx,
		`SELECT access_token, shop_url, subscription_id FROM merchants WHERE shop_url = $1`, shopURL)
	err = row.Scan(&profile.AccessToken, &profile.ShopURL, &profile.SubscriptionID)
	return profile, err
}

func (db *Database) UpdateToken(ctx context.Context, merchantID int64, accessToken string) (int64, error) {
	conn, err := db.Conn(ctx)
	if err != nil {
		return merchantID, err
	}
	defer conn.Release()
	row := conn.QueryRow(ctx,
		`UPDATE merchants SET (updated_at, access_token) = (current_timestamp, $2)
			WHERE id = $1
			RETURNING id`, merchantID, accessToken)

	err = row.Scan(&merchantID)
	if err != nil {
		return merchantID, err
	}
	return merchantID, err
}

func (db *Database) AddVariantID(ctx context.Context, shopURL string, variantID int64) (int64, error) {
	conn, err := db.Conn(ctx)
	if err != nil {
		return 0, err
	}
	var merchantID int64
	defer conn.Release()
	row := conn.QueryRow(ctx,
		`UPDATE merchants SET (updated_at, offset_variant_id) = (current_timestamp, $2)
			WHERE shop_url = $1
			RETURNING id`, shopURL, variantID)

	err = row.Scan(&merchantID)
	if err != nil {
		return merchantID, err
	}
	return merchantID, err
}

// GetProductVariantID returns the product variant id for the carbon offset product
func (db *Database) GetProductVariantID(ctx context.Context, shopURL string) (variantID int64, err error) {
	conn, err := db.Conn(ctx)
	if err != nil {
		return variantID, err
	}
	defer conn.Release()

	row := conn.QueryRow(ctx,
		`SELECT offset_variant_id FROM merchants WHERE shop_url = $1`, shopURL)
	err = row.Scan(&variantID)
	if err != nil {
		return variantID, err
	}
	return variantID, err

}

// SaveSubscriptionID persist a subscription id to a merchant profile
func (db *Database) SaveSubscriptionID(ctx context.Context, shopURL string, subscriptionID int64) error {
	conn, err := db.Conn(ctx)
	if err != nil {
		return err
	}
	var merchantID int64
	defer conn.Release()
	row := conn.QueryRow(ctx,
		`UPDATE merchants SET (updated_at, subscription_id) = (current_timestamp, $2)
			WHERE shop_url = $1
			RETURNING id`, shopURL, subscriptionID)

	err = row.Scan(&merchantID)
	return err

}

// Updates the value of should_render, used by the extension to render conditionally
func (db *Database) UpdateShouldRender(ctx context.Context, shopURL string, shouldRender bool) error {
	conn, err := db.Conn(ctx)
	if err != nil {
		return err
	}
	var merchantID int64
	defer conn.Release()
	row := conn.QueryRow(ctx,
		`UPDATE merchants SET (updated_at, should_render) = (current_timestamp, $2)
			WHERE shop_url = $1
			RETURNING id`, shopURL, shouldRender)

	err = row.Scan(&merchantID)
	return err
}
