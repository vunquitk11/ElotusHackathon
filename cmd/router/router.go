package router

import (
	"context"
	"net/http"

	"github.com/elotus_hackathon/handler/authenticated"
	"github.com/elotus_hackathon/handler/public"
	"github.com/elotus_hackathon/pkg/httpserv"
	"github.com/go-chi/chi/v5"
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
	r.Route(prefix, func(r chi.Router) {
		r.Get("/test-auth", rtr.authenticatedHandler.TestAuth())
	})
}

func (rtr Router) public(r chi.Router) {
	const prefix = "/public"
	r.Route(prefix, func(r chi.Router) {
		r.Get("/test-public", rtr.publicHandler.TestPublic())
		r.Post("/register", rtr.publicHandler.Register())
		r.Post("/login", rtr.publicHandler.Login())
	})
}
