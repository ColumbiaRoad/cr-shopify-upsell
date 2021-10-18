package api

import (
	"context"
	"io"
	"text/template"

	"github.com/ColumbiaRoad/cr-shopify-upsell/backend/app/merchant"

	"github.com/labstack/echo/v4"
)

type testTemplate struct {
	templates *template.Template
}

func (t *testTemplate) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

type fixture struct {
	api      *Server
	merchant *testMerchant
}

type testMerchant struct {
	handleInstall       func(ctx context.Context, shopURL, accessToken string) (int64, error)
	addVariantID        func(ctx context.Context, shopURL string, variantID int64) (int64, error)
	getVariantIDForShop func(ctx context.Context, shopURL string) (variantID int64, err error)
	getProfileByURL     func(ctx context.Context, shopURL string) (merchant.Profile, error)
	addSubscriptionID   func(ctx context.Context, shopURL string, subscriptionID int64) error
}

func setTestFixture() *fixture {
	srv := New("", "", "")
	srv.Routes()
	t := &testTemplate{}
	srv.Router.Renderer = t
	merchant := &testMerchant{}
	srv.Merchant = merchant
	return &fixture{
		api:      srv,
		merchant: merchant,
	}
}

func (m *testMerchant) HandleInstall(ctx context.Context, shopURL, accessToken string) (int64, error) {
	return m.handleInstall(ctx, shopURL, accessToken)
}

func (m *testMerchant) AddVariantID(ctx context.Context, shopURL string, variantID int64) (int64, error) {
	return m.addVariantID(ctx, shopURL, variantID)
}
func (m *testMerchant) GetVariantIDForShop(ctx context.Context, shopURL string) (variantID int64, err error) {
	return m.getVariantIDForShop(ctx, shopURL)
}
func (m *testMerchant) GetShopByURL(ctx context.Context, shopURL string) (profile merchant.Profile, err error) {
	return m.getProfileByURL(ctx, shopURL)
}

func (m *testMerchant) AddSubscriptionID(ctx context.Context, shopURL string, subscriptionID int64) error {
	return m.addSubscriptionID(ctx, shopURL, subscriptionID)
}
