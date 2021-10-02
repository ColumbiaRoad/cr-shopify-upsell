package merchant

import "context"

// Storage is the interface to interact with the database, the methods are implemented in db/
type Storage interface {
	PersistToken(ctx context.Context, shop_url, accessToken string) (int64, error)
	UpdateToken(ctx context.Context, merchantID int64, accessToken string) (int64, error)
	CheckMerchantByShopURL(ctx context.Context, shopURL string) (int64, error)
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
	HandleInstall(ctx context.Context, shop_url, accessToken string) (int64, error)
}

func (m *merchant) HandleInstall(ctx context.Context, shop_url, accessToken string) (int64, error) {
	existsID, err := m.storage.CheckMerchantByShopURL(ctx, shop_url)
	if err != nil {
		return 0, err
	}
	if existsID > 0 {
		return m.storage.UpdateToken(ctx, existsID, accessToken)
	}
	merchantID, err := m.storage.PersistToken(ctx, shop_url, accessToken)
	return merchantID, err
}
