package file

import (
	"context"
	model2 "github.com/petme/api/internal/model"
)

// UploadFile saves new image file to db
func (i impl) UploadFile(ctx context.Context, username string, input model2.File) (model2.File, error) {
	uploader, err := i.repo.User().GetUserByUsername(ctx, username)
	if err != nil {
		return model2.File{}, err
	}

	if uploader.ID == 0 {
		return model2.File{}, model2.ErrUserNotFound
	}

	input.UserID = uploader.ID
	file, err := i.repo.File().InsertFile(ctx, input)
	if err != nil {
		return model2.File{}, err
	}

	return file, nil
}
