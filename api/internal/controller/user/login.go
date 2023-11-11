package user

import (
	"context"
	"errors"
	model2 "github.com/petme/api/internal/model"

	"golang.org/x/crypto/bcrypt"
)

// Login logins user to system, return jwt token
func (i impl) Login(ctx context.Context, input model2.User) (model2.User, error) {
	// get user by username
	user, err := i.repo.User().GetUserByUsername(ctx, input.Username)
	if err != nil {
		return model2.User{}, err
	}
	if user.ID == 0 {
		return model2.User{}, model2.ErrUserNotFound
	}

	// Comparing the password with the hash
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return model2.User{}, errors.New("incorrect password")
	}

	return user, nil
}
