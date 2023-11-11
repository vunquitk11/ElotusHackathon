package public

import (
	"encoding/json"
	"github.com/petme/api/internal/model"
	httpserv2 "github.com/petme/api/pkg/httpserv"
	"net/http"
)

type userRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r userRequest) Validate() error {
	if r.Username == "" {
		return &httpserv2.Error{
			Status: http.StatusBadRequest,
			Code:   "empty_username",
			Desc:   "empty username",
		}
	}

	if r.Password == "" {
		return &httpserv2.Error{
			Status: http.StatusBadRequest,
			Code:   "empty_password",
			Desc:   "empty password",
		}
	}
	return nil
}

// Register is handler func for register new user
func (h Handler) Register() http.HandlerFunc {
	return httpserv2.ErrHandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
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

		httpserv2.RespondJSON(ctx, w, httpserv2.Success{Message: "success"})
		return nil
	})
}
