package api

import (
	"context"
	"encoding/json"

	"net/http"
	"testing"

	"github.com/ColumbiaRoad/cr-shopify-upsell/backend/app/merchant"
	"github.com/ColumbiaRoad/cr-shopify-upsell/backend/lib/server/servertest"

	"github.com/matryer/is"
)

func TestOffer(t *testing.T) {
	is := is.NewRelaxed(t)
	fx := setTestFixture()
	fx.merchant.getVariantIDForShop = func(ctx context.Context, shopURL string) (int64, error) {
		return 123123, nil
	}
	fx.merchant.getProfileByURL = func(ctx context.Context, shopURL string) (merchant.Profile, error) {
		var testProfile merchant.Profile
		testProfile.ShopURL = "offset-demo.myshopify.com"
		return testProfile, nil
	}
	r := servertest.Get(fx.api, "/v1/offer?shop=offset-demo.myshopify.com")
	is.Equal(r.Code, http.StatusOK)
	var offerResponse RenderOffer
	err := json.NewDecoder(r.Body).Decode(&offerResponse)
	is.NoErr(err)
	is.Equal(offerResponse.ProductTitle, "🌍 Take climate action!")
	is.Equal(offerResponse.VariantId, int64(123123))
}

func TestOfferWithoutShopQueryParam(t *testing.T) {
	is := is.NewRelaxed(t)
	fx := setTestFixture()
	r := servertest.Get(fx.api, "/v1/offer")
	is.Equal(r.Code, http.StatusBadRequest)
}

func TestOfferForUnknownShop(t *testing.T) {
	is := is.NewRelaxed(t)
	fx := setTestFixture()
	fx.merchant.getVariantIDForShop = func(ctx context.Context, shopURL string) (int64, error) {
		return 0, nil
	}
	r := servertest.Get(fx.api, "/v1/offer?shop=offset-demo.myshopify.com")
	is.Equal(r.Code, http.StatusNotFound)
}
