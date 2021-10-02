package api

import (
	"encoding/json"

	"net/http"
	"testing"

	"github.com/ColumbiaRoad/cr-shopify-upsell/backend/lib/server/servertest"

	"github.com/matryer/is"
)

func TestOfferEndpoint(t *testing.T) {
	is := is.NewRelaxed(t)
	fx := setTestFixture()
	r := servertest.Get(fx.api, "/v1/offer")
	is.Equal(r.Code, http.StatusOK)
	var offerResponse OfferResponse
	err := json.NewDecoder(r.Body).Decode(&offerResponse)
	is.NoErr(err)
	is.Equal(offerResponse.ProductTitle, "Climate Compensation")
}
