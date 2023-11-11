package user

import (
	"context"
	"errors"
	"github.com/petme/api/internal/model"
	"github.com/petme/api/internal/repository"
	"github.com/petme/api/internal/repository/user"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// TestController_Register means unit test of function Register in user controller
func TestController_Register(t *testing.T) {
	type mockGeneratePassword struct {
		call bool
		out  string
		err  error
	}
	type mockRepoInsert struct {
		call bool
		out  model.User
		err  error
	}
	type mockRepoGetUser struct {
		call bool
		out  model.User
		err  error
	}
	type arg struct {
		givenInput           model.User
		mockRepoGetUser      mockRepoGetUser
		mockRepoInsert       mockRepoInsert
		mockGeneratePassword mockGeneratePassword
		expRes               model.User
		expErr               error
	}

	validInput := model.User{
		Username: "username1",
		Password: "123456",
	}

	tcs := map[string]arg{
		"error when GetUserByUsername": {
			givenInput: validInput,
			mockRepoGetUser: mockRepoGetUser{
				call: true,
				err:  errors.New("something went wrong"),
			},
			expErr: errors.New("something went wrong"),
		},
		"error username existing in db": {
			givenInput: validInput,
			mockRepoGetUser: mockRepoGetUser{
				call: true,
				out: model.User{
					ID:       1,
					Username: "username1",
				},
			},
			expErr: errors.New("user already exists"),
		},
		"error hash password": {
			givenInput: validInput,
			mockRepoGetUser: mockRepoGetUser{
				call: true,
				out:  model.User{},
			},
			mockGeneratePassword: mockGeneratePassword{
				call: true,
				err:  errors.New("something went wrong"),
			},
			expErr: errors.New("something went wrong"),
		},
		"error when insert user": {
			givenInput: validInput,
			mockRepoGetUser: mockRepoGetUser{
				call: true,
				out:  model.User{},
			},
			mockGeneratePassword: mockGeneratePassword{
				call: true,
				out:  "hashed_password",
			},
			mockRepoInsert: mockRepoInsert{
				call: true,
				err:  errors.New("something went wrong"),
			},
			expErr: errors.New("something went wrong"),
		},
		"success": {
			givenInput: validInput,
			mockRepoGetUser: mockRepoGetUser{
				call: true,
				out:  model.User{},
			},
			mockGeneratePassword: mockGeneratePassword{
				call: true,
				out:  "hashed_password",
			},
			mockRepoInsert: mockRepoInsert{
				call: true,
				out: model.User{
					ID:        1,
					Username:  "username1",
					Password:  "hashed_password",
					CreatedAt: time.Date(2023, 12, 30, 7, 35, 48, 0, time.UTC),
					UpdatedAt: time.Date(2023, 12, 30, 7, 35, 48, 0, time.UTC),
				},
			},
			expRes: model.User{
				ID:        1,
				Username:  "username1",
				Password:  "hashed_password",
				CreatedAt: time.Date(2023, 12, 30, 7, 35, 48, 0, time.UTC),
				UpdatedAt: time.Date(2023, 12, 30, 7, 35, 48, 0, time.UTC),
			},
		},
	}

	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {

			defer func() {
				generatePasswordFunc = generatePassword
			}()

			if tc.mockGeneratePassword.call {
				generatePasswordFunc = func(password string) (string, error) {
					require.Equal(t, tc.givenInput.Password, password)
					return tc.mockGeneratePassword.out, tc.mockGeneratePassword.err
				}
			}

			// Given
			mockRegistry := new(repository.MockRegistry)
			mockUserRepo := new(user.MockRepository)
			if tc.mockRepoGetUser.call {
				mockRegistry.On("User").Return(mockUserRepo)
				mockUserRepo.On("GetUserByUsername", mock.Anything, tc.givenInput.Username).
					Return(tc.mockRepoGetUser.out, tc.mockRepoGetUser.err)
			}

			if tc.mockRepoInsert.call {
				mockRegistry.On("User").Return(mockUserRepo)
				mockUserRepo.On("InsertUser", mock.Anything, model.User{
					Username: tc.givenInput.Username,
					Password: tc.mockGeneratePassword.out,
				}).
					Return(tc.mockRepoInsert.out, tc.mockRepoInsert.err)
			}

			// When
			instance := New(mockRegistry)
			result, err := instance.Register(context.Background(), tc.givenInput)

			if tc.expErr != nil || err != nil {
				require.Equal(t, tc.expErr, err)
			} else {
				require.Equal(t, tc.expRes, result)
			}
		})
	}
}
