package file

import (
	"context"

	"github.com/elotus_hackathon/model"
)

// UploadFile saves new image file to db
func (i impl) UploadFile(ctx context.Context, input model.File) (model.File, error) {
	file, err := i.repo.File().InsertFile(ctx, input)
	if err != nil {
		return model.File{}, err
	}

	return file, nil
}
