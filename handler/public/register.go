package public

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/elotus_hackathon/model"
	"github.com/elotus_hackathon/pkg/httpserv"
)

type registerRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r registerRequest) Validate() error {
	if r.Username == "" {
		return errors.New("empty username")
	}

	if r.Password == "" {
		return errors.New("empty password")
	}
	return nil
}

// Register is handler func for register new user
func (h Handler) Register() http.HandlerFunc {
	return httpserv.ErrHandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()

		var req registerRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			return err
		}

		if err := req.Validate(); err != nil {
			return err
		}

		_, err := h.userCtrl.Register(ctx, model.User{
			Username: req.Username,
			Password: req.Password,
		})
		if err != nil {
			return err
		}

		httpserv.RespondJSON(ctx, w, httpserv.Success{Message: "success"})
		return nil
	})
}
