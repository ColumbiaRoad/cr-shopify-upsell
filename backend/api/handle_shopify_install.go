package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// InstallResponse is returned by the health check
type InstallResponse struct {
	Msg string `json:"message"`
}

// @Summary Shopify Oauth install endpoint
// @Description This is the starting point of the app install flow
// @Accept json
// @Produce json
// @Success 200 {object} InstallResponse "ok"
// @Router /v1/install [get]
// @Tags shopify
func (s *Server) handleInstall() echo.HandlerFunc {
	return func(c echo.Context) error {
		shopName := c.QueryParams().Get("shop")
		state := "nonce"
		authUrl := s.App.AuthorizeUrl(shopName, state)
		fmt.Println("sssss", authUrl)
		return s.Redirect(c, http.StatusFound, authUrl)
	}
}
