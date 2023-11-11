package public

import (
	"encoding/json"
	"github.com/petme/api/internal/model"
	httpserv2 "github.com/petme/api/pkg/httpserv"
	"github.com/petme/api/pkg/jwt"
	"net/http"
)

// Login is handler func for login to system
func (h Handler) Login() http.HandlerFunc {
	return httpserv2.ErrHandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()

		var req userRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			return err
		}

		if err := req.Validate(); err != nil {
			return err
		}

		user, err := h.userCtrl.Login(ctx, model.User{
			Username: req.Username,
			Password: req.Password,
		})
		if err != nil {
			return err
		}

		tokenString, expirationTime, err := jwt.GenerateToken(user.Username)
		// we set the client cookie for "token" as the JWT we just generated
		// we also set an expiry time which is the same as the token itself
		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Path:    "/",
			Expires: expirationTime,
		})

		httpserv2.RespondJSON(ctx, w, httpserv2.Success{Message: "success"})
		return nil
	})
}
