package authenticated

import (
	"github.com/letsvote/api/internal/controller/file"
	"github.com/letsvote/api/internal/controller/user"
)

// Handler is the web handler for this pkg
type Handler struct {
	userCtrl user.Controller
	fileCtrl file.Controller
}

// New instantiates a new Handler and returns it
func New(userCtrl user.Controller, fileCtrl file.Controller) Handler {
	return Handler{
		userCtrl: userCtrl,
		fileCtrl: fileCtrl,
	}
}
