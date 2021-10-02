package api

import (
	"encoding/json"

	"net/http"
	"testing"

	"github.com/ColumbiaRoad/cr-shopify-upsell/backend/lib/server/servertest"

	"github.com/matryer/is"
)

type OrderInput struct {
	Id        string      `json:"id"`
	LineItems [2]LineItem `json:"line_item"`
}

func TestWebhook(t *testing.T) {
	is := is.NewRelaxed(t)
	fx := setTestFixture()

	var lineItems [2]LineItem
	lineItems[0] = LineItem{
		VariantId: "abcdef",
		Quantity:  1,
	}
	lineItems[1] = LineItem{
		VariantId: "fedcba",
		Quantity:  2,
	}

	postQuery := OrderInput{
		Id:        "abc",
		LineItems: lineItems,
	}

	r := servertest.Post(fx.api, "/v1/shopify/webhook", postQuery)
	is.Equal(r.Code, http.StatusOK)
	var healthResponse HealthResponse
	err := json.NewDecoder(r.Body).Decode(&healthResponse)
	is.NoErr(err)
	is.Equal(healthResponse.Msg, "ok")
}

/*
func WebhookShouldRespondOKUnnecessaryTopics(t *testing.T) {
	is := is.NewRelaxed(t)
	fx := setTestFixture()

	postQuery := s{
		field: "",
	}

	r := servertest.Post(fx.api, "/v1/shopify/webhook", postQuery)
	is.Equal(r.Code, http.StatusOK)
	var healthResponse HealthResponse
	err := json.NewDecoder(r.Body).Decode(&healthResponse)
	is.NoErr(err)
	is.Equal(healthResponse.Msg, "ok")
}
*/
