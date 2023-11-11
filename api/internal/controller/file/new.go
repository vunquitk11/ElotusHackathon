package file

import (
	"context"
	"github.com/petme/api/internal/model"
	"github.com/petme/api/internal/repository"
)

// The Controller interface provides specification related to order functionality.
type Controller interface {
	// UploadFile saves new image file to db
	UploadFile(ctx context.Context, username string, input model.File) (model.File, error)
	// GetFilesByUsername return list file by username
	GetFilesByUsername(ctx context.Context, username string) ([]model.File, error)
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
