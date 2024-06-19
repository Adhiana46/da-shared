package token

import (
	Cache "github.com/adhiana46/da-shared/pkg/cache"
	"github.com/golang-jwt/jwt/v5"
)

type Option func(*TokenHandler)

func Issuer(issuer string) Option {
	return func(th *TokenHandler) {
		th.issuer = issuer
	}
}

func SigningMethod(signingMethod *jwt.SigningMethodHMAC) Option {
	return func(th *TokenHandler) {
		th.signingMethod = signingMethod
	}
}

func SecretKey(secretKey string) Option {
	return func(th *TokenHandler) {
		th.secretKey = secretKey
	}
}

func WithCache(cache Cache.Cache) Option {
	return func(th *TokenHandler) {
		th.c = cache
	}
}
