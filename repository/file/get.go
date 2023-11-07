package file

import (
	"context"

	"github.com/petme/model"
	"github.com/petme/repository/orm"
	pkgerrors "github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// GetFilesByUserID return list file by user id
func (i impl) GetFilesByUserID(ctx context.Context, userID int64) ([]model.File, error) {
	ormSlice, err := orm.Files(
		orm.FileWhere.UserID.EQ(userID),
		qm.OrderBy(orm.FileColumns.CreatedAt+" DESC"),
	).All(ctx, i.dbConn)
	if err != nil {
		return nil, pkgerrors.WithStack(err)
	}
	return toFiles(ormSlice), nil
}

func toFiles(models orm.FileSlice) []model.File {
	result := make([]model.File, len(models))
	for idx, item := range models {
		result[idx] = model.File{
			ID:        item.ID,
			UserID:    item.UserID,
			Name:      item.Name,
			Type:      item.Type,
			Size:      item.Size,
			Data:      item.Data,
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		}
	}
	return result
}
