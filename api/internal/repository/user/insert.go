package user

import (
	"context"
	"github.com/petme/api/internal/model"
	"github.com/petme/api/internal/repository/orm"

	pkgerrors "github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// InsertUser saves new user to DB
func (i impl) InsertUser(ctx context.Context, input model.User) (model.User, error) {
	ormModel := orm.User{
		Username: input.Username,
		Password: input.Password,
	}
	if err := ormModel.Insert(ctx, i.dbConn, boil.Infer()); err != nil {
		return model.User{}, pkgerrors.WithStack(err)
	}

	input.ID = ormModel.ID
	input.CreatedAt = ormModel.CreatedAt
	input.UpdatedAt = ormModel.UpdatedAt

	return input, nil
}
