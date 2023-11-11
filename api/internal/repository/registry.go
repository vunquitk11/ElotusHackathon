package repository

import (
	"github.com/petme/api/internal/repository/file"
	"github.com/petme/api/internal/repository/user"
	"github.com/petme/api/pkg/db/pg"
)

// Registry is the registry of all the domain specific repositories and also provides transaction capabilities.
type Registry interface {
	// User returns the user repo
	User() user.Repository
	// File returns the file repo
	File() file.Repository
}

// New returns a new instance of Registry
func New(dbConn pg.BeginnerExecutor) Registry {
	return impl{
		dbConn: dbConn,
		user:   user.New(dbConn),
		file:   file.New(dbConn),
	}
}

type impl struct {
	dbConn pg.BeginnerExecutor // Only used to start DB txns
	user   user.Repository
	file   file.Repository
}

// User returns the user repo
func (i impl) User() user.Repository {
	return i.user
}

// File returns the file repo
func (i impl) File() file.Repository {
	return i.file
}
