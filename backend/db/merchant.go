package db

import (
	"context"
)

func (db *Database) PersistToken(ctx context.Context, shopUrl, accessToken string) (merchantID int64, err error) {
	conn, err := db.Conn(ctx)
	if err != nil {
		return merchantID, err
	}
	defer conn.Release()
	row := conn.QueryRow(ctx,
		`INSERT INTO merchants(created_at, updated_at, shop_url, access_token) 
				VALUES (current_timestamp, current_timestamp, $1, $2)
			RETURNING id`, shopUrl, accessToken)

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
		return merchantID, err
	}
	defer conn.Release()

	row := conn.QueryRow(ctx,
		`SELECT id FROM merchants WHERE shop_url = $1`, shopURL)
	err = row.Scan(&merchantID)
	if err != nil {
		return merchantID, err
	}
	return merchantID, err
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
