package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// ErrorResponse wraps go errors into an object
type ErrorResponse struct {
	Error string `json:"error"`
}

// ErrBadRequest responds with a 400 error and given error message
func (s *Server) ErrBadRequest(message string) error {
	return echo.NewHTTPError(http.StatusBadRequest, message)
}

// ErrUnauthorized responds witha 401 error and given error message
func (s *Server) ErrUnauthorized(message string) error {
	return echo.NewHTTPError(http.StatusUnauthorized, message)
}

// ErrPaymentRequired responds with a 402 error and given error message
func (s *Server) ErrPaymentRequired(message string) error {
	return echo.NewHTTPError(http.StatusPaymentRequired, message)
}

// ErrForbidden responds with a 403 error and given error message
func (s *Server) ErrForbidden(message string) error {
	return echo.NewHTTPError(http.StatusForbidden, message)
}

// ErrNotFound responds with a 404 error and given error message
func (s *Server) ErrNotFound(message string) error {
	return echo.NewHTTPError(http.StatusNotFound, message)
}

// ErrConflict responds with a 409 error and given error message
func (s *Server) ErrConflict(message string) error {
	return echo.NewHTTPError(http.StatusConflict, message)
}
