package echomiddleware

import (
	"strings"

	Constants "github.com/adhiana46/da-shared/constants"
	Errors "github.com/adhiana46/da-shared/errors"
	PkgToken "github.com/adhiana46/da-shared/pkg/token"
	"github.com/labstack/echo/v4"
)

func JWTAuthMiddleware(tokenHandler *PkgToken.TokenHandler) func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return Errors.NewErrorInvalidToken()
			}

			authHeaderChunk := strings.Split(authHeader, " ")
			if len(authHeaderChunk) != 2 {
				return Errors.NewErrorInvalidToken()
			}

			tokenType, tokenStr := authHeaderChunk[0], authHeaderChunk[1]
			if strings.ToLower(tokenType) != "bearer" {
				return Errors.NewErrorInvalidToken()
			}

			token, claims, err := tokenHandler.Parse(tokenStr)
			if err != nil || !token.Valid {
				return Errors.NewErrorInvalidToken()
			}

			if claims.RegisteredClaims.Subject != "access-token" {
				return Errors.NewErrorInvalidToken()
			}

			// Set user
			c.Set(Constants.USER_KEY_CTX, claims)

			return next(c)
		}
	}
}
