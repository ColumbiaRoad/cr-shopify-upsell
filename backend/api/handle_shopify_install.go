package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/shopspring/decimal"

	goshopify "github.com/bold-commerce/go-shopify"

	"github.com/labstack/gommon/log"

	"github.com/labstack/echo/v4"
)

const productImageURL = "https://placekitten.com/2048/2048"

// @Summary Shopify admin app start page and starting point for app installs
// @Description This is the starting point of the app install flow but also serves the admin app for already installed merchants
// @Accept html
// @Produce html
// @Success 302
// @Router /v1/shopify/ [get]
// @Tags shopify
func (s *Server) handleShopify() echo.HandlerFunc {
	return func(c echo.Context) error {
		shopName := c.QueryParams().Get("shop")
		ctx := c.Request().Context()
		state := "nonce"
		if ok, _ := s.Shopify.VerifyAuthorizationURL(c.Request().URL); !ok {
			log.Warn("failed to validate signature")
			return s.Respond(c, http.StatusUnauthorized, ErrorResponse{Error: "invalid Signature"})
		}
		profile, err := s.Merchant.GetShopByURL(ctx, shopName)
		if err != nil {
			if err.Error() != "no rows in result set" {
				log.Errorf("failed to check profile: %v", err)
				return s.Respond(c, http.StatusBadRequest, "error when checking for profile:")
			}
		}
		if profile.AccessToken == "" {
			authUrl := s.Shopify.AuthorizeUrl(shopName, state)
			return s.Redirect(c, http.StatusFound, authUrl)
		}
		return c.Render(http.StatusOK, "index.html", map[string]interface{}{
			"shop":   shopName,
			"apiKey": s.Shopify.ApiKey,
		})

	}
}

// @Summary Shopify Callback handler
// @Description Once the merchant has approved access we continue the installation process
// @Accept html
// @Produce html
// @Success 200
// @Router /v1/shopify/callback [get]
// @Tags shopify
func (s *Server) handleCallback() echo.HandlerFunc {
	return func(c echo.Context) error {
		if ok, _ := s.Shopify.VerifyAuthorizationURL(c.Request().URL); !ok {
			log.Warn("failed to validate signature")
			return s.Respond(c, http.StatusUnauthorized, ErrorResponse{Error: "invalid signature"})
		}
		ctx := c.Request().Context()
		shopURL := c.QueryParams().Get("shop")
		code := c.QueryParams().Get("code")
		accessToken, err := s.Shopify.GetAccessToken(shopURL, code)
		if err != nil {
			log.Warnf("failed to get token: %s", accessToken)
			return s.Respond(c, http.StatusBadRequest, ErrorResponse{Error: "failed to get token"})
		}
		_, err = s.Merchant.HandleInstall(ctx, shopURL, accessToken)
		if err != nil {
			errValue := fmt.Sprintf("failed to create merchant: %v", err)
			return s.Respond(c, http.StatusBadRequest, ErrorResponse{Error: errValue})
		}
		shopifyClient := goshopify.NewClient(*s.Shopify, shopURL, accessToken)
		img := goshopify.Image{
			Src: productImageURL,
		}
		var images []goshopify.Image
		images = append(images, img)
		p := goshopify.Product{
			Vendor:   "Climate Action",
			Title:    "Take Climate Action",
			BodyHTML: "Help fight climate change by donating a small amount of money to the Compensate non-profit climate fund",
			Images:   images,
		}
		product, err := shopifyClient.Product.Create(p)
		if err != nil {
			log.Warn(err)
			return s.Respond(c, http.StatusInternalServerError, ErrorResponse{Error: "failed to create product variant"})
		}
		_, err = s.Merchant.AddVariantID(ctx, shopURL, product.ID)
		if err != nil {
			log.Warn(err)
			return s.Respond(c, http.StatusInternalServerError, ErrorResponse{Error: "failed to persist product variant"})
		}
		returnURL := AppURL + "/v1/shopify?" + c.Request().URL.RawQuery
		return s.Redirect(c, http.StatusSeeOther, returnURL)
	}
}

