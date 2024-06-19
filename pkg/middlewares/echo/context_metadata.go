package echomiddleware

import (
	"context"

	"github.com/labstack/echo/v4"
)

type contextKey string

const (
	keyIPAddress contextKey = "ip_address"
	keyUserAgent contextKey = "user_agent"
)

// Store metadata like ip-address, user-agent, etc to request context
func ContextMetadataMiddleware(keyValues ...map[contextKey]any) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := c.Request().Context()
			values := map[contextKey]any{
				keyIPAddress: c.RealIP(),
				keyUserAgent: c.Request().UserAgent(),
			}
			for key, value := range values {
				ctx = context.WithValue(ctx, key, value)
			}

			// Set value from keyValues
			if len(keyValues) > 0 {
				for key, value := range keyValues[0] {
					ctx = context.WithValue(ctx, key, value)
				}
			}

			// SET CONTEXT
			c.SetRequest(c.Request().WithContext(ctx))

			return next(c)
		}
	}
}
