package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4/middleware"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

const UserAgent = "Columbia Road"

type Server struct {
	Router   *echo.Echo
	validate *validator.Validate
}

func New() *Server {
	router := echo.New()
	router.Use(middleware.Logger())
	router.Logger.SetLevel(log.WARN)
	router.Use(middleware.Recover())
	server := Server{
		Router:   router,
		validate: validator.New(),
	}
	return &server
}

// Bind parameters from request context c into struct v
// Respects request's Content-Type header and uses respective struct tags in v for binding.
func (s *Server) Bind(c echo.Context, v interface{}) error {
	if err := json.NewDecoder(c.Request().Body).Decode(v); err != nil {
		log.Errorf("bind parameters: %v", err)
		return s.ErrBadRequest(fmt.Sprint("Invalid request body:", err))
	}
	if err := s.validate.Struct(v); err != nil {
		return s.ErrBadRequest(fmt.Sprintf("Invalid request: %v", err))
	}
	return nil
}

// Respond with give HTTP status code and value v.
// Always sends a JSON response, which means that v must be serializable as JSON.
func (s *Server) Respond(c echo.Context, code int, v interface{}) error {
	c.Response().Header().Set("User-Agent", UserAgent)
	return c.JSON(code, v)
}

// Run the server at address
func (s *Server) Run() error {
	return s.Router.Start(":" + port())
}

// Shutdown gracefully waiting for timeout to finish with all active requests
func (s *Server) Shutdown(timeout time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	log.Info(ctx, "server shutting down...")
	if err := s.Router.Shutdown(ctx); err != nil {
		log.Fatalf("shutdown timeout: %v", err)
	}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Router.ServeHTTP(w, r)
}

func port() string {
	if port, exists := os.LookupEnv("PORT"); exists {
		return port
	}
	return "80"
}
