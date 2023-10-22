package file

import (
	"context"

	"github.com/elotus_hackathon/model"
	"github.com/elotus_hackathon/pkg/db/pg"
)

// Repository provides the specification of the functionality provided by this pkg
type Repository interface {
	// InsertFile saves new file to DB
	InsertFile(ctx context.Context, input model.File) (model.File, error)
}

// New returns an implementation instance satisfying Repository
func New(dbConn pg.ContextExecutor) Repository {
	return impl{dbConn: dbConn}
}

type impl struct {
	dbConn pg.ContextExecutor
}
