package public

import (
	"encoding/json"
	"net/http"

	"github.com/elotus_hackathon/model"
	"github.com/elotus_hackathon/pkg/httpserv"
)

type userRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r userRequest) Validate() error {
	if r.Username == "" {
		return &httpserv.Error{
			Status: http.StatusBadRequest,
			Code:   "empty_username",
			Desc:   "empty username",
		}
	}

	if r.Password == "" {
		return &httpserv.Error{
			Status: http.StatusBadRequest,
			Code:   "empty_password",
			Desc:   "empty password",
		}
	}
	return nil
}

// Register is handler func for register new user
func (h Handler) Register() http.HandlerFunc {
	return httpserv.ErrHandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()

		var req userRequest
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
