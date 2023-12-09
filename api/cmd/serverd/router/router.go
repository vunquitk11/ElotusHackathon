package router

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/letsvote/api/internal/handler/authenticated"
	"github.com/letsvote/api/internal/handler/public"
	"github.com/letsvote/api/pkg/httpserv"
	"github.com/letsvote/api/pkg/jwt"
)

// Router defines the routes & handlers of the app
type Router struct {
	ctx                  context.Context
	corsOrigins          []string
	authenticatedHandler authenticated.Handler
	publicHandler        public.Handler
}

// Handler returns the Handler for use by the server
func (rtr Router) Handler() http.Handler {
	return httpserv.Handler(
		rtr.routes,
	)
}

func (rtr Router) routes(r chi.Router) {
	r.Group(rtr.authenticated)
	r.Group(rtr.public)
}

func (rtr Router) authenticated(r chi.Router) {
	const prefix = "/authenticated"
	r.Use(jwt.Authenticator)
	r.Route(prefix, func(r chi.Router) {
		r.Post("/upload", rtr.authenticatedHandler.UploadFile())
		r.Post("/logout", rtr.authenticatedHandler.Logout())
		r.Get("/files", rtr.authenticatedHandler.GetFilesByUser())
	})
}

func (rtr Router) public(r chi.Router) {
	const prefix = "/public"
	r.Route(prefix, func(r chi.Router) {
		r.Post("/register", rtr.publicHandler.Register())
		r.Post("/login", rtr.publicHandler.Login())
	})
}
