package api

import (
	"log"
	"os"

	"github.com/ColumbiaRoad/cr-shopify-upsell/backend/app/merchant"
	"github.com/ColumbiaRoad/cr-shopify-upsell/backend/lib/server"
	goshopify "github.com/bold-commerce/go-shopify"
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
	Shopify  *goshopify.App
	Merchant merchant.Merchants
}

var AppURL string

// New creates a new Server with an HTTP Router
// @title Carbon offset - Shopify upsell extension
// @version 1.0
// @description Enable merchants to help fight climate change
// @termsOfService http://swagger.io/terms/
// @host tba.com
// @BasePath /v1
func New(apiKey, apiSecret, redirectUrl string) *Server {
	// Create an app somewhere.
	appURL, found := os.LookupEnv("BACKEND_URL")
	if !found {
		log.Fatalf("variable BACKEND_URL not defined")
	}
	AppURL = appURL
	app := goshopify.App{
		ApiKey:      apiKey,
		ApiSecret:   apiSecret,
		RedirectUrl: redirectUrl,
		Scope:       "read_products,write_products,read_orders",
	}
	return &Server{
		Server:  server.New(),
		Shopify: &app,
	}
}
