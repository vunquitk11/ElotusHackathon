package httpserv

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Handler(
	routerFunc func(chi.Router)) http.Handler {
	r := chi.NewRouter()
	// TODO: add middleware here
	r.Use(
		// render.SetContentType(render.ContentTypeJSON), // set content-type headers as application/json
		middleware.Logger, // log api request calls
		// middleware.DefaultCompress, // compress results, mostly gzipping assets and json
		middleware.StripSlashes, // match paths with a trailing slash, strip it, and continue routing through the mux
		middleware.Recoverer,    // recover from panics without crashing server
	)

	r.Group(routerFunc)

	return r
}
