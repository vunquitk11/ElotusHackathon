package authenticated

import (
	httpserv2 "github.com/petme/api/pkg/httpserv"
	"net/http"
	"time"
)

// Logout is handler func for logout of system
func (h Handler) Logout() http.HandlerFunc {
	return httpserv2.ErrHandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()

		// immediately clear the token from cookie
		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   "",
			Path:    "/",
			Expires: time.Now(),
		})

		httpserv2.RespondJSON(ctx, w, httpserv2.Success{Message: "success"})
		return nil
	})
}
