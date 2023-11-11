package file

import (
	"context"

	"github.com/petme/api/internal/model"
	"github.com/petme/api/internal/repository/orm"
	pkgerrors "github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// InsertFile saves new file to DB
func (i impl) InsertFile(ctx context.Context, input model.File) (model.File, error) {
	ormModel := orm.File{
		UserID: input.UserID,
		Name:   input.Name,
		Type:   input.Type,
		Size:   input.Size,
		Data:   input.Data,
	}
	if err := ormModel.Insert(ctx, i.dbConn, boil.Infer()); err != nil {
		return model.File{}, pkgerrors.WithStack(err)
	}

	input.ID = ormModel.ID
	input.CreatedAt = ormModel.CreatedAt
	input.UpdatedAt = ormModel.UpdatedAt

	return input, nil
}
