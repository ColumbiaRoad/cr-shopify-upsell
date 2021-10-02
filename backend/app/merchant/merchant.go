package merchant

import "context"

// Storage is the interface to interact with the database, the methods are implemented in db/
type Storage interface {
	PersistToken(ctx context.Context, shopUrl, accessToken string) (int64, error)
	UpdateToken(ctx context.Context, merchantID int64, accessToken string) (int64, error)
	CheckMerchantByShopURL(ctx context.Context, shopURL string) (int64, error)
	AddVariantID(ctx context.Context, shopURL string, variantID int64) (int64, error)
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
	HandleInstall(ctx context.Context, shopUrl, accessToken string) (int64, error)
	AddVariantID(ctx context.Context, shopUrl string, variantID int64) (int64, error)
}

// HandleInstall makes it possible to add or update the merchants shopify access token
func (m *merchant) HandleInstall(ctx context.Context, shopUrl, accessToken string) (int64, error) {
	existsID, err := m.storage.CheckMerchantByShopURL(ctx, shopUrl)
	if err != nil {
		return 0, err
	}
	if existsID > 0 {
		return m.storage.UpdateToken(ctx, existsID, accessToken)
	}
	merchantID, err := m.storage.PersistToken(ctx, shopUrl, accessToken)
	return merchantID, err
}

// AddVariantID as a part of the onboarding we create a product in the merchants store and we need
// to save that id
func (m *merchant) AddVariantID(ctx context.Context, shopUrl string, variantID int64) (int64, error) {
	merchantID, err := m.storage.AddVariantID(ctx, shopUrl, variantID)
	return merchantID, err
}
