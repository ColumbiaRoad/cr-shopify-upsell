package api

import (
	_ "github.com/ColumbiaRoad/cr-shopify-upsell/backend/api/docs"
	"github.com/labstack/echo/v4/middleware"

	echoSwagger "github.com/swaggo/echo-swagger"
)

// Routes sets up Server routes
//go:generate swag init -g server.go
func (s *Server) Routes() {
	s.Router.Use(middleware.CORS())
	s.Router.GET("/health", s.handleHealth())
	v1 := s.Router.Group("/v1")
	// Swagger docs
	v1.GET("/docs/*", echoSwagger.WrapHandler)
	v1.GET("/shopify", s.handleShopify())
	v1.GET("/shopify/callback", s.handleCallback())
	v1.GET("/shopify/billing/create", s.handleCreateRecurringApplicationCharge())
	v1.GET("/offer", s.handleOffer())
	v1.POST("/sign-changeset", s.handleSignChangeSet())
}