type Response struct {
	ConfirmationURL string `json:"return_url"`
}

// @Summary Initiate shopify billing
// @Description Setup a application charge and redirect the merchant to the Shopify billing approval page
// @Accept html
// @Produce html
// @Success 200 {object} Response "ok"
// @Router /v1/shopify/billing/create [post]
// @Tags shopify
func (s *Server) handleCreateRecurringApplicationCharge() echo.HandlerFunc {
	return func(c echo.Context) error {
		shopURL := c.QueryParams().Get("shop")
		if shopURL == "" {
			return s.Respond(c, http.StatusBadRequest, "missing parameter: shop")
		}
		ctx := c.Request().Context()
		profile, err := s.Merchant.GetShopByURL(ctx, shopURL)
		if err != nil {
			if err.Error() != "no rows in result set" {
				log.Errorf("failed to check profile: %v", err)
				return s.Respond(c, http.StatusBadRequest, "error when checking for profile:")
			}
		}
		shopifyClient := goshopify.NewClient(*s.Shopify, profile.ShopURL, profile.AccessToken)
		cappedAmount := decimal.NewFromInt(10000)
		price := decimal.NewFromInt(0)
		testCharge := true
		var appCharge = goshopify.RecurringApplicationCharge{
			CappedAmount: &cappedAmount,
			Price:        &price,
			Name:         "Climate action",
			ReturnURL:    AppURL + "/v1/shopify/billing/return?shop=" + shopURL,
			Terms:        "We will only charge you the amount you have already collected from your customers",
			Test:         &testCharge,
		}
		chargeResponse, err := shopifyClient.RecurringApplicationCharge.Create(appCharge)
		if err != nil {
			log.Errorf("failed to initiate application charge: %v %v", err, chargeResponse)
			return s.Respond(c, http.StatusBadRequest, "failed to initiate application charge:")
		}
		return s.Respond(c, http.StatusOK, Response{ConfirmationURL: chargeResponse.ConfirmationURL})
	}
}

// @Summary Complete shopify billing
// @Description Handle the callback from Shopify when a merchant accepts the billing charges
// @Accept html
// @Produce html
// @Success 303
// @Router /v1/shopify/billing/return [get]
// @Tags shopify
func (s *Server) handleCompleteRecurringApplicationCharge() echo.HandlerFunc {
	return func(c echo.Context) error {
		shopURL := c.QueryParams().Get("shop")
		chargeIDParam := c.QueryParams().Get("charge_id")

		if shopURL == "" {
			return s.Respond(c, http.StatusBadRequest, "missing parameter: shop")
		}
		if chargeIDParam == "" {
			return s.Respond(c, http.StatusBadRequest, "missing parameter: charge_id")
		}
		chargeID, err := strconv.ParseInt(chargeIDParam, 10, 64)

		ctx := c.Request().Context()
		profile, err := s.Merchant.GetShopByURL(ctx, shopURL)
		if err != nil {
			if err.Error() != "no rows in result set" {
				log.Errorf("failed to check profile: %v", err)
				return s.Respond(c, http.StatusBadRequest, "error when checking for profile:")
			}
		}
		shopifyClient := goshopify.NewClient(*s.Shopify, profile.ShopURL, profile.AccessToken)

		charge, err := shopifyClient.RecurringApplicationCharge.Get(chargeID, nil)
		if err != nil {
			log.Errorf("failed to get application charge: %v: charge_id: %i", err, chargeID)
			return s.Respond(c, http.StatusBadRequest, "failed to activate application charge:")
		}
		activateChargeResponse, err := shopifyClient.RecurringApplicationCharge.Activate(*charge)
		if err != nil {
			log.Errorf("failed to activate application charge: %v %v", err, activateChargeResponse)
			return s.Respond(c, http.StatusBadRequest, "failed to activate application charge:")
		}
		return c.Render(http.StatusOK, "success.html", map[string]interface{}{
			"subscription_id": activateChargeResponse.ID,
		})

	}
}
