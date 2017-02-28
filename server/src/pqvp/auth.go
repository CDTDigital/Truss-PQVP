package main

import (
	"errors"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	request "github.com/dgrijalva/jwt-go/request"
	"go.uber.org/zap"
)

type customClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

// A signing key should typically be a private key that is stored
// in a secure secrets management system. This would change in a
// non-prototype, production system.
var signingKey = []byte("truss-pqvp-demo")

// CreateJwt takes a User and returns a jwt token that has custom claims set.
func CreateJwt(u User) (string, error) {

	claims := customClaims{
		u.Email,
		/// expire the json web token after 15 minutes
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(signingKey)
	if err != nil {
		logger, _ := zap.NewProduction()
		logger.Error("could not sign the token with our key",
			zap.Error(err),
		)
		return "", err
	}
	return ss, nil
}

// Allowed takes a request and verifies if the user is allowed to access
// the requested resource
func Allowed(r *http.Request) (string, error) {
	var claims customClaims
	token, err := request.ParseFromRequestWithClaims(r, request.AuthorizationHeaderExtractor, &claims, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		logger, _ := zap.NewProduction()
		logger.Error("could not parse authorization request",
			zap.Error(err),
		)
	}

	if err == nil && token.Valid {
		return claims.Email, nil
	}

	return claims.Email, errors.New("user is not allowed")
}
