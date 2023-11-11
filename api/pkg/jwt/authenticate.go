package jwt

import (
	"context"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

func Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// We can obtain the session token from the requests cookies, which come with every request
		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			// For any other type of error, return a bad request status
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Get the JWT string from the cookie
		tokenStr := c.Value

		claims := &Claims{}
		// Parse the JWT string and store the result in `claims`.
		tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (any, error) {
			return jwtKey, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if !tkn.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// push username to context
		ctx := r.Context()
		req := r.WithContext(context.WithValue(ctx, "userName", claims.Username))
		*r = *req

		// Token is authenticated, pass it through
		next.ServeHTTP(w, r)
	})
}
