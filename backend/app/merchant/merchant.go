package merchant

import "context"

// Storage is the interface to interact with the database, the methods are implemented in db/
type Storage interface {
	PersistToken(ctx context.Context, shop_url, access_token string) (int64, error)
	CheckMerchantByShop(ctx context.Context, shopURL string) (bool, error)
}

type merchant struct {
	storage Storage
}

// New is a new User
func New(storage Storage) Merchants {
	return &merchant{
		storage: storage,
	}
}

// Merchants represents the interface to "Users"
type Merchants interface {
	HandleInstall(ctx context.Context, shop_url, access_token string) (int64, error)
}

func (m *merchant) HandleInstall(ctx context.Context, shop_url, access_token string) (int64, error) {
	merchantID, err := m.storage.PersistToken(ctx, shop_url, access_token)
	return merchantID, err
}
