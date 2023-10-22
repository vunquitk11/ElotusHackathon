package jwt

import (
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(username string) (string, time.Time, error) {
	// Declare the expiration time of the token
	minutes, err := strconv.Atoi(os.Getenv("EXPIRATION_MINUTES"))
	if err != nil {
		return "", time.Time{}, err
	}

	expirationTime := time.Now().Add(time.Duration(rand.Int31n(int32(minutes))) * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", time.Time{}, err
	}

	return tokenString, expirationTime, nil
}
