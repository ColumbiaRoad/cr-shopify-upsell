package api

import (
	"context"
)

type fixture struct {
	api      *Server
	merchant *testMerchant
}

type testMerchant struct {
	handleInstall       func(ctx context.Context, shopUrl, accessToken string) (int64, error)
	addVariantID        func(ctx context.Context, shopUrl string, variantID int64) (int64, error)
	getVariantIDForShop func(ctx context.Context, shopURL string) (variantID int64, err error)
}

func setTestFixture() *fixture {
	srv := New("", "", "")
	srv.Routes()
	merchant := &testMerchant{}
	srv.Merchant = merchant
	return &fixture{
		api:      srv,
		merchant: merchant,
	}
}

func (m *testMerchant) HandleInstall(ctx context.Context, shopUrl, accessToken string) (int64, error) {
	return m.handleInstall(ctx, shopUrl, accessToken)
}

func (m *testMerchant) AddVariantID(ctx context.Context, shopUrl string, variantID int64) (int64, error) {
	return m.addVariantID(ctx, shopUrl, variantID)
}
func (m *testMerchant) GetVariantIDForShop(ctx context.Context, shopURL string) (variantID int64, err error) {
	return m.getVariantIDForShop(ctx, shopURL)
}
