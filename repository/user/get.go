package user

import (
	"context"
	"database/sql"
	"errors"

	"github.com/petme/model"
	"github.com/petme/repository/orm"
	pkgerrors "github.com/pkg/errors"
)

// GetUserByUsername return user by username
func (i impl) GetUserByUsername(ctx context.Context, username string) (model.User, error) {
	ormModel, err := orm.Users(
		orm.UserWhere.Username.EQ(username),
	).One(ctx, i.dbConn)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.User{}, nil
		}
		return model.User{}, pkgerrors.WithStack(err)
	}
	return model.User{
		ID:       ormModel.ID,
		Username: ormModel.Username,
		Password: ormModel.Password,
	}, nil
}
