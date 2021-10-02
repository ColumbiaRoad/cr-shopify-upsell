package db

import (
	"context"
)

func (db *Database) PersistToken(ctx context.Context, shop_url, access_token string) (merchantID int64, err error) {
	conn, err := db.Conn(ctx)
	if err != nil {
		return merchantID, err
	}
	defer conn.Release()
	row := conn.QueryRow(ctx,
		`INSERT INTO merchants(created_at, updated_at, shop_url, access_token) 
				VALUES (current_timestamp, current_timestamp, $1, $2)
			RETURNING id`, shop_url, access_token)

	err = row.Scan(&merchantID)
	if err != nil {
		return merchantID, err
	}
	return merchantID, err
}
func (db *Database) CheckMerchantByShop(ctx context.Context, shopURL string) (bool, error) {
	return true, nil
}
