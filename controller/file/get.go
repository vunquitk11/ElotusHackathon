package file

import (
	"context"
	"errors"

	"github.com/elotus_hackathon/model"
)

// GetFilesByUsername return list file by username
func (i impl) GetFilesByUsername(ctx context.Context, username string) ([]model.File, error) {
	uploader, err := i.repo.User().GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	if uploader.ID == 0 {
		return nil, errors.New("user not found")
	}

	files, err := i.repo.File().GetFilesByUserID(ctx, uploader.ID)
	if err != nil {
		return nil, err
	}

	return files, nil
}
