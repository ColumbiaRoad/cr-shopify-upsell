package api

import (
	"net/http"

	"github.com/ColumbiaRoad/cr-shopify-upsell/backend/services"
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
		// 1. Verify webhookshopifyApp.VerifyWebhookRequest
		/*validated := s.App.VerifyWebhookRequest(c.Request())
		if !validated {
			return s.ErrForbidden("Forbidden")
		}
		// 2. Check type (orders/paid)
		topic := c.Request().Header.Get("X-Shopify-Topic")
		if topic != OrderPaid {
			// Webhook not relevant.
			return s.Respond(c, http.StatusOK, WebhookResponse{Msg: "Ok"})
		}*/

		// 3. Check if line items with correct variantID exists.
		var requestData Order

		if err := s.Bind(c, &requestData); err != nil {
			return s.ErrBadRequest(err.Error())
		}

		var compensationLineItems []LineItem

		for _, s := range requestData.LineItems {
			if s.VariantId == CompensationVariantId {
				compensationLineItems = append(compensationLineItems, s)
			}
		}

		if len(compensationLineItems) == 0 {
			// No compensation items in this order. Nothing to process.
			return s.Respond(c, http.StatusOK, WebhookResponse{Msg: "Ok"})
		}

		// 4. Some items found, let's send a billing request.
		// 4.1. Store in DB
		// TODO
		// 4.2. Fire another thread to make billing request async
		go services.MakeShopifyBillingRequest()
		// 4.3. Respond OK.
		return s.Respond(c, http.StatusOK, WebhookResponse{Msg: requestData.LineItems[1].VariantId})
	}
}

// there can be multiple quantity of the compensation product.
//
