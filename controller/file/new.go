package file

import "github.com/elotus_hackathon/repository"

// The Controller interface provides specification related to order functionality.
type Controller interface {
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
