package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/gommon/log"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
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

type signParams struct {
	ReferenceId string `json:"referenceId"`
	Changes     string `json:"changes"`
	Token       string `json:"token"`
}

type signResponse struct {
	Token string `json:"token"`
}

// @Summary Returns the upsell offer product
// @Description Showcasing how to use the application architecture
// @Accept json
// @Produce json
// @Success 200 {object} Offer "ok"
// @Router /v1/offer [get]
// @Param shop path string true "The myshopify.com url for the shop"
// @Tags offer
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

// @Summary Signs the changeset from Shopify
// @Description The sign-changeset endpoint uses JWT tokens for the following reasons:
// @Description Verifies that the request comes from Shopify.
// @DescriptionSigns changes your app wants to apply to an order, for example, adding a product. Shopify uses the signature to verify that the changes come from your app.
// @Accept json
// @Param request body signParams true "Required data to be able to sign a request"
// @Produce json
// @Success 200 {object} signResponse "ok"
// @Router /v1/sign-changeset [post]
// @Tags offer
func (s *Server) handleSignChangeSet() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req signParams
		hmacSecret := []byte(s.ApiSecret)
		if err := s.Bind(c, &req); err != nil {
			return s.ErrBadRequest(err.Error())
		}
		token, err := jwt.Parse(req.Token, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return hmacSecret, nil
		})
		if err != nil {
			return s.Respond(c, http.StatusUnprocessableEntity, err.Error())
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok && !token.Valid {
			return s.Respond(c, http.StatusUnauthorized, err.Error())
		}
		// TODO, this need to be verified once we can see which claims we get in the request.
		fmt.Println("got claims: ", claims)
		if claims["referenceId"] != nil {
			log.Warn("missing claim...")
		}
		signedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"jti":     uuid.New(),
			"iss":     hmacSecret,
			"iat":     time.Now().UTC().UnixNano() / 1e6,
			"sub":     req.ReferenceId, // TODO verify that referenceID is the same as in the claim["decodedToken.input_data.initialPurchase.referenceId"]",
			"changes": req.Changes,
		})

		// Sign and get the complete encoded token as a string using the secret
		tokenString, err := signedToken.SignedString(hmacSecret)
		if err != nil {
			return s.Respond(c, http.StatusInternalServerError, err.Error())
		}
		return s.Respond(c, http.StatusOK, signResponse{Token: tokenString})
	}
}
