package echomiddleware

import (
	"log/slog"

	"github.com/labstack/echo/v4"
)

func RequestLoggingMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			details := map[string]any{
				"method":      c.Request().Method,
				"endpoint":    c.Request().URL.String(),
				"client_ip":   c.RealIP(),
				"user_agenet": c.Request().UserAgent(),
			}
			slog.Info("HTTP Request Received", slog.Any("details", details))

			return next(c)
		}
	}
}
