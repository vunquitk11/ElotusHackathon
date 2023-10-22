package router

import (
	"context"

	"github.com/elotus_hackathon/controller/file"
	"github.com/elotus_hackathon/controller/user"
	"github.com/elotus_hackathon/handler/authenticated"
	"github.com/elotus_hackathon/handler/public"
)

// New creates and returns a new Router instance
func New(
	ctx context.Context,
	userCtrl user.Controller,
	fileCtrl file.Controller,
) Router {
	return Router{
		ctx:                  ctx,
		corsOrigins:          nil,
		authenticatedHandler: authenticated.New(userCtrl, fileCtrl),
		publicHandler:        public.New(userCtrl),
	}
}
