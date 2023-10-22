package file

import (
	"context"
	"testing"

	"github.com/elotus_hackathon/model"
	"github.com/elotus_hackathon/pkg/db/pg"
	"github.com/elotus_hackathon/pkg/testutil"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/require"
)

func Test_InsertFile(t *testing.T) {
	t.Setenv("DB_URL", "postgres://postgres:postgres@localhost:5432/elotus?sslmode=disable")
	type arg struct {
		givenInput model.File
		expResult  model.File
		expErr     error
	}

	tcs := map[string]arg{
		"success": {
			givenInput: model.File{
				UserID: 1,
				Name:   "filename",
				Type:   "filetype",
				Size:   1024,
				Data:   "abc",
			},
			expResult: model.File{
				ID:     1,
				UserID: 1,
				Name:   "filename",
				Type:   "filetype",
				Size:   1024,
				Data:   "abc",
			},
		},
	}
	for scenario, tc := range tcs {
		t.Run(scenario, func(t *testing.T) {
			testutil.WithTxDB(t, func(tx pg.BeginnerExecutor) {
				// When
				instance := New(tx)
				result, err := instance.InsertFile(context.Background(), tc.givenInput)

				// Then
				if tc.expErr != nil {
					require.EqualError(t, err, tc.expErr.Error())
				} else {
					require.NoError(t, err)
					require.NotNil(t, result.ID)
					if !cmp.Equal(tc.expResult, result,
						cmpopts.IgnoreFields(model.File{}, "CreatedAt", "UpdatedAt", "ID")) {
						t.Errorf("\n File mismatched. \n expected: %+v \n got: %+v \n diff: %+v", tc.expResult, result,
							cmp.Diff(tc.expResult, result, cmpopts.IgnoreFields(model.File{}, "CreatedAt", "UpdatedAt", "ID")))
						t.FailNow()
					}
				}
			})
		})
	}
}
