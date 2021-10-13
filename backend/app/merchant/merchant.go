package merchant

import "context"

// Storage is the interface to interact with the database, the methods are implemented in db/
type Storage interface {
	PersistToken(ctx context.Context, shopURL, accessToken string) (int64, error)
	UpdateToken(ctx context.Context, merchantID int64, accessToken string) (int64, error)
	CheckMerchantByShopURL(ctx context.Context, shopURL string) (int64, error)
	GetProfileByURL(ctx context.Context, shopURL string) (Profile, error)
	AddVariantID(ctx context.Context, shopURL string, variantID int64) (int64, error)
	GetProductVariantID(ctx context.Context, shopURL string) (int64, error)
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
	HandleInstall(ctx context.Context, shopURL, accessToken string) (int64, error)
	AddVariantID(ctx context.Context, shopURL string, variantID int64) (int64, error)
	GetVariantIDForShop(ctx context.Context, shopURL string) (variantID int64, err error)
	GetShopByURL(ctx context.Context, shopURL string) (Profile, error)
}

// HandleInstall makes it possible to add or update the merchants shopify access token
func (m *merchant) HandleInstall(ctx context.Context, shopURL, accessToken string) (int64, error) {
	existsID, err := m.storage.CheckMerchantByShopURL(ctx, shopURL)
	if err != nil {
		if err.Error() != "no rows in result set" {
			return 0, err
		}
	}
	if existsID > 0 {
		return m.storage.UpdateToken(ctx, existsID, accessToken)
	}
	merchantID, err := m.storage.PersistToken(ctx, shopURL, accessToken)
	return merchantID, err
}

// AddVariantID as a part of the onboarding we create a product in the merchants store and we need
// to save that id
func (m *merchant) AddVariantID(ctx context.Context, shopURL string, variantID int64) (int64, error) {
	merchantID, err := m.storage.AddVariantID(ctx, shopURL, variantID)
	return merchantID, err
}

// GetVariantIDForShop returns the product variant id for the carbon offset product
func (m *merchant) GetVariantIDForShop(ctx context.Context, shopURL string) (variantID int64, err error) {
	return m.storage.GetProductVariantID(ctx, shopURL)
}

// GetShopByURL returns a Profile for the shop if found
func (m *merchant) GetShopByURL(ctx context.Context, shopURL string) (Profile, error) {
	return m.storage.GetProfileByURL(ctx, shopURL)
}
