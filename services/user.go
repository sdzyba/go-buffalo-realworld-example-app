package services

import (
	"github.com/pkg/errors"
	"github.com/sdzyba/go-buffalo-realworld-example-app/models"
	"golang.org/x/crypto/bcrypt"
)

type UserCreateParams struct {
	User struct {
		Username string `json:"username" validate:"required"`
		Email    string `json:"email"    validate:"required,email"`
		Password string `json:"password" validate:"required"`
	} `json:"user"`
}

type User struct {
}

func (u *User) Create(params *UserCreateParams) (*models.User, error) {
	err := validate.Struct(params)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	password, err := bcrypt.GenerateFromPassword([]byte(params.User.Password), bcrypt.MinCost)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	user := &models.User{
		Email:             params.User.Email,
		EncryptedPassword: string(password),
		Username:          params.User.Username,
	}

	return userRepository.Create(user)
}
