package file

import (
	"context"
	"github.com/petme/api/internal/model"
	"github.com/petme/api/pkg/db/pg"
)

// Repository provides the specification of the functionality provided by this pkg
type Repository interface {
	// InsertFile saves new file to DB
	InsertFile(ctx context.Context, input model.File) (model.File, error)
	// GetFilesByUserID return list file by user id
	GetFilesByUserID(ctx context.Context, userID int64) ([]model.File, error)
}

// New returns an implementation instance satisfying Repository
func New(dbConn pg.ContextExecutor) Repository {
	return impl{dbConn: dbConn}
}

type impl struct {
	dbConn pg.ContextExecutor
}
