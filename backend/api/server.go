package api

import (
	"html/template"
	"io"
	"os"

	"github.com/ColumbiaRoad/cr-shopify-upsell/backend/app/merchant"
	"github.com/ColumbiaRoad/cr-shopify-upsell/backend/lib/server"
	goshopify "github.com/bold-commerce/go-shopify"
	"github.com/labstack/echo/v4"
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

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

var AppURL string

// New creates a new Server with an HTTP Router
// @title Carbon offset - Shopify upsell extension
// @version 1.0
// @description Enable merchants to help fight climate change
// @termsOfService http://swagger.io/terms/
// @host tba.com
// @BasePath /v1
func New(apiKey, apiSecret, backendURL string) *Server {
	// Create an app somewhere.
	AppURL = backendURL
	app := goshopify.App{
		ApiKey:      apiKey,
		ApiSecret:   apiSecret,
		RedirectUrl: backendURL + "/v1/shopify/callback",
		Scope:       "read_products,write_products,read_orders",
	}

	templatePath, found := os.LookupEnv("TEMPLATE_PATH")
	if !found {
		templatePath = "/app/templates"
	}
	templates := template.Must(template.ParseGlob(templatePath + "/*.html"))
	t := &Template{
		templates: templates,
	}
	s := server.New()

	s.Router.Static("/build", templatePath + "/build")

	s.Router.Renderer = t
	return &Server{
		Server:  s,
		Shopify: &app,
	}
}
