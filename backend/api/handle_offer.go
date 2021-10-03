package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Offer describes the offered upsell product. See https://shopify.dev/apps/checkout/post-purchase/update-an-order-for-a-checkout-post-purchase-app-extension
type Offer struct {
	VariantId          int64  `json:"variantId"`
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
// @Success 200 {object} Offer "ok"
// @Router /v1/offer [get]
// @Param shop path string true "The myshopify.com url for the shop"
// @Tags things
func (s *Server) handleOffer() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		shopURL := c.QueryParams().Get("shop")
		if shopURL == "" {
			return s.Respond(c, http.StatusBadRequest, ErrorResponse{Error: "missing required parameter: shop_url"})
		}
		variantID, err := s.Merchant.GetVariantIDForShop(ctx, shopURL)
		if err != nil {
			return s.Respond(c, http.StatusNotFound, ErrorResponse{Error: "could not find product variant id for shop"})
		}
		if variantID == 0 {
			return s.Respond(c, http.StatusNotFound, ErrorResponse{Error: "could not find shop"})
		}
		responseBody := Offer{
			VariantId:          variantID,
			ProductTitle:       "Climate Compensation",
			ProductImageURL:    "https://cataas.com/cat/cute/says/hello",
			ProductDescription: "Please save the world.",
			OriginalPrice:      2,
			DiscountedPrice:    2,
		}
		return s.Respond(c, http.StatusOK, responseBody)
	}
}
