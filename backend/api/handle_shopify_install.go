package api

import (
	"fmt"
	"net/http"

	goshopify "github.com/bold-commerce/go-shopify"

	"github.com/labstack/gommon/log"

	"github.com/labstack/echo/v4"
)

const productImageURL = "https://placekitten.com/2048/2048"

// @Summary Shopify Oauth install endpoint
// @Description This is the starting point of the app install flow
// @Accept html
// @Produce html
// @Success 302
// @Router /v1/shopify/ [get]
// @Tags shopify
func (s *Server) handleInstall() echo.HandlerFunc {
	return func(c echo.Context) error {
		shopName := c.QueryParams().Get("shop")
		state := "nonce"
		authUrl := s.Shopify.AuthorizeUrl(shopName, state)
		return s.Redirect(c, http.StatusFound, authUrl)
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
			return s.Respond(c, http.StatusUnauthorized, ErrorResponse{Error: "invalid Signature"})
		}
		ctx := c.Request().Context()
		shopURL := c.QueryParams().Get("shop")
		code := c.QueryParams().Get("code")
		accessToken, err := s.Shopify.GetAccessToken(shopURL, code)
		if err != nil {
			log.Warnf("failed to get token: %s", accessToken)
			return s.Respond(c, http.StatusBadRequest, ErrorResponse{Error: "failed to get token"})
		}
		merchantID, err := s.Merchant.HandleInstall(ctx, shopURL, accessToken)
		if err != nil {
			return s.Respond(c, http.StatusBadRequest, ErrorResponse{Error: "failed to create merchant"})
		}
		shopifyClient := goshopify.NewClient(*s.Shopify, shopURL, accessToken)

		img := goshopify.Image{
			Src: productImageURL,
		}
		var images []goshopify.Image
		images = append(images, img)
		p := goshopify.Product{
			Vendor:   "Compensate",
			Title:    "Carbon offset",
			BodyHTML: "Help fight climate change by donating a small amount of money to the Compensate non-profit climate fund",
			Images:   images,
		}

		product, err := shopifyClient.Product.Create(p)
		if err != nil {
			log.Warn(err)
			return s.Respond(c, http.StatusInternalServerError, ErrorResponse{Error: "failed to create product cariant"})
		}
		merchantID, err = s.Merchant.AddVariantID(ctx, shopURL, product.ID)
		if err != nil {
			log.Warn(err)
			return s.Respond(c, http.StatusInternalServerError, ErrorResponse{Error: "failed to persist product variant"})
		}
		fmt.Println("My produist", product, merchantID)
		// TODO: render the admin template
		log.Warn("merchant id ", merchantID)
		return s.Respond(c, http.StatusOK, ErrorResponse{Error: " Looks good to me!"})
	}
}
