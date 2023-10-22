package user

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

func Test_InsertUser(t *testing.T) {
	t.Setenv("DB_URL", "postgres://postgres:postgres@localhost:5432/elotus?sslmode=disable")
	type arg struct {
		givenInput model.User
		expResult  model.User
		expErr     error
	}

	tcs := map[string]arg{
		"success": {
			givenInput: model.User{
				Username: "username1",
				Password: "password1",
			},
			expResult: model.User{
				ID:       1,
				Username: "username1",
				Password: "password1",
			},
		},
	}
	for scenario, tc := range tcs {
		t.Run(scenario, func(t *testing.T) {
			testutil.WithTxDB(t, func(tx pg.BeginnerExecutor) {
				// When
				instance := New(tx)
				result, err := instance.InsertUser(context.Background(), tc.givenInput)

				// Then
				if tc.expErr != nil {
					require.EqualError(t, err, tc.expErr.Error())
				} else {
					require.NoError(t, err)
					require.NotNil(t, result.ID)
					if !cmp.Equal(tc.expResult, result,
						cmpopts.IgnoreFields(model.User{}, "CreatedAt", "UpdatedAt", "ID")) {
						t.Errorf("\n User mismatched. \n expected: %+v \n got: %+v \n diff: %+v", tc.expResult, result,
							cmp.Diff(tc.expResult, result, cmpopts.IgnoreFields(model.User{}, "CreatedAt", "UpdatedAt", "ID")))
						t.FailNow()
					}
				}
			})
		})
	}
}
