package api

import (
	"github.com/ColumbiaRoad/cr-shopify-upsell/backend/lib/server"
)

// ErrorResponse wraps go errors into an object
type ErrorResponse struct {
	Error string `json:"error"`
}

// SuccessResponse wraps successful responses into an object
type SuccessResponse struct {
	Payload string `json:"payload"`
}

// Server contains all the necessary dependencies for running the service
// this is where you would add your "app's" which has the business logic
type Server struct {
	*server.Server
	JWTSecret string
}

// New creates a new Server with an HTTP Router
// @title The Service
// @version 1.0
// @description This is a sample server
// @termsOfService http://swagger.io/terms/
// @host things-host.com
// @BasePath /v1
func New() *Server {
	return &Server{
		Server: server.New(),
	}
}
