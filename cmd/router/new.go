package router

import (
	"context"

	"github.com/petme/controller/file"
	"github.com/petme/controller/user"
	"github.com/petme/handler/authenticated"
	"github.com/petme/handler/public"
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
