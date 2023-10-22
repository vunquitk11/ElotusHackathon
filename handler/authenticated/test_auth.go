package authenticated

import (
	"net/http"

	"github.com/elotus_hackathon/pkg/httpserv"
)

func (h Handler) TestAuth() http.HandlerFunc {
	return httpserv.ErrHandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()

		httpserv.RespondJSON(ctx, w, httpserv.Success{Message: "TEST AUTHENTICATED SUCCESS"})
		return nil
	})
}
