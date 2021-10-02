package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

const OrderPaid = "orders/paid"
const CompensationVariantId = "fedcba"

type WebhookResponse struct {
	Msg string `json:"message"`
}

type WebhookForbidden struct {
	Msg string `json:"message"`
}

type LineItem struct {
	VariantId string `json:"variant_id"`
	Quantity  int    `json:"quantity"`
	Price     string `json:"price"`
}

type Order struct {
	Id        string     `json:"id"`
	LineItems []LineItem `json:"line_item"`
}

// @Summary Just demoing things
// @Description Showcasing how to use the application architecture
// @Accept json
// @Produce json
// @Success 200 {object} WebhookResponse "ok"
// @Router /v1/webhook [get]
// @Tags things
func (s *Server) handleWebhook() echo.HandlerFunc {
	return func(c echo.Context) error {
		validated := s.Shopify.VerifyWebhookRequest(c.Request())
		if !validated {
			return s.ErrForbidden("Forbidden")
		}
		// Is webhook type relevant? (orders/paid)
		topic := c.Request().Header.Get("X-Shopify-Topic")
		if topic != OrderPaid {
			// Webhook not relevant.
			return s.Respond(c, http.StatusOK, WebhookResponse{Msg: "Ok"})
		}

		// Check if line items with compensation variant id exists.
		var requestData Order

		if err := s.Bind(c, &requestData); err != nil {
			return s.ErrBadRequest(err.Error())
		}

		var compensationLineItems []LineItem

		for _, s := range requestData.LineItems {
			if s.VariantId == CompensationVariantId {
				// TODO: Get real variant ID for a given merchant. Maybe from a in memory cache?
				compensationLineItems = append(compensationLineItems, s)
			}
		}

		if len(compensationLineItems) == 0 {
			// No compensation item in this order. Nothing to process.
			return s.Respond(c, http.StatusOK, WebhookResponse{Msg: "Ok"})
		}

		i := compensationLineItems[0] // There should always be only one.

		// Compensation item found, let's store the webhook and send an async billing request.
		if err := s.Merchant.HandleIncomingWebhook(c.Request().Context(), "shopurl", requestData.Id, i.Price, i.Quantity); err != nil {
			return s.ErrBadRequest("problem.")
		}

		return s.Respond(c, http.StatusOK, WebhookResponse{Msg: "Ok"})
	}
}
