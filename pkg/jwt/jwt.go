package jwt

import (
	"github.com/golang-jwt/jwt/v5"
)

// Create the JWT key used to create the signature
var jwtKey = []byte("my_secret_key")

// Claims is a struct that will be encoded to a JWT.
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}
