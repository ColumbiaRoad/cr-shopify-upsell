package db

import "context"

func (db *Database) PersistUser(ctx context.Context, merchant string) (int64, error) {
	return 123, nil
}
func (db *Database) CheckMerchantByShop(ctx context.Context, shopURL string) (bool, error) {
	return true, nil
}
