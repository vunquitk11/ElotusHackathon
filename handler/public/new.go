package public

import (
	"github.com/petme/controller/user"
)

// Handler is the web handler for this pkg
type Handler struct {
	userCtrl user.Controller
}

// New instantiates a new Handler and returns it
func New(userCtrl user.Controller) Handler {
	return Handler{
		userCtrl: userCtrl,
	}
}
