package authenticated

import (
	"net/http"
	"time"

	"github.com/letsvote/api/pkg/httpserv"
)

// Logout is handler func for logout of system
func (h Handler) Logout() http.HandlerFunc {
	return httpserv.ErrHandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()

		// immediately clear the token from cookie
		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   "",
			Path:    "/",
			Expires: time.Now(),
		})

		httpserv.RespondJSON(ctx, w, httpserv.Success{Message: "success"})
		return nil
	})
}
