package user

import (
	"context"
	"errors"

	"github.com/petme/api/internal/model"
	"golang.org/x/crypto/bcrypt"
)

// Login logins user to system, return jwt token
func (i impl) Login(ctx context.Context, input model.User) (model.User, error) {
	// get user by username
	user, err := i.repo.User().GetUserByUsername(ctx, input.Username)
	if err != nil {
		return model.User{}, err
	}
	if user.ID == 0 {
		return model.User{}, model.ErrUserNotFound
	}

	// Comparing the password with the hash
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return model.User{}, errors.New("incorrect password")
	}

	return user, nil
}
