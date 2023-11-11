package authenticated

import (
	"github.com/petme/api/internal/model"
	httpserv2 "github.com/petme/api/pkg/httpserv"
	"net/http"
)

// GetFilesByUser return list file by login user
func (h Handler) GetFilesByUser() http.HandlerFunc {
	return httpserv2.ErrHandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()

		// pull username from context
		username := ctx.Value("userName").(string)
		if username == "" {
			return &httpserv2.Error{
				Status: http.StatusUnauthorized,
				Code:   "user_not_found",
				Desc:   model.ErrUserNotFound.Error(),
			}
		}

		files, err := h.fileCtrl.GetFilesByUsername(ctx, username)
		if err != nil {
			return err
		}

		httpserv2.RespondJSON(ctx, w, files)
		return nil
	})
}
