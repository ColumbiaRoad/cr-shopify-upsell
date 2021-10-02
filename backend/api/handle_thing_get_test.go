package api

import (
	"encoding/json"

	"net/http"
	"testing"

	"github.com/ColumbiaRoad/cr-shopify-upsell/backend/lib/server/servertest"

	"github.com/matryer/is"
)

func TestGetThing(t *testing.T) {
	is := is.NewRelaxed(t)
	fx := setTestFixture()
	r := servertest.Get(fx.api, "/health")
	is.Equal(r.Code, http.StatusOK)
	var healthResponse HealthResponse
	err := json.NewDecoder(r.Body).Decode(&healthResponse)
	is.NoErr(err)
	is.Equal(healthResponse.Msg, "ok")
}
