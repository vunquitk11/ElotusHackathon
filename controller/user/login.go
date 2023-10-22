package user

import (
	"context"
	"github.com/elotus_hackathon/model"
	"golang.org/x/crypto/bcrypt"
)

// Login logins user to system, return jwt token
func (i impl) Login(ctx context.Context, input model.User) (string, error) {
	// get user by username
	user, err := i.repo.User().GetUserByUsername(ctx, input.Username)
	if err != nil {
		return "", err
	}

	if user.ID == 0 {
		return "", nil
	}

	// Comparing the password with the hash
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return "", err
	}

	return "default-jwt", nil
}
