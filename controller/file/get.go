package file

import (
	"context"

	"github.com/petme/model"
)

// GetFilesByUsername return list file by username
func (i impl) GetFilesByUsername(ctx context.Context, username string) ([]model.File, error) {
	uploader, err := i.repo.User().GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	if uploader.ID == 0 {
		return nil, model.ErrUserNotFound
	}

	files, err := i.repo.File().GetFilesByUserID(ctx, uploader.ID)
	if err != nil {
		return nil, err
	}

	return files, nil
}
