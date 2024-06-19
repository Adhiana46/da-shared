package token

import (
	"crypto/md5"
	"fmt"
	"io"
	"time"

	Entity "github.com/adhiana46/da-shared/entity"
	Errors "github.com/adhiana46/da-shared/errors"
	Cache "github.com/adhiana46/da-shared/pkg/cache"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
)

const (
	_cacheInvalidateTokenKey = "token-blacklist:%s"
)

type TokenHandler struct {
	issuer        string
	signingMethod *jwt.SigningMethodHMAC
	secretKey     string
	// Cache to store invalidated token
	c Cache.Cache
}

func NewTokenHandler(opts ...Option) *TokenHandler {
	instance := TokenHandler{
		issuer:        "token",
		signingMethod: jwt.SigningMethodHS256,
		secretKey:     "secret",
		c:             nil,
	}

	// apply options
	for _, opt := range opts {
		opt(&instance)
	}

	return &instance
}

func (t *TokenHandler) Generate(claims *Entity.AuthClaims) (string, error) {
	path := "TokenHandler:Generate"

	// Set issuer
	claims.Issuer = t.issuer

	token := jwt.NewWithClaims(t.signingMethod, claims)
	tokenStr, err := token.SignedString([]byte(t.secretKey))
	if err != nil {
		return "", errors.Wrap(err, path)
	}

	return tokenStr, nil
}

func (t *TokenHandler) Parse(tokenStr string) (*jwt.Token, *Entity.AuthClaims, error) {
	path := "TokenHandler:Parse"
	secretKey := t.secretKey

	// VALIDATE tokenStr Format
	_, _, err := new(jwt.Parser).ParseUnverified(tokenStr, jwt.MapClaims{})
	if err != nil {
		return nil, nil, errors.Wrap(Errors.NewErrorInvalidToken(), path)
	}

	// CHECK IF TOKEN ALREADY INVALIDATED
	if t.c != nil {
		tokenHasher := md5.New()
		if _, err := io.WriteString(tokenHasher, tokenStr); err != nil {
			return nil, nil, errors.Wrap(err, path)
		}
		tokenSum := tokenHasher.Sum(nil)
		tokenKey := fmt.Sprintf(_cacheInvalidateTokenKey, string(tokenSum))
		if v, _ := t.c.Get(tokenKey); v == "1" {
			return nil, nil, errors.Wrap(Errors.NewErrorInvalidToken(), path)
		}
	}

	token, err := jwt.ParseWithClaims(tokenStr, &Entity.AuthClaims{}, func(t *jwt.Token) (interface{}, error) {
		// validate signing algo
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, nil, errors.Wrap(err, path)
	}

	if !token.Valid {
		return nil, nil, errors.Wrap(Errors.NewErrorInvalidToken(), path)
	}

	claims, ok := token.Claims.(*Entity.AuthClaims)
	if !ok {
		return nil, nil, errors.Wrap(Errors.NewErrorInvalidToken("invalid token: error parsing token"), path)
	}

	if claims.Issuer != t.issuer {
		return nil, nil, errors.Wrap(Errors.NewErrorInvalidToken("invalid token: invalid issuer"), path)
	}

	return token, claims, nil
}

// Invalidate token
// store the token in cache until it expire
func (t *TokenHandler) Invalidate(tokenStr string) error {
	path := "TokenHandler:Invalidate"

	if t.c == nil {
		return nil
	}

	token, _, err := new(jwt.Parser).ParseUnverified(tokenStr, jwt.MapClaims{})
	if err != nil {
		return errors.Wrap(Errors.NewErrorInvalidToken(), path)
	}

	tokenExpTime, err := token.Claims.GetExpirationTime()
	if err != nil {
		return errors.Wrap(err, path)
	}

	// calc expire time in seconds
	now := time.Now()
	duration := tokenExpTime.Sub(now)
	expSecond := int32(duration.Seconds())

	fmt.Println("EXP SECOND", expSecond)

	// Token already expired
	if expSecond <= 0 {
		return nil
	}

	tokenHasher := md5.New()
	if _, err := io.WriteString(tokenHasher, tokenStr); err != nil {
		return errors.Wrap(err, path)
	}
	tokenSum := tokenHasher.Sum(nil)
	tokenKey := fmt.Sprintf(_cacheInvalidateTokenKey, string(tokenSum))
	if v, _ := t.c.Get(tokenKey); v == "1" {
		return errors.Wrap(Errors.NewErrorInvalidToken(), path)
	}

	if err := t.c.Set(tokenKey, "1", expSecond); err != nil {
		return errors.Wrap(err, path)
	}

	return nil
}
