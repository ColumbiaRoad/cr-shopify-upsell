package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// OfferResponse describes the offered upsell product. See https://shopify.dev/apps/checkout/post-purchase/update-an-order-for-a-checkout-post-purchase-app-extension
type OfferResponse struct {
	VariantId          string `json:"variantId"`
	ProductTitle       string `json:"productTitle"`
	ProductImageURL    string `json:"productImageURL"`
	ProductDescription string `json:"productDescription"`
	OriginalPrice      int    `json:"originalPrice"`
	DiscountedPrice    int    `json:"discountedPrice"`
}

// @Summary Returns the upsell offer product
// @Description Showcasing how to use the application architecture
// @Accept json
// @Produce json
// @Success 200 {object} OfferResponse "ok"
// @Router /v1/offer [get]
// @Tags things
func (s *Server) handleOffer() echo.HandlerFunc {
	return func(c echo.Context) error {
		responseBody := OfferResponse{
			VariantId:          "abcdefg",
			ProductTitle:       "Climate Compensation",
			ProductImageURL:    "https://cataas.com/cat/cute/says/hello",
			ProductDescription: "Please save the world.",
			OriginalPrice:      2,
			DiscountedPrice:    2,
		}

		return s.Respond(c, http.StatusOK, responseBody)
	}
}
