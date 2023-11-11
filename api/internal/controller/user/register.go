package user

import (
	"context"
	"errors"
	"github.com/petme/api/internal/model"

	"golang.org/x/crypto/bcrypt"
)

var (
	// use for unit test
	generatePasswordFunc = generatePassword
)

func generatePassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

// Register create new user
func (i impl) Register(ctx context.Context, input model.User) (model.User, error) {
	// check if username exist in db
	existingUser, err := i.repo.User().GetUserByUsername(ctx, input.Username)
	if err != nil {
		return model.User{}, err
	}
	if existingUser.ID != 0 {
		return model.User{}, errors.New("user already exists")
	}

	// Hashing the password with the default cost of 10
	hashedPassword, err := generatePasswordFunc(input.Password)
	if err != nil {
		return model.User{}, err
	}

	input.Password = hashedPassword
	user, err := i.repo.User().InsertUser(ctx, input)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
