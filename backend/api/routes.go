package api

import (
	_ "github.com/ColumbiaRoad/cr-shopify-upsell/backend/api/docs"

	//_ "goStarter/demo/docs"

	_ "github.com/labstack/echo/v4/middleware"

	echoSwagger "github.com/swaggo/echo-swagger"
)

// Routes sets up Server routes
//go:generate swag init -g server.go
func (s *Server) Routes() {
	s.Router.GET("/health", s.handleHealth())
	v1 := s.Router.Group("/v1")

	// Swagger docs
	v1.GET("/docs/*", echoSwagger.WrapHandler)
}
