package authenticated

import (
	"errors"
	"net/http"

	"github.com/elotus_hackathon/pkg/httpserv"
)

// GetFilesByUser return list file by login user
func (h Handler) GetFilesByUser() http.HandlerFunc {
	return httpserv.ErrHandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()

		// pull username from context
		username := ctx.Value("userName").(string)
		if username == "" {
			return errors.New("user not found")
		}

		files, err := h.fileCtrl.GetFilesByUsername(ctx, username)
		if err != nil {
			return err
		}

		httpserv.RespondJSON(ctx, w, files)
		return nil
	})
}
