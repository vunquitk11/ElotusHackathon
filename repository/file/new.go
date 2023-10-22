package file

import (
	"github.com/elotus_hackathon/pkg/db/pg"
)

// Repository provides the specification of the functionality provided by this pkg
type Repository interface {
}

// New returns an implementation instance satisfying Repository
func New(dbConn pg.ContextExecutor) Repository {
	return impl{dbConn: dbConn}
}

type impl struct {
	dbConn pg.ContextExecutor
}
