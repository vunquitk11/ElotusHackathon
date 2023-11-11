package file

import (
	"context"
	model2 "github.com/petme/api/internal/model"
)

// GetFilesByUsername return list file by username
func (i impl) GetFilesByUsername(ctx context.Context, username string) ([]model2.File, error) {
	uploader, err := i.repo.User().GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	if uploader.ID == 0 {
		return nil, model2.ErrUserNotFound
	}

	files, err := i.repo.File().GetFilesByUserID(ctx, uploader.ID)
	if err != nil {
		return nil, err
	}

	return files, nil
}
