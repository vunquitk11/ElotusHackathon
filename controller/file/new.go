package file

import (
	"context"

	"github.com/elotus_hackathon/model"
	"github.com/elotus_hackathon/repository"
)

// The Controller interface provides specification related to order functionality.
type Controller interface {
	// UploadFile saves new image file to db
	UploadFile(ctx context.Context, username string, input model.File) (model.File, error)
}

type impl struct {
	repo repository.Registry
}

// New returns an implementation instance satisfying controller impl
func New(repo repository.Registry) Controller {
	return impl{
		repo: repo,
	}
}
