package api

import (
	"net/http"

	"github.com/labstack/gommon/log"

	"github.com/labstack/echo/v4"
)

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
// @Description This is the starting point of the app install flow
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
		shopName := c.QueryParams().Get("shop")
		code := c.QueryParams().Get("code")
		token, err := s.Shopify.GetAccessToken(shopName, code)
		if err != nil {
			log.Warnf("failed to get token: %s", token)
			return s.Respond(c, http.StatusBadRequest, ErrorResponse{Error: "failed to get token"})
		}
		// TODO: Persist the token and the merchant
		// TODO: render the admin template
		return s.Respond(c, http.StatusOK, ErrorResponse{Error: "Looks good to me!" + token})
	}
}
