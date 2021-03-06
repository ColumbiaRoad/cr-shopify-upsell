package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// HealthResponse is returned by the health check
type HealthResponse struct {
	Msg string `json:"message"`
}

// @Summary Health endpoint
// @Description Nice to have for health checks
// @Accept json
// @Produce json
// @Success 200 {object} HealthResponse "ok"
// @Router /health [get]
// @Tags things
func (s *Server) handleHealth() echo.HandlerFunc {
	return func(c echo.Context) error {
		return s.Respond(c, http.StatusOK, HealthResponse{Msg: "ok"})
	}
}
