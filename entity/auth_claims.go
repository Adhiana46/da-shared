package entity

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// JWT TOKEN PAYLOAD
type AuthClaims struct {
	Id           string    `json:"id"`
	Email        string    `json:"email"`
	Package      string    `json:"package"`
	LastActiveAt time.Time `json:"last_active_at"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	jwt.RegisteredClaims
}
