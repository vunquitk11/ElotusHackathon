package public

import (
	"encoding/json"
	"net/http"

	"github.com/elotus_hackathon/model"
	"github.com/elotus_hackathon/pkg/httpserv"
)

// Login is handler func for login to system
func (h Handler) Login() http.HandlerFunc {
	return httpserv.ErrHandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()

		var req userRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			return err
		}

		if err := req.Validate(); err != nil {
			return err
		}

		result, err := h.userCtrl.Login(ctx, model.User{
			Username: req.Username,
			Password: req.Password,
		})
		if err != nil {
			return err
		}

		httpserv.RespondJSON(ctx, w, httpserv.Success{Message: result})
		return nil
	})
}
