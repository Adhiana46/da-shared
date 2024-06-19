package echomiddleware

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/labstack/echo/v4"
)

// Store request body into request context
// only store json request
func ContextRequestBodyMiddleware(keyValues ...map[contextKey]any) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// IF request is not application/json
			if c.Request().Header.Get("Content-Type") != "application/json" {
				return next(c)
			}

			// Read the request body
			reqBody, err := io.ReadAll(c.Request().Body)
			if err != nil {
				return err
			}

			// Restore the request body to the context
			c.Request().Body = io.NopCloser(bytes.NewBuffer(reqBody))

			reqBodyMap := map[string]any{}
			if err := json.Unmarshal(reqBody, &reqBodyMap); err != nil {
				return err
			}

			c.Set("request_body", reqBodyMap)

			return next(c)
		}
	}
}
