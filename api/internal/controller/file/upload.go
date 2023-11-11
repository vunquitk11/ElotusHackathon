package file

import (
	"context"

	"github.com/petme/api/internal/model"
)

// UploadFile saves new image file to db
func (i impl) UploadFile(ctx context.Context, username string, input model.File) (model.File, error) {
	uploader, err := i.repo.User().GetUserByUsername(ctx, username)
	if err != nil {
		return model.File{}, err
	}

	if uploader.ID == 0 {
		return model.File{}, model.ErrUserNotFound
	}

	input.UserID = uploader.ID
	file, err := i.repo.File().InsertFile(ctx, input)
	if err != nil {
		return model.File{}, err
	}

	return file, nil
}
