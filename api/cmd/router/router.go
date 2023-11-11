package router

import (
	"context"
	"github.com/petme/api/internal/handler/authenticated"
	"github.com/petme/api/internal/handler/public"
	"github.com/petme/api/pkg/httpserv"
	"github.com/petme/api/pkg/jwt"
	"net/http"
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